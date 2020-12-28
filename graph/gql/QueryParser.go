package gql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"strings"
	"context"
)

func GetPaginationTargetAndColumns(ctx context.Context) (string, []string, error)  {

	target := ""
	columns := []string{}

	fields:=graphql.CollectFieldsCtx(ctx,nil)
	preloads := GetPreloads(ctx)
	targets := []string{}
	for _,field:=range fields{
		if(field.Name != "page" && field.Name != "total" && field.Name != "last_page"){
			targets = append(targets, field.Name)
		}
	}
	if(len(targets) >= 2){
		return target, columns, gqlerror.Errorf("Pagination not allowed multiple pagination")
	}

	for _, preload := range preloads{
		if(preload != "page" && preload != "total" && preload != "last_page" && preload != targets[0]){
			columns = append(columns, getColumn(targets[0], preload))
		}
	}

	return targets[0], columns, nil
}
func getColumn(target string ,preload string) string{
	return strings.ReplaceAll(preload, target+".", "")
}