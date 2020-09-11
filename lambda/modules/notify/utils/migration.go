package utils

import (
	"github.com/khankhulgun/khankhulgun/DB"
	notifyModels "github.com/khankhulgun/khankhulgun/lambda/modules/notify/models"
)

func AutoMigrateSeed() {
	db := DB.DB

	db.AutoMigrate(
		&notifyModels.Notification{},
		&notifyModels.NotificationStatus{},
		&notifyModels.NotificationTarget{},
	)

}

