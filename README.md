# ross.gg

Personal site built with [Hugo](https://gohugo.io/) + [PaperMod](https://github.com/adityatelange/hugo-PaperMod), deployed on [Cloudflare Workers](https://workers.cloudflare.com/).

Theme is installed via Hugo Modules.

## Development

```sh
npm run dev       # hugo server -D
npm run build     # hugo --minify
npm run deploy    # npx wrangler deploy
npm run preview   # build + wrangler dev
```
