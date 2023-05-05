package routers

import (
	"ginweb/app/controllers"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routing
	r.GET("/", controllers.ShowIndex)
	r.GET("/about", controllers.ShowAbout)
	r.GET("/contact", controllers.ShowContact)

	// Routing Todo
	r.GET("/todo", controllers.GetTodoList)
	r.GET("/todo/create", controllers.GetCreateTodo)
	r.POST("/todo/create", controllers.PostCreateTodo)
	r.GET("/todo/:id", controllers.GetDetailTodo)
	r.GET("/todo/:id/edit", controllers.GetUpdateTodo)
	r.POST("/todo/update", controllers.PostUpdateTodo)
	r.POST("/todo/delete", controllers.PostDeleteTodo)

	// Load templates for page with partial html
	r.HTMLRender = renderTemplates()

	return r
}

func renderTemplates() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("index", "app/templates/layout.tmpl", "app/views/index.html")
	render.AddFromFiles("about", "app/templates/layout.tmpl", "app/views/about.html")
	render.AddFromFiles("contact", "app/templates/layout.tmpl", "app/views/contact.html")
	render.AddFromFiles("todocreate", "app/templates/layout.tmpl", "app/views/todo/todoCreate.html")
	render.AddFromFiles("todolist", "app/templates/layout.tmpl", "app/views/todo/todoList.html")
	render.AddFromFiles("todoupdate", "app/templates/layout.tmpl", "app/views/todo/todoUpdate.html")
	render.AddFromFiles("tododetail", "app/templates/layout.tmpl", "app/views/todo/todoDetail.html")
	render.AddFromFiles("errorpage", "app/templates/layout.tmpl", "app/views/error.html")
	return render
}
