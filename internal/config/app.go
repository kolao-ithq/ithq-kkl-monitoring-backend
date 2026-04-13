package config

import (
	"codelab/go-gin-template/internal/database/mysql"
	"codelab/go-gin-template/internal/database/redis"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
)

// "mega-erp-backend-svc/internal/adapters/secondary/postgres"
// "mega-erp-backend-svc/internal/adapters/secondary/redis"

// Config holds the application configuration
type Config struct {
	Server   ServerConfig
	Database mysql.Config
	Postgres postgres.Config
	Redis    redis.Config
	JWT      JWTConfig
	// Minio    minio.Config
	// Mailer   mailer.Config
	APIKey APIKeyConfig
	File   FileConfig
}

type FileConfig struct {
	Host   string
	Port   string
	Key    string
	Bucket string
	PathIp string
}

// ServerConfig holds the HTTP server configuration
type ServerConfig struct {
	Port         string
	Env          string
	ReadTimeout  int
	WriteTimeout int
}

type APIKeyConfig struct {
	Value string
}

// JWTConfig holds the JWT configuration
type JWTConfig struct {
	Secret        string
	RefreshSecret string
}

// var MinioGlobal minio.Config
var JWTGlobal JWTConfig
var APIKeyGlobal APIKeyConfig
var FileGlobal FileConfig
var ServerGlobal ServerConfig

// LoadConfig loads the application configuration from environment variables and config files
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// Set default values
	setDefaults()

	// Read from config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("❌ failed to read config file: %w", err)
		}
		// Config file not found, will use defaults and environment variables
	}

	// Read from environment variables
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("❌ failed to unmarshal config: %w", err)
	}
	// MinioGlobal = config.Minio
	JWTGlobal = config.JWT
	APIKeyGlobal = config.APIKey
	FileGlobal = config.File
	ServerGlobal = config.Server
	return &config, nil
}

// setDefaults sets the default values for the configuration
func setDefaults() {
	// Server defaults
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.env", "dev")
	viper.SetDefault("server.readTimeout", 10)
	viper.SetDefault("server.writeTimeout", 10)

	// Database defaults
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.user", "root")
	viper.SetDefault("database.password", "rootpassword")
	viper.SetDefault("database.dbname", "go-app")
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("database.timezone", "UTC")

	// Postgres defaults
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.user", "root")
	viper.SetDefault("postgres.password", "root")
	viper.SetDefault("postgres.dbname", "local")
	viper.SetDefault("postgres.sslmode", "disable")
	viper.SetDefault("postgres.timezone", "UTC")

	// Redis defaults
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)

	// JWT defaults
	viper.SetDefault("jwt.secret", "your-secret-key")
	viper.SetDefault("jwt.refreshSecret", "your-refresh-secret-key")

	// Minio defaults
	viper.SetDefault("minio.endpoint", "localhost:2087")
	viper.SetDefault("minio.accessKey", "minio")
	viper.SetDefault("minio.secretKey", "minio")
	viper.SetDefault("minio.bucketName", "sabaix")
	viper.SetDefault("minio.baseUrl", "http://localhost:2087/sabaix")

	// Mailer defaults
	viper.SetDefault("mailer.host", "smtp.example.com")
	viper.SetDefault("mailer.username", "username")
	viper.SetDefault("mailer.password", "password")
	viper.SetDefault("mailer.port", 000)
	viper.SetDefault("mailer.from", "from")
}
