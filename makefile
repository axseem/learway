diff:
	atlas migrate diff $(NAME) \
		--dir "file://internal/database/migrations" \
		--to "file://internal/database/schema/" \
		--dev-url "sqlite://dev?mode=memory"

migrate:
	atlas migrate apply \
		--dir "file://internal/database/migrations" \
		--url "sqlite://dev.db"

gen: 
	sqlc generate