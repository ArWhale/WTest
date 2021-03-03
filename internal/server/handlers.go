package server

import (
	"github.com/SArtemJ/WTest/internal/config"
	"github.com/SArtemJ/WTest/internal/customer"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CustomerHandlers struct {
	customerRepo customer.Repository
}

func NewCustomerHandlers(cr customer.Repository) *CustomerHandlers {
	return &CustomerHandlers{
		customerRepo: cr,
	}
}

func (ch *CustomerHandlers) GetAllCustomers(c *gin.Context) {
	all, err := ch.customerRepo.GetAllCustomers()
	if err != nil {
		//todo
	}

	c.JSON(http.StatusCreated, all)
}

func (ch *CustomerHandlers) CreateCustomer(c *gin.Context) {
	var newCustomer customer.Customer
	if err := c.ShouldBind(&newCustomer); err == nil {
		//todo
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	parsedBD, err := time.Parse(config.DefaultDateLayout, newCustomer.BirthDate.String())
	if err != nil {
		//todo
	}
	newCustomer.BirthDate = parsedBD

	//TODO check exists customer

	_, err = ch.customerRepo.CreateCustomer(&newCustomer)
	if err != nil {
		//TODO
	}

	c.JSON(http.StatusCreated, newCustomer)
}
