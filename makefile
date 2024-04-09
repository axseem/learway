diff:
	atlas migrate diff $(NAME) \
		--dir "file://storage/sqlite/migrations" \
		--to "file://storage/sqlite/schema/" \
		--dev-url "sqlite://dev?mode=memory"

migrate:
	atlas migrate apply \
		--dir "file://storage/sqlite/migrations" \
		--url "sqlite://dev.db"

gen: 
	sqlc generate