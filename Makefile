run:
	export DB_USER="odin" && \
	export DB_PASS="devops" && \
	npx tailwindcss -i ./views/css/input.css -o ./views/css/output.css && \
	templ generate && \
	go run cmd/main.go
