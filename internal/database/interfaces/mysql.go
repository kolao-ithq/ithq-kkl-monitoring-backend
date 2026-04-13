package interfaces

import "gorm.io/gorm"

// MySql Repository interface defines the methods for the MySQL repository
type MySQLRepository interface {
	GetDB() *gorm.DB
	Close() error
}
