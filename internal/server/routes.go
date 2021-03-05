package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitServiceRoutes(router *gin.Engine, ch *CustomerHandlers) {
	router.GET("/customers/create", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"view.html",
			gin.H{
				"title":        "Create new customer",
				"isCreateOnly": true,
				"isInternal":   false,
			},
		)
	})
	router.GET("/customers/search", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"search.html",
			gin.H{
				"title": "Search customers",
			},
		)
	})
	router.GET("/", ch.GetAllCustomers)
	router.GET("/customers/view/:id", ch.GetOneCustomer)

	router.POST("/customers/create", ch.CreateCustomer)
	router.POST("/customers/search", ch.GetAllCustomers)

	router.POST("/customers/update", ch.UpdateCustomer)
	router.DELETE("/customers", ch.DeleteCustomer)
}
