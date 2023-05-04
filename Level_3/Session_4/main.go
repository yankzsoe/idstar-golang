package main

import (
	"ginweb/routers"

	"github.com/gin-contrib/multitemplate"
)

func main() {
	// init router
	router := routers.SetupRouter()

	// Set static file for load all assets
	router.Static("/assets", "./assets")

	// Load templates for page with partial html
	router.HTMLRender = renderTemplates()

	// Start App on port 5001
	router.Run(":5001")

}

func renderTemplates() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("index", "app/templates/layout.tmpl", "app/views/index.html")
	render.AddFromFiles("about", "app/templates/layout.tmpl", "app/views/about.html")
	render.AddFromFiles("contact", "app/templates/layout.tmpl", "app/views/contact.html")
	return render
}
