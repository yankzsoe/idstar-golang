package controllers

import (
	"net/http"
	"strconv"
	"time"

	"ginweb/app/helper"
	model "ginweb/app/models"

	"github.com/gin-gonic/gin"
)

// Show multiple template (index)
func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Index",
		"year":  time.Now().Year(),
	})
}

// Show multiple template (about)
func ShowAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about", gin.H{
		"title": "About",
		"year":  time.Now().Year(),
	})
}

// Show multiple template (contact)
func ShowContact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact", gin.H{
		"title": "Contact",
		"year":  time.Now().Year(),
	})
}

// Show Todo List Page
func GetTodoList(ctx *gin.Context) {
	err, todos := helper.GetList()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "errorpage", gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "get list from db",
			"message": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "todolist", gin.H{
		"title": "Todo List",
		"year":  time.Now().Year(),
		"todos": todos,
	})
}

// Show AddNew todo page
func GetCreateTodo(c *gin.Context) {
	c.HTML(http.StatusOK, "todocreate", gin.H{
		"title": "Create New Todo",
		"year":  time.Now().Year(),
	})
}

// Save Todo and back to todo list page
func PostCreateTodo(ctx *gin.Context) {
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	todo := model.Todo{
		Title:       title,
		Description: description,
	}

	err := helper.Save(todo)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "errorpage", gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "saving data on db",
			"message": err.Error(),
		})
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/todo")
}

// Show Todo Detail page
func GetDetailTodo(c *gin.Context) {
	id := c.Param("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "errorpage", gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "Parsing param Id",
			"message": err.Error(),
		})
		return
	}

	err, todo := helper.Get(val)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "errorpage", gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "Get Todo from db",
			"message": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "tododetail", gin.H{
		"title": "Detail - " + todo.Title,
		"year":  time.Now().Year(),
		"todo":  todo,
	})
}

// Show Update todo page
func GetUpdateTodo(c *gin.Context) {
	id := c.Param("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "errorpage", gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "Parsing param Id",
			"message": err.Error(),
		})
		return
	}

	err, todo := helper.Get(val)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "errorpage", gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "Get Todo from db",
			"message": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "todoupdate", gin.H{
		"title": "Update - " + todo.Title,
		"year":  time.Now().Year(),
		"todo":  todo,
	})
}

// Update todo and back to todo list page
func PostUpdateTodo(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	id := c.PostForm("todoId")

	val, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "errorpage", gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "Parsing param Id",
			"message": err.Error(),
		})
		return
	}

	todo := model.Todo{
		Title:       title,
		Description: description,
	}
	err = helper.Update(val, todo)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "errorpage", gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "Update data on db",
			"message": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/todo")
}

// Delete Todo from todo list page
func PostDeleteTodo(c *gin.Context) {
	id := c.PostForm("todoId")

	val, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "errorpage", gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "Parsing param Id",
			"message": err.Error(),
		})
		return
	}

	err = helper.Delete(val)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "errorpage", gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "Delete data on db",
			"message": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/todo")
}
