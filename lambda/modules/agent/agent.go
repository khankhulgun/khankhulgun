package agent

import (
	"github.com/khankhulgun/khankhulgun/config"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent/agentMW"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent/handlers"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent/utils"
	vpUtils "github.com/khankhulgun/khankhulgun/tools"
	"github.com/labstack/echo/v4"
	"html/template"
)

func Set(e *echo.Echo) {

	if config.Config.App.Migrate == "true"{
		utils.AutoMigrateSeed()
	}
	templates := vpUtils.GetTemplates(e)

	/* REGISTER VIEWS */
	AbsolutePath := config.AbsolutePath()
	templates["login.html"] = template.Must(template.ParseFiles(AbsolutePath+"lambda/modules/agent/templates/login.html"))
	templates["forgot.html"] = template.Must(template.ParseFiles(AbsolutePath+"lambda/modules/agent/templates/email/forgot.html"))

	/* ROUTES */
	a :=e.Group("/auth")
	a.GET("/", handlers.LoginPage)
	a.GET("/login", handlers.LoginPage)
	a.POST("/login", handlers.Login)
	a.POST("/logout", handlers.Logout)

	/*PASSWORD RESET*/
	a.POST("/send-forgot-mail", handlers.SendForgotMail)
	a.POST("/password-reset", handlers.PasswordReset)

	u :=e.Group("/agent")
	u.GET("/users", handlers.GetUsers, agentMW.IsLoggedInCookie, agentMW.IsAdmin)
	u.GET("/search/:q", handlers.SearchUsers, agentMW.IsLoggedInCookie, agentMW.IsAdmin)
	u.GET("/users/deleted", handlers.GetDeletedUsers, agentMW.IsLoggedInCookie, agentMW.IsAdmin)
	u.GET("/delete/:id", handlers.DeleteUser, agentMW.IsLoggedInCookie, agentMW.IsAdmin)
	u.GET("/roles", handlers.GetRoles,  agentMW.IsLoggedInCookie, agentMW.IsAdmin)


}

