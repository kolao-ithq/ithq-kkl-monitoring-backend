package application

import (
	"codelab/go-gin-template/internal/config"
	"codelab/go-gin-template/internal/core/modules/users/ports/output"
	repositories "codelab/go-gin-template/internal/core/modules/users/repository"
	"codelab/go-gin-template/internal/database/interfaces"
	"codelab/go-gin-template/internal/database/mysql"
	"codelab/go-gin-template/internal/database/redis"
	"log"
)

// Repositories holds all data storage interfaces
type AppRepositories struct {
	// System repositories
	mysql interfaces.MySQLRepository
	redis interfaces.RedisRepository
	// minio    *minio.MinIORepository
	// mailer   *mailer.MailerRepository
	// postgres *postgres.PostgresRepository
	// Domain repositories
	users output.UserRepository
}

// setupRepositories initializes all data storage connections
func SetupRepositories(cfg *config.Config) *AppRepositories {
	// MySQL setup
	mysqlRepo, err := mysql.NewMySQLRepository(cfg.Database)
	if err != nil {
		log.Fatalf("❌ Failed to initialize MySQL repository: %v", err)
	}

	// PostgreSQL setup
	// postgresRepo, err := postgres.NewPostgresRepository(cfg.Postgres)
	// if err != nil {
	// 	log.Fatalf("❌ Failed to initialize PostgreSQL repository: %v", err)
	// }

	// Redis setup
	redisRepo, err := redis.NewRedisRepository(cfg.Redis)
	if err != nil {
		log.Fatalf("❌ Failed to initialize Redis repository: %v", err)
	}

	// MinIO setup
	// minioRepo, err := minio.NewMinIORepository(cfg.Minio)
	// if err != nil {
	// 	log.Fatalf("❌ Failed to initialize MinIO repository: %v", err)
	// }
	// Mailer setup
	// mailerRepo := mailer.NewMailerRepository(cfg.Mailer)

	// Domain repositories
	userRepo := repositories.NewUserRepository(mysqlRepo.GetDB())

	return &AppRepositories{
		// mysql:  mysqlRepo,
		redis: redisRepo,
		// minio:    minioRepo,
		users: userRepo,
		// mailer:   mailerRepo,
		// postgres: postgresRepo,
	}

}

// closeRepositories gracefully closes all connections
func CloseRepositories(repos *AppRepositories) {
	if repos.mysql != nil {
		repos.mysql.Close()
	}
	if repos.redis != nil {
		repos.redis.Close()
	}
	// if repos.postgres != nil {
	// 	repos.postgres.Close()
	// }
}
