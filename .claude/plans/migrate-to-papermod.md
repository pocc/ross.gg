# Plan: Fresh PaperMod on Cloudflare Worker

## Objective
Strip the repo down to content markdown, static assets, and the Cloudflare Worker, then rebuild as a default PaperMod Hugo site. Remove all custom layouts, Tailwind, old theme references, and dead API endpoints. Social icons (email, GitHub, LinkedIn, Stack Overflow, RSS) appear on the profile-mode homepage. Site continues to deploy via the existing Cloudflare Worker.

## Context
- **Project**: `/Users/rj/My Drive/projects/ross.gg`
- **Keep**:
  - `content/**` — all markdown + `resume.pdf`
  - `static/**` — images, favicons, manifests
  - `src/worker.ts` — stripped to static serving + `/api/subscribe` only
  - `wrangler.toml`, `tsconfig.json` — Worker build/deploy config
  - `.git/`, `.claude/`, `.gitignore`
- **Delete**:
  - `layouts/` (30+ custom template files)
  - `assets/` (Tailwind CSS)
  - `themes/` (empty `even/` submodule)
  - `config/` (split config dir — replaced by single `hugo.toml`)
  - `data/` (`books.yaml`, `games.yaml`, `tea.yaml` — only consumed by deleted Bento cards)
  - `docs/` (stale architecture docs referencing old design)
  - `.gitmodules` (only entry is `themes/even`)
  - `tailwind.config.js`, `postcss.config.js`
  - `Makefile` (replaced by `package.json` scripts)
  - `LICENSE.md` (user can re-add if wanted)
- **Rewrite**:
  - `package.json` — keep only `wrangler`, `@cloudflare/workers-types`, update scripts
  - `README.md` — brief description
  - `src/worker.ts` — remove `/api/github`, `/api/stackoverflow`, `/api/status` handlers
- **Social URLs** (from current `config/_default/params.toml`):
  - email: `mailto:rj@swit.sh`
  - github: `https://github.com/pocc`
  - linkedin: `https://www.linkedin.com/in/rossbjacobs/`
  - stackoverflow: `https://stackoverflow.com/users/1596750/ross-jacobs`
  - rss: `/index.xml`

## Constraints
- Delete everything not strictly needed for: PaperMod rendering, content, static assets, or Cloudflare Worker deployment.
- No custom layouts or CSS. Default PaperMod only.
- No Tailwind, PostCSS, or Node dev dependencies beyond `wrangler` and `@cloudflare/workers-types`.
- Keep `wrangler.toml` build command (`hugo --minify`) and assets config unchanged.
- Front-matter edits are minimal: only fix what breaks PaperMod (`toc` → `ShowToc`, remove `layout: now`). Leave unknown keys (`stage`, `related`) alone.
- Clean `_index.md` files of non-Hugo task-tracking fields.
- Zero trace of hugo-theme-even: no `.gitmodules`, no `themes/even/`, no references in config.
- Do not commit or push.

## Approach
Delete-first, build-second. Remove all non-essential files, write a single `hugo.toml` with PaperMod module import + full config, strip the Worker to essentials, add search/archives stubs, fix front matter, verify build.

## Steps

1. **Delete non-essential files and directories**
   - Remove directories: `layouts/`, `assets/`, `themes/`, `config/`, `data/`, `docs/`
   - Remove files: `.gitmodules`, `tailwind.config.js`, `postcss.config.js`, `Makefile`
   - Keep: `content/`, `static/`, `src/`, `.git/`, `.claude/`, `.gitignore`, `wrangler.toml`, `tsconfig.json`, `package.json`, `README.md`, `LICENSE.md`

2. **Strip Worker to essentials** — `src/worker.ts`
   - Remove: `handleStatusGet`, `handleStatusUpdate`, `handleGithub`, `handleStackOverflow` functions and their switch cases
   - Remove `STATUS_AUTH_TOKEN` from `Env` interface (keep `ASSETS`, `KV`, `BUTTONDOWN_API_KEY`)
   - Simplify `handleApi` router to only `/api/subscribe` + 404 default
   - Keep: `handleSubscribe`, `json` helper, `corsPreflightResponse`, static asset serving (`env.ASSETS.fetch(request)`)
   - Keep CORS since `/api/subscribe` uses form POST

3. **Rewrite `package.json`** — Worker deps only
   - Keep `name`, `private`
   - Scripts: `dev` → `hugo server -D`, `build` → `hugo --minify`, `deploy` → `npx wrangler deploy`, `preview` → `hugo --minify && npx wrangler dev`
   - devDependencies: only `wrangler` and `@cloudflare/workers-types`. Remove `tailwindcss`, `@tailwindcss/typography`, `postcss`, `autoprefixer`.

4. **Initialize Hugo Module** — repo root
   - Run `hugo mod init github.com/rossja/ross.gg` → creates `go.mod`
   - Run `hugo mod get github.com/adityatelange/hugo-PaperMod` → fetches theme, creates `go.sum`

5. **Write `hugo.toml`** — single file, full PaperMod config
   - Core: `baseURL = "https://ross.gg/"`, `title = "Ross Jacobs"`, `languageCode = "en-US"`, `enableRobotsTXT = true`, `enableEmoji = true`, `enableGitInfo = true`, `buildDrafts = false`
   - Module import: `[module] [[module.imports]] path = "github.com/adityatelange/hugo-PaperMod"`
   - Outputs: `[outputs] home = ["HTML", "RSS", "JSON"]`
   - Pagination: `[pagination] pagerSize = 10`
   - Markup: `[markup.goldmark.renderer] unsafe = true`, `[markup.highlight] style = "dracula"`, lineNos, guessSyntax, `[markup.tableOfContents]` startLevel 2, endLevel 4
   - Menu: About (10), Essays (20), Notes (30), Interests (40), Now (50), Search (60), Archives (70)
   - Params: `env = "production"`, description, author, `defaultTheme = "auto"`, `ShowReadingTime = true`, `ShowShareButtons = false`, `ShowPostNavLinks = true`, `ShowBreadCrumbs = true`, `ShowCodeCopyButtons = true`, `ShowWordCount = true`, `ShowToc = true`, `TocOpen = false`
   - Profile mode: enabled, title "Ross Jacobs", subtitle "Customer Solutions Engineer · Digital Gardener", imageUrl "/profile_pic_eff.jpg", imageWidth 160, imageHeight 160, buttons: Essays → /essays/, Notes → /notes/, About → /about/
   - Social icons: `[[params.socialIcons]]` for github, linkedin, stackoverflow, email, rss
   - Assets: favicon, favicon16x16, favicon32x32, apple_touch_icon
   - Fuse search: default opts

6. **Create `content/search.md`**
   - ```yaml
     ---
     title: "Search"
     layout: "search"
     summary: "search"
     placeholder: "Search posts"
     ---
     ```

7. **Create `content/archives.md`**
   - ```yaml
     ---
     title: "Archives"
     layout: "archives"
     url: "/archives/"
     summary: "archives"
     ---
     ```

8. **Fix content front matter** — minimal surgical edits
   - `content/essays/dns-caches.md`: `toc: true` → `ShowToc: true`
   - `content/now.md`: remove `layout: now` line (PaperMod renders as regular page)
   - `content/_index.md`: strip task-tracking fields (`created`, `due`, `wait_until`, `completed`, `recurs`, `priority`, `dependencies`, `source`), keep `title` + `description`
   - `content/essays/_index.md`: same cleanup
   - `content/notes/_index.md`: same cleanup
   - Leave `stage`, `related`, `featured`, `draft`, `tags` untouched everywhere.

9. **Update `.gitignore`**
   - Ensure has: `public/`, `resources/`, `.hugo_build.lock`, `node_modules/`, `.wrangler/`, `.dev.vars`, `.DS_Store`
   - Remove any stale entries for things that no longer exist.

10. **Update `README.md`**
    - Brief: "ross.gg — personal site. Hugo + PaperMod, deployed on Cloudflare Workers."
    - Commands: `npm run dev`, `npm run build`, `npm run deploy`
    - Theme via Hugo Modules.

11. **Build and verify**
    - Run `hugo --minify` — must exit 0 with no ERROR lines.
    - Confirm `public/index.html`, `public/search/index.html`, `public/archives/index.html`, `public/index.json` all exist.

## Instructions
- Write `hugo.toml` as a single file at repo root (not split `config/_default/`).
- PaperMod social icon names: `github`, `linkedin`, `stackoverflow`, `email`, `rss`.
- Do not create any files under `layouts/`. Default PaperMod only.
- `wrangler.toml` stays exactly as-is — its `[build] command = "hugo --minify"` and `[assets] directory = "./public"` work with the new setup unchanged.
- After stripping `src/worker.ts`, verify it type-checks: `npx tsc --noEmit`.
- `hugo mod get` requires Go installed. If it fails, surface the error.

## Evaluation Contract

### Criterion: Hugo build succeeds
- **Assertion**: `hugo --minify` exits 0, no ERROR lines, `public/index.html` exists.
- **Method**: command
- **Command**: `cd "/Users/rj/My Drive/projects/ross.gg" && hugo --minify 2>&1 | tail -5; echo "EXIT:$?"; test -f public/index.html && echo "OK"`
- **Expected**: EXIT:0, OK
- **Evidence**: stdout
- **Priority**: critical

### Criterion: PaperMod markup present
- **Assertion**: Build output contains PaperMod-specific HTML classes.
- **Method**: grep
- **Command**: `grep -rlE "entry-header|post-entry|class=\"profile\"" "/Users/rj/My Drive/projects/ross.gg/public/" --include="*.html" | head -5`
- **Expected**: at least one file
- **Evidence**: file list
- **Priority**: critical

### Criterion: Profile-mode homepage with social icons
- **Assertion**: `public/index.html` has profile markup and all 5 social links.
- **Method**: command
- **Command**: `cd "/Users/rj/My Drive/projects/ross.gg" && grep -c "profile" public/index.html && grep -oE "(github\.com/pocc|linkedin\.com/in/rossbjacobs|stackoverflow\.com/users/1596750|mailto:rj@swit\.sh|index\.xml)" public/index.html | sort -u | wc -l`
- **Expected**: profile count > 0; social link count = 5
- **Evidence**: stdout counts
- **Priority**: critical

### Criterion: Search page + JSON index
- **Assertion**: `public/search/index.html` and `public/index.json` exist, JSON non-empty.
- **Method**: command
- **Command**: `cd "/Users/rj/My Drive/projects/ross.gg" && test -f public/search/index.html && python3 -c "import json; d=json.load(open('public/index.json')); print(len(d))"`
- **Expected**: number > 0
- **Evidence**: stdout
- **Priority**: critical

### Criterion: Archives page exists
- **Assertion**: `public/archives/index.html` exists.
- **Method**: command
- **Command**: `test -f "/Users/rj/My Drive/projects/ross.gg/public/archives/index.html"`
- **Expected**: exit 0
- **Evidence**: exit code
- **Priority**: critical

### Criterion: Content preserved
- **Assertion**: All original content files still exist.
- **Method**: command
- **Command**: `cd "/Users/rj/My Drive/projects/ross.gg" && test -f content/essays/dns-caches.md && test -f content/notes/installing-julia.md && test -f content/notes/julia-binary.md && test -f content/notes/julia-versions.md && test -f content/notes/powershell-profile.md && test -f content/about.md && test -f content/now.md && test -f content/resume.pdf && test -f content/interests/_index.md`
- **Expected**: exit 0
- **Evidence**: exit code
- **Priority**: critical

### Criterion: Static assets preserved
- **Assertion**: Favicons, profile image, and content images in `static/`.
- **Method**: command
- **Command**: `cd "/Users/rj/My Drive/projects/ross.gg" && test -f static/favicon.ico && test -f static/profile_pic_eff.jpg && test -d static/img`
- **Expected**: exit 0
- **Evidence**: exit code
- **Priority**: critical

### Criterion: Dead API endpoints removed
- **Assertion**: `src/worker.ts` has no references to `/api/github`, `/api/stackoverflow`, or `/api/status`.
- **Method**: grep
- **Command**: `! grep -E "api/github|api/stackoverflow|api/status" "/Users/rj/My Drive/projects/ross.gg/src/worker.ts"`
- **Expected**: exit 0 (no matches)
- **Evidence**: empty output
- **Priority**: critical

### Criterion: Subscribe endpoint preserved
- **Assertion**: `src/worker.ts` still handles `/api/subscribe`.
- **Method**: grep
- **Command**: `grep -F "api/subscribe" "/Users/rj/My Drive/projects/ross.gg/src/worker.ts"`
- **Expected**: at least one match
- **Evidence**: matched lines
- **Priority**: critical

### Criterion: Worker TypeScript compiles
- **Assertion**: `npx tsc --noEmit` exits 0.
- **Method**: command
- **Command**: `cd "/Users/rj/My Drive/projects/ross.gg" && npx tsc --noEmit`
- **Expected**: exit 0, no error output
- **Evidence**: exit code
- **Priority**: critical

### Criterion: No old theme/layout/Tailwind artifacts
- **Assertion**: None of these exist: `layouts/`, `assets/`, `themes/`, `config/`, `data/`, `docs/`, `.gitmodules`, `tailwind.config.js`, `postcss.config.js`, `Makefile`
- **Method**: command
- **Command**: `cd "/Users/rj/My Drive/projects/ross.gg" && ! test -e layouts && ! test -e assets && ! test -e themes && ! test -e config && ! test -e data && ! test -e docs && ! test -e .gitmodules && ! test -e tailwind.config.js && ! test -e postcss.config.js && ! test -e Makefile`
- **Expected**: exit 0
- **Evidence**: exit code
- **Priority**: critical

### Criterion: wrangler.toml unchanged
- **Assertion**: `wrangler.toml` has no modifications.
- **Method**: command
- **Command**: `cd "/Users/rj/My Drive/projects/ross.gg" && git diff --stat wrangler.toml`
- **Expected**: empty output
- **Evidence**: git diff output
- **Priority**: critical

### Criterion: Hugo Module installed
- **Assertion**: `go.mod` exists and references PaperMod.
- **Method**: grep
- **Command**: `grep -F "adityatelange/hugo-PaperMod" "/Users/rj/My Drive/projects/ross.gg/go.mod"`
- **Expected**: at least one match
- **Evidence**: matched line
- **Priority**: critical

### Criterion: No Tailwind in package.json
- **Assertion**: `package.json` has no tailwindcss/postcss/autoprefixer deps.
- **Method**: grep
- **Command**: `! grep -E "tailwindcss|postcss|autoprefixer|typography" "/Users/rj/My Drive/projects/ross.gg/package.json"`
- **Expected**: exit 0 (no matches)
- **Evidence**: empty output
- **Priority**: critical

### Criterion: Menu links in nav
- **Assertion**: Homepage nav links to search and archives.
- **Method**: grep
- **Command**: `grep -oE 'href="/search/"|href="/archives/"' "/Users/rj/My Drive/projects/ross.gg/public/index.html"`
- **Expected**: two matches
- **Evidence**: matched strings
- **Priority**: nice-to-have

## Verification Sequence
1. `! test -e layouts && ! test -e assets && ! test -e themes && ! test -e config && ! test -e .gitmodules` — old artifacts deleted
2. `test -f go.mod && grep -F "adityatelange/hugo-PaperMod" go.mod` — module installed
3. `test -f content/essays/dns-caches.md && test -f content/about.md && test -f static/favicon.ico` — content/static preserved
4. `! grep -E "api/github|api/stackoverflow|api/status" src/worker.ts && grep -F "api/subscribe" src/worker.ts` — Worker cleaned up
5. `npx tsc --noEmit` — Worker compiles
6. `hugo --minify` — site builds
7. `test -f public/index.html && test -f public/search/index.html && test -f public/archives/index.html && test -f public/index.json` — all pages generated
8. `grep -E "profile" public/index.html` — profile mode active
9. `grep -oE "(github\.com/pocc|linkedin\.com|stackoverflow\.com|mailto:rj|index\.xml)" public/index.html | wc -l` — 5 social icons
