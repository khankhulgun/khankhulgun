package dataform

import (
	"fmt"
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
	"os"
	"path/filepath"
	"strings"
	"github.com/thedevsaddam/govalidator"
	"time"
)



func CheckFileExist(filepath string, fileName string, fileType string, ext string, i int) string{

	newFileName := ""
	if i > 0{
		newFileName = fmt.Sprintf("%v", i)+"-"+fileName+ext
	} else {
		newFileName = fileName+ext
	}
	_, err := os.Stat(filepath+newFileName)

	if os.IsNotExist(err) {
		return newFileName
	} else {
		i = i+1
		return CheckFileExist(filepath, fileName,fileType, ext, i)
	}
}

func makeUploadable(src io.Reader, fileType string, ext string, fileName string) map[string]string  {
	var name = strings.TrimRight(fileName, ext)
	currentTime := time.Now()
	year := fmt.Sprintf("%v", currentTime.Year())
	month := fmt.Sprintf("%v", currentTime.Month())

	var publicPath string = "public"
	var uploadPath string = "/uploaded/"+fileType+"/"+year+"/"+month+"/"
	var fullPath string = publicPath+uploadPath

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		os.MkdirAll(fullPath, 0755)
		// Create your file
	}

	var i int = 0
	newFileName := CheckFileExist(fullPath, name, fileType, ext, i)
	// Destination
	dst, err := os.Create(fullPath+newFileName)
	if err != nil {
		return map[string]string{
			"httpPath":"",
			"basePath":"",
			"fileName":"",
		}
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return map[string]string{
			"httpPath":"",
			"basePath":"",
			"fileName":"",
		}
	}




	return map[string]string{
		"httpPath":uploadPath+newFileName,
		"basePath":fullPath,
		"fileName":newFileName,
	}

}

func Upload(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
			"message": "file not found",
		})
	}



	//
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
			"message": "server error",
		})
	}
	defer src.Close()
	//srcMime := src



	var ext_ = filepath.Ext(file.Filename)
	ext := strings.ToLower(strings.TrimPrefix(ext_, "."))
	var fileType string = "images"
	rules := govalidator.MapData{
		//"file:file": []string{"ext:jpg,png,jpeg,svg,JPG,PNG,JPEG,SVG", "size:100000", "mime:jpg,png,jpeg,svg,JPG,PNG,JPEG,SVG", "required"},
		"file:file": []string{"ext:jpg,png,jpeg,svg,gif,JPG,PNG,JPEG,SVG,GIF", "size:100000000",  "required"},
	}
	mimeTypes := []string{
		"image/svg+xml",
		"image/jpeg",
		"image/png",
		"image/gif",
	}


	if ext == "dwg" || ext == "pdf" || ext == "swf" || ext == "doc" || ext == "docx" || ext == "csv" || ext == "xls" || ext == "xlsx" || ext == "ppt" || ext == "pptx" {
		rules = govalidator.MapData{
			"file:file": []string{"ext:xls,xlsx,doc,docx,pdf,ppt,pptx,csv,XLS,XLSX,DOC,DOCX,PDF,PPT,PPTX,CSV", "size:4000000",  "required"},
		}
		mimeTypes = []string{
			"application/acad",
			"application/pdf",
			"application/x-shockwave-flash",
			"application/x-shockwave-flash2-preview",
			"application/msword",
			"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
			"application/vnd.ms-excel",
			"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
			"application/vnd.ms-powerpoint",
			"application/vnd.openxmlformats-officedocument.presentationml.presentation",
			"text/csv",
		}
		fileType = "documents"
	}
	if ext == "mp4" ||ext == "m4v" || ext == "avi"{
		rules = govalidator.MapData{
			"file:file": []string{"ext:mp4,m4v,avi,MP4,M4V,AVI", "size:40000000",  "required"},
		}
		mimeTypes = []string{
			"video/mp4",
			"video/x-m4v",
			"video/x-msvideo",
		}
		fileType = "videos"
	}
	if ext == "mp3" || ext == "wav" {
		rules = govalidator.MapData{
			"file:file": []string{"ext:mp3,wav,MP3,WAV", "size:400000",  "required"},
		}
		mimeTypes = []string{
			"audio/mpeg",
			"audio/wav",
		}
		fileType = "audios"
	}


	//mimeType, _, err  := mimetype.DetectReader(srcMime)

	mimeType :="1"


	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "false",
			"message": "can't parse file mime, server error",
		})
	}
	mimeAllowed := false
	for _, m := range mimeTypes{
		if m == mimeType{
			mimeAllowed = true
		}
	}

	if mimeAllowed == false {
		//return c.JSON(http.StatusBadRequest, map[string]string{
		//	"status": "false",
		//	"message": "file mime not allowed",
		//})
	}


	messages := govalidator.MapData{
		"file:file": []string{"ext:file not allowed", "required:File required",  "size:File size too big"},
	}
	opts := govalidator.Options{
		Request:c.Request(),     // request object
		Rules:   rules, // rules map,
		Messages: messages,
	}
	v := govalidator.New(opts)
	e := v.Validate()

	if len(e) >= 1{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "false",
			"message": e,
		})
	}
	upload := makeUploadable(src, fileType, ext_, file.Filename)
	return c.String(http.StatusOK, upload["httpPath"])
}
