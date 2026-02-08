# Ross.gg Redesign: Digital Garden Recommendations

## Table of Contents

1. [Design Philosophy](#design-philosophy)
2. [Task 1: Content Architecture](#task-1-content-architecture)
3. [Task 2: Visual & UI Modernization](#task-2-visual--ui-modernization)
4. [Task 3: Functionality](#task-3-functionality)
5. [Hugo Folder Structure](#hugo-folder-structure)
6. [Tailwind CSS Configuration](#tailwind-css-configuration)
7. [Bento Homepage Roadmap](#bento-homepage-roadmap)

---

## Design Philosophy

**"Sophisticated Tech-Noir"** — The site should feel like walking into a well-lit
workshop: clean surfaces, deliberate tool placement, and evidence of ongoing work.
Not sterile, not cluttered. A space where a Cloudflare engineer's technical depth
coexists naturally with tea ceremony and board game strategy.

**Core Principles:**
- Content-first: Typography and whitespace do the heavy lifting, not decoration
- Progressive disclosure: The homepage teases depth; each click rewards curiosity
- Living site: Visible signals that the site is actively tended (last-updated dates, growth indicators, status)
- No JavaScript frameworks: Hugo templates + vanilla JS + Cloudflare Workers for dynamic bits

---

## Task 1: Content Architecture

### 1.1 Digital Garden Layout — Three Content Tiers

Structure content into three distinct types, each with its own visual treatment
and Hugo section:

#### Tier 1: Polished Essays (`content/essays/`)

Long-form, finished pieces. These are the "front door" content — linked from
the homepage, shared on social media, indexed for SEO.

```yaml
# Frontmatter for essays
---
title: "Why FTP Still Matters in 2026"
date: 2026-02-07
description: "A defense of the protocol everyone loves to hate"
tags: ["networking", "protocols", "cloudflare"]
series: "Protocol Deep Dives"       # optional grouping
stage: "evergreen"                  # always "evergreen" for essays
cover: "/img/essays/ftp-cover.webp" # optional hero image
toc: true
---
```

**Visual treatment:** Full-width layout, hero image optional, prominent
typography, table of contents sidebar on desktop, reading progress bar.

#### Tier 2: Technical Notes (`content/notes/`)

Obsidian-style working notes. Shorter, less polished, interconnected via
backlinks. These are the "workshop" — visible but understood to be in progress.

```yaml
# Frontmatter for notes
---
title: "Cloudflare Workers Socket API Quirks"
date: 2026-02-07
lastmod: 2026-02-07                 # show "tended on" date
tags: ["cloudflare", "workers", "sockets"]
stage: "seedling"                   # seedling | budding | evergreen
related:                            # explicit backlinks (supplement auto-detected)
  - "/notes/ftp-passive-mode"
  - "/essays/why-ftp-still-matters"
---
```

**Growth stages** (shown as small icons next to titles):
- **Seedling** — Just planted. Rough, possibly incomplete.
- **Budding** — Has structure and substance, but still growing.
- **Evergreen** — Mature, reliable, regularly maintained.

**Visual treatment:** Narrower reading column, visible metadata bar showing
stage + last tended date, backlinks section at bottom, lighter styling to
signal "working document."

#### Tier 3: Current Interests (`content/interests/`)

Non-technical, personal content. Board games, tea, home automation. These
humanize the site and create unexpected entry points.

```yaml
# Frontmatter for interests
---
title: "Learning Brass: Birmingham"
date: 2026-01-15
tags: ["board-games", "strategy", "economic"]
category: "board-games"             # board-games | tea | home-automation
rating: 9                           # optional, 1-10
status: "currently-playing"         # currently-playing | completed | on-shelf
---
```

**Visual treatment:** Card-based grid layout, cover images prominent,
shorter-form content, filter by category.

### 1.2 The "Now" Page (`content/now.md`)

Inspired by [Derek Sivers' /now movement](https://nownownow.com/about). A
single page answering "What is Ross focused on right now?" Updated manually
every 2-4 weeks.

**Structure:**

```markdown
---
title: "Now"
description: "What I'm focused on right now"
layout: "now"
lastmod: 2026-02-07
---

## Work
Building FTP protocol support for Port of Call using Cloudflare Workers
Sockets API. Debugging passive mode data connections.

## Learning
- Deep-diving into the FTP RFC (959) and EPSV extensions
- Studying Hive (board game) opening theory

## Reading
- "Designing Data-Intensive Applications" by Martin Kleppmann (re-read)
- "The Classic of Tea" by Lu Yu (translation by Francis Ross Carpenter)

## Playing
- **Brass: Birmingham** — 4-player campaign with weekly group
- **Hive Pocket** — Online ranked matches

## Drinking
- Tie Guan Yin oolong from Anxi, Fujian province
- Experimenting with gongfu brewing parameters

## Home Automation
- Migrating Zigbee devices from Hue Bridge to Zigbee2MQTT
- Building a tea water temperature controller with ESP32

---
*This is a [/now page](https://nownownow.com/about). Last updated February 7, 2026.*
```

**Hugo layout** (`layouts/page/now.html`): Render with a distinct template that
shows `lastmod` prominently at the top ("Updated 3 weeks ago"), uses a
two-column layout on desktop (work/learning on left, personal on right), and
includes a subtle "subscribe to updates" link.

### 1.3 Bookshelf / Game Shelf Component

A data-driven component using Hugo's data templates. Store structured data in
`data/` files, render with a partial.

**Data files:**

```yaml
# data/games.yaml
- title: "Brass: Birmingham"
  designer: "Gavan Brown, Matt Tolman, Martin Wallace"
  bgg_id: 224517
  cover: "/img/shelf/brass-birmingham.webp"
  status: "currently-playing"
  rating: 9.5
  plays: 23
  notes: "Best economic game ever made. Canal era strategy is underrated."
  tags: ["economic", "heavy", "2-4 players"]

- title: "Hive Pocket"
  designer: "John Yianni"
  bgg_id: 154597
  cover: "/img/shelf/hive.webp"
  status: "currently-playing"
  rating: 8
  plays: 100+
  notes: "Chess-like depth in a bag. Perfect travel game."
  tags: ["abstract", "2-player", "portable"]

- title: "Spirit Island"
  designer: "R. Eric Reuss"
  bgg_id: 162886
  cover: "/img/shelf/spirit-island.webp"
  status: "on-shelf"
  rating: 9
  plays: 15
  notes: "Cooperative anti-colonial tower defense. Incredible theming."
  tags: ["cooperative", "heavy", "1-4 players"]
```

```yaml
# data/books.yaml
- title: "Designing Data-Intensive Applications"
  author: "Martin Kleppmann"
  cover: "/img/shelf/ddia.webp"
  status: "reading"          # reading | finished | want-to-read
  rating: 10
  category: "technical"
  started: 2026-01-10
  notes: "The networking chapters are especially relevant to my work."

- title: "The Classic of Tea"
  author: "Lu Yu (trans. Francis Ross Carpenter)"
  cover: "/img/shelf/classic-of-tea.webp"
  status: "reading"
  rating: 8
  category: "personal"
  started: 2026-01-20
```

**Hugo partial** (`layouts/partials/shelf.html`): Renders a filterable grid
of cards. Each card shows cover image, title, status badge, and rating.
Clicking opens a detail view with notes. Filter tabs: "Currently
Playing/Reading" | "Completed" | "Want to Play/Read."

No JavaScript framework needed — use Hugo to render all states as CSS-toggled
sections, or a tiny vanilla JS filter (~20 lines).

### 1.4 Obsidian Integration

**Recommended approach: `obsidian-export` + Hugo render hooks**

This is the most reliable pipeline for Hugo specifically. Quartz is the other
major option, but it replaces Hugo entirely — since you want to stay on Hugo,
the render-hook approach is better.

#### Pipeline

```
Obsidian Vault → obsidian-export (Rust CLI) → content/notes/ → Hugo build
```

1. **[obsidian-export](https://github.com/zoni/obsidian-export)** — Rust CLI
   tool that converts Obsidian vault files to standard Markdown:
   - Resolves `[[wikilinks]]` to standard `[text](url)` links
   - Handles embeds (`![[note]]`)
   - Strips Obsidian-specific syntax
   - Preserves frontmatter

   ```bash
   # In Makefile
   export-notes:
       obsidian-export /path/to/vault/publish content/notes/ \
           --frontmatter=always \
           --hard-linebreaks=false
   ```

2. **Selective publishing** — In your Obsidian vault, use a `publish/`
   subfolder or a frontmatter flag (`publish: true`) to control which notes
   get exported. Only notes you explicitly mark will appear on the site.

3. **Hugo render hooks for backlinks** — Create a custom render hook to
   auto-generate a "Backlinks" section at the bottom of each note:

   ```html
   <!-- layouts/partials/backlinks.html -->
   {{ $currentPage := . }}
   {{ $backlinks := slice }}
   {{ range where .Site.RegularPages "Section" "notes" }}
     {{ if and (ne .Permalink $currentPage.Permalink) (findRE $currentPage.RelPermalink .RawContent) }}
       {{ $backlinks = $backlinks | append . }}
     {{ end }}
   {{ end }}

   {{ if gt (len $backlinks) 0 }}
   <aside class="backlinks">
     <h3>Pages that link here</h3>
     <ul>
       {{ range $backlinks }}
       <li>
         <a href="{{ .RelPermalink }}">{{ .Title }}</a>
         <span class="stage-badge">{{ .Params.stage }}</span>
       </li>
       {{ end }}
     </ul>
   </aside>
   {{ end }}
   ```

4. **Wikilink render hook** (fallback for any `[[links]]` that
   `obsidian-export` misses):

   ```html
   <!-- layouts/_default/_markup/render-link.html -->
   {{ $url := .Destination }}
   {{ $text := .Text }}

   {{/* Convert wiki-style paths to Hugo paths */}}
   {{ if hasPrefix $url "notes/" }}
     {{ $url = printf "/%s/" $url }}
   {{ end }}

   <a href="{{ $url | safeURL }}">{{ $text }}</a>
   ```

#### Workflow Summary

```
1. Write notes in Obsidian normally (use [[wikilinks]])
2. Tag notes for publishing (move to publish/ folder or set publish: true)
3. Run `make export-notes` (obsidian-export converts to Hugo-compatible MD)
4. Run `make build` (Hugo builds with backlink detection)
5. Push to deploy
```

This can be automated into a single `make publish` target.

---

## Task 2: Visual & UI Modernization

### 2.1 Bento Grid Homepage

The homepage is a CSS Grid of cards ("bento boxes") at varying sizes. Each
card is a window into a different facet of the site. The grid is responsive:
4 columns on desktop, 2 on tablet, 1 on mobile.

**Grid Layout (Desktop — 4 columns, auto rows):**

```
┌─────────────────────────┬───────────┬───────────┐
│                         │  Latest   │   /now     │
│   Hero / Identity       │  Essay    │  snippet   │
│   "Ross Jacobs"         │           │            │
│   Subtitle + status     ├───────────┼───────────┤
│                         │  GitHub   │  Tea of    │
│                         │  Activity │  the Week  │
├────────────┬────────────┼───────────┴───────────┤
│  Featured  │  Board     │                        │
│  Technical │  Game of   │  Recent Notes          │
│  Note      │  the Month │  (3-4 seedling links)  │
│            │            │                        │
├────────────┴────────────┼───────────┬───────────┤
│                         │  Stack    │  Reading   │
│  Newsletter signup      │  Overflow │  Now       │
│                         │  Rep      │            │
└─────────────────────────┴───────────┴───────────┘
```

**Card types and their data sources:**

| Card | Size | Data Source | Update Frequency |
|------|------|-------------|------------------|
| Hero/Identity | 2x2 | `config.toml` + CF Worker status | Static + live |
| Latest Essay | 1x1 | Hugo (latest from `essays/`) | On publish |
| /now snippet | 1x1 | Hugo (excerpt from `now.md`) | Manual |
| GitHub Activity | 1x1 | CF Worker → GitHub API | Live (cached 1hr) |
| Tea of the Week | 1x1 | `data/tea.yaml` | Manual |
| Featured Note | 1x1 | Hugo (pinned note) | On publish |
| Board Game of Month | 1x1 | `data/games.yaml` | Manual |
| Recent Notes | 2x1 | Hugo (latest 4 from `notes/`) | On publish |
| Newsletter | 2x1 | Static HTML form | Static |
| SO Reputation | 1x1 | CF Worker → SO API | Live (cached 6hr) |
| Currently Reading | 1x1 | `data/books.yaml` | Manual |

**Hugo template** (`layouts/index.html`):

```html
{{ define "main" }}
<section class="bento-grid">

  <!-- Hero: 2 cols, 2 rows -->
  <article class="bento-card bento-hero col-span-2 row-span-2">
    <h1>Ross Jacobs</h1>
    <p class="subtitle">{{ .Site.Params.subtitle }}</p>
    <div id="live-status" data-worker="/api/status">
      <!-- Hydrated by CF Worker -->
    </div>
    <nav class="hero-links">
      <a href="/about">About</a>
      <a href="/essays">Essays</a>
      <a href="/notes">Notes</a>
      <a href="/now">Now</a>
    </nav>
  </article>

  <!-- Latest Essay -->
  {{ with (index (where .Site.RegularPages "Section" "essays") 0) }}
  <article class="bento-card bento-essay">
    <span class="card-label">Latest Essay</span>
    <h3><a href="{{ .RelPermalink }}">{{ .Title }}</a></h3>
    <time>{{ .Date.Format "Jan 2, 2006" }}</time>
    <p>{{ .Description }}</p>
  </article>
  {{ end }}

  <!-- /now snippet -->
  {{ with .Site.GetPage "/now" }}
  <article class="bento-card bento-now">
    <span class="card-label">/now</span>
    <p>{{ .Summary | truncate 120 }}</p>
    <a href="/now">What I'm up to &rarr;</a>
    <time class="last-updated">Updated {{ .Lastmod.Format "Jan 2" }}</time>
  </article>
  {{ end }}

  <!-- GitHub Activity (hydrated by Worker) -->
  <article class="bento-card bento-github" data-worker="/api/github">
    <span class="card-label">Building</span>
    <div id="github-activity">
      <noscript>See my work on <a href="https://github.com/pocc">GitHub</a></noscript>
    </div>
  </article>

  <!-- Tea of the Week -->
  {{ with index .Site.Data.tea 0 }}
  <article class="bento-card bento-tea">
    <span class="card-label">Drinking</span>
    <h3>{{ .name }}</h3>
    <p class="tea-origin">{{ .origin }}</p>
    <p class="tea-notes">{{ .notes }}</p>
  </article>
  {{ end }}

  <!-- Featured Note -->
  {{ range first 1 (where (where .Site.RegularPages "Section" "notes") ".Params.featured" true) }}
  <article class="bento-card bento-note">
    <span class="card-label">{{ .Params.stage | default "note" }}</span>
    <h3><a href="{{ .RelPermalink }}">{{ .Title }}</a></h3>
    <p>{{ .Description }}</p>
  </article>
  {{ end }}

  <!-- Board Game of the Month -->
  {{ range first 1 (where .Site.Data.games "status" "currently-playing") }}
  <article class="bento-card bento-game">
    <span class="card-label">Playing</span>
    {{ with .cover }}<img src="{{ . }}" alt="{{ $.title }}" loading="lazy">{{ end }}
    <h3>{{ .title }}</h3>
    <p>{{ .plays }} plays &middot; {{ .rating }}/10</p>
  </article>
  {{ end }}

  <!-- Recent Notes -->
  <article class="bento-card bento-recent-notes col-span-2">
    <span class="card-label">Garden</span>
    <ul>
      {{ range first 4 (where .Site.RegularPages "Section" "notes") }}
      <li>
        <span class="stage-icon stage-{{ .Params.stage }}"></span>
        <a href="{{ .RelPermalink }}">{{ .Title }}</a>
        <time>{{ .Lastmod.Format "Jan 2" }}</time>
      </li>
      {{ end }}
    </ul>
    <a href="/notes">All notes &rarr;</a>
  </article>

  <!-- Newsletter -->
  <article class="bento-card bento-newsletter col-span-2">
    {{ partial "newsletter-form.html" . }}
  </article>

  <!-- SO Reputation (hydrated by Worker) -->
  <article class="bento-card bento-so" data-worker="/api/stackoverflow">
    <span class="card-label">Stack Overflow</span>
    <div id="so-rep">
      <noscript><a href="https://stackoverflow.com/users/1596750">Profile</a></noscript>
    </div>
  </article>

  <!-- Currently Reading -->
  {{ range first 1 (where .Site.Data.books "status" "reading") }}
  <article class="bento-card bento-reading">
    <span class="card-label">Reading</span>
    {{ with .cover }}<img src="{{ . }}" alt="{{ $.title }}" loading="lazy">{{ end }}
    <h3>{{ .title }}</h3>
    <p class="book-author">{{ .author }}</p>
  </article>
  {{ end }}

</section>
{{ end }}
```

**CSS for the Bento Grid** (Tailwind utility classes, shown as equivalent CSS
for clarity):

```css
.bento-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 1rem;
  max-width: 72rem;
  margin: 0 auto;
  padding: 2rem 1rem;
}

/* Tablet */
@media (max-width: 1024px) {
  .bento-grid { grid-template-columns: repeat(2, 1fr); }
}

/* Mobile */
@media (max-width: 640px) {
  .bento-grid { grid-template-columns: 1fr; }
}

.bento-card {
  background: var(--card-bg);
  border: 1px solid var(--card-border);
  border-radius: 1rem;
  padding: 1.5rem;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.bento-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
}

.col-span-2 { grid-column: span 2; }
.row-span-2 { grid-row: span 2; }

/* On mobile, nothing spans multiple columns */
@media (max-width: 640px) {
  .col-span-2, .row-span-2 {
    grid-column: span 1;
    grid-row: span 1;
  }
}
```

### 2.2 Typography

**Recommended pairing: Fraunces (headings) + Inter (body)**

**Why Fraunces over Playfair Display:**
- Fraunces is a variable font with an "optical size" axis — it adapts from
  delicate display sizes to sturdy text sizes automatically
- It has a distinctive "wonky" axis that adds subtle personality at large sizes
  (the soft irregularity of the serifs feels handcrafted, not corporate)
- It pairs better with geometric sans-serifs than Playfair, which fights for
  attention

**Why Inter for body:**
- Designed specifically for screens; excellent legibility at small sizes
- Has tabular and proportional number variants (useful for data-heavy notes)
- Variable font — one file covers all weights
- Massive language support

**Alternative body option:** If Inter feels too common, consider **IBM Plex Sans**
(more distinctive character, similar readability, good monospace companion in
IBM Plex Mono for code blocks).

**Font loading strategy:**

```html
<!-- In <head>, preload critical fonts -->
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Fraunces:opsz,wght@9..144,300;700;900&family=Inter:wght@400;500;600&display=swap" rel="stylesheet">
```

Or better, self-host via Cloudflare Pages for zero external requests:

```
static/
  fonts/
    fraunces-variable.woff2
    inter-variable.woff2
```

```css
@font-face {
  font-family: 'Fraunces';
  src: url('/fonts/fraunces-variable.woff2') format('woff2');
  font-weight: 300 900;
  font-display: swap;
}

@font-face {
  font-family: 'Inter';
  src: url('/fonts/inter-variable.woff2') format('woff2');
  font-weight: 400 600;
  font-display: swap;
}
```

**Type scale** (using a 1.25 ratio — "Major Third"):

| Element | Font | Size | Weight | Tracking |
|---------|------|------|--------|----------|
| Hero h1 | Fraunces | 3.815rem (61px) | 900 | -0.02em |
| Page h1 | Fraunces | 3.052rem (49px) | 700 | -0.02em |
| h2 | Fraunces | 2.441rem (39px) | 700 | -0.01em |
| h3 | Fraunces | 1.953rem (31px) | 700 | 0 |
| h4 | Inter | 1.563rem (25px) | 600 | 0 |
| Body | Inter | 1.125rem (18px) | 400 | 0 |
| Small / metadata | Inter | 0.889rem (14px) | 500 | 0.01em |
| Code | JetBrains Mono | 0.95rem (15px) | 400 | 0 |

### 2.3 Cloudflare Workers for Live Data

Use Cloudflare Workers to serve lightweight JSON endpoints that the homepage
cards fetch on load. This gives you live data without rebuilding the site.

**Worker: `/api/status`**

Returns your current status message (set via a simple KV store or D1 database,
updatable from a Shortcuts/CLI command).

```typescript
// workers/status.ts
export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    const status = await env.KV.get('current-status');
    const location = await env.KV.get('current-location');

    return Response.json({
      status: status || 'Building things at Cloudflare',
      location: location || 'San Francisco, CA',
      updatedAt: await env.KV.get('status-updated-at'),
    }, {
      headers: {
        'Access-Control-Allow-Origin': 'https://ross.gg',
        'Cache-Control': 'public, max-age=300', // 5 min cache
      },
    });
  },
};

// Update status from CLI:
// curl -X PUT https://ross.gg/api/status \
//   -H "Authorization: Bearer $TOKEN" \
//   -d '{"status": "Debugging FTP passive mode", "location": "Home office"}'
```

**Worker: `/api/github`**

Returns recent GitHub activity (cached 1 hour).

```typescript
// workers/github.ts
export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    const cached = await env.KV.get('github-activity', 'json');
    if (cached) return Response.json(cached);

    const res = await fetch('https://api.github.com/users/pocc/events?per_page=5', {
      headers: { 'User-Agent': 'ross.gg-worker' },
    });
    const events = await res.json();

    const summary = events
      .filter((e: any) => ['PushEvent', 'CreateEvent', 'PullRequestEvent'].includes(e.type))
      .slice(0, 3)
      .map((e: any) => ({
        type: e.type,
        repo: e.repo.name,
        date: e.created_at,
      }));

    await env.KV.put('github-activity', JSON.stringify(summary), { expirationTtl: 3600 });

    return Response.json(summary, {
      headers: {
        'Access-Control-Allow-Origin': 'https://ross.gg',
        'Cache-Control': 'public, max-age=3600',
      },
    });
  },
};
```

**Client-side hydration** (tiny vanilla JS, inlined in the Hugo template):

```html
<script>
  document.querySelectorAll('[data-worker]').forEach(async (card) => {
    try {
      const res = await fetch(card.dataset.worker);
      if (!res.ok) return;
      const data = await res.json();
      card.dispatchEvent(new CustomEvent('worker-data', { detail: data }));
    } catch { /* graceful degradation — noscript content stays */ }
  });

  // Status card
  document.getElementById('live-status')?.addEventListener('worker-data', (e) => {
    const { status, location } = e.detail;
    e.target.innerHTML = `<p class="status-text">${status}</p><p class="status-location">${location}</p>`;
  });

  // GitHub card
  document.getElementById('github-activity')?.addEventListener('worker-data', (e) => {
    const html = e.detail.map(ev =>
      `<div class="gh-event"><span class="gh-repo">${ev.repo.split('/')[1]}</span><span class="gh-type">${ev.type.replace('Event','')}</span></div>`
    ).join('');
    e.target.innerHTML = html;
  });
</script>
```

This approach means:
- The Hugo build produces a fully functional static page (noscript fallbacks)
- Workers provide opt-in live data with aggressive caching
- Zero build-time dependencies on external APIs
- You can update your status from your phone via a simple PUT request

---

## Task 3: Functionality

### 3.1 Newsletter Integration

**Recommended: Cloudflare Workers + D1 (SQLite) + Buttondown**

Avoid embedding third-party JavaScript (Mailchimp, ConvertKit). Instead, use
a CF Worker as a proxy to keep the frontend clean and the signup fast.

**Frontend** (`layouts/partials/newsletter-form.html`):

```html
<div class="newsletter">
  <h3>Occasional dispatches</h3>
  <p>Technical notes on networking, protocols, and building things.
     No spam. Unsubscribe anytime.</p>
  <form action="/api/subscribe" method="POST" class="newsletter-form">
    <input
      type="email"
      name="email"
      placeholder="you@example.com"
      required
      autocomplete="email"
      class="newsletter-input"
    >
    <button type="submit" class="newsletter-button">Subscribe</button>
  </form>
  <p class="newsletter-meta">
    <span id="subscriber-count"></span> subscribers &middot; RSS also available
  </p>
</div>
```

**Worker** (`/api/subscribe`):

```typescript
export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    if (request.method !== 'POST') {
      return new Response('Method not allowed', { status: 405 });
    }

    const form = await request.formData();
    const email = form.get('email')?.toString().trim().toLowerCase();

    if (!email || !email.includes('@')) {
      return Response.redirect('https://ross.gg/subscribe/error', 303);
    }

    // Forward to Buttondown API (or store in D1 directly)
    const res = await fetch('https://api.buttondown.email/v1/subscribers', {
      method: 'POST',
      headers: {
        'Authorization': `Token ${env.BUTTONDOWN_API_KEY}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, type: 'regular' }),
    });

    if (res.ok || res.status === 409) { // 409 = already subscribed
      return Response.redirect('https://ross.gg/subscribe/thanks', 303);
    }

    return Response.redirect('https://ross.gg/subscribe/error', 303);
  },
};
```

**Why Buttondown:** It is minimalist, Markdown-native, has a generous free tier
(100 subscribers), and the API is simple. Alternatives: Listmonk (self-hosted)
or raw D1 + Resend/SES if you want full control.

The form degrades gracefully without JavaScript (POST redirect pattern). With
JS, you can enhance it to show inline success/error messages.

### 3.2 Reading Progress Bar

A thin, non-intrusive progress bar at the very top of the viewport. Only
appears on essay and note pages (not the homepage or shelf pages).

**Implementation** (vanilla JS, ~25 lines):

```html
<!-- layouts/partials/reading-progress.html -->
{{ if or (eq .Section "essays") (eq .Section "notes") }}
<div class="reading-progress" aria-hidden="true">
  <div class="reading-progress-bar" id="reading-progress"></div>
</div>

<style>
  .reading-progress {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 3px;
    z-index: 9999;
    background: transparent;
  }
  .reading-progress-bar {
    height: 100%;
    width: 0%;
    background: var(--accent);
    transition: width 50ms linear;
  }
</style>

<script>
  (function() {
    const bar = document.getElementById('reading-progress');
    if (!bar) return;

    const article = document.querySelector('article') || document.querySelector('main');
    if (!article) return;

    function updateProgress() {
      const rect = article.getBoundingClientRect();
      const total = rect.height - window.innerHeight;
      if (total <= 0) { bar.style.width = '100%'; return; }
      const progress = Math.min(Math.max(-rect.top / total, 0), 1);
      bar.style.width = (progress * 100) + '%';
    }

    let ticking = false;
    window.addEventListener('scroll', function() {
      if (!ticking) {
        requestAnimationFrame(function() { updateProgress(); ticking = false; });
        ticking = true;
      }
    }, { passive: true });

    updateProgress();
  })();
</script>
{{ end }}
```

**Design notes:**
- 3px height — visible but not distracting
- Uses `var(--accent)` — ties into the site color system
- `aria-hidden="true"` — decorative, not announced by screen readers
- `requestAnimationFrame` + passive scroll listener — no jank
- Only loads on content pages (Hugo conditional)

---

## Hugo Folder Structure

The complete proposed structure for the redesigned site:

```
ross.gg/
├── config/
│   └── _default/
│       ├── config.toml         # Base Hugo config (replaces root config.toml)
│       ├── params.toml         # Site parameters (social links, metadata)
│       ├── menus.toml          # Navigation menus
│       └── markup.toml         # Markdown rendering settings
│
├── content/
│   ├── _index.md               # Homepage content/metadata
│   ├── now.md                  # /now page
│   ├── about.md                # /about page (migrate existing)
│   ├── uses.md                 # /uses page (tools, hardware, setup)
│   ├── essays/
│   │   ├── _index.md           # Essays listing page
│   │   ├── why-ftp-matters.md
│   │   └── dns-caches.md       # Migrated from post/2021/
│   ├── notes/
│   │   ├── _index.md           # Notes garden listing page
│   │   ├── cloudflare-workers-sockets.md
│   │   ├── ftp-passive-mode.md
│   │   └── hugo-obsidian-pipeline.md
│   ├── interests/
│   │   ├── _index.md           # Interests listing page
│   │   ├── board-games/
│   │   │   ├── _index.md
│   │   │   ├── brass-birmingham.md
│   │   │   └── hive.md
│   │   ├── tea/
│   │   │   ├── _index.md
│   │   │   └── gongfu-brewing.md
│   │   └── home-automation/
│   │       ├── _index.md
│   │       └── zigbee2mqtt-migration.md
│   └── subscribe/
│       ├── thanks.md           # Post-subscribe confirmation
│       └── error.md            # Subscribe error page
│
├── data/
│   ├── games.yaml              # Board game collection data
│   ├── books.yaml              # Reading list data
│   ├── tea.yaml                # Tea rotation data
│   └── uses.yaml               # Tools/hardware for /uses page
│
├── layouts/
│   ├── _default/
│   │   ├── baseof.html         # Base template (fonts, meta, progress bar)
│   │   ├── single.html         # Default single page
│   │   ├── list.html           # Default list page
│   │   └── _markup/
│   │       └── render-link.html  # Wikilink resolver
│   ├── index.html              # Homepage (Bento grid)
│   ├── essays/
│   │   ├── single.html         # Essay layout (TOC sidebar, progress bar)
│   │   └── list.html           # Essays archive
│   ├── notes/
│   │   ├── single.html         # Note layout (backlinks, stage badge)
│   │   └── list.html           # Garden view (filterable by stage)
│   ├── interests/
│   │   ├── single.html         # Interest detail
│   │   └── list.html           # Card grid with category filter
│   ├── page/
│   │   └── now.html            # /now page layout
│   └── partials/
│       ├── head.html           # <head> with fonts, meta, OG tags
│       ├── header.html         # Site header/nav
│       ├── footer.html         # Site footer
│       ├── backlinks.html      # Backlinks section for notes
│       ├── shelf.html          # Bookshelf/game shelf component
│       ├── newsletter-form.html
│       ├── reading-progress.html
│       ├── bento/
│       │   ├── hero.html
│       │   ├── essay-card.html
│       │   ├── now-card.html
│       │   ├── github-card.html
│       │   ├── tea-card.html
│       │   ├── game-card.html
│       │   ├── notes-card.html
│       │   ├── newsletter-card.html
│       │   ├── so-card.html
│       │   └── reading-card.html
│       └── meta/
│           ├── opengraph.html
│           └── structured-data.html
│
├── assets/
│   ├── css/
│   │   ├── main.css            # Tailwind entry point (@tailwind directives)
│   │   └── components/
│   │       ├── bento.css       # Bento grid styles
│   │       ├── typography.css  # Type scale and prose styles
│   │       ├── cards.css       # Card component styles
│   │       └── garden.css      # Digital garden specific (stage badges, backlinks)
│   └── js/
│       └── main.js             # Worker hydration + progress bar (bundled by Hugo Pipes)
│
├── static/
│   ├── fonts/
│   │   ├── fraunces-variable.woff2
│   │   └── inter-variable.woff2
│   ├── img/
│   │   ├── shelf/              # Book/game cover images
│   │   ├── essays/             # Essay hero images
│   │   ├── og/                 # OpenGraph preview images
│   │   └── ...                 # (existing image directories)
│   ├── favicon.ico
│   └── ...                     # (existing favicon suite)
│
├── workers/
│   ├── status.ts               # /api/status — live status message
│   ├── github.ts               # /api/github — recent activity
│   ├── stackoverflow.ts        # /api/stackoverflow — reputation
│   └── subscribe.ts            # /api/subscribe — newsletter signup
│
├── Makefile                    # Build, export, deploy commands
├── tailwind.config.js          # Tailwind configuration
├── postcss.config.js           # PostCSS config (Tailwind + autoprefixer)
├── package.json                # Tailwind + PostCSS deps only
└── docs/
    ├── ARCHITECTURE.md
    ├── CONTENT.md
    └── PROMPT_RECOMMENDATIONS.md  # This file
```

**Key differences from current structure:**
- `config/` directory replaces single `config.toml` (Hugo's config directory feature)
- `content/` reorganized into semantic sections (essays, notes, interests)
- `layouts/` is fully custom (no longer dependent on a third-party theme)
- `assets/css/` uses Hugo Pipes for Tailwind processing
- `data/` holds structured YAML for shelves
- `workers/` holds Cloudflare Worker source (deployed separately)
- `package.json` added but only for Tailwind/PostCSS tooling (no JS framework)

**Migrating existing content:**
- `content/post/2021/dns_caches.md` → `content/essays/dns-caches.md`
- `content/post/2019/*.md` → `content/notes/` (they're more note-like in length)
- `content/about.md` → stays at `content/about.md`
- `content/projects.md` → integrated into `content/about.md` or homepage cards

---

## Tailwind CSS Configuration

```javascript
// tailwind.config.js

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './layouts/**/*.html',
    './content/**/*.md',
    './assets/js/**/*.js',
  ],

  theme: {
    fontFamily: {
      display: ['"Fraunces"', 'Georgia', '"Times New Roman"', 'serif'],
      body: ['"Inter"', 'system-ui', '-apple-system', 'sans-serif'],
      mono: ['"JetBrains Mono"', '"Fira Code"', 'Consolas', 'monospace'],
    },

    fontSize: {
      // Major Third scale (1.25 ratio), base 18px
      'xs':   ['0.72rem',  { lineHeight: '1.4' }],   // 11.5px — fine print
      'sm':   ['0.889rem', { lineHeight: '1.5' }],   // 14px — metadata, labels
      'base': ['1.125rem', { lineHeight: '1.7' }],   // 18px — body text
      'lg':   ['1.25rem',  { lineHeight: '1.6' }],   // 20px — lead paragraphs
      'xl':   ['1.563rem', { lineHeight: '1.4' }],   // 25px — h4
      '2xl':  ['1.953rem', { lineHeight: '1.3' }],   // 31px — h3
      '3xl':  ['2.441rem', { lineHeight: '1.2' }],   // 39px — h2
      '4xl':  ['3.052rem', { lineHeight: '1.1' }],   // 49px — h1
      '5xl':  ['3.815rem', { lineHeight: '1.05' }],  // 61px — hero
    },

    colors: {
      // Tech-Noir palette — dark with warm accents
      transparent: 'transparent',
      current: 'currentColor',

      // Backgrounds
      noir: {
        950: '#0a0a0b',  // Deepest background
        900: '#111113',  // Primary background
        850: '#18181b',  // Elevated surface (cards in dark mode)
        800: '#1f1f23',  // Card backgrounds
        700: '#2a2a30',  // Borders, dividers
        600: '#3a3a42',  // Subtle borders
      },

      // Light mode backgrounds
      stone: {
        50:  '#fafaf9',  // Page background
        100: '#f5f5f4',  // Card backgrounds
        200: '#e7e5e4',  // Borders
        300: '#d6d3d1',  // Muted borders
        400: '#a8a29e',  // Placeholder text
        500: '#78716c',  // Muted text
      },

      // Text
      text: {
        primary:   'var(--text-primary)',    // Near-white or near-black
        secondary: 'var(--text-secondary)',  // Muted
        muted:     'var(--text-muted)',      // Very muted
      },

      // Accent — warm amber/gold (like aged paper under warm light)
      accent: {
        DEFAULT: '#e0a458',
        light:   '#f0c078',
        dark:    '#c08030',
        muted:   '#e0a45833',  // 20% opacity for backgrounds
      },

      // Semantic
      seedling: '#6ec46e',   // Green — new growth
      budding:  '#e0a458',   // Amber — developing
      evergreen:'#4a90d9',   // Blue — mature

      // Status
      success: '#6ec46e',
      warning: '#e0a458',
      error:   '#d94a4a',
      info:    '#4a90d9',
    },

    extend: {
      spacing: {
        '18': '4.5rem',
        '88': '22rem',
        '128': '32rem',
      },

      maxWidth: {
        'prose': '68ch',        // Optimal reading line length
        'wide':  '72rem',       // Bento grid max width
      },

      borderRadius: {
        'card': '1rem',
        'card-sm': '0.75rem',
      },

      boxShadow: {
        'card': '0 1px 3px rgba(0,0,0,0.08), 0 1px 2px rgba(0,0,0,0.06)',
        'card-hover': '0 8px 24px rgba(0,0,0,0.12)',
        'card-dark': '0 1px 3px rgba(0,0,0,0.3), 0 1px 2px rgba(0,0,0,0.2)',
        'card-dark-hover': '0 8px 24px rgba(0,0,0,0.4)',
      },

      gridTemplateColumns: {
        'bento': 'repeat(4, 1fr)',
        'bento-md': 'repeat(2, 1fr)',
      },

      typography: ({ theme }) => ({
        DEFAULT: {
          css: {
            '--tw-prose-body': theme('colors.text.primary'),
            '--tw-prose-headings': theme('colors.text.primary'),
            '--tw-prose-links': theme('colors.accent.DEFAULT'),
            '--tw-prose-code': theme('colors.accent.light'),
            fontFamily: theme('fontFamily.body').join(', '),
            fontSize: theme('fontSize.base[0]'),
            lineHeight: theme('fontSize.base[1].lineHeight'),
            maxWidth: '68ch',
            h1: {
              fontFamily: theme('fontFamily.display').join(', '),
              fontWeight: '700',
              letterSpacing: '-0.02em',
            },
            h2: {
              fontFamily: theme('fontFamily.display').join(', '),
              fontWeight: '700',
              letterSpacing: '-0.01em',
            },
            h3: {
              fontFamily: theme('fontFamily.display').join(', '),
              fontWeight: '700',
            },
            a: {
              color: theme('colors.accent.DEFAULT'),
              textDecoration: 'underline',
              textDecorationColor: theme('colors.accent.muted'),
              textUnderlineOffset: '3px',
              transition: 'color 0.15s ease, text-decoration-color 0.15s ease',
              '&:hover': {
                color: theme('colors.accent.light'),
                textDecorationColor: theme('colors.accent.DEFAULT'),
              },
            },
            code: {
              fontFamily: theme('fontFamily.mono').join(', '),
              fontSize: '0.95em',
              fontWeight: '400',
            },
            'code::before': { content: 'none' },
            'code::after': { content: 'none' },
          },
        },
      }),
    },
  },

  plugins: [
    require('@tailwindcss/typography'),
  ],

  darkMode: 'class',
};
```

**PostCSS configuration:**

```javascript
// postcss.config.js
module.exports = {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
};
```

**Tailwind entry point:**

```css
/* assets/css/main.css */
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  :root {
    --text-primary: theme('colors.noir.950');
    --text-secondary: theme('colors.stone.500');
    --text-muted: theme('colors.stone.400');
    --card-bg: theme('colors.stone.100');
    --card-border: theme('colors.stone.200');
    --accent: theme('colors.accent.DEFAULT');
  }

  .dark {
    --text-primary: #e8e6e3;
    --text-secondary: #a0a0a8;
    --text-muted: #6a6a72;
    --card-bg: theme('colors.noir.800');
    --card-border: theme('colors.noir.700');
    --accent: theme('colors.accent.light');
  }
}
```

**Hugo Pipes integration** (`layouts/partials/head.html`):

```html
{{ $css := resources.Get "css/main.css" }}
{{ $css = $css | postCSS }}
{{ if hugo.IsProduction }}
  {{ $css = $css | minify | fingerprint }}
{{ end }}
<link rel="stylesheet" href="{{ $css.RelPermalink }}" {{ if hugo.IsProduction }}integrity="{{ $css.Data.Integrity }}"{{ end }}>
```

**`package.json`** (minimal — only build tooling):

```json
{
  "name": "ross.gg",
  "private": true,
  "scripts": {
    "dev": "hugo server --buildDrafts --environment development",
    "build": "hugo --minify --environment production"
  },
  "devDependencies": {
    "tailwindcss": "^3.4",
    "@tailwindcss/typography": "^0.5",
    "postcss": "^8.4",
    "postcss-cli": "^11",
    "autoprefixer": "^10"
  }
}
```

---

## Bento Homepage Roadmap

A phased plan for building the redesigned site incrementally. Each phase
produces a deployable site — no big-bang rewrite.

### Phase 0: Foundation (Week 1)

**Goal:** Scaffold the new Hugo structure, install Tailwind, deploy a
skeleton to Cloudflare Pages.

- [ ] Create `config/_default/` directory, split `config.toml` into
      `config.toml`, `params.toml`, `menus.toml`, `markup.toml`
- [ ] Initialize `package.json` with Tailwind + PostCSS dependencies
- [ ] Create `tailwind.config.js` and `postcss.config.js`
- [ ] Set up `assets/css/main.css` with Tailwind directives
- [ ] Create `layouts/_default/baseof.html` with font loading, dark mode toggle, head partial
- [ ] Self-host Fraunces + Inter variable fonts in `static/fonts/`
- [ ] Create minimal `layouts/index.html` with a single "coming soon" card
- [ ] Update `Makefile`: `build` target runs `hugo --minify`
- [ ] Verify Cloudflare Pages deployment works
- [ ] **Deliverable:** Blank site with correct fonts and Tailwind running

### Phase 1: Bento Homepage (Week 2)

**Goal:** Build the homepage grid with static content cards.

- [ ] Create `layouts/index.html` with the full Bento grid template
- [ ] Build each card partial in `layouts/partials/bento/`
- [ ] Start with static-only cards: hero, latest essay, /now snippet, newsletter
- [ ] Create `content/_index.md` with homepage metadata
- [ ] Create `content/now.md` with initial /now content
- [ ] Implement responsive grid (4 → 2 → 1 columns)
- [ ] Add hover micro-interactions (translateY + shadow)
- [ ] Style card labels with accent color
- [ ] **Deliverable:** Functional Bento homepage with static cards

### Phase 2: Content Migration (Week 3)

**Goal:** Migrate existing content into the new section structure.

- [ ] Create `content/essays/_index.md` and `content/notes/_index.md`
- [ ] Migrate `dns_caches.md` → `content/essays/dns-caches.md` (update frontmatter)
- [ ] Migrate Julia/PowerShell posts → `content/notes/` (add `stage` field)
- [ ] Update `content/about.md` (new layout, current role at Cloudflare)
- [ ] Build `layouts/essays/single.html` (TOC sidebar, progress bar)
- [ ] Build `layouts/notes/single.html` (stage badge, backlinks)
- [ ] Build `layouts/essays/list.html` and `layouts/notes/list.html`
- [ ] Add reading progress bar partial
- [ ] Add backlinks partial (scan content for internal links)
- [ ] **Deliverable:** All existing content accessible in new layouts

### Phase 3: Data-Driven Shelves (Week 4)

**Goal:** Add the personal content — games, books, tea.

- [ ] Create `data/games.yaml`, `data/books.yaml`, `data/tea.yaml`
- [ ] Build `layouts/partials/shelf.html` (shared card grid component)
- [ ] Create `content/interests/` section with board-games, tea, home-automation subsections
- [ ] Build interest detail pages and list views
- [ ] Wire data files into Bento homepage cards (game of month, tea of week, reading)
- [ ] Add cover images to `static/img/shelf/`
- [ ] **Deliverable:** Personal interests visible on homepage and in dedicated sections

### Phase 4: Cloudflare Workers (Week 5)

**Goal:** Add live data to the homepage.

- [ ] Create `workers/` directory with Worker source files
- [ ] Build `/api/status` Worker with KV store
- [ ] Build `/api/github` Worker with caching
- [ ] Build `/api/stackoverflow` Worker with caching
- [ ] Build `/api/subscribe` Worker with Buttondown integration
- [ ] Add client-side hydration script to homepage
- [ ] Configure Workers routes in Cloudflare dashboard
- [ ] Set up KV namespaces for caching
- [ ] Build CLI/Shortcut for updating status from phone
- [ ] **Deliverable:** Live GitHub activity, status, and newsletter on homepage

### Phase 5: Obsidian Pipeline (Week 6)

**Goal:** Publish Obsidian notes to the site.

- [ ] Install `obsidian-export` (Rust CLI)
- [ ] Configure Obsidian vault with a `publish/` folder
- [ ] Add `export-notes` target to Makefile
- [ ] Build Hugo render hook for wikilink resolution
- [ ] Test backlink detection across exported notes
- [ ] Add garden view (`/notes/`) with stage filtering (seedling/budding/evergreen)
- [ ] Document the publish workflow in `docs/`
- [ ] **Deliverable:** Obsidian → Hugo pipeline working end-to-end

### Phase 6: Polish & Dark Mode (Week 7)

**Goal:** Refine the experience.

- [ ] Implement dark mode toggle (persisted via `localStorage`, respects `prefers-color-scheme`)
- [ ] Audit typography at all breakpoints
- [ ] Add OpenGraph images for each section (use Hugo to generate or create static templates)
- [ ] Add structured data (JSON-LD) for articles
- [ ] Performance audit: target Lighthouse 95+ on all metrics
- [ ] Accessibility audit: keyboard navigation, ARIA labels, color contrast
- [ ] Update `docs/ARCHITECTURE.md` with final structure
- [ ] Clean up unused images from `static/img/unused/`
- [ ] **Deliverable:** Production-ready, polished site

### Phase 7: Ongoing Tending

After launch, the site becomes a living garden:

- Write 1-2 notes per week from Obsidian
- Update /now page every 2-4 weeks
- Rotate game/tea/book data as interests shift
- Promote mature notes to polished essays when ready
- Monitor Worker analytics for API health
