package handlers

import (
	"fmt"
	"net/http"
	"github.com/khankhulgun/khankhulgun/lambda/modules/agent/models"
	"github.com/khankhulgun/khankhulgun/lambda/config"
	"reflect"
	"strings"
	"time"
	"math/rand"
	"github.com/khankhulgun/khankhulgun/lambda/plugins/mailer"
	"github.com/labstack/echo/v4"
	"github.com/khankhulgun/khankhulgun/DB"
	agentUtils "github.com/khankhulgun/khankhulgun/lambda/modules/agent/utils"
)

type mailPost struct {
	Email string    `json:"email" `
	Lang string    `json:"lang"`
}
type passwordResetPost struct {
	Email string    `json:"email" `
	Code string    `json:"code"`
	Lang string    `json:"lang"`
	Password string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
}
func SendForgotMail(c echo.Context) error {


	data := new(mailPost)
	if err := c.Bind(data); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"error": "Post data error ",
			"msg": "Post data error ",
		})
	}

	if data.Lang == ""{
		data.Lang = "mn"
	}

	StaticWords := reflect.ValueOf(config.Config.StaticWords[data.Lang]).Interface().(map[string]interface{})
	if data.Email == ""{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"error": StaticWords["emailRequired"],
			"msg": StaticWords["emailRequired"],
		})
	}

	foundUser := agentUtils.AuthUserObjectByEmail(data.Email)

	if len(foundUser) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": StaticWords["userNotFound"],
			"msg": StaticWords["userNotFound"],
			"status":false,
		})
	}

	pReset := models.PasswordReset{}
	DB.DB.Where("email = ?", data.Email).Delete(pReset)


	permittedChars := strings.Join(shuffle([]string{"0","1","2","3","4","5","6","7","8","9"}),"")
	//permittedChars := strings.Join(shuffle([]string{"0","1","2","3","4","5","6","7","8","9","a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}),"")
	tokenPre := string([]rune(permittedChars)[0:6])
	token, _ := agentUtils.Hash(tokenPre)

	pReset.Email = data.Email
	pReset.Token = token
	pReset.CreatedAt = time.Now()

	DB.DB.NewRecord(pReset)
	DB.DB.Create(pReset)



	mail := mailer.NewRequest([]string{data.Email}, StaticWords["passwordResetCode"].(string))
	mailSent := mail.Send("github.com/khankhulgun/khankhulgun/lambda/modules/agent/templates/email/forgot.html", map[string]string{
		"keyword":           tokenPre,
		"passwordReset":     StaticWords["passwordReset"].(string),
		"passwordResetCode": StaticWords["passwordResetCode"].(string),
		"title":             config.Config.Title,
		"noReply":           StaticWords["noReply"].(string),
		"copyright":         config.Config.Copyright,
	})

	if mailSent {

		delete(foundUser, "password")

		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": StaticWords["passwordResetCodeSent"],
			"status": true,
			"data": foundUser,
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"error": "Mail error please contact system administrator",
			"msg": "Mail error please contact system administrator",
		})
	}


}
func PasswordReset(c echo.Context) error {

	data := new(passwordResetPost)
	if err := c.Bind(data); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"error": "Post data error ",
			"msg": "Post data error ",
		})
	}

	if data.Lang == ""{
		data.Lang = "mn"
	}


	StaticWords := reflect.ValueOf(config.Config.StaticWords[data.Lang]).Interface().(map[string]interface{})
	if data.Email == "" || data.Code == "" || data.Password == "" || data.PasswordConfirm == ""{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"error": StaticWords["passwordResetCodeRequired"],
			"msg": StaticWords["passwordResetCodeRequired"],
		})
	}

	foundUser := models.User{}
	pReset := models.PasswordReset{}
	PasswordResetTimeOut := config.Config.PasswordResetTimeOut

	errU := DB.DB.Where("email = ?",data.Email).First(&foundUser).Error
	errR := DB.DB.Where("email = ?",data.Email).First(&pReset).Error

	if errU != nil || foundUser.Login == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": StaticWords["userNotFound"],
			"msg": StaticWords["userNotFound"],
			"status": false,
		})
	}
	if errR != nil || pReset.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": StaticWords["userNotFound"],
			"msg": StaticWords["userNotFound"],
			"status": false,
		})
	}

	now := time.Now()
	diff := now.Sub(pReset.CreatedAt)

	mins := int(diff.Minutes())

	if PasswordResetTimeOut >= mins && pReset.Wrong <= 2{

		if agentUtils.IsSame(data.Code, pReset.Token) {

			if data.Password == data.PasswordConfirm {

				newPassword, _ := agentUtils.Hash(data.Password)

				foundUser.Password = newPassword
				err := DB.DB.Save(foundUser).Error
				if err != nil {

					return c.JSON(http.StatusOK, map[string]interface{}{
						"status": false,
					})
				} else {
					DB.DB.Where("email = ?", data.Email).Delete(pReset)
					return c.JSON(http.StatusOK, map[string]interface{}{
						"msg": StaticWords["passwordResetSuccess"],
						"status": true,
					})
				}

			} else {

				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"error": StaticWords["passwordConfirmError"],
					"msg": StaticWords["passwordConfirmError"],
					"status": false,
				})
			}

		} else {
			pReset.Wrong = pReset.Wrong+1
			DB.DB.Save(pReset)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": StaticWords["passwordResetCodeIncorrect"],
				"msg": StaticWords["passwordResetCodeIncorrect"],
				"status": false,
			})
		}
	} else {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": StaticWords["passwordResetCodeTimeout"],
			"msg": StaticWords["passwordResetCodeTimeout"],
			"status": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": false,
	})
}
func shuffle(src []string) []string {
	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return final
}