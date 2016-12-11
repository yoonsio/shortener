package shortener

// Config is configuration structure
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig contains server-related configs
type ServerConfig struct {
	Port int
}

// DatabaseConfig contains database-related configs
type DatabaseConfig struct {
	URI    string
	DBName string
}
