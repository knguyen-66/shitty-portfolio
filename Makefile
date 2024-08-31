.PHONY: tailwind-watch
tailwind-watch:
	rm -f ./internal/static/css/styles.css
	./tailwindcss -i ./internal/static/css/input.css -o ./internal/static/css/styles.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./internal/static/css/input.css -o ./internal/static/css/styles.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: prepare
prepare:
	./scripts/dev-install-packages.sh
	go mod tidy

.PHONY: dev
dev:
	make templ-generate
	air