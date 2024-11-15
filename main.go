package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/assets", "./assets")

	// router.LoadHTMLGlob("templates/*")
	// Load HTML templates, including subfolders
	router.LoadHTMLGlob("templates/**/*")

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/index.tmpl", gin.H{
			"Title": "Home"})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/about.tmpl", gin.H{
			"Title": "About Us"})
	})

	router.GET("/service", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/service.tmpl", gin.H{
			"Title": "Services"})
	})

	router.GET("/industries", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/industries.tmpl", gin.H{
			"Title": "Industries"})
	})

	router.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/contact.tmpl", gin.H{
			"Title": "Contact Us"})
	})

	router.GET("/career", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/career.tmpl", gin.H{
			"Title": "Career"})
	})

	router.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/blog.tmpl", gin.H{
			"Title": "Blogs"})
	})

	// router.GET("/subfolder1/page1", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "subfolder1/page1.html", nil)
	// })

	// router.GET("/subfolder2/page2", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "subfolder2/page2.html", nil)
	// })

	// Start the server
	router.Run(":8080")
}
