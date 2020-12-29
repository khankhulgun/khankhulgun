package generator

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/dbToStruct"
	"github.com/khankhulgun/khankhulgun/graph/generator/models"
	"github.com/khankhulgun/khankhulgun/graph/generator/plugin/resolvergen"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/DBSchema"
	puzzleModels "github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/models"
	khankhulgunConfig "github.com/khankhulgun/khankhulgun/config"
	"github.com/otiai10/copy"
	"github.com/volatiletech/sqlboiler/strmangle"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
)

func Generate(projectName string) {
	cfg, err := config.LoadConfig("graph/gqlgen.yml")

	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	err = api.Generate(cfg,
		api.AddPlugin(resolvergen.New(projectName+"/graph/resolvers")),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

}
func GenerateSchema(projectName string) {

	GqlTables := []models.GqlTable{}

	preTables := []puzzleModels.VBSchema{}

	DB.DB.Where("type = 'graphql'").Find(&preTables)

	for _, preTable := range preTables {
		GqlTable := models.GqlTable{}
		json.Unmarshal([]byte(preTable.Schema), &GqlTable)
		GqlTables = append(GqlTables, GqlTable)
	}

	resolverTmplate := `package resolvers

import (
	"context"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/graph/gql"
	"%s/graph/model"
	"%s/graph/models"
	%s
)

func %s(ctx context.Context, sorts []*model.Sort, filters []*model.Filter%s) ([]*models.%s, error) {
	%s
	result := []*models.%s{}
	requestColumns, %s:= gql.GetColumns(ctx, "%s")
	requestColumns = append(requestColumns, "%s")
	requestColumns = append(requestColumns, []string{%s}...)
	requestColumns = gql.RemoveDuplicate(requestColumns)
	query := DB.DB.Select(requestColumns)
	columns := %sColumns()
	query, errorFilter := gql.Filter(filters, query,columns)
	if(errorFilter != nil){
		return result, errorFilter 
	}
	query, errorOrder := gql.Order(sorts, query, columns)
	if(errorOrder != nil){
		return result, errorOrder
	}
	err := query.Find(&result).Error

	%s
}

func %sColumns() []string {
	return []string{%s}
}
`
	subTemp := `var %sSubs = map[string]model.SubTable{
	%s
}
func %sSub(table string) model.SubTable {
	return %sSubs[table]
}
`
	setSubTemplate := `
func Set%sSubs(ctx context.Context, parents []*models.%s, subs[]gql.Sub, subSorts []*model.SubSort, subFilters []*model.SubFilter) ([]*models.%s, error) {
	parentIds := ""
	for _, parent := range parents{
		if(parentIds == ""){
			parentIds = strconv.Itoa(parent.%s)
		} else {
			parentIds = parentIds + ","+strconv.Itoa(parent.%s)
		}
	}
	for _, sub := range subs {
		%s
	}

	return  parents, nil
}`
	subSetTemp := `if (sub.Table == "%s"){
			subItem := %sSub("%s")
			sorts := []*model.Sort{}
			filters := []*model.Filter{}
			for _, sort := range subSorts {
				if sort.Table == "%s" {
					newSort := model.Sort{
						Column: sort.Column,
						Order: sort.Order,
					}
					sorts = append(sorts, &newSort)
				}
			}
			for _, filter := range subFilters {
				if filter.Table == "%s" {
					newFilter := model.Filter{
						Column: filter.Column,
						Condition: filter.Condition,
						Value: filter.Value,
					}
					filters = append(filters, &newFilter)
				}
			}
			parentFilter := model.Filter{}

			parentFilter.Condition = "whereIn"
			parentFilter.Column = subItem.ConnectionField
			parentFilter.Value = parentIds
			filters = append(filters, &parentFilter)

			sub.Columns = append(sub.Columns, subItem.ConnectionField)
			SubItems, err  := %s
			if err != nil {
				return parents, err
			}
			for _, SubItemData := range SubItems{
				for i, _ := range parents{
					if(parents[i].%s == SubItemData.%s){
						parents[i].%s = append(parents[i].%s, SubItemData)
					}
				}
			}
		}`

	paginationTmplate := `package resolvers

import (
	"context"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/tools"
	"github.com/khankhulgun/khankhulgun/graph/gql"
	"%s/graph/model"
	"%s/graph/models"
)

func Paginate(ctx context.Context, sorts []*model.Sort, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, page int, size int) (*model.Paginate, error) {

	target, _, err := gql.GetPaginationTargetAndColumns(ctx)
	requestColumns, %s := gql.GetColumns(ctx, target)

	Paginate := model.Paginate{
		Page: 0,
		Total:0,
		LastPage:0,
	}
	if(err != nil){
		return &Paginate, err
	}
	query := DB.DB
`
	QueryContent := "type Query {\n"
	Pagination := "\ntype paginate  {\n    page: Int!\n    total: Int!\n    last_page: Int!\n"

	paginationSub := "_"

	for _, table := range GqlTables {
		modelAlias := DBSchema.GetModelAlias(table.Table)
		Identity := DBSchema.GetModelAlias(table.Identity)
		subTables := []string{}
		subTablesMap := ""
		subSetTemps := ""

		if len(table.Subs) >= 1 {
			paginationSub = "subs"
		}
		for _, sub := range table.Subs {
			subAlias := DBSchema.GetModelAlias(sub.Table)
			subTables = append(subTables, sub.Table)
			subTablesMap = subTablesMap + fmt.Sprintf(`"%s": model.SubTable{
	Table:"%s",
	ParentIdentity:"%s",
	ConnectionField:"%s",
},
`,
				sub.Table,
				sub.Table,
				sub.ParentIdentity,
				sub.ConnectionField,
			)
			subCaller := subAlias + "(ctx, sorts, filters)"
			subHasSub := false
			for _, tableCheck := range GqlTables {
				if tableCheck.Table == sub.Table {
					if len(tableCheck.Subs) >= 1 {
						subHasSub = true
					}
				}
			}
			if subHasSub {
				subCaller = subAlias + "(ctx, sorts, filters, subSorts, subFilters)"
			}
			subSetTemps = subSetTemps + fmt.Sprintf(subSetTemp,
				sub.Table,
				modelAlias,
				sub.Table,
				sub.Table,
				sub.Table,
				subCaller,
				Identity,
				DBSchema.GetModelAlias(sub.ConnectionField),
				subAlias,
				subAlias,
			)

		}
		parentConnectsions := ""
		for _, tableCheck := range GqlTables {
			for _, sub := range tableCheck.Subs {
				if table.Table == sub.Table {
					if parentConnectsions == "" {
						parentConnectsions = "\"" + sub.ConnectionField + "\""
					} else {
						parentConnectsions = parentConnectsions + ",\"" + sub.ConnectionField + "\""
					}
				}
			}
		}

		structStr := dbToStruct.TableToStruct(table.Table, table.HiddenColumns, "models", subTables)
		schema := dbToStruct.TableToGraphql(table.Table, table.HiddenColumns, subTables, false)
		colunms := dbToStruct.TableColumns(table.Table, table.HiddenColumns)
		//schemaOrderBy := dbToStruct.TableToGraphqlOrderBy(table.Table, table.HiddenColumns)

		if len(table.Subs) >= 1 {
			QueryContent = QueryContent + "    " + strmangle.Singular(table.Table) + "(sorts:[sort], filters:[filter], subSorts:[subSort], subFilters:[subFilter]): [" + modelAlias + "!]\n"
		} else {
			QueryContent = QueryContent + "    " + strmangle.Singular(table.Table) + "(sorts:[sort], filters:[filter]): [" + modelAlias + "!]\n"
		}
		Pagination = Pagination + "    " + strmangle.Singular(table.Table) + ":[" + modelAlias + "!]\n"



		WriteFile(structStr, "graph/models/"+modelAlias+".go")

		authCheck := ""
		if table.CheckAuth.IsLoggedIn {
			authCheck = `_, authErr := gql.CheckAuth(ctx, []int{`+strings.Trim(strings.Replace(fmt.Sprint(table.CheckAuth.Roles), " ", ",", -1), "[]")+`})
	if authErr != nil {
		return nil, authErr
	}`
		}

		subFilterOrders := ""
		subFromCtx := "_"
		resolverReturn := `return result, err`
		importStrconv := ``
		if len(table.Subs) >= 1 {
			subFilterOrders = ", subSorts []*model.SubSort, subFilters []*model.SubFilter"
			subFromCtx = "subs"
			resolverReturn = fmt.Sprintf(`if(len(subs) >= 1){
		resultWithSubs, errorsub := Set%sSubs(ctx, result, subs, subSorts, subFilters)
		return resultWithSubs, errorsub
	}else{
		return result, err
	}`, modelAlias)
			importStrconv = "\"strconv\""
		}

		resolver := fmt.Sprintf(resolverTmplate,
			projectName,
			projectName,
			importStrconv,
			modelAlias,
			subFilterOrders,
			modelAlias,
			authCheck,
			modelAlias,
			subFromCtx,
			table.Table,
			table.Identity,
			parentConnectsions,
			modelAlias,
			resolverReturn,
			modelAlias,
			colunms,


		)

		if len(table.Subs) >= 1 {
			resolver = resolver + fmt.Sprintf(subTemp,
				modelAlias,
				subTablesMap,
				modelAlias,
				modelAlias)

			resolver = resolver + fmt.Sprintf(setSubTemplate,
				modelAlias,
				modelAlias,
				modelAlias,
				Identity,
				Identity,
				subSetTemps,
			)

		}

		actions := createActions(table, modelAlias, colunms)

		resolver = resolver +actions
		formattedResolver, err := format.Source([]byte(resolver))

		if(err != nil){
			fmt.Println(err)
			fmt.Println(resolver)
		}

		if err == nil {
			WriteFile(string(formattedResolver), "graph/resolvers/"+modelAlias+".go")
		}

		WriteFile(schema, "graph/schemas/"+modelAlias+".graphql")

		paginationReturn := "return &Paginate, nil"

		if len(table.Subs) >= 1 {
			paginationReturn = fmt.Sprintf(`if len(subs) >= 1 {
				resultWithSubs, errorsub := Set%sSubs(ctx, Paginate.%s, subs, subSorts, subFilters)
				Paginate.%s = resultWithSubs
				return &Paginate, errorsub
			} else {
				return &Paginate, nil
			}`, modelAlias, modelAlias, modelAlias)
		}

		paginationTmplate = paginationTmplate + fmt.Sprintf(`if(target == "%s"){
		%s
		requestColumns = append(requestColumns, "%s")
		requestColumns = append(requestColumns, []string{%s}...)
		query = query.Select(requestColumns)
		data := []*models.%s{}
		
		TabeColumns := %sColumns()
		query, errorFilter := gql.Filter(filters, query,TabeColumns)
		if(errorFilter != nil){
			return &Paginate, errorFilter
		}
		query, errorOrder := gql.Order(sorts, query, TabeColumns)
		if(errorOrder != nil){
			return &Paginate, errorOrder
		}
		
		pagination := tools.Paging(&tools.Param{
			DB:    query,
			Page:  page,
			Limit: size,
		}, &data)
		Paginate.%s = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = pagination.Total
		%s
	}`, strings.ToLower(modelAlias), authCheck, table.Identity, parentConnectsions, modelAlias, modelAlias, modelAlias, paginationReturn) + "\n"

	}

	Pagination = Pagination + "}\n"

	QueryContent = QueryContent + "    paginate(sorts: [sort], filters:[filter], subSorts:[subSort], subFilters:[subFilter], page:Int!, size:Int!): paginate!\n}\n" + Pagination + "\n"

	paginationTmplate = fmt.Sprintf(paginationTmplate,
		projectName,
		projectName,
		paginationSub,
	)
	paginationTmplate = paginationTmplate + "return &Paginate, nil\n}"

	subscriptionSchema, subscriptions:= createSubscription(GqlTables)

	QueryContent = QueryContent + "\n"+subscriptionSchema

	WriteFile(QueryContent, "graph/schemas/schemas.graphql")

	formattedPagination, err := format.Source([]byte(paginationTmplate))
	if err == nil {
		WriteFile(string(formattedPagination), "graph/resolvers/Paginate.go")
	}


	createGraphqlFile(projectName, subscriptions)

	createActionUpdateActions(GqlTables)


}
func createActionUpdateActions(GqlTables []models.GqlTable){

	mutations := `type Mutation {
`
	mutationTemp := `	%s
	%s
    %s`

	for _, table := range GqlTables {
		if(table.Actions.Create || table.Actions.Update){
			modelAlias := DBSchema.GetModelAlias(table.Table)
			schema := dbToStruct.TableToGraphql(table.Table, []string{"created_at", "created_at", "deleted_at", table.Identity}, []string{}, true)


			createMutation := ""
			if(table.Actions.Create){
				descrition := "mutation-create"
				if(table.Subscription){
					descrition = descrition+":subscription-"+modelAlias
				}
				createMutation = fmt.Sprintf("\n    \"%s\"\n    create%s(input: %sInput!):%s!", descrition, modelAlias, modelAlias, modelAlias)
			}
			updateMutation := ""
			if(table.Actions.Update){
				descrition := "mutation-update"
				if(table.Subscription){
					descrition = descrition+":subscription-"+modelAlias
				}
				updateMutation = fmt.Sprintf("\n    \"%s\"\n    update%s(id: ID!, input:%sInput!):%s!", descrition, modelAlias, modelAlias, modelAlias)
			}
			deleteMutation := ""
			if(table.Actions.Delete){
				descrition := "mutation-delete"
				if(table.Subscription){
					descrition = descrition+":subscription-"+modelAlias
				}
				deleteMutation = fmt.Sprintf("\n    \"%s\"\n    delete%s(id: ID!):%s!", descrition, modelAlias, modelAlias)
			}
			mutations = mutations+fmt.Sprintf(mutationTemp, createMutation, updateMutation, deleteMutation)

			WriteFile(schema, "graph/schemas/"+modelAlias+"Input.graphql")
		}
	}
	mutations = mutations +"\n}"
	WriteFile(mutations, "graph/schemas/mutations.graphql")
}
func createGraphqlFile(projectName string, subscriptions []map[string]string){
	temp := `package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/khankhulgun/khankhulgun/graph/gql"
	"github.com/labstack/echo/v4"
	"%s/graph/generated"
	%s
	"net/http"
	"sync"
	"time"
	"math/rand"
)

type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

func Set(e *echo.Echo) {

	e.Use(gql.Process)


	config := generated.Config{Resolvers: &Resolver{
		%s
		mutex:           sync.Mutex{},
	}}
	graphqlHandler := handler.New(generated.NewExecutableSchema(config))
	playgroundHandler := playground.Handler("GraphQL playground", "/query")

	graphqlHandler.AddTransport(transport.POST{})
	graphqlHandler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})


	e.Any("/query", func(c echo.Context) error {
		cc := c.(*gql.CustomContext)
		req := cc.Request()
		res := cc.Response()
		graphqlHandler.ServeHTTP(res, req)

		return nil
	})

	e.GET("/play", func(c echo.Context) error {
		cc := c.(*gql.CustomContext)
		req := cc.Request()
		res := cc.Response()
		playgroundHandler.ServeHTTP(res, req)
		return nil
	})
}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	s := string(b)
	return s
}

`
	tempResolver := `package graph

import (
	%s
	"sync"
)

type Resolver struct {
	%s
	mutex          sync.Mutex
}

`

	subs := ""
	subsR := ""
	for _, subscription := range subscriptions{
		subs = subs + fmt.Sprintf("%s%sChannel: map[string]chan *models.%s{},\n", subscription["model"], subscription["action"], subscription["model"])
		subsR = subsR + fmt.Sprintf("%s%sChannel map[string]chan *models.%s\n", subscription["model"], subscription["action"], subscription["model"])

	}
	modelsImport := ""
	if(len(subscriptions) >= 1){
		modelsImport = fmt.Sprintf("\"%s/graph/models\"",
			projectName,
		)
	}
	graphql := fmt.Sprintf(temp,
		projectName,
		modelsImport,
		subs,
	)
	resolver := fmt.Sprintf(tempResolver,
		modelsImport,
		subsR,
	)

	formatted, err := format.Source([]byte(graphql))

	if err == nil {
		WriteFile(string(formatted), "graph/graphql.go")
	}

	formattedR, errR := format.Source([]byte(resolver))
	if errR == nil {
		WriteFile(string(formattedR), "graph/resolver.go")
	}
}
func createSubscription(GqlTables []models.GqlTable) (string, []map[string]string){

	Subscription := `type Subscription {
`
	SubscriptionTmp := `
    "%s"
    %sCreated: %s!
`
	SubscriptionTmpUpdated := `
    "%s"
    %sUpdated: %s!
`
	SubscriptionTmpDeleted := `
    "%s"
    %sDeleted: %s!
`
	Subscriptions := []map[string]string{}
	SubscriptionFound := false
	for _, table := range GqlTables {

		if(table.Subscription && table.Actions.Create){
			SubscriptionFound = true
			modelAlias := DBSchema.GetModelAlias(table.Table)
			Subscription = Subscription+fmt.Sprintf(SubscriptionTmp, "subscription-created:"+modelAlias, modelAlias, modelAlias)

			Subscriptions = append(Subscriptions, map[string]string{
				"model":modelAlias,
				"action":"Created",
			})
		}

		if(table.Subscription && table.Actions.Update){
			SubscriptionFound = true
			modelAlias := DBSchema.GetModelAlias(table.Table)
			Subscription = Subscription+fmt.Sprintf(SubscriptionTmpUpdated, "subscription-updated:"+modelAlias, modelAlias, modelAlias)

			Subscriptions = append(Subscriptions, map[string]string{
				"model":modelAlias,
				"action":"Updated",
			})
		}

		if(table.Subscription && table.Actions.Delete){
			SubscriptionFound = true
			modelAlias := DBSchema.GetModelAlias(table.Table)
			Subscription = Subscription+fmt.Sprintf(SubscriptionTmpDeleted, "subscription-deleted:"+modelAlias, modelAlias, modelAlias)

			Subscriptions = append(Subscriptions, map[string]string{
				"model":modelAlias,
				"action":"Deleted",
			})
		}

	}
	Subscription = Subscription +"\n}"
	if(SubscriptionFound){
		return Subscription, Subscriptions
	} else  {
		return "", Subscriptions
	}

}
func createActions(table models.GqlTable, modelAlias string, colunms string) string{

	createTemp := `
func Create%s(ctx context.Context, input model.%sInput) (*models.%s, error) {
			row := models.%s{}
			%s
			DB.DB.NewRecord(row)
			err := DB.DB.Create(&row).Error
			return &row, err
		}`
	updateTemp := `
func Update%s(ctx context.Context, id string, input model.%sInput) (*models.%s, error) {
			row := models.%s{}
			DB.DB.Where("%s = ?", id).Find(&row)
			%s
			err := DB.DB.Save(&row).Error

			return &row, err
		}`
	deleteTemp := `
func Delete%s(ctx context.Context, id string) (*models.%s, error) {
			row := models.%s{}
			DB.DB.Where("%s = ?", id).Find(&row)
			err := DB.DB.Where("%s = ?", id).Delete(&models.%s{}).Error
			return &row, err
		}`

	actions := ""

	columnsWithInput := ""
	colunms = strings.ReplaceAll(colunms, "\"", "")
	colunms = strings.ReplaceAll(colunms, " ", "")
	for _, column :=range strings.Split(colunms, ","){
		if(column != table.Identity && column != "created_at" && column != "created_at" && column != "deleted_at"){
			columnReady := DBSchema.GetModelAlias(column)
			columnsWithInput = columnsWithInput + fmt.Sprintf("row.%s = *input.%s\n",
				columnReady,
				columnReady,
			)
		}
	}

		if(table.Actions.Create){

			actions = actions + fmt.Sprintf(createTemp,
				modelAlias,
				modelAlias,
				modelAlias,
				modelAlias,
				columnsWithInput,
			)
		}

		if(table.Actions.Update){

			actions = actions + fmt.Sprintf(updateTemp,
				modelAlias,
				modelAlias,
				modelAlias,
				modelAlias,
				table.Identity,
				columnsWithInput,
			)
		}

		if(table.Actions.Delete){

			actions = actions + fmt.Sprintf(deleteTemp,
				modelAlias,
				modelAlias,
				modelAlias,
				table.Identity,
				table.Identity,
				modelAlias,
			)
		}

		return actions

}
func GQLInit(projectPath string, projectName string) {

	dir := projectPath
	AbsolutePath := khankhulgunConfig.AbsolutePath()

	modelsPatch := dir + "/graph/models"
	schemaPatch := dir + "/graph/schemas"
	resolversPatch := dir + "/graph/resolvers"
	schemaCommonPatch := dir + "/graph/schemas-common"
	if _, err := os.Stat(modelsPatch); os.IsNotExist(err) {

		os.MkdirAll(modelsPatch, 0755)
		os.MkdirAll(schemaPatch, 0755)
		os.MkdirAll(resolversPatch, 0755)
		os.MkdirAll(schemaCommonPatch, 0755)

	} else {

		os.RemoveAll(modelsPatch)
		os.RemoveAll(schemaPatch)
		os.RemoveAll(resolversPatch)
		os.RemoveAll(schemaCommonPatch)
		os.MkdirAll(modelsPatch, 0755)
		os.MkdirAll(schemaPatch, 0755)
		os.MkdirAll(resolversPatch, 0755)
		os.MkdirAll(schemaCommonPatch, 0755)
	}
	copy.Copy(AbsolutePath+"graph/schemas-common/", dir+"/graph/schemas-common/")

	gqlgenFile, _ := ioutil.ReadFile(AbsolutePath + "/graph/gqlgen.yml.example")
	gqlgenFileContent := strings.ReplaceAll(string(gqlgenFile), "PROJECTNAME", projectName)
	WriteFile(gqlgenFileContent, dir+"/graph/gqlgen.yml")

	//graphqlFile, _ := ioutil.ReadFile(AbsolutePath + "/graph/graphql.go.exmaple")
	//graphqlFileContent := strings.ReplaceAll(string(graphqlFile), "PROJECTNAME", projectName)
	//WriteFile(graphqlFileContent, dir+"/graph/graphql.go")
	GenerateSchema(projectName)
	Generate(projectName)

}
func WriteFile(fileContent string, path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err, f)
	}

	l2, err := f.WriteString(fileContent)
	if err != nil {
		fmt.Println(err, l2)
		f.Close()
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)

	}
}
