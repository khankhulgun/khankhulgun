package generator

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/khankhulgun/khankhulgun/dbToStruct"
	"github.com/khankhulgun/khankhulgun/graph/generator/models"
	"github.com/khankhulgun/khankhulgun/graph/generator/plugin/resolvergen"
	"github.com/khankhulgun/khankhulgun/lambda/modules/puzzle/DBSchema"
	khankhulgunConfig "github.com/khankhulgun/khankhulgun/config"
	"github.com/otiai10/copy"
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
		api.AddPlugin(resolvergen.New(projectName+"/graph/resolvers")), // This is the magic line
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

}
func GenerateSchema(projectName string)  {

	GqlTables := []models.GqlTable{}

	jsonstr := `[
  {
    "table": "users",
    "checkAuth": {
      "isLoggedIn": true,
      "roles": [1]
    },
    "hidden_columns": ["password", "status"]
  },
  {
    "table": "aimag",
    "checkAuth": {
      "isLoggedIn": false,
      "roles": [1]
    },
    "hidden_columns": []
  }
]`
	json.Unmarshal([]byte(jsonstr), &GqlTables)



	resolverTmplate := `package resolvers

import (
	"context"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/graph/builder"
	"%s/graph/model"
	"%s/graph/models"
)

func %s(ctx context.Context, sorts []*model.Sort, filters []*model.Filter) ([]*models.%s, error) {
	%s
	result := []*models.%s{}
	preloads := builder.GetPreloads(ctx)
	query := DB.DB.Select(preloads)
	columns := %sColumns()
	query, errorFilter := builder.Filter(filters, query,columns)
	if(errorFilter != nil){
		return result, errorFilter 
	}
	query, errorOrder := builder.Order(sorts, query, columns)
	if(errorOrder != nil){
		return result, errorOrder
	}
	err := query.Find(&result).Error

	return result, err
}

func %sColumns() []string {
	return []string{%s}
}
`
	paginationTmplate := `package resolvers

import (
	"context"
	"github.com/khankhulgun/khankhulgun/DB"
	"github.com/khankhulgun/khankhulgun/tools"
	"github.com/khankhulgun/khankhulgun/graph/builder"
	"%s/graph/model"
	"%s/graph/models"
)

func Paginate(ctx context.Context, sorts []*model.Sort, filters []*model.Filter, page int, size int) (*model.Paginate, error) {


	target, columns, err := builder.GetPaginationTargetAndColumns(ctx)

	Paginate := model.Paginate{
		Page: 0,
		Total:0,
		LastPage:0,
	}
	if(err != nil){
		return &Paginate, err
	}
	query := DB.DB.Select(columns)
`
	QueryContent := "type Query {\n"
	Pagination := "\ntype paginate  {\n    page: Int!\n    total: Int!\n    last_page: Int!\n"

	paginationTmplate = fmt.Sprintf(paginationTmplate,
		projectName,
		projectName)

	for _, table := range GqlTables{


		structStr := dbToStruct.TableToStruct(table.Table, table.HiddenColumns, "models")
		schema := dbToStruct.TableToGraphql(table.Table, table.HiddenColumns)
		colunms := dbToStruct.TableColumns(table.Table, table.HiddenColumns)
		//schemaOrderBy := dbToStruct.TableToGraphqlOrderBy(table.Table, table.HiddenColumns)
		modelAlias := DBSchema.GetModelAlias(table.Table)

		QueryContent = QueryContent + "    "+strings.ToLower(modelAlias)+"(sorts:[sort], filters:[filter]): ["+modelAlias+"!]\n"
		Pagination = Pagination + "    "+strings.ToLower(modelAlias)+":["+modelAlias+"!]\n"

		WriteFile(structStr, "graph/models/" + modelAlias + ".go")

		authCheck := ""
		if(table.CheckAuth.IsLoggedIn){
			authCheck = `_, authErr := builder.CheckAuth(ctx)
	if authErr != nil {
		return nil, authErr
	}`
		}
		resolver := fmt.Sprintf(resolverTmplate,
			projectName,
			projectName,
			modelAlias,
			modelAlias,
			authCheck,
			modelAlias,
			modelAlias,
			modelAlias,
			colunms)

		WriteFile(resolver, "graph/resolvers/" + modelAlias + ".go")
		WriteFile(schema, "graph/schemas/" + modelAlias + ".graphql")


		paginationTmplate = paginationTmplate + fmt.Sprintf(`if(target == "%s"){
		%s
		data := []*models.%s{}
		
		TabeColumns := %sColumns()
		query, errorFilter := builder.Filter(filters, query,TabeColumns)
		if(errorFilter != nil){
			return &Paginate, errorFilter
		}
		query, errorOrder := builder.Order(sorts, query, TabeColumns)
		if(errorOrder != nil){
			return &Paginate, errorOrder
		}
		errDB := query.Find(&data).Error
		pagination := tools.Paging(&tools.Param{
			DB:    query,
			Page:  page,
			Limit: size,
		}, &data)
		Paginate.%s = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = pagination.Total
		return &Paginate, errDB
	}`, strings.ToLower(modelAlias),authCheck, modelAlias, modelAlias, modelAlias) +"\n"

	}

	Pagination = Pagination + "}\n"

	QueryContent = QueryContent + "    paginate(sorts: [sort], filters:[filter], page:Int!, size:Int!): paginate!\n}\n"+Pagination+"\n"
	paginationTmplate = paginationTmplate+"return &Paginate, nil\n}"
	WriteFile(QueryContent, "graph/schemas/schemas.graphql")

	formattedPagination, err := format.Source([]byte(paginationTmplate))
	if err == nil {
		WriteFile(string(formattedPagination), "graph/resolvers/Paginate.go")
	}


}

func GQLInit(projectPath string, projectName string) {

	dir := projectPath
	AbsolutePath := khankhulgunConfig.AbsolutePath()

	modelsPatch := dir+"/graph/models"
	schemaPatch := dir+"/graph/schemas"
	resolversPatch := dir+"/graph/resolvers"
	schemaCommonPatch := dir+"/graph/schemas-common"
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

	gqlgenFile, _ := ioutil.ReadFile(AbsolutePath+"/graph/gqlgen.yml.example")
	gqlgenFileContent := strings.ReplaceAll(string(gqlgenFile), "PROJECTNAME", projectName)
	WriteFile(gqlgenFileContent, dir+"/graph/gqlgen.yml")

	graphqlFile, _ := ioutil.ReadFile(AbsolutePath+"/graph/graphql.go.exmaple")
	graphqlFileContent := strings.ReplaceAll(string(graphqlFile), "PROJECTNAME", projectName)
	WriteFile(graphqlFileContent, dir+"/graph/graphql.go")
	GenerateSchema(projectName)
	Generate(projectName)

}

func WriteFile(fileContent string, path string){
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
