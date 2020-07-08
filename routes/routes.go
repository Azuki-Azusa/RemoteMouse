package routes

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)


func Routes() {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.Delims("{{{","}}}")
    router.LoadHTMLGlob("templates/*.html")

    router.Static("/assets", "./assets")
    router.StaticFile("/favicon.ico", "./favicon.ico")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "controller.html", gin.H{})
    })
    router.GET("/controller", func(c *gin.Context) {
        c.HTML(http.StatusOK, "controller.html", gin.H{})
    })
    
    router.POST("/press", func(c *gin.Context) { 
        code := c.PostForm("code")
        fmt.Println(code)
        Press(code)
        c.JSON(200, gin.H{
            "action": "按下",
            "code": code,
        })
    })

    router.POST("/release", func(c *gin.Context) { 
        code := c.PostForm("code")
        fmt.Println(code)
        Release(code)
        c.JSON(200, gin.H{
            "action": "松开",
            "code": code,
        })
    })
    
    router.Run(":2434") // listen and serve on 0.0.0.0:2434

}