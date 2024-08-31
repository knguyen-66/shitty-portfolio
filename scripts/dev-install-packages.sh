#!/bin/bash

curl -L "https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.10/tailwindcss-linux-x64" > tailwindcss && chmod +x ./tailwindcss
curl "https://unpkg.com/htmx.org@2.0.2/dist/htmx.js" > ./internal/static/script/htmx.js
curl "https://unpkg.com/htmx.org@2.0.2" > ./internal/static/script/htmx.min.js  # for production later 
 
# curl -o ./internal/static/css/daisy.full.min.css https://cdn.jsdelivr.net/npm/daisyui@4.12.10/dist/full.min.css