package db

// Config is used to connect to database
type Config struct {
	Addr     string
	Port     int
	User     string
	Password string
	DBName   string
}
