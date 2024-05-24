package initializers

import "github.com/quilo-bikcodes/Go-JWT/models"

func SyncData() {
	DB.AutoMigrate(&models.User{})
}