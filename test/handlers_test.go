package test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/SArtemJ/WTest/internal/customer"
	"github.com/SArtemJ/WTest/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//TODO run docker-compose up postgres-test before run tests
type RepoForTests struct{}

func (r *RepoForTests) CreateCustomer(c *customer.DbCustomer) (*customer.DbCustomer, error) {
	return &customer.DbCustomer{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Gender:    c.Gender,
		Email:     c.Email,
		Address:   c.Address,
		Birthdate: c.Birthdate,
	}, nil
}

func (r *RepoForTests) UpdateCustomer(c *customer.DbCustomer) error {
	return nil
}

func (r *RepoForTests) GetAllCustomers(limit, offset *int64) ([]*customer.DbCustomer, error) {
	return []*customer.DbCustomer{
		&customer.DbCustomer{
			FirstName: fmt.Sprintf("uniqueCustomerName%d", 1),
			LastName:  fmt.Sprintf("uniqueCustomerLast%d", 1),
			Gender:    "Male",
			Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 1, "@gmail.com"),
			Address:   fmt.Sprintf("uniqueCustomerAddress%d", 1),
			Birthdate: time.Date(1995, time.Month(10), 01, 0, 0, 0, 0, time.UTC),
		},
		&customer.DbCustomer{
			FirstName: fmt.Sprintf("uniqueCustomerName%d", 2),
			LastName:  fmt.Sprintf("uniqueCustomerLast%d", 2),
			Gender:    "Female",
			Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 2, "@gmail.com"),
			Address:   fmt.Sprintf("uniqueCustomerAddress%d", 2),
			Birthdate: time.Date(1996, time.Month(10), 02, 0, 0, 0, 0, time.UTC),
		},
		&customer.DbCustomer{
			FirstName: fmt.Sprintf("uniqueCustomerName%d", 3),
			LastName:  fmt.Sprintf("uniqueCustomerLast%d", 3),
			Gender:    "Male",
			Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 3, "@gmail.com"),
			Address:   fmt.Sprintf("uniqueCustomerAddress%d", 3),
			Birthdate: time.Date(1997, time.Month(10), 03, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}

func (r *RepoForTests) GetCustomerByID(id int64) (*customer.DbCustomer, error) {
	return &customer.DbCustomer{
		ID:        id,
		FirstName: fmt.Sprintf("uniqueCustomerName%d", 4),
		LastName:  fmt.Sprintf("uniqueCustomerLast%d", 4),
		Gender:    "Male",
		Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 4, "@gmail.com"),
		Address:   fmt.Sprintf("uniqueCustomerAddress%d", 4),
		Birthdate: time.Date(1998, time.Month(10), 04, 0, 0, 0, 0, time.UTC),
	}, nil
}

func (r *RepoForTests) SearchCustomers(firstName, lastName string, limit, offset *int64) ([]*customer.DbCustomer, error) {
	return []*customer.DbCustomer{
		&customer.DbCustomer{
			FirstName: fmt.Sprintf("uniqueCustomerName%d", 5),
			LastName:  fmt.Sprintf("uniqueCustomerLast%d", 5),
			Gender:    "Female",
			Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 5, "@gmail.com"),
			Address:   fmt.Sprintf("uniqueCustomerAddress%d", 5),
			Birthdate: time.Date(1999, time.Month(10), 05, 0, 0, 0, 0, time.UTC),
		},
		&customer.DbCustomer{
			FirstName: fmt.Sprintf("uniqueCustomerName%d", 6),
			LastName:  fmt.Sprintf("uniqueCustomerLast%d", 6),
			Gender:    "Male",
			Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 6, "@gmail.com"),
			Address:   fmt.Sprintf("uniqueCustomerAddress%d", 6),
			Birthdate: time.Date(2000, time.Month(10), 06, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}

func (r *RepoForTests) SearchCustomersByEmail(email string) (*customer.DbCustomer, error) {
	return &customer.DbCustomer{
		FirstName: fmt.Sprintf("uniqueCustomerName%d", 5),
		LastName:  fmt.Sprintf("uniqueCustomerLast%d", 5),
		Gender:    "Female",
		Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 5, "@gmail.com"),
		Address:   fmt.Sprintf("uniqueCustomerAddress%d", 5),
		Birthdate: time.Date(1999, time.Month(10), 05, 0, 0, 0, 0, time.UTC),
	}, nil
}

func (r *RepoForTests) DeleteCustomerByID(id int64) error {
	return nil
}

func NewRepoForTest() *RepoForTests {
	return &RepoForTests{}
}

func BuildSuite(t *testing.T, withTemplates bool) (*gin.Engine, *server.CustomerHandlers) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("genderCustom", customer.GenderValidation())
		v.RegisterValidation("birthdateCustom", customer.BirthDateValidation())
	}
	r.LoadHTMLGlob(".././web/templates/*")

	repo := NewRepoForTest()
	ch := server.NewCustomerHandlers(repo, logrus.NewEntry(logrus.StandardLogger()))
	return r, ch
}

/*
	router.GET("/customers/view/:id", ch.GetOneCustomer)

	router.POST("/customers/create", ch.CreateCustomer)
	router.POST("/customers/search", ch.GetAllCustomers)

	router.PUT("/customers/update", ch.UpdateCustomer)
	router.DELETE("/customers", ch.DeleteCustomer)

router.GET("/customers/create", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"view.html",
			gin.H{
				"createOnly": true,
				"title":      "Create new customer",
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

*/

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
	if err != nil || strings.Index(string(p), "<FisrtName> ") < 0 {
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
