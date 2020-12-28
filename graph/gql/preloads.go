package gql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}

func GetColumns(ctx context.Context, prefix string) ([]string, []Sub) {
	fields:=graphql.CollectFieldsCtx(ctx,nil)

	columns := []string{}
	Subs := []Sub{}

	for _, column := range fields {
		if(column.Selections == nil){
			columns = append(columns, column.Name)
		} else {
			SubsNested := GetSubTables(ctx, column, column.Name)


			Subs = append(Subs, SubsNested...)

		}
	}

	for _, Sub := range Subs {
		if(Sub.Table == prefix){
			return Sub.Columns, Subs
		}
	}

	return  columns, Subs
}
func GetSubTables(ctx context.Context, Field graphql.CollectedField, subTable string)[]Sub{
	fields := graphql.CollectFields(graphql.GetOperationContext(ctx), Field.Selections, nil)

	columns := []string{}
	Subs := []Sub{}

	for _, column := range fields {
		if(column.Selections == nil){
			columns = append(columns, column.Name)
		} else {
			SubsNested := GetSubTables(ctx, column, column.Name)
			Subs = append(Subs, SubsNested...)

		}
	}
	subTableNew := Sub{
		Table: subTable,
		Columns: columns,
	}
	Subs = append(Subs, subTableNew)
	return Subs
}

type Sub struct {
	Table string
	Columns []string
}

func RemoveDuplicate(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}