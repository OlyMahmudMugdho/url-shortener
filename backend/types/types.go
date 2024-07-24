package types

type PostgresConfig struct {
	Username string
	Password string
	Db       string
	Sslmode  string
}

type User struct {
	Id        string `json:id`
	Name      string `json: "name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
