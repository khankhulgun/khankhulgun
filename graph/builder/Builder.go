package builder

import (
	"encoding/json"
	"errors"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"strings"
)

func Order(sortsPre interface{}, query *gorm.DB, columns []string) (*gorm.DB, error) {
	var sorts []map[string]string
	order, err := json.Marshal(sortsPre)
	if err != nil {
		return query, errors.New("Please insert correct sort values")
	}
	err2 := json.Unmarshal(order, &sorts)
	if err2 != nil {
		return query, errors.New("Please insert correct sort values")
	}
	for _, sort := range sorts {

		errCol := CheckColumns(sort["column"], columns)
		if errCol != nil {
			return query, errCol
		}
		query = query.Order(sort["column"] + " " + sort["order"])

	}
	return query, nil

}
func Filter(filtersPre interface{}, query *gorm.DB, columns []string) (*gorm.DB, error) {

	var filters []map[string]string
	order, err := json.Marshal(filtersPre)
	if err != nil {
		return query, errors.New("Please insert correct filter value")
	}
	err2 := json.Unmarshal(order, &filters)

	if err2 != nil {
		return query, errors.New("Please insert correct filter value")
	}

	if len(filters) >= 1 {
		for _, filter := range filters {

			errCol := CheckColumns(filter["column"], columns)

			fmt.Println(filter["condition"])
			if errCol != nil {
				return query, errCol
			}
			k := filter["column"]
			v := filter["value"]

			switch filter["condition"] {

			case "equals":
				query = query.Where(k+" = ?", v)
			case "notEqual":
				query = query.Where(k+" != ?", v)
			case "contains":
				query = query.Where("LOWER("+k+") LIKE ?", "%"+strings.ToLower(v)+"%")
			case "notContains":
				query = query.Where("LOWER("+k+") not LIKE ?", "%"+strings.ToLower(v)+"%")
			case "startsWith":
				query = query.Where("LOWER("+k+")  LIKE ?", strings.ToLower(v)+"%")
			case "endsWith":
				query = query.Where("LOWER("+k+")  LIKE ?", "%"+strings.ToLower(v))
			case "greaterThan":
				query = query.Where(k+" >= ?", v)
			case "greaterThanOrEqual":
				query = query.Where(k+" > ?", v)
			case "lessThan":
				query = query.Where(k+" < ?", v)
			case "lessThanOrEqual":
				query = query.Where(k+" <= ?", v)
			case "isNull":
				query = query.Where(k+" IS NULL")
			case "notNull":
				query = query.Where(k+" IS NOT NULL")

			default:
				return query, errors.New(filter["condition"] + ": is wrong condition")
			}

		}

	}

	return query, nil
}
func CheckColumns(column string, columns []string) error {
	for _, value := range columns {
		if value == column {
			return nil
		}
	}
	return errors.New(column + ": Column not found")
}

type CustomContext struct {
	echo.Context
	ctx    context.Context
}
func Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "EchoContextKey", c)
		c.SetRequest(c.Request().WithContext(ctx))

		cc := &CustomContext{c, ctx}

		return next(cc)
	}
}