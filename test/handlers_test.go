package test

import (
	"encoding/json"
	"fmt"
	"github.com/ArWhale/WTest/internal/customer"
	"github.com/ArWhale/WTest/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func BuildSuite(t *testing.T, withTemplates bool) (*gin.Engine, *server.CustomerHandlers) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("genderCustom", customer.GenderValidation())
		v.RegisterValidation("birthdateCustom", customer.BirthDateValidation())
	}
	r.LoadHTMLGlob(".././web/templates/*")

	//only success repo actions
	create := func(c *customer.Customer) (*customer.Customer, error) {
		return c, nil
	}

	update := func(*customer.Customer) error {
		return nil
	}

	getAll := func(*int64, *int64) ([]*customer.Customer, error) {
		return nil, nil
	}

	getByID := func(int64) (*customer.Customer, error) {
		return nil, nil
	}

	search := func(*string, *string, *int64, *int64) ([]*customer.Customer, error) {
		return nil, nil
	}

	del := func(id int64) error {
		return nil
	}

	repo := NewRepoMock(create, update, getAll, getByID, search, del)
	ch := server.NewCustomerHandlers(repo, logrus.NewEntry(logrus.StandardLogger()))
	return r, ch
}

func TestCreateCustomer(t *testing.T) {
	w := httptest.NewRecorder()

	r, h := BuildSuite(t, true)
	r.POST("/customers/create", h.CreateCustomer)

	payload := createBodyHappy(t)
	req, _ := http.NewRequest("POST", "/customers/create", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<FirstName> ") < 0 {
		t.Fail()
	}
}

func createBodyHappy(t *testing.T) string {
	example := &customer.WebCustomer{
		FirstName: fmt.Sprintf("uniqueCustomerName%d", 10001),
		LastName:  fmt.Sprintf("uniqueCustomerLast%d", 10001),
		Gender:    "male",
		Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 10001, "@gmail.com"),
		Address:   fmt.Sprintf("uniqueCustomerAddress%d", 10001),
		Birthdate: "1998-10-01",
	}

	b, err := json.Marshal(example)
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}
