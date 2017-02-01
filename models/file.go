package models

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"recipes-api/config"
)

type File struct {
	Filename string
	Path     string
	Body     string
}

func DownloadFile(c *gin.Context) {
	path := config.Get().OutputDir + "/" + c.Params.ByName("name")

	c.Header("Content-Disposition", "attachment; filename="+c.Params.ByName("name"))
	c.File(path)
}

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	path := config.Get().OutputDir + "/" + header.Filename

	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadInFile(RelativePath string) string {
	absPath, _ := filepath.Abs(RelativePath)

	b, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Println(err)
	}
	return string(b)
}

func CreateDirectory(AbsolutePath string) {
	//TODO: update permissions to be stricter?
	var err = os.Mkdir(AbsolutePath, 0777)
	checkError(err)
}

func CreateFile(AbsolutePath string) {
	//detect if file exists
	var _, err = os.Stat(AbsolutePath)

	//if not, create it
	if os.IsNotExist(err) {
		var file, err = os.Create(AbsolutePath)
		checkError(err)
		defer file.Close()
	}
}

func WriteFile(AbsolutePath string, Contents string) {
	CreateFile(AbsolutePath)

	// open file using READ & WRITE permission
	var file, err = os.OpenFile(AbsolutePath, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	_, err = file.WriteString(Contents)
	checkError(err)
	// save changes
	err = file.Sync()
	checkError(err)

}

func fileExists(AbsolutePath string) bool {
	if _, err := os.Stat(AbsolutePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func checkError(err error) {
	//todo: pass error up?
	if err != nil {
		log.Println(err.Error())
	}
}
