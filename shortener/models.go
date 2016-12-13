package shortener

import "time"

// ShortenRequest is json request structure for /shorten
type ShortenRequest struct {
	URL string
}

// ShortenResponse is json response structure for /shorten
type ShortenResponse struct {
	Short string
}

// OriginalRequest is json request structure for /original
type OriginalRequest struct {
	Short string
}

// OriginalResponse is json response structure for /original
type OriginalResponse struct {
	Original string
}

// URIMap is database object that represents
// mapping between original URI and shortened URI
type URIMap struct {
	Original  string
	Short     string
	Timestamp time.Time
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
