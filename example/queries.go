package queries

const (
	
	// UsersCreate is imported from example/users_create.sql
	UsersCreate = `INSERT INTO users (name) VALUES ($1)`
	
	// UsersSelect is imported from example/users_select.sql
	UsersSelect = `SELECT * FROM users`
	
)