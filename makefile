diff:
	atlas migrate diff $(NAME) \
		--dir "file://database/migrations" \
		--to "file://database/schema/schema.sql" \
		--dev-url "sqlite://dev?mode=memory"

migrate:
	atlas migrate apply \
		--dir "file://database/migrations" \
		--url "sqlite://dev.db"

gen: 
	sqlc generate && templ generate