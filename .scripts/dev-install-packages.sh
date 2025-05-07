#!/bin/bash

# run from root directory

curl -L "https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-linux-x64" -o tailwindcss --create-dirs && chmod +x ./tailwindcss
curl "https://unpkg.com/htmx.org@2.0.2/dist/htmx.min.js" -o ./internal/static/script/htmx-2_0_2.min.js --create-dirs
curl "https://cdn.jsdelivr.net/npm/alpinejs@3.14.1/dist/cdn.min.js" -o ./internal/static/script/alpinejs-3_14_1.min.js --create-dirs

curl "https://unpkg.com/@highlightjs/cdn-assets@11.9.0/styles/github-dark-dimmed.min.css" -o ./internal/static/css/highlightjs-dark-11_9_0.min.css --create-dirs
curl "https://unpkg.com/@highlightjs/cdn-assets@11.9.0/styles/github.min.css" -o ./internal/static/css/highlightjs-light-11_9_0.min.css --create-dirs
curl "https://unpkg.com/@highlightjs/cdn-assets@11.9.0/highlight.min.js" -o ./internal/static/script/highlightjs-11_9_0.min.js --create-dirs
curl "https://unpkg.com/feather-icons@4.29.2/dist/feather-sprite.svg" -o ./internal/static/imgs/svg/feather-sprite.svg --create-dirs

# curl "https://unpkg.com/htmx.org@2.0.2/dist/htmx.js" > ./internal/static/script/htmx-2_0_2.js
# curl -o ./internal/static/css/daisy.full.min.css https://cdn.jsdelivr.net/npm/daisyui@4.12.10/dist/full.min.css
