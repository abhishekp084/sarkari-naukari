// main.go
package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router instance
	r := gin.Default()

	// Load and parse the template
	templ := template.Must(template.ParseFiles("template.html"))
	r.SetHTMLTemplate(templ)

	// Define a route for the home page
	r.GET("/", serveHome)

	// Start the server on port 8000
	r.Run(":8000")
}

func serveHome(c *gin.Context) {
	// Load data (if any)
	data := gin.H{
		// Pass any data to the template here
		"title": "Home Page",
	}

	// Render the template with data
	c.HTML(http.StatusOK, "template.html", data)
}
