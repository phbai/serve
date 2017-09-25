package main

import (
    "github.com/gin-gonic/gin"
    // "encoding/json"
)

type FileInfo struct {  
    Name string `json:"name"`
    Size int64 `json:"size"`
    IsDir bool `json:"isDir"`
}

func main() {
    r := gin.Default()
    
    // data, _ := json.Marshal(lists)
   
    // fmt.Printf("%s\n", data)
    r.GET("/list", func(c *gin.Context) {
        lists, _ := GetList(".");

        c.JSON(200, gin.H{
            "result": lists,
        })
    })

    r.GET("/list/:path", func(c *gin.Context) {
        path := c.Param("path")

        lists, _ := GetList(path);

        c.JSON(200, gin.H{
            "result": lists,
        })
    })

    r.GET("/read/:path", func(c *gin.Context) {
        path := c.Param("path")

        c.File(path)
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}