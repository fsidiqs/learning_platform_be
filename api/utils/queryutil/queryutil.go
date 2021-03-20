package queryutil

const (
	INSERT_USER = "INSERT INTO users" +
		"(name, email, password)" +
		"VALUES(?,?,?)"
)
