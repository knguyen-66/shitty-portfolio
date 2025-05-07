SHELL := /bin/bash

.PHONY: tailwind-watch
tailwind-watch:
	rm -f ./internal/static/css/styles.css; ./tailwindcss -i ./internal/static/css/input.css -o ./internal/static/css/styles.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./internal/static/css/input.css -o ./internal/static/css/styles.css
	./tailwindcss -i ./internal/static/css/input.css -o ./internal/static/css/styles.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: prepare
prepare:
	if [[ ! -e .env ]]; then\
		cp .env.example .env;\
	fi
	./.scripts/dev-install-packages.sh
	go mod tidy
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/air-verse/air@latest
	./.scripts/goose.sh up
	sqlc generate

.PHONY: dev
dev:
	make tailwind-build && make templ-generate
	trap 'kill %1; kill %2' SIGINT
	make templ-watch & make tailwind-watch & air

.PHONY: build-prod
build-prod:
	make prepare
	make tailwind-build
	make templ-generate
	go build -o .prod/main .

.PHONY: prod
prod:
	if [[ ! -e .prod/main ]]; then\
		make build-prod;\
	fi
	.prod/main
