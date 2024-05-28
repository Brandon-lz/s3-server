package miniorouters

import (
	"fmt"
	"go_learn/funcs"
	"go_learn/http_service/core"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/minio")
	group.GET("list-path/*path", ListPath)
	group.GET("getfile/*path", GetFile)
	group.DELETE("delete/*path",RemoveFile)
	group.POST("/upload", func(c *gin.Context) {
        // single file
        file, _ := c.FormFile("file")
        log.Println(file.Filename)

		filinfo,err :=file.Open()
		if err!=nil{
			panic(err)
		}
		
		if err:=funcs.UploadFile(filinfo,file.Size,file.Filename);err!=nil{
			panic("upload fail")
		}
		core.SuccessHandler(c,nil)
    })
}


// http://localhost:8080/api/v1/minio/list-path/test
func ListPath(c *gin.Context) {
	path := c.Param("path")
	fmt.Println(path)
	core.SuccessHandler(c, funcs.ListFiles(path))
}

// http://localhost:8080/api/v1/minio/getfile/test/testdata.txt
func GetFile(c *gin.Context) {
	path := c.Param("path")
	data,err := funcs.Get_file(path)
	if err!=nil{
		panic(err)
	}
	var res = make(map[string]interface{})
	res["data"] = data
	res["filename"] = path
	core.SuccessHandler(c, res)
}


func RemoveFile(c *gin.Context){
	path := c.Param("path")
	if funcs.DeleteFile(path)!=nil{
		panic("delete failed")
	}
	core.SuccessHandler(c,nil)
}
