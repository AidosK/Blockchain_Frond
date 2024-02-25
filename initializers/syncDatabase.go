package initializers

import (
	"github.com/AidosK/golang-blockchain/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
