interface Env {
  ASSETS: Fetcher;
  KV?: KVNamespace;
  BUTTONDOWN_API_KEY?: string;
}

export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    const url = new URL(request.url);

    if (url.pathname.startsWith("/api/")) {
      return handleApi(url.pathname, request, env);
    }

    return env.ASSETS.fetch(request);
  },
} satisfies ExportedHandler<Env>;

function json(data: unknown, status = 200): Response {
  return new Response(JSON.stringify(data), {
    status,
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*",
      "Cache-Control": "no-store",
    },
  });
}

function corsPreflightResponse(): Response {
  return new Response(null, {
    headers: {
      "Access-Control-Allow-Origin": "https://ross.gg",
      "Access-Control-Allow-Methods": "POST, OPTIONS",
      "Access-Control-Allow-Headers": "Content-Type",
    },
  });
}

async function handleApi(
  path: string,
  request: Request,
  env: Env
): Promise<Response> {
  if (request.method === "OPTIONS") {
    return corsPreflightResponse();
  }

  try {
    switch (path) {
      case "/api/subscribe":
        return handleSubscribe(request, env);
      default:
        return json({ error: "Not found" }, 404);
    }
  } catch (err) {
    const message = err instanceof Error ? err.message : "Internal error";
    return json({ error: message }, 500);
  }
}

async function handleSubscribe(request: Request, env: Env): Promise<Response> {
  if (request.method !== "POST") {
    return json({ error: "Method not allowed" }, 405);
  }

  const formData = await request.formData();
  const email = formData.get("email")?.toString().trim().toLowerCase();

  if (!email || !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
    return Response.redirect("https://ross.gg/subscribe/error/", 303);
  }

  if (env.BUTTONDOWN_API_KEY) {
    const res = await fetch("https://api.buttondown.email/v1/subscribers", {
      method: "POST",
      headers: {
        Authorization: `Token ${env.BUTTONDOWN_API_KEY}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, type: "regular" }),
    });

    if (!res.ok && res.status !== 409) {
      return Response.redirect("https://ross.gg/subscribe/error/", 303);
    }
  } else if (env.KV) {
    await env.KV.put(`subscriber:${email}`, new Date().toISOString());
  } else {
    return Response.redirect("https://ross.gg/subscribe/error/", 303);
  }

  return Response.redirect("https://ross.gg/subscribe/thanks/", 303);
}
