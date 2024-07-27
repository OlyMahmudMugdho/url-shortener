package types

type PostgresConfig struct {
	Username string
	Password string
	Db       string
	Sslmode  string
}

type ContextKey string
