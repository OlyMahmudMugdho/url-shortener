package types

type PostgresConfig struct {
	Username string
	Password string
	Db       string
	Sslmode  string
}

type ContextKey string

type Error struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
