package utils

import (
	"time"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/config"
	agentModels "github.com/khankhulgun/khankhulgun/lambda/modules/agent/models"

)

func AutoMigrateSeed() {
	db := DB.DB

	db.AutoMigrate(
		&agentModels.Role{},
		&agentModels.User{},
		&agentModels.PasswordReset{},
	)

	if config.Config.App.Seed == "true" {
		var roles []agentModels.Role
		db.Find(&roles)

		if len(roles) <= 0 {
			seedData()
		}
	}
}
func seedData() {
	/*SUPER ADMIN ROLE*/
	role := agentModels.Role{
		Name:"super-admin",
		DisplayName:"Систем админ",
	}

	db := DB.DB
	db.Create(&role)

	/*SUPER ADMIN USER*/
	password, _ := Hash(config.Config.SuperAdmin.Password)
	user := agentModels.User{
		Role:1,
		Login:config.Config.SuperAdmin.Login,
		Email:config.Config.SuperAdmin.Email,
		Password:password,
		Status:"2",
		Birthday:time.Now(),
		Gender:"m",
	}


	db.Create(&user)



}
