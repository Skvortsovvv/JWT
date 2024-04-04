migrate-up:
	goose -dir ./migrations postgres "postgres://postgres@localhost:5432/mydb?sslmode=disable"  up