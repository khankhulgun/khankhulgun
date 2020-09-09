package config

import (
	"encoding/json"
	"fmt"
	"github.com/khankhulgun/khankhulgun/config"
	"os"
	"sync"
)

var onceConfig sync.Once
type LambdaConfig struct {
	Theme       string `json:"theme"`
	Domain      string `json:"domain"`
	Title       string `json:"title"`
	SubTitle    string `json:"subTitle"`
	Copyright   string `json:"copyright"`
	Favicon     string `json:"favicon"`
	Bg          string `json:"bg"`
	Logo        string `json:"logo"`
	LogoText    string `json:"logoText"`
	SuperURL    string `json:"super_url"`
	AppURL      string `json:"app_url"`
	HasLanguage bool   `json:"has_language"`
	KrudPublic bool   `json:"krud_public"`
	ControlPanel struct {
		LogoLight   string `json:"logoLight"`
		LogoDark    string `json:"logoDark"`
		BrandBtnURL string `json:"brandBtnUrl"`
	} `json:"controlPanel"`
	Languages   []struct {
		Label string `json:"label"`
		Code  string `json:"code"`
	} `json:"languages"`
	DefaultLanguage string `json:"default_language"`
	RoleRedirects   []struct {
		RoleID int64    `json:"role_id"`
		URL    string `json:"url"`
	} `json:"role-redirects"`
	UserDataFields         []string `json:"user_data_fields"`
	DataFormCustomElements []struct {
		Element string `json:"element"`
	} `json:"data_form_custom_elements"`
	PasswordResetTimeOut int `json:"password_reset_time_out"`
	StaticWords map[string]interface{} `json:"static_words"`
	Notify struct {
		FirebaseConfig struct {
			APIKey            string `json:"apiKey"`
			PublicKey            string `json:"publicKey"`
			AuthDomain        string `json:"authDomain"`
			DatabaseURL       string `json:"databaseURL"`
			ProjectID         string `json:"projectId"`
			StorageBucket     string `json:"storageBucket"`
			MessagingSenderID string `json:"messagingSenderId"`
			AppID             string `json:"appId"`
			MeasurementID     string `json:"measurementId"`
		} `json:"firebaseConfig"`
		ServerKey string `json:"serverKey"`
		Sound     string `json:"sound"`
		Icon      string `json:"icon"`

	} `json:"notify"`
}
var Config LambdaConfig
func init()  {
	onceConfig.Do(func() {


		configFiel := config.Config.LambdaConfig.ConfigFile
		configFile, err := os.Open(configFiel)
		defer configFile.Close()
		if err != nil{
			fmt.Println("CONFIG FILE NOT FOUND")
		}
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&Config)
	})

}
