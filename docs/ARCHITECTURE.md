# Architecture

## Overview

ross.gg is a personal blog/portfolio site for Ross Jacobs, built with Hugo (a static site generator). It is hosted at https://ross.gg.

## Tech Stack

- **Static site generator:** [Hugo](https://gohugo.io/)
- **Theme:** [hugo-theme-even](https://github.com/olOwOlo/hugo-theme-even) (git submodule)
- **Syntax highlighting:** Chroma (server-side, via Hugo)
- **Analytics:** Google Analytics (UA-143876903-1)
- **License:** Apache 2.0

## Repository

- **Source:** https://github.com/pocc/ross.gg (GitHub username: `pocc`)
- **Branch:** `master`
- **Commits:** ~83 total

## Build & Deploy

```bash
# Build the site (outputs to public/)
make build          # runs: hugo --theme=even -D

# Deploy (copies to a web server document root — legacy, not currently used)
make deploy         # runs: sudo cp -r public/* /var/www/html
```

The `public/` directory is gitignored. Current hosting/deployment mechanism is unclear from the repo alone (no CI config present; the Makefile `deploy` target suggests a VPS with nginx/apache, but the config references tools.ross.gg as a separate subdomain).

## Project Structure

```
ross.gg/
├── config.toml              # Hugo configuration (site title, menus, params)
├── Makefile                 # Build and deploy commands
├── content/
│   ├── about.md             # About page
│   ├── projects.md          # Projects showcase
│   ├── manifest.json        # Web app manifest (Android icons)
│   ├── resume.pdf           # Downloadable resume
│   └── post/
│       ├── 2019/
│       │   ├── julia_exe.md              # Making a Julia Binary
│       │   ├── julia_install.md          # Installing Julia
│       │   ├── julia_versions.md         # Julia Versions comparison
│       │   └── setup_powershell_profile.md # PowerShell profile setup
│       └── 2021/
│           └── dns_caches.md             # DNS Caches (browser & OS)
├── static/
│   ├── favicon.ico, favicon-*.png, etc.  # Favicons
│   ├── browserconfig.xml                 # Windows tile config
│   ├── site.webmanifest                  # PWA manifest
│   └── img/
│       ├── profile_pic_eff.jpg           # Profile photo
│       ├── dns_caches/                   # DNS cache article images
│       ├── julia/                        # Julia article images
│       ├── powershell/                   # PowerShell article images
│       ├── projects/                     # Project screenshots
│       ├── wireshark/                    # Wireshark contribution images
│       └── unused/                       # Unused images (kept in repo)
├── themes/
│   └── even/                # Git submodule (hugo-theme-even) — NOT initialized
└── docs/                    # Documentation (this folder)
```

## Navigation (config.toml)

The site has four main menu items:
1. **About** — `/about`
2. **Tools** — External link to `https://tools.ross.gg/`
3. **Archives** — `/post/` (paginated blog posts)
4. **Tags** — `/tags/`

## Hugo Configuration Notes

- `baseURL`: https://ross.gg/
- `enableGitInfo`: true (shows last modified from git)
- `paginate`: 5 posts per page
- `buildDrafts`: false
- Syntax highlighting via Chroma (not client-side highlight.js)
- Table of contents enabled, starting at heading level 1
- Goldmark renderer with `unsafe: true` (allows raw HTML in markdown)

## Social Links

Configured in `config.toml` under `[params.social]`:
- Email: rj@swit.sh
- Stack Overflow: /users/1596750/ross-jacobs
- LinkedIn: /in/rossbjacobs/
- GitHub: github.com/pocc

## Theme Submodule

The theme is `hugo-theme-even` referenced as a git submodule but not currently initialized in the working copy. To initialize:

```bash
git submodule update --init --recursive
```
