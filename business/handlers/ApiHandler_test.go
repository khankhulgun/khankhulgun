package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"projects/business/models"
	"vp/DB"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
)


func TestCompanyBySections(t *testing.T)  {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/CompanyBySections")
	
	jsonData := []models.AanSalbarCount{}
	DB.DB.Table("ds_company_sub").Select("count(*) as too, ds_company_sub.tp_section, salbar_id").Group("salbar_id").Find(&jsonData)

	jsonString, _ := json.Marshal(jsonData)

	// Assertions
	if assert.NoError(t, CompanyBySections(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestTaniltsuulga(t *testing.T) {

	//Set up code
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/Info")

	//API data caller
	jsonData := models.Taniltsuulga{}
	DB.DB.Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)


	// Testing
	if assert.NoError(t, Taniltsuulga(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestCountByCompany(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/CountByCompany")

	var tcom int
	DB.DB.Table("company_register").Count(&tcom)

	if assert.NoError(t, CountByCompany(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)


		body, _ := ioutil.ReadAll(rec.Body)
		preParse := strings.Replace(string(body), "\n", "", 1)
		parsedInt, _ := strconv.Atoi(preParse)
		assert.Equal(t, tcom, parsedInt)

	}
}
func TestCountByProduct (t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/CountProduct")

	var tproduct int
	DB.DB.Table("buteegdehuun").Count(&tproduct)

	if assert.NoError(t, CountByProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		body, _ := ioutil.ReadAll(rec.Body)
		preParse := strings.Replace(string(body), "\n", "", 1)
		parsedInt, _ := strconv.Atoi(preParse)

		assert.Equal(t, tproduct, parsedInt)
	}
}
func TestCountByService(t *testing.T)  {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/CountService")

	var tser int
	DB.DB.Table("uilchilgee").Count(&tser)

	if assert.NoError(t, CountByService(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body, _ := ioutil.ReadAll(rec.Body)
		preParse := strings.Replace(string(body), "\n", "", 1)
		parsedInt, _ := strconv.Atoi(preParse)

		assert.Equal(t, tser, parsedInt)
	}

}
func TestCountByAimag(t *testing.T)  {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/CountByAimag")

	var aimag int
	DB.DB.Table("ds_company_register").Group("aimag_id").Count(&aimag)

	if assert.NoError(t, CountByAimag(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body, _ := ioutil.ReadAll(rec.Body)
		preParse := strings.Replace(string(body), "\n", "", 1)
		parsedInt, _ := strconv.Atoi(preParse)

		assert.Equal(t, aimag, parsedInt)
	}
}
func TestCompanyById(t *testing.T)  {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()


	c := e.NewContext(req, rec)
	c.SetPath("/api/CompanyById/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")


	id := "1"
	com := models.DSCompanyRegisterDetail{}
	DB.DB.Table("ds_company_register").Where("id = ?",id).Find(&com)
	if(com.ID >= 1){
		chiglels := []models.DSCompanySub{}
		Buteegdehuuns := []models.DSButeegdehuun{}
		Uilchilgeenuud := []models.DSUilchilgee{}
		DB.DB.Where("company_id = ?",com.ID).Find(&chiglels)
		DB.DB.Where("company_id = ?",com.ID).Find(&Buteegdehuuns)
		DB.DB.Where("company_id = ?",com.ID).Find(&Uilchilgeenuud)

		com.UilAjilgaaChigleluud = chiglels
		com.Buteegdehuuns = Buteegdehuuns
		com.Uilchilgeenuud = Uilchilgeenuud
	}

	jsonString, _ := json.Marshal(com)

	if assert.NoError(t, CompanyById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestAanAimagCompany(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/AanAimagCompany")

	jsonData := []models.AanAimagCounts{}
	DB.DB.Table("ds_company_register").Select("count(*) as too ,ds_company_register.aimagname, aimag_id").Group("aimag_id").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)
	if assert.NoError(t, AanAimagCompany(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestCompanyByProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/CompanyByProduct")

	jsonData := []models.ProductsByCompany{}
	DB.DB.Table("ds_buteegdehuun").Select("count(*) as too ,ds_buteegdehuun.company_name, company_id").Group("company_id").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)
	if assert.NoError(t, CompanyByProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestButeegdehuunDetails(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ProductDetails/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	id := "1"
	buteeid := models.DSButeegdehuun{}

	DB.DB.Table("ds_buteegdehuun").Where("id = ?",id).Find(&buteeid)
	jsonString, _ := json.Marshal(buteeid)

	if assert.NoError(t, ButeegdehuunDetails(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}

}
func TestCompanyDetails(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/CompanyDetails/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	id := "1"

	compan := models.DSCompanyRegister{}
	DB.DB.Where("id = ?",id).Find(&compan)
	jsonString, _ := json.Marshal(compan)
	if assert.NoError(t, CompanyDetails(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}

}
func TestUilchilgeeDetails(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ServiceDetails/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	id := "1"

	uilchilgee := models.DSUilchilgee{}
	DB.DB.Where("id = ?",id).Find(&uilchilgee)
	jsonString, _ := json.Marshal(uilchilgee)
	if assert.NoError(t, UilchilgeeDetails(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}


}
func TestProductByTorol(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ProductByTorol")
	jsonData := []models.ButeegdehuunTorol{}
	DB.DB.Table("ds_buteegdehuun").Select("count(*) as too ,ds_buteegdehuun.torol, torol_id").Group("torol_id").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)

	if assert.NoError(t, ProductByTorol(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}

}
func TestProductByDedTorol(t *testing.T){
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ProductByDedTorol")

	jsonData := []models.ButeegdehuunDedTorol{}
	DB.DB.Table("ds_buteegdehuun").Select("count(*) as too ,ds_buteegdehuun.ded_torol, ded_torol_id").Group("ded_torol_id").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)

	if assert.NoError(t, ProductByDedTorol(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestProductById(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ProductById/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	id := "1"
	pro := models.DSButeegdehuun{}
	DB.DB.Where("id = ?",id).Find(&pro)
	jsonString, _ := json.Marshal(pro)
	if assert.NoError(t, ProductById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}

}
func TestProductByAimag(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ProductByAimag")
	jsonData := []models.ButeegdehuunAimag{}
	DB.DB.Table("ds_buteegdehuun").Select("count(*) as too ,ds_buteegdehuun.aimagname, aimag_id").Group("aimag_id").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)

	if assert.NoError(t,ProductByAimag(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestProductBySalbar(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ProductBySalbar")
	jsonData := []models.ButeegdehuunSalbar{}
	DB.DB.Table("ds_buteegdehuun").Select("count(*) as too, tp_section, salbar_id").Group("salbar_id, tp_section").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)
	if assert.NoError(t,ProductBySalbar(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}

}
func TestServiceByAngilal(t *testing.T){
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ServiceByAngilal")
	jsonData := []models.ServiceAngilalCount{}
	DB.DB.Table("ds_uilchilgee").Select("count(*) as too, ds_uilchilgee.torol,ds_uilchilgee.torol_id").Group("torol_id").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)
	if assert.NoError(t,ServiceByAngilal(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}

}
func TestServiceBySalbar(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ServiceBySalbar")
	jsonData := []models.ServiceSalbarCount{}
	DB.DB.Table("ds_uilchilgee").Select("count(*) as too, tp_section, salbar_id").Group("salbar_id, tp_section").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)
	if assert.NoError(t,ServiceBySalbar(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}

}
func TestServiceById(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ServiceById/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	id := "1"
	uil := models.DSUilchilgee{}
	DB.DB.Where("id = ?",id).Find(&uil)
	jsonString, _ := json.Marshal(uil)
	if assert.NoError(t, ServiceById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestSalbarChart(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/SalbarChart")

	jsonData := []models.ButeegdehuunSalbar{}
	DB.DB.Limit(5).Order("Too desc").Table("ds_buteegdehuun").Select("count(*) as too, ds_company_sub.tp_section, ds_company_sub.salbar_id").Joins("LEFT JOIN ds_company_sub on ds_company_sub.id = ds_buteegdehuun.company_id").Group("ds_company_sub.salbar_id").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)

	if assert.NoError(t,SalbarChart(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}

}
func TestProductByTorolChart(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ProductTypeChart")
	jsonData := []models.ButeegdehuunTorol{}
	DB.DB.Limit(5).Order("Too desc").Table("ds_buteegdehuun").Select("count(*) as too ,ds_buteegdehuun.torol, torol_id").Group("torol_id").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)
	if assert.NoError(t,ProductByTorolChart(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestTravels(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/Travels")
	jsonData := []models.DSJuulchlal{}
	DB.DB.Table("ds_juulchlal").Select("count(*) as too ,ds_juulchlal.aimagname, aimag_id").Group("aimag_id").Find(&jsonData)
	jsonString, _ := json.Marshal(jsonData)
	if assert.NoError(t,Travels(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}

}
func TestTravelById(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/TravelById/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	id := "1"
	travel := models.DSJuulchlal{}
	DB.DB.Table("ds_juulchlal").Where("id = ?",id).Find(&travel)
	if(travel.ID >= 1){
		TravelProduct := []models.DSJuulchinButeegdehuun{}
		TravelSer := []models.DSJuulchinUilchilgee{}
		DB.DB.Where("jiilchin_id = ?",travel.ID).Find(&TravelProduct)
		DB.DB.Where("juulchin_id = ?",travel.ID).Find(&TravelSer)

		travel.TravelProduct = TravelProduct
		travel.TravelSer = TravelSer
	}
	jsonString, _ := json.Marshal(travel)
	if assert.NoError(t, TravelById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestTravelByCompany(t *testing.T){
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/TravelByCompany/:company_id")
	c.SetParamNames("company_id")
	c.SetParamValues("1")
	companyid := "1"

	company := models.DSJuulchlal{}
	DB.DB.Table("ds_juulchlal").Select("ds_juulchlal.ner, ds_company_register.name as companyname").Joins("left join business.ds_company_register on business.ds_juulchlal.company_id=ds_company_register.id").Where("company_id = ?",companyid).Find(&company)
	jsonString, _ := json.Marshal(company)


	if assert.NoError(t, TravelByCompany(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}

}
func TestProductCategory(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/product-category/:aimag_id")
	c.SetParamNames("aimag_id")
	c.SetParamValues("1")
	aimag_id := "1"

	BTorols:=[]models.BTorol{}
	for i := 0; i < len(BTorols); i++{
		var too int;
		var lastMonthToo int;
		if(aimag_id != ""){
			DB.DB.Table("buteegdehuun").Where("torol_id = ? AND aimag_id = ?", BTorols[i].ID, aimag_id).Count(&too)


			DB.DB.Table("buteegdehuun").Where("torol_id = ? AND aimag_id = ? AND created_at BETWEEN (CURDATE() - INTERVAL 30 DAY) AND CURDATE()", BTorols[i].ID).Count(&lastMonthToo)
		} else {
			DB.DB.Table("buteegdehuun").Where("torol_id = ?", BTorols[i].ID).Count(&too)

			DB.DB.Table("buteegdehuun").Where("torol_id = ? AND created_at BETWEEN (CURDATE() - INTERVAL 30 DAY) AND CURDATE()", BTorols[i].ID).Count(&lastMonthToo)
		}

		BTorols[i].Too = too
		BTorols[i].LastMonthToo = lastMonthToo
	}
	jsonString, _ := json.Marshal(BTorols)


	if assert.NoError(t, ProductCategory(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}
}
func TestServiceLastMonthsNumbers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/ServiceLastMonthsNumbers/:aimag_id")
	c.SetParamNames("aimag_id")
	c.SetParamValues("1")
	aimagid := "1"

	UTorol := []models.UTorol{}
	DB.DB.Order("torol ASC").Find(&UTorol)

	for i := 0; i <len(UTorol); i++{
		var too int;
		var lastmonthtoo int;
		if(aimagid != ""){
			DB.DB.Table("uilchilgee").Where("torol_id = ? AND aimag_id = ?",UTorol[i].ID,aimagid).Count(&too)
			//DB.DB.Table("uilchilgee").Where("torol_id = ? AND aimag_id = ? AND created_at BETWEEN (CURDATE() - INTERVAL 30 DAY) AND CURDATE()",UTorol[i].ID).Count(&lastmonthtoo)
		} else {
			DB.DB.Table("uilchilgee").Where("torol_id =?",UTorol[i].ID).Count(&too)
			DB.DB.Table("uilchilgee").Where("torol_id = ? AND aimag_id = ? AND created_at BETWEEN (CURDATE() - INTERVAL 30 DAY) AND CURDATE()",UTorol[i].ID).Count(&lastmonthtoo)
		}
		UTorol[i].Too = too
		UTorol[i].LastMonthToo = lastmonthtoo
	}

	jsonString, _ := json.Marshal(UTorol)


	if assert.NoError(t,ServiceLastMonthsNumbers(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(jsonString), rec.Body.String())
	}



}