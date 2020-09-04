package handlers

import (
	//"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	//"os/exec"
	//"github.com/khankhulgun/khankhulgun/tools"

)

func BuildMe(c echo.Context) (err error) {

	//AbsolutePath  := utils.AbsolutePath()
	//
	//cmd, err := exec.Command("/bin/sh", "-c",AbsolutePath+"start.sh").Output()
	//if err != nil {
	//	fmt.Printf("error %s", err)
	//}
	//output := string(cmd)
	//
	//fmt.Println(output)
	return c.Redirect(http.StatusSeeOther, "/auth/login")

}