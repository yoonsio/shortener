package shortener

// ShortenRequest is json request structure for /shorten
type ShortenRequest struct {
	URL string
}

// OriginalReq is json request structure for /original
type OriginalReq struct {
	Short string
}

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
