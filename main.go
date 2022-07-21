package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
// videoSrv service.VideoService = service.New()
// videoCtrl controller.VideoController = controller.New(videoSrv)
)

func main() {

	// const mar = "ॐ"
	// fmt.Printf("%T", mar)

	// var mar9 float64 = 15.33
	// var mar9 = 15.33
	// const mar9 = 15.33
	// log.SetFlags(log.Ldate | log.Lshortfile)
	// fmt.Printf("%T\n", mar9)
	// fmt.Printf("%v", mar9)

	// fmt.Println("Hello MithiMili")

	s := gin.Default()
	s.Static("/assets", "./assets")
	s.LoadHTMLGlob("./templates/*.html")
	s.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "ॐ",
		})
		// 	// c.JSON(http.StatusOK, gin.H{
		// 	// 	"message": "ॐ",
		// 	// })
	})
	s.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "ॐ About Page",
		})
		// 	// c.JSON(http.StatusOK, gin.H{
		// 	// 	"message": "ॐ",
		// 	// })
	})
	s.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "ॐ Contact Page",
		})
		// 	// c.JSON(http.StatusOK, gin.H{
		// 	// 	"message": "ॐ",
		// 	// })
	})

	// s.GET("/videos", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, videoCtrl.FindAll())
	// })

	// s.POST("/videos", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, videoCtrl.Save(ctx))
	// })

	s.Run(":8080")

}
