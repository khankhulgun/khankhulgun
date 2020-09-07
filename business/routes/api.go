package routes

import (
	echo "github.com/labstack/echo/v4"
	"project/business/handlers"
)

func Api(e *echo.Echo) {
	a := e.Group("/api")
	a.GET("/home", handlers.HomeMobile)                     //gar utasnii

	// other page
	a.GET("/Info", handlers.Taniltsuulga)
	a.GET("/CountCompany", handlers.CountByCompany)
	a.GET("/CountProduct", handlers.CountByProduct)
	a.GET("/CountService", handlers.CountByService)
	a.GET("/CountByAimag", handlers.CountByAimag)
	a.GET("/product-category", handlers.ProductCategory)
	a.GET("/product-category/:aimag_id", handlers.ProductCategory)
	a.GET("/ServiceLastMonthsNumbers", handlers.ServiceLastMonthsNumbers)
	a.GET("/ServiceLastMonthsNumbers/:aimag_id", handlers.ServiceLastMonthsNumbers)

	//aimag
	a.GET("/AimagSalbar", handlers.AanAimagSalbar)

	//baiguullaga
	a.GET("/AanAimagCompany", handlers.AanAimagCompany)
	a.GET("/CompanyBySections", handlers.CompanyBySections)
	a.POST("/company/:page", handlers.CompanyFilter)
	//a.GET("/CompanyPaginate", handlers.CompanyPaginate)
	a.GET("/CompanyById/:id", handlers.CompanyById)
	a.GET("/CompanyByProduct", handlers.CompanyByProduct)

	// ann details
	a.GET("/ProductDetails/:id", handlers.ButeegdehuunDetails)
	a.GET("/CompanyDetails/:id", handlers.CompanyDetails)
	a.GET("/ServiceDetails/:id", handlers.UilchilgeeDetails)

	//chart
	a.GET("/SalbarChart", handlers.SalbarChart)
	a.GET("/ProductTypeChart", handlers.ProductByTorolChart)

	//buteegdehuun
	a.GET("/ProductByTorol", handlers.ProductByTorol)
	a.GET("/ProductByDedTorol", handlers.ProductByDedTorol)
	a.GET("/ProductById/:id", handlers.ProductById)
	a.GET("/ProductByAimag", handlers.ProductByAimag)
	a.GET("/ProductBySalbar", handlers.ProductBySalbar)
	a.POST("/ProductFilter/:page", handlers.ProductFilter)

	//uilchilgee

	a.GET("/ServiceByAngilal", handlers.ServiceByAngilal)
	a.GET("/ServiceBySalbar", handlers.ServiceBySalbar)

	//a.GET("/ServicePaginate", handlers.ServicePaginate)

	a.GET("/ServiceById/:id", handlers.ServiceById)
	a.GET("/ServiceByAimag", handlers.ServiceByAimag)
	a.POST("/ServiceFilter/:page", handlers.ServiceFilter)

	//POST

	a.POST("/comments", handlers.Comments)

	// aylal
	a.GET("/Travels", handlers.Travels)
	a.GET("/TravelById/:id", handlers.TravelById)
	a.GET("/TravelByCompany/:company_id", handlers.TravelByCompany)
	a.POST("/TravelByFilter/:page", handlers.TravelByFilter)

	// updated tested
}

