package server

import (
	"fmt"
	"github.com/SArtemJ/WTest/internal/consts"
	"github.com/SArtemJ/WTest/internal/customer"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CustomerHandlers struct {
	customerRepo customer.Repository
	logger       logrus.FieldLogger
}

func NewCustomerHandlers(cr customer.Repository, logger logrus.FieldLogger) *CustomerHandlers {
	return &CustomerHandlers{
		customerRepo: cr,
		logger:       logger,
	}
}

func (ch *CustomerHandlers) CreateCustomer(c *gin.Context) {
	webModel := ch.ValidateIncomingModel(c)
	if webModel == nil {
		return
	}

	in, err := webModel.ToDb()
	if err != nil {
		LoggingActionError(ch.logger, "web model to db", err)
		InternalError(c, consts.ActionCreate, webModel)
		return
	}

	exist, err := ch.customerRepo.SearchCustomersByEmail(webModel.Email)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionSearch, err)
		InternalError(c, consts.ActionCreate, webModel)
		return
	}

	if exist != nil {
		LoggingActionMessage(ch.logger, fmt.Sprintf("customer with email %v already exists", in.Email))
		CreateAlreadyExists(c, webModel)
		return
	}

	created, err := ch.customerRepo.CreateCustomer(in)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionCreate, err)
		InternalError(c, consts.ActionCreate, webModel)
		return
	}

	Success(c, http.StatusCreated, created.ToWeb(), consts.ActionCreate)
	return
}

func (ch *CustomerHandlers) UpdateCustomer(c *gin.Context) {
	webModel := ch.ValidateIncomingModel(c)
	if webModel == nil {
		return
	}

	in, err := webModel.ToDb()
	if err != nil {
		LoggingActionError(ch.logger, "web model to db", err)
		InternalError(c, consts.ActionCreate, webModel)
		return
	}

	exist, err := ch.customerRepo.GetCustomerByID(webModel.ID)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionGetByID, err)
		InternalError(c, consts.ActionUpdate, webModel)
		return
	}

	if exist == nil {
		LoggingActionMessage(ch.logger, "no customers for update")
		NotFound(c, consts.ActionUpdate, webModel)
		return
	}

	err = ch.customerRepo.UpdateCustomer(in)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionUpdate, err)
		InternalError(c, consts.ActionUpdate, webModel)
		return
	}

	Success(c, http.StatusOK, webModel, consts.ActionUpdate)
	return
}

func (ch *CustomerHandlers) DeleteCustomer(c *gin.Context) {
	id, err := ch.ReadHiddenID(c)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionDelete, err)
		InternalError(c, consts.ActionDelete, nil)
		return
	}

	exist, err := ch.customerRepo.GetCustomerByID(*id)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionGetByID, err)
		InternalError(c, consts.ActionDelete, nil)
		return
	}

	if exist == nil {
		LoggingActionMessage(ch.logger, "no customers for delete")
		NotFound(c, consts.ActionUpdate, nil)
		return
	}

	err = ch.customerRepo.DeleteCustomerByID(*id)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionDelete, err)
		InternalError(c, consts.ActionDelete, nil)
		return
	}

	Success(c, http.StatusOK, nil, consts.ActionDelete)
	return
}

func (ch *CustomerHandlers) GetAllCustomers(c *gin.Context) {
	all, err := ch.customerRepo.GetAllCustomers(nil, nil)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionGetAll, err)
		InternalError(c, consts.ActionGetAll, nil)
		return
	}

	var webView []*customer.WebCustomer
	for _, item := range all {
		webView = append(webView, item.ToWeb())
	}
	SuccessList(c, http.StatusOK, webView)
	return
}

func (ch *CustomerHandlers) GetOneCustomer(c *gin.Context) {
	id, err := ch.ReadHiddenID(c)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionGetByID, err)
		InternalError(c, consts.ActionGetByID, nil)
		return
	}

	dbModel, err := ch.customerRepo.GetCustomerByID(*id)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionGetByID, err)
		InternalError(c, consts.ActionGetByID, nil)
		return
	}

	Success(c, http.StatusOK, dbModel.ToWeb(), consts.ActionGetByID)
	return
}

func (ch *CustomerHandlers) SearchCustomers(c *gin.Context) {
	l := c.Query("limit")
	if l == "" {
		LoggingActionMessage(ch.logger, "wrong limit value")
		InternalError(c, consts.ActionSearch, nil)
		return
	}

	limitValue, err := strconv.ParseInt(l, 0, 64)
	if err != nil {
		LoggingActionMessage(ch.logger, "wrong limit value")
		InternalError(c, consts.ActionSearch, nil)
		return
	}

	o := c.Query("offset")
	if o == "" {
		LoggingActionMessage(ch.logger, "wrong offset value")
		InternalError(c, consts.ActionSearch, nil)
		return
	}

	offsetValue, err := strconv.ParseInt(o, 0, 64)
	if err != nil {
		LoggingActionMessage(ch.logger, "wrong offset value")
		InternalError(c, consts.ActionSearch, nil)
		return
	}

	var in customer.SearchCustomer
	if err := c.ShouldBind(&in); err == nil {
		LoggingActionMessage(ch.logger, "customer search model wrong value")
		InternalError(c, consts.ActionSearch, nil)
		return
	}

	if in.FirstName == "" || in.LastName == "" {
		LoggingActionMessage(ch.logger, "customers operation update failed - not enough data")
		InternalError(c, consts.ActionSearch, nil)
		return
	}

	list, err := ch.customerRepo.SearchCustomers(in.FirstName, in.LastName, &limitValue, &offsetValue)
	if err != nil {
		LoggingRepoError(ch.logger, consts.ActionGetByID, err)
		InternalError(c, consts.ActionGetByID, nil)
		return
	}

	var webView []*customer.WebCustomer
	for _, item := range list {
		webView = append(webView, item.ToWeb())
	}

	SuccessList(c, http.StatusOK, webView)
	return
}

func (ch *CustomerHandlers) ValidateIncomingModel(c *gin.Context) *customer.WebCustomer {
	var webModel customer.WebCustomer
	if err := c.ShouldBind(&webModel); err != nil {
		LoggingActionError(ch.logger, "validation", err)
		CreateValidationError(c, &webModel, err)
		return nil
	}
	return &webModel
}

func (ch *CustomerHandlers) ReadHiddenID(c *gin.Context) (*int64, error) {
	id := c.Param("id")
	if id == "" {
		return nil, fmt.Errorf("wrong customer id param")
	}

	idValue, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		return nil, err
	}
	return &idValue, nil
}
