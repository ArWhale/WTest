package server

import (
	"fmt"
	"github.com/ArWhale/WTest/internal/consts"
	"github.com/ArWhale/WTest/internal/customer"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func CreateValidationError(c *gin.Context, data *customer.WebCustomer, err error) {
	var re = regexp.MustCompile(`Field validation for '(\w+)'`)
	var detail string
	m := re.FindStringSubmatch(err.Error())
	if m != nil {
		detail = m[0]
	}

	c.HTML(http.StatusBadRequest, "view.html", gin.H{
		"isCreateOnly": true,
		"isInternal":   false,
		"payload":      data,
		"validation":   detail,
	})
}

func InternalError(c *gin.Context, action string, data *customer.WebCustomer) {
	switch action {
	case consts.ActionCreate:
		c.HTML(http.StatusInternalServerError, "view.html", gin.H{
			"isCreateOnly": true,
			"isInternal":   true,
			"message":      "failed to create customer",
		})
	case consts.ActionUpdate:
		c.HTML(http.StatusInternalServerError, "view.html", gin.H{
			"isCreateOnly": false,
			"isInternal":   true,
			"message":      "failed to update customer",
			"payload":      data,
		})
	}
}

func CreateAlreadyExists(c *gin.Context, data *customer.WebCustomer) {
	c.HTML(http.StatusInternalServerError, "view.html", gin.H{
		"isCreateOnly": true,
		"isInternal":   true,
		"message":      "customer already exists",
		"payload":      data,
	})
}

func Success(c *gin.Context, code int, data *customer.WebCustomer, action string) {
	c.HTML(code, "view.html", gin.H{
		"isCreateOnly": false,
		"success":      fmt.Sprintf("customer was %sd", action),
		"payload":      data,
	})
}

func SuccessList(c *gin.Context, code int, data []*customer.WebCustomer) {
	c.HTML(code, "index.html", gin.H{
		"isCreateOnly": false,
		"success":      "founded customers",
		"payload":      data,
	})
}

func NotFound(c *gin.Context, action string, data *customer.WebCustomer) {
	c.HTML(http.StatusNotFound, "view.html", gin.H{
		"isCreateOnly": false,
		"isInternal":   true,
		"message":      fmt.Sprintf("failed to %s customer", action),
		"payload":      data,
	})
}
