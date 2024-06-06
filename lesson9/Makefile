run:build
	@./bin/tienloc
templ:
	@templ generate -watch -proxy:=http://localhost:8080

tailwind:
	@tailwindcss -i internals/static/css/input.css -o internals/static/css/output.css
install:
	@go install github.com/a-h/templ/cmd/templ@lastest
	@go get ./...
	@go mod tidy
build:
	tailwindcss -i internals/static/css/input.css -o internals/static/css/output.css --watch
	@templ generate view
