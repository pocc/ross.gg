interface Env {
  ASSETS: Fetcher;
  KV?: KVNamespace;
  BUTTONDOWN_API_KEY?: string;
  STATUS_AUTH_TOKEN?: string;
}

export default {
  async fetch(request: Request, env: Env, ctx: ExecutionContext): Promise<Response> {
    const url = new URL(request.url);

    if (url.pathname.startsWith('/api/')) {
      return handleApi(url.pathname, request, env, ctx);
    }

    return env.ASSETS.fetch(request);
  },
} satisfies ExportedHandler<Env>;

// --- Helpers ---

function json(data: unknown, status = 200, cacheSeconds = 0, origin = '*'): Response {
  return new Response(JSON.stringify(data), {
    status,
    headers: {
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin': origin,
      'Cache-Control': cacheSeconds > 0 ? `public, max-age=${cacheSeconds}` : 'no-store',
    },
  });
}

function corsPreflightResponse(): Response {
  return new Response(null, {
    headers: {
      'Access-Control-Allow-Origin': 'https://ross.gg',
      'Access-Control-Allow-Methods': 'GET, POST, PUT, OPTIONS',
      'Access-Control-Allow-Headers': 'Content-Type, Authorization',
    },
  });
}

// --- Router ---

async function handleApi(
  path: string,
  request: Request,
  env: Env,
  ctx: ExecutionContext,
): Promise<Response> {
  if (request.method === 'OPTIONS') {
    return corsPreflightResponse();
  }

  try {
    switch (path) {
      case '/api/status':
        if (request.method === 'PUT') return handleStatusUpdate(request, env);
        return handleStatusGet(env);
      case '/api/subscribe':
        return handleSubscribe(request, env);
      case '/api/github':
        return handleGithub(env, ctx);
      case '/api/stackoverflow':
        return handleStackOverflow(env, ctx);
      default:
        return json({ error: 'Not found' }, 404);
    }
  } catch (err) {
    const message = err instanceof Error ? err.message : 'Internal error';
    return json({ error: message }, 500);
  }
}

// --- /api/status ---

async function handleStatusGet(env: Env): Promise<Response> {
  if (!env.KV) {
    return json({ status: 'Building things at Cloudflare', location: null, updatedAt: null }, 200, 300);
  }

  const [status, location, updatedAt] = await Promise.all([
    env.KV.get('current-status'),
    env.KV.get('current-location'),
    env.KV.get('status-updated-at'),
  ]);

  return json({
    status: status || 'Building things at Cloudflare',
    location,
    updatedAt,
  }, 200, 300);
}

async function handleStatusUpdate(request: Request, env: Env): Promise<Response> {
  const auth = request.headers.get('Authorization');
  if (!env.STATUS_AUTH_TOKEN || auth !== `Bearer ${env.STATUS_AUTH_TOKEN}`) {
    return json({ error: 'Unauthorized' }, 401);
  }

  if (!env.KV) {
    return json({ error: 'Storage unavailable' }, 503, 0, 'https://ross.gg');
  }

  let body: { status?: string; location?: string };
  try {
    body = await request.json() as { status?: string; location?: string };
  } catch {
    return json({ error: 'Invalid JSON' }, 400, 0, 'https://ross.gg');
  }

  const writes: Promise<void>[] = [];
  if (body.status) writes.push(env.KV.put('current-status', body.status));
  if (body.location) writes.push(env.KV.put('current-location', body.location));
  writes.push(env.KV.put('status-updated-at', new Date().toISOString()));
  await Promise.all(writes);

  return json({ success: true }, 200, 0, 'https://ross.gg');
}

// --- /api/subscribe ---

async function handleSubscribe(request: Request, env: Env): Promise<Response> {
  if (request.method !== 'POST') {
    return json({ error: 'Method not allowed' }, 405);
  }

  const formData = await request.formData();
  const email = formData.get('email')?.toString().trim().toLowerCase();

  if (!email || !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
    return Response.redirect('https://ross.gg/subscribe/error/', 303);
  }

  if (env.BUTTONDOWN_API_KEY) {
    const res = await fetch('https://api.buttondown.email/v1/subscribers', {
      method: 'POST',
      headers: {
        Authorization: `Token ${env.BUTTONDOWN_API_KEY}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, type: 'regular' }),
    });

    if (!res.ok && res.status !== 409) {
      return Response.redirect('https://ross.gg/subscribe/error/', 303);
    }
  } else if (env.KV) {
    await env.KV.put(`subscriber:${email}`, new Date().toISOString());
  } else {
    return Response.redirect('https://ross.gg/subscribe/error/', 303);
  }

  return Response.redirect('https://ross.gg/subscribe/thanks/', 303);
}

// --- /api/github ---

async function handleGithub(env: Env, ctx: ExecutionContext): Promise<Response> {
  if (env.KV) {
    const cached = await env.KV.get('github-activity', 'json');
    if (cached) return json(cached, 200, 3600);
  }

  const res = await fetch('https://api.github.com/users/pocc/events?per_page=10', {
    headers: { 'User-Agent': 'ross.gg-worker' },
  });

  if (!res.ok) {
    return json({ error: 'GitHub API error' }, 502);
  }

  const events = (await res.json()) as Array<{
    type: string;
    repo: { name: string };
    created_at: string;
  }>;

  const summary = events
    .filter((e) => ['PushEvent', 'CreateEvent', 'PullRequestEvent'].includes(e.type))
    .slice(0, 3)
    .map((e) => ({
      type: e.type.replace('Event', ''),
      repo: e.repo.name,
      date: e.created_at,
    }));

  if (env.KV) {
    ctx.waitUntil(env.KV.put('github-activity', JSON.stringify(summary), { expirationTtl: 3600 }));
  }

  return json(summary, 200, 3600);
}

// --- /api/stackoverflow ---

async function handleStackOverflow(env: Env, ctx: ExecutionContext): Promise<Response> {
  if (env.KV) {
    const cached = await env.KV.get('so-reputation', 'json');
    if (cached) return json(cached, 200, 21600);
  }

  const res = await fetch(
    'https://api.stackexchange.com/2.3/users/1596750?site=stackoverflow',
    { headers: { 'User-Agent': 'ross.gg-worker', 'Accept-Encoding': 'gzip' } },
  );

  if (!res.ok) {
    return json({ error: 'Stack Overflow API error' }, 502);
  }

  const data = (await res.json()) as {
    items: Array<{
      reputation: number;
      badge_counts: { gold: number; silver: number; bronze: number };
      display_name: string;
    }>;
  };

  const user = data.items?.[0];
  if (!user) return json({ error: 'User not found' }, 404);

  const summary = {
    reputation: user.reputation,
    badges: user.badge_counts,
    name: user.display_name,
  };

  if (env.KV) {
    ctx.waitUntil(env.KV.put('so-reputation', JSON.stringify(summary), { expirationTtl: 21600 }));
  }

  return json(summary, 200, 21600);
}
