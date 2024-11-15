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
			"Title": "About Us", "Img": "hero1"})
	})

	router.GET("/service", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/service.tmpl", gin.H{
			"Title": "Services", "Img": "hero2"})
	})

	router.GET("/industries", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/industries.tmpl", gin.H{
			"Title": "Industries", "Img": "hero3"})
	})

	router.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/contact.tmpl", gin.H{
			"Title": "Contact Us", "Img": "hero2"})
	})

	router.GET("/career", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/career.tmpl", gin.H{
			"Title": "Career", "Img": "overlay"})
	})

	router.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main/blog.tmpl", gin.H{
			"Title": "Blogs", "Img": "overlay"})
	})

	router.Run(":8080")
}
