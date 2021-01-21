package krudMW

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/khankhulgun/khankhulgun/DB"
	arcGIS "github.com/khankhulgun/khankhulgun/lambda/modules/arcGIS/handlers"
	"github.com/khankhulgun/khankhulgun/lambda/modules/krud/models"
	"github.com/khankhulgun/khankhulgun/lambda/modules/notify/handlers"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
)
type (
	// Skipper defines a function to skip middleware. Returning true skips processing
	// the middleware.
	Skipper func(echo.Context) bool

	// BeforeFunc defines a function which is executed just before the middleware.
	BeforeFunc func(echo.Context)
)
func DefaultSkipper(echo.Context) bool {
	return false
}
type crudResponse struct {
	Data struct{
		ID        int     `gorm:"column:id;" json:"id"`
	} `json:"data"`
}

type (
	// BodyDumpConfig defines the config for BodyDump middleware.
	BodyDumpConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper Skipper

		// Handler receives request and response payload.
		// Required.
		Handler BodyDumpHandler
	}

	// BodyDumpHandler receives the request and response payload.
	BodyDumpHandler func(echo.Context, []byte, []byte)

	bodyDumpResponseWriter struct {
		io.Writer
		http.ResponseWriter
	}
)

var (
	// DefaultBodyDumpConfig is the default BodyDump middleware config.
	DefaultBodyDumpConfig = BodyDumpConfig{
		Skipper: DefaultSkipper,
	}
)

// BodyDump returns a BodyDump middleware.
//
// BodyDump middleware captures the request and response payload and calls the
// registered handler.


// BodyDumpWithConfig returns a BodyDump middleware with config.
// See: `BodyDump()`.
func CrudLoggerNew(next echo.HandlerFunc, useNotify bool, UseArcGISConnection bool, GetMODEL func(schema_id string) (string, interface{}), GetGridMODEL func(schema_id string) (interface{}, interface{}, string, string, interface{}, string), isDeleteAction bool) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			action := c.Param("action")

			if(isDeleteAction){
				action = "delete"
			}
			if(action == "store" || action == "update" || action == "delete" || action == "edit"){
				// Request
				reqBody := []byte{}
				if c.Request().Body != nil { // Read
					reqBody, _ = ioutil.ReadAll(c.Request().Body)
				}
				c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset

				// Response
				resBody := new(bytes.Buffer)
				mw := io.MultiWriter(c.Response().Writer, resBody)
				writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
				c.Response().Writer = writer

				if err = next(c); err != nil {
					c.Error(err)
				}
				req := c.Request()
				user := c.Get("user").(*jwt.Token)
				claims := user.Claims.(jwt.MapClaims)
				userID := claims["id"].(float64)
				schemaId, _ := strconv.ParseInt(c.Param("schemaId"), 10, 64)
				RowId := c.Param("id")

				Log := models.CrudLog{
					UserId: int64(userID),
					Ip: c.RealIP(),
					UserAgent: req.UserAgent(),
					Action: action,
					SchemaId: schemaId,
					RowId: RowId,
					Input: string(resBody.Bytes()),
				}

				if(action == "store"){

					var response crudResponse
					if err := json.Unmarshal(resBody.Bytes(), &response); err != nil {
						panic(err)
					}
					Log.RowId = strconv.Itoa(response.Data.ID)

				}

				DB.DB.Create(&Log)

				if(useNotify){
					if(action == "store" || action == "update" || action == "delete"){
						handlers.BuildNotification(reqBody, schemaId, action, int64(userID))
					}
				}

				if(UseArcGISConnection){
					if(action == "store" || action == "update"){
						arcGIS.SAVEGIS(reqBody, schemaId, action, Log.RowId, GetMODEL)
					}else if(action == "delete"){
						arcGIS.DELTEGIS(reqBody, schemaId, action, Log.RowId, GetGridMODEL)
					}
				}

				return
			} else {
				return next(c)
			}



		}
	}


func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *bodyDumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

