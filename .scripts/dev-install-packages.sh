#!/bin/bash

# RUN FROM ROOT DIR

curl -L "https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.10/tailwindcss-linux-x64" > tailwindcss && chmod +x ./tailwindcss
curl "https://unpkg.com/htmx.org@2.0.2/dist/htmx.min.js" > ./internal/static/script/htmx-2_0_2.min.js
curl "https://cdn.jsdelivr.net/npm/alpinejs@3.14.1/dist/cdn.min.js" > ./internal/static/script/alpinejs-3_14_1.min.js
 
# curl "https://unpkg.com/htmx.org@2.0.2/dist/htmx.js" > ./internal/static/script/htmx-2_0_2.js
# curl -o ./internal/static/css/daisy.full.min.css https://cdn.jsdelivr.net/npm/daisyui@4.12.10/dist/full.min.css