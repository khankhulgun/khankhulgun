package handlers

import (
	echo "github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"vp/DB"
	"project/business/models"
	"vp/utils"
	"github.com/foolin/goview/supports/echoview-v4"
)

//home page
func HomeProduction (c echo.Context) error {

	return echoview.Render(c, http.StatusOK, "home_production", map[string]interface{}  {})

}
func HomeDevelopment (c echo.Context) error {

	return echoview.Render(c, http.StatusOK, "home_development", map[string]interface{}  {})

}
func HomeMobile (c echo.Context) error {
	data := HomeData()
	return c.JSON(http.StatusOK, data)
}

func HomeData()map[string]interface{}  {
	contact := Contact()
	slides := NuurSlide()
	buteegdehuun := Buteegdehuun()
	hamtragchBaiguullaga := HamtragchBaiguullaga()
	uilchilgee := Uilchilgee()
	companyBySection := CompanyBySection()
	companyByAimag := CompanyByAimag()
	SalbarCount := SalbarCount()
	AimagCount := AimagCount()
	Zone := Zone()
	Aimag := Aimag()
	Aylalcount := AylalCount()
	CountLastMonthProducts := CountLastMonthProducts()
	CountLastMonthService := CountLastMonthService()
	CountLastMonthTravel := CountLastMonthTravel()

	return  map[string]interface{}{
		"contact":contact,
		"slides":slides,
		"buteegdehuun":buteegdehuun,
		"hamtragchBaiguullaga":hamtragchBaiguullaga,
		"uilchilgee":uilchilgee,
		"companyBySection":companyBySection,
		"companyByAimag":companyByAimag,
		"SalbarCount":SalbarCount,
		"AimagCount":AimagCount,
		"zone":Zone,
		"aimag":Aimag,
		"aylalcount":Aylalcount,
		"CountLastMonthProducts":CountLastMonthProducts,
		"CountLastMonthService":CountLastMonthService,
		"CountLastMonthTravel":CountLastMonthTravel,
	}
}

func Contact () models.Contact {
	contact := models.Contact{}
	DB.DB.Order("id DESC").Find(&contact)
	return contact
}
func NuurSlide() []models.NuurSlide {
	slide := []models.NuurSlide{}
	DB.DB.Order("id DESC").Find(&slide)
	return slide

}
func Buteegdehuun () []models.DSButeegdehuun {
	buteegdehuun := []models.DSButeegdehuun{}
	DB.DB.Order("id DESC").Limit(9).Find(&buteegdehuun)
	return buteegdehuun
}
func HamtragchBaiguullaga () []models.HamtragchBaiguullaga{
	hamtragch := []models.HamtragchBaiguullaga{}
	DB.DB.Find(&hamtragch)
	return hamtragch
}
func Uilchilgee () []models.DSUilchilgee{
	uilchilgee := []models.DSUilchilgee{}
	DB.DB.Order("id DESC").Limit(9).Find(&uilchilgee)
	return uilchilgee
}
func Comments (c echo.Context) (err error) {
	ilgeelt := new(models.Comments)
	if err = c.Bind(ilgeelt); err !=nil {
		return c.JSON(http.StatusOK,map[string]interface{}{
			"status":false,
		})
	}
	DB.DB.NewRecord(&ilgeelt)
	DB.DB.Create(ilgeelt)
	return c.JSON(http.StatusOK,map[string]interface{}{
		"status":true,
	})

}
func SalbarCount() []models.SalbarCountAll {

	salbarCount := []models.SalbarCountAll{}


	DB.DB.Table("tp_section").Find(&salbarCount)

	for i := 0; i <len(salbarCount);i++ {
		var AanCount int
		DB.DB.Table("ds_company_sub").Where("salbar_id = ?",salbarCount[i].TpSectionId).Count(&AanCount)
		salbarCount[i].AanToo = AanCount

		var ButeegdehuunCount int
		DB.DB.Table("ds_buteegdehuun").Joins("LEFT JOIN ds_company_sub on ds_company_sub.company_id = ds_buteegdehuun.company_id").Where("ds_company_sub.salbar_id = ?",salbarCount[i].TpSectionId).Count(&ButeegdehuunCount)
		salbarCount[i].ButeegdehuunToo = ButeegdehuunCount

		var UilchilgeeCount int
		DB.DB.Table("ds_uilchilgee").Joins("LEFT JOIN ds_company_sub on ds_company_sub.company_id = ds_uilchilgee.company_id").Where("ds_company_sub.salbar_id = ?",salbarCount[i].TpSectionId).Count(&UilchilgeeCount)
		salbarCount[i].UilchilgeeToo = UilchilgeeCount
	}


	return salbarCount
}
func Zone () []models.Zone {
	bus := []models.Zone{}
	DB.DB.Order("id DESC").Find(&bus)
	return bus
}
func Aimag ()[] models.Aimag {
	aimag := []models.Aimag{}
	DB.DB.Order("id DESC").Find(&aimag)
	return aimag
}
func AylalCount () int {
	var too int;
	DB.DB.Table("budg_burtgel").Count(&too)
	return too
}
func CountLastMonthProducts () int {
	var too int;
	DB.DB.Table("buteegdehuun").Where("created_at BETWEEN (CURDATE() - INTERVAL 30 DAY) AND CURDATE()").Count(&too)
	return too
}
func CountLastMonthService () int {
	var too int;
	DB.DB.Table("uilchilgee").Where("created_at BETWEEN (CURDATE() - INTERVAL 30 DAY) AND CURDATE()").Count(&too)
	return too
}
func CountLastMonthTravel () int {
	var too int;
	DB.DB.Table("budg_burtgel").Where("created_at BETWEEN (CURDATE() - INTERVAL 30 DAY) AND CURDATE()").Count(&too)
	return too
}

//count
func CompanyBySection() []models.AanSalbarCount {

	aanSalbarCount := []models.AanSalbarCount{}

	DB.DB.Limit(9).Table("ds_company_sub").Select("count(*) as too, ds_company_sub.tp_section, salbar_id").Group("salbar_id").Find(&aanSalbarCount)


	return aanSalbarCount
}
func CompanyByAimag () []models.AnnAimagCount {
	annAimagCount := []models.AnnAimagCount{}
	DB.DB.Table("ds_company_register").Select("count(*) as too, ds_company_register.aimagname,aimag_id").Group("aimag_id").Find(&annAimagCount)

	return annAimagCount
}

 //other api
func Taniltsuulga(c echo.Context) error  {
		taniltsuuga := models.Taniltsuulga{}
		DB.DB.Find(&taniltsuuga)
		return c.JSON(http.StatusOK,taniltsuuga)

}
func CountByCompany (c echo.Context) error {
	var tcom int;
	DB.DB.Table("company_register").Count(&tcom)
	return c.JSON(http.StatusOK,tcom)
}
func CountByProduct (c echo.Context) error{
	var tproduct int;
	DB.DB.Table("buteegdehuun").Count(&tproduct)
	return c.JSON(http.StatusOK,tproduct)
}
func CountByService(c echo.Context) error  {
	var tser int;
	DB.DB.Table("uilchilgee").Count(&tser)
	return c.JSON(http.StatusOK,tser)

}
func CountByAimag (c echo.Context) error {
	var aimag int;
	DB.DB.Table("ds_company_register").Group("aimag_id").Count(&aimag)
	return c.JSON(http.StatusOK,aimag)
}
 // Company

func AanAimagSalbar(c echo.Context) error  {

	aanAimagSalbaraar := []models.AanAimagSalbaraar{}
	DB.DB.Table("aimag").Find(&aanAimagSalbaraar)

	for i := 0; i < len(aanAimagSalbaraar); i++  {

		aanSalbarCount := []models.AanSalbarCount{}
		DB.DB.Table("ds_company_sub").Select("count(*) as too, ds_company_sub.tp_section, .ds_company_subsalbar_id").Joins("left join company_register on company_register.id=company_sub.company_id").Joins("left join aimag on company_register.aimag_id=aimag.id").Where("company_register.aimag_id = ?", aanAimagSalbaraar[i].AimagId).Group("salbar_id").Find(&aanSalbarCount)

		aanAimagSalbaraar[i].Salbars = aanSalbarCount

	}

	return c.JSON(http.StatusOK,aanAimagSalbaraar)
}   //test harah
func CompanyFilter(c echo.Context) error  {
	Page := c.Param("page")
	//QUERY Prepare
	companies := []models.DSCompanyRegister{}
	query := DB.DB



	//Parse post request data by struct
	request := new(models.RequestCompany)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
		})
	}

	if(request.OrderType == "date"){
		if(request.Order == "desc"){
			query = query.Order("id DESC")
		} else {
			query = query.Order("id ASC")
		}

	} else if(request.OrderType == "name"){
		if(request.Order == "desc"){
			query = query.Order("name DESC")                                          // orderloh arga
		} else {
			query = query.Order("name ASC")
		}
	} else {
		query = query.Order("id DESC")
	}




	// Starting Conditions
	if(len(request.AiamagIDS) >= 1){
		query = query.Where("aimag_id IN (?)", request.AiamagIDS)
	}
	if(len(request.SalbarIDs) >= 1){

		subs := []models.CompanySub{}
		DB.DB.Where("salbar_id IN (?)", request.SalbarIDs).Find(&subs)

		companyIDsBySabar := []int{}
		for _, sub := range subs{
			companyIDsBySabar = append(companyIDsBySabar, sub.CompanyID)
		}
		query = query.Where("id in (?)", companyIDsBySabar)
	}
	if(request.Ner != ""){
		query = query.Where("name LIKE ?", "%"+request.Ner+"%").Or("description LIKE ?", "%"+request.Ner+"%")
	}
	PageNumber := 1
	if i, err := strconv.Atoi(Page); err == nil {
		PageNumber = i
	}
	// End Conditions

	//Pagination
	data := utils.Paging(&utils.Param{
		DB:    query,
		Page:  PageNumber,
		Limit: 9,
	}, &companies)
	return c.JSON(http.StatusOK,data)
}
//func CompanyPaginate (c echo.Context)error {
//	cpage := []models.CompanyRegister{}
//	query := DB.DB.Table("company_register").Find(&cpage)
//	data := utils.Paging(&utils.Param{
//		DB:    query,
//		Page:  1,
//		Limit: 9,
//	}, &cpage)
//
//	return c.JSON(http.StatusOK,data)
//}
func CompanyById (c echo.Context) error {
	id := c.Param("id")
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

	return c.JSON(http.StatusOK,com)
}

func CompanyBySections(c echo.Context) error {

	aanSalbarCount := []models.AanSalbarCount{}

	DB.DB.Table("ds_company_sub").Select("count(*) as too, ds_company_sub.tp_section, salbar_id").Group("salbar_id").Find(&aanSalbarCount)

	return  c.JSON(http.StatusOK,aanSalbarCount)
}
func AanAimagCompany (c echo.Context)error {
	aanAimagCount := []models.AanAimagCounts{}
	DB.DB.Table("ds_company_register").Select("count(*) as too ,ds_company_register.aimagname, aimag_id").Group("aimag_id").Find(&aanAimagCount)
	return  c.JSON(http.StatusOK,aanAimagCount)

}
func CompanyByProduct (c echo.Context)error {

	productbycompany := []models.ProductsByCompany{}
	DB.DB.Table("ds_buteegdehuun").Select("count(*) as too ,ds_buteegdehuun.company_name, company_id").Group("company_id").Find(&productbycompany)
	return  c.JSON(http.StatusOK,productbycompany)
}



//ann details

func ButeegdehuunDetails (c echo.Context) error {
	id := c.Param("id")
	buteeid := models.DSButeegdehuun{}

	DB.DB.Table("ds_buteegdehuun").Where("id = ?",id).Find(&buteeid)
	return c.JSON(http.StatusOK,buteeid)
}
func CompanyDetails (c echo.Context) error {
	id := c.Param("id")
	compan := models.DSCompanyRegister{}
	DB.DB.Where("id = ?",id).Find(&compan)
	return c.JSON(http.StatusOK,compan)
}
func UilchilgeeDetails (c echo.Context) error {
	id := c.Param("id")
	uilchilgee := models.DSUilchilgee{}
	DB.DB.Where("id = ?",id).Find(&uilchilgee)
	return c.JSON(http.StatusOK,uilchilgee)
}

//buteegdehuun aimag

func ProductByTorol(c echo.Context) error {
	torolcount := []models.ButeegdehuunTorol{}
	DB.DB.Table("ds_buteegdehuun").Select("count(*) as too ,ds_buteegdehuun.torol, torol_id").Group("torol_id").Find(&torolcount)
	return  c.JSON(http.StatusOK,torolcount)
}
func ProductByDedTorol(c echo.Context) error {
	torolcount := []models.ButeegdehuunDedTorol{}
	DB.DB.Table("ds_buteegdehuun").Select("count(*) as too ,ds_buteegdehuun.ded_torol, ded_torol_id").Group("ded_torol_id").Find(&torolcount)
	return  c.JSON(http.StatusOK,torolcount)
}
func ProductById (c echo.Context) error {
	id := c.Param("id")
	pro := models.DSButeegdehuun{}
	DB.DB.Where("id = ?",id).Find(&pro)
	return c.JSON(http.StatusOK,pro)
}
func ProductByAimag(c echo.Context) error{
	aimagcount := []models.ButeegdehuunAimag{}
	DB.DB.Table("ds_buteegdehuun").Select("count(*) as too ,ds_buteegdehuun.aimagname, aimag_id").Group("aimag_id").Find(&aimagcount)
	return  c.JSON(http.StatusOK,aimagcount)
}
func ProductBySalbar(c echo.Context) error{
	salbarcount := []models.ButeegdehuunSalbar{}
	DB.DB.Table("ds_buteegdehuun").Select("count(*) as too, tp_section, salbar_id").Group("salbar_id, tp_section").Find(&salbarcount)
	return c.JSON(http.StatusOK,salbarcount)
}
func ProductFilter(c echo.Context) error  {

	Page := c.Param("page")

	//QUERY Prepare
	products := []models.DSButeegdehuun{}
	query := DB.DB


	//Parse post request data by struct
	request := new(models.RequestProduct)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
		})
	}

	//order

	if(request.OrderType == "date"){
		if(request.Order == "desc"){
			query = query.Order("id DESC")
		} else {
			query = query.Order("id ASC")
		}

	} else if(request.OrderType == "buteegdehuun"){
		if(request.Order == "desc"){
			query = query.Order("buteegdehuun DESC")
		} else {
			query = query.Order("buteegdehuun ASC")
		}
	} else {
		query = query.Order("id DESC")
	}

	//end order

	// Starting Conditions
	if(len(request.AiamagIDS) >= 1){
		query = query.Where("aimag_id IN (?)", request.AiamagIDS)
	}
	if(len(request.TorolIDs) >= 1){
		query = query.Where("torol_id IN (?)", request.TorolIDs)
	}
	if(len(request.SalbarIDs) >= 1){



		query = query.Where("salbar_id in (?)", request.SalbarIDs)
	}
	if(request.Buteegdehuun != ""){
		query = query.Where("buteegdehuun LIKE ?", "%"+request.Buteegdehuun+"%").Or("tailbar LIKE ?", "%"+request.Buteegdehuun+"%")
	}
	PageNumber := 1
	if i, err := strconv.Atoi(Page); err == nil {
		PageNumber = i
	}
	// End Conditions

	//Pagination
	data := utils.Paging(&utils.Param{
		DB:    query,
		Page:  PageNumber,
		Limit: 9,
	}, &products)

	return c.JSON(http.StatusOK,data)
}      //test hiigeegu

// uilchilgee

func ServiceByAngilal (c echo.Context) error {
	serviceCount := []models.ServiceAngilalCount{}
	DB.DB.Table("ds_uilchilgee").Select("count(*) as too, ds_uilchilgee.torol,ds_uilchilgee.torol_id").Group("torol_id").Find(&serviceCount)


	return c.JSON(http.StatusOK,serviceCount)

}
func ServiceBySalbar (c echo.Context) error {
	serviceCount := []models.ServiceSalbarCount{}
	DB.DB.Table("ds_uilchilgee").Select("count(*) as too, tp_section, salbar_id").Group("salbar_id, tp_section").Find(&serviceCount)


	return c.JSON(http.StatusOK,serviceCount)

}

//func ServicePaginate (c echo.Context)error {
//	upage := []models.Uilchilgee{}
//	query := DB.DB.Table("uilchilgee").Find(&upage)
//	data := utils.Paging(&utils.Param{
//		DB:    query,
//		Page:  1,
//		Limit: 9,
//	}, &upage)
//
//	return c.JSON(http.StatusOK,data)
//}


func ServiceById (c echo.Context) error {
	id := c.Param("id")
	uil := models.DSUilchilgee{}
	DB.DB.Where("id = ?",id).Find(&uil)
	return c.JSON(http.StatusOK,uil)
}
func ServiceByAimag(c echo.Context) error{
	aimagcount := []models.ServiceAimagCount{}
	DB.DB.Table("ds_uilchilgee").Select("count(*) as too ,ds_uilchilgee.aimagname, aimag_id").Group("aimag_id").Find(&aimagcount)
	return  c.JSON(http.StatusOK,aimagcount)
}
func ServiceFilter(c echo.Context) error  {

	Page := c.Param("page")

	//QUERY Prepare
	servies := []models.DSUilchilgee{}
	query := DB.DB


	//Parse post request data by struct
	request := new(models.RequestService)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
		})
	}

	//order

	if(request.OrderType == "date"){
		if(request.Order == "desc"){
			query = query.Order("id DESC")
		} else {
			query = query.Order("id ASC")
		}

	} else if(request.OrderType == "name"){
		if(request.Order == "desc"){
			query = query.Order("name DESC")
		} else {
			query = query.Order("name ASC")
		}
	} else {
		query = query.Order("id DESC")
	}

	//end order

	// Starting Conditions
	if(len(request.AiamagIDS) >= 1){
		query = query.Where("aimag_id IN (?)", request.AiamagIDS)
	}
	if(len(request.TorolIDs) >= 1){
		query = query.Where("torol_id IN (?)", request.TorolIDs)
	}
	if(len(request.SalbarIDs) >= 1){
		if(len(request.SalbarIDs) >= 1){

			query = query.Where("salbar_id in (?)", request.SalbarIDs)
		}
	}
	if(request.Name != ""){
		query = query.Where("name LIKE ?", "%"+request.Name+"%").Or("tailbar LIKE ?", "%"+request.Name+"%")
	}
	PageNumber := 1
	if i, err := strconv.Atoi(Page); err == nil {
		PageNumber = i
	}
	// End Conditions

	//Pagination
	data := utils.Paging(&utils.Param{
		DB:    query,
		Page:  PageNumber,
		Limit: 9,
	}, &servies)

	return c.JSON(http.StatusOK,data)
}

  // aimag

func AimagCount() []models.AimagCountAll {

	aimagCount := []models.AimagCountAll{}
	DB.DB.Table("aimag").Select("id as aimag_id, aimagname,zurag").Find(&aimagCount)
	for i := 0; i <len(aimagCount);i++ {
		var Aanaimag int
		DB.DB.Table("ds_company_register").Where("aimag_id = ?",aimagCount[i].AimagId).Count(&Aanaimag)
		aimagCount[i].AanToo = Aanaimag

		var ButeegdehuunCount int
		DB.DB.Table("ds_buteegdehuun").Where("aimag_id = ?",aimagCount[i].AimagId).Count(&ButeegdehuunCount)
		aimagCount[i].ButeegdehuunToo = ButeegdehuunCount

		var UilchilgeeCount int
		DB.DB.Table("ds_uilchilgee").Where("aimag_id = ?",aimagCount[i].AimagId).Count(&UilchilgeeCount)
		aimagCount[i].UilchilgeeToo = UilchilgeeCount

		var AimagSalbarCount int
		DB.DB.Table("company_sub").Joins("left join company_register on company_register.id=company_sub.company_id").Joins("left join aimag on company_register.aimag_id=aimag.id").Where("company_register.aimag_id = ?",aimagCount[i].AimagId).Group("company_sub.salbar_id").Count(&AimagSalbarCount)
		aimagCount[i].SalbarToo = AimagSalbarCount

	}
	return aimagCount
}    //asuuh

// chart
func SalbarChart(c echo.Context) error{
	salbarcount := []models.ButeegdehuunSalbar{}
	DB.DB.Limit(5).Order("Too desc").Table("ds_buteegdehuun").Select("count(*) as too, ds_company_sub.tp_section, ds_company_sub.salbar_id").Joins("LEFT JOIN ds_company_sub on ds_company_sub.id = ds_buteegdehuun.company_id").Group("ds_company_sub.salbar_id").Find(&salbarcount)
	return c.JSON(http.StatusOK,salbarcount)
}
func ProductByTorolChart(c echo.Context) error {
	torolcount := []models.ButeegdehuunTorol{}
	DB.DB.Limit(5).Order("Too desc").Table("ds_buteegdehuun").Select("count(*) as too ,ds_buteegdehuun.torol, torol_id").Group("torol_id").Find(&torolcount)
	return  c.JSON(http.StatusOK,torolcount)
}


// aylal juulchlal
//func Travels(c echo.Context)error  {
//	travel := []models.DSJuulchlal{}
//	DB.DB.Table("ds_juulchlal").Find(&travel)
//	return c.JSON(http.StatusOK,travel)
//}
func Travels(c echo.Context) error{
	aimagcount := []models.DSJuulchlal{}
	DB.DB.Table("ds_juulchlal").Select("count(*) as too ,ds_juulchlal.aimagname, aimag_id").Group("aimag_id").Find(&aimagcount)
	return  c.JSON(http.StatusOK,aimagcount)
}
func TravelById(c echo.Context) error{
	id := c.Param("id")
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
	return c.JSON(http.StatusOK,travel)

}

//func TravelById (c echo.Context) error {
//	id := c.Param("id")
//	main := models.DSJuulchlal{}
//	DB.DB.Table("ds_juulchlal").Select("ds_juulchlal.ner,ds_juulchlal.taniltsuulga,ds_juulchlal.zurag,ds_juulchlal.zuragnuud,ds_juulchlal.aimag_id,ds_juulchlal.aimagname,ds_juulchlal.created_at,ds_juulchlal.updated_at, ds_company_register.name as companyname").Joins("left join business.ds_company_register on business.ds_juulchlal.company_id=ds_company_register.id").Where("company_id = ?",id).Find(&main)
//	return c.JSON(http.StatusOK,main)
//}
//
//


func TravelByFilter (c echo.Context)error{
	Page :=c.Param("")
	travel := []models.DSJuulchlal{}
	query := DB.DB
	request := new(models.RequestTravel)
	if err := c.Bind(request); err !=nil {
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status": "false",
		})
	}
	if(len(request.AiamagIDS) >= 1){
		query = query.Where("aimag_id IN (?)", request.AiamagIDS)
	}
	if(request.Ner != ""){
		query = query.Where("ner LIKE ?", "%"+request.Ner+"%").Or("taniltsuulga LIKE ?", "%"+request.Ner+"%")
	}

	// gaiham ehlel
	//if(request.Taniltsuulga != "taniltsuulga"){
	//	query = query.Where("taniltsuulga LIKE ?","%"+request.Taniltsuulga)
	//}
	// togosgol
	PageNumber := 1
	if i, err := strconv.Atoi(Page); err == nil {
		PageNumber = i
	}
	data := utils.Paging(&utils.Param{
		DB:    query,
		Page:  PageNumber,
		Limit: 9,
	}, &travel)

	return c.JSON(http.StatusOK,data)
}
func TravelByCompany (c echo.Context) error{
	companyid := c.Param("company_id")
	company := models.DSJuulchlal{}
	DB.DB.Table("ds_juulchlal").Select("ds_juulchlal.ner, ds_company_register.name as companyname").Joins("left join business.ds_company_register on business.ds_juulchlal.company_id=ds_company_register.id").Where("company_id = ?",companyid).Find(&company)
	return c.JSON(http.StatusOK,company)
}



//suuliin sariin too
func ProductCategory (c echo.Context) error {

	BTorols := []models.BTorol{}

	DB.DB.Order("torol ASC").Find(&BTorols)
	aimag_id := c.Param("aimag_id")


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

	return c.JSON(http.StatusOK, BTorols)
}

func ServiceLastMonthsNumbers (c echo.Context) error {
	UTorol := []models.UTorol{}
	DB.DB.Order("torol ASC").Find(&UTorol)
	aimagid := c.Param("aimag_id")

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
	return c.JSON(http.StatusOK,UTorol)

}
