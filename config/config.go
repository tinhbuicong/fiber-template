package config

import (
	"os"
	"strconv"
	"time"
)

// Config represents application configuration
type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	JWT       JWTConfig
	Redis     RedisConfig
	Mail      MailConfig
	Log       LogConfig
	Metrics   MetricsConfig
	RateLimit RateLimitConfig
}

// MetricsConfig for Prometheus /metrics endpoint (App → Prometheus → Alertmanager)
type MetricsConfig struct {
	Port string // METRICS_PORT, default 9091
}

// RateLimitConfig for limiter middleware (per-zone)
type RateLimitConfig struct {
	Enabled          bool // enable/disable all rate limiters
	PublicMax        int  // Public (/, /health) – default 100 req/min
	PublicExpiration time.Duration
	AuthMax          int // Auth (/api/v1/auth/*) – default 5 req/min
	AuthExpiration   time.Duration
	LogsMax          int // Logs (/api/v1/logs) – default 2 req/min
	LogsExpiration   time.Duration
	// Legacy single-zone (optional)
	Max        int
	Expiration time.Duration
}

// LogConfig represents file logging configuration (Zap + Lumberjack)
type LogConfig struct {
	Dir        string // directory for log files
	Filename   string // base filename, e.g. app.log
	MaxDays    int    // keep N days of history
	MaxSizeMB  int    // max size per file in MB before rotate
	Compress   bool   // compress rotated files
	MaxBackups int    // max number of old files to retain
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Port         string
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// JWTConfig represents JWT configuration
type JWTConfig struct {
	SecretKey          string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
	Issuer             string
}

// RedisConfig represents Redis configuration
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// MailConfig represents mail configuration
type MailConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
	FromName string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "3000"),
			Host:         getEnv("SERVER_HOST", "0.0.0.0"),
			ReadTimeout:  getDuration("SERVER_READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getDuration("SERVER_WRITE_TIMEOUT", 10*time.Second),
			IdleTimeout:  getDuration("SERVER_IDLE_TIMEOUT", 120*time.Second),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "auth_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey:          getEnv("JWT_SECRET_KEY", "your-super-secret-key-change-in-production"),
			AccessTokenExpiry:  getDuration("JWT_ACCESS_EXPIRY", 15*time.Minute),
			RefreshTokenExpiry: getDuration("JWT_REFRESH_EXPIRY", 7*24*time.Hour),
			Issuer:             getEnv("JWT_ISSUER", "fiber-template"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvInt("REDIS_DB", 0),
		},
		Mail: MailConfig{
			Host:     getEnv("MAIL_HOST", "smtp.mailtrap.io"),
			Port:     getEnvInt("MAIL_PORT", 587),
			User:     getEnv("MAIL_USER", ""),
			Password: getEnv("MAIL_PASSWORD", ""),
			From:     getEnv("MAIL_FROM", "noreply@example.com"),
			FromName: getEnv("MAIL_FROM_NAME", "Fiber Template"),
		},
		Log: LogConfig{
			Dir:        getEnv("LOG_DIR", "./storage/logs"),
			Filename:   getEnv("LOG_FILENAME", "app.log"),
			MaxDays:    getEnvInt("LOG_MAX_DAYS", 7),
			MaxSizeMB:  getEnvInt("LOG_MAX_SIZE_MB", 100),
			Compress:   getEnv("LOG_COMPRESS", "true") == "true",
			MaxBackups: getEnvInt("LOG_MAX_BACKUPS", 7),
		},
		Metrics: MetricsConfig{
			Port: getEnv("METRICS_PORT", "9091"),
		},
		RateLimit: RateLimitConfig{
			Enabled:          getEnv("RATELIMIT_ENABLED", "true") == "true",
			PublicMax:        getEnvInt("RATELIMIT_PUBLIC_MAX", 100),
			PublicExpiration: getDuration("RATELIMIT_PUBLIC_EXPIRATION", time.Minute),
			AuthMax:          getEnvInt("RATELIMIT_AUTH_MAX", 5),
			AuthExpiration:   getDuration("RATELIMIT_AUTH_EXPIRATION", time.Minute),
			LogsMax:          getEnvInt("RATELIMIT_LOGS_MAX", 2),
			LogsExpiration:   getDuration("RATELIMIT_LOGS_EXPIRATION", time.Minute),
			Max:              getEnvInt("RATELIMIT_MAX", 60),
			Expiration:       getDuration("RATELIMIT_EXPIRATION", time.Minute),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
