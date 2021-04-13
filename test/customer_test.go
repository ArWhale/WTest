package test

import (
	"fmt"
	"github.com/ArWhale/WTest/internal/consts"
	"github.com/ArWhale/WTest/internal/customer"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateCustomerHappy(t *testing.T) {

	create := func(c *customer.Customer) (*customer.Customer, error) {
		return c, nil
	}

	c := &customer.Customer{
		FirstName: fmt.Sprintf("uniqueCustomerName%d", 1000),
		LastName:  fmt.Sprintf("uniqueCustomerLast%d", 1000),
		Gender:    "Male",
		Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 1000, "@gmail.com"),
		Address:   fmt.Sprintf("uniqueCustomerAddress%d", 1000),
		Birthdate: time.Date(1989, time.Month(10), 01, 0, 0, 0, 0, time.UTC),
	}

	getById := func(id int64) (*customer.Customer, error) {
		c.ID = id
		return c, nil
	}

	repo := NewRepoMock(create, nil, nil, getById, nil, nil)

	created, err := repo.CreateCustomer(c)
	assert.Nil(t, err)
	assert.NotNil(t, created.ID)
	assert.NotEqual(t, "", created.ID)

	founded, err := repo.GetCustomerByID(created.ID)
	assert.Nil(t, err)
	assert.NotNil(t, founded)
	assert.Equal(t, founded.Birthdate.Format(consts.DefaultDateLayout), created.Birthdate.Format(consts.DefaultDateLayout))
}

/*
func TestCreateCustomerError(t *testing.T) {
	conn, err := InitConnDbForTest()
	if err != nil {
		t.Fatalf("db connection error %#v", err)
	}

	c := &customer.Customer{
		FirstName: fmt.Sprintf("uniqueCustomerName%d", 1001),
		LastName:  fmt.Sprintf("uniqueCustomerLast%d", 1001),
		Gender:    "Male",
		Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 1001, "@gmail.com"),
		Address:   fmt.Sprintf("uniqueCustomerAddress%d", 1001),
		Birthdate: time.Date(1990, time.Month(10), 01, 0, 0, 0, 0, time.UTC),
	}

	repo := adapters.NewCustomerRepository(conn)

	created, err := repo.CreateCustomer(c)
	assert.Nil(t, err)
	assert.NotNil(t, created.ID)
	assert.NotEqual(t, "", created.ID)

	created, err = repo.CreateCustomer(c)
	assert.NotNil(t, err)
}

func TestGetAllCustomerError(t *testing.T) {
	conn, err := InitConnDbForTest()
	if err != nil {
		t.Fatalf("db connection error %#v", err)
	}
	repo := adapters.NewCustomerRepository(conn)

	all, err := repo.GetAllCustomers(nil, nil)
	assert.Nil(t, err)
	assert.NotEqual(t, int(0), len(all))
}

func TestUpdateCustomer(t *testing.T) {
	conn, err := InitConnDbForTest()
	if err != nil {
		t.Fatalf("db connection error %#v", err)
	}
	repo := adapters.NewCustomerRepository(conn)

	c, err := repo.GetCustomerByID(1)
	assert.Nil(t, err)

	c.FirstName = "random_name"
	c.LastName = "random_last"
	var setedG string
	switch strings.ToLower(c.Gender) {
	case "male":
		c.Gender = "Female"
		setedG = "Female"
	case "female":
		c.Gender = "Male"
		setedG = "Male"
	}
	c.Email = fmt.Sprintf("uniqueCustomerEmail%d", 1002)
	c.Birthdate = time.Date(1991, time.Month(10), 01, 0, 0, 0, 0, time.UTC)

	err = repo.UpdateCustomer(c)
	assert.Nil(t, err)

	founded, err := repo.GetCustomerByID(c.ID)
	assert.Nil(t, err)
	assert.NotNil(t, founded)
	assert.Equal(t, founded.FirstName, "random_name")
	assert.Equal(t, founded.LastName, "random_last")
	assert.Equal(t, founded.Gender, setedG)
	assert.Equal(t, founded.Email, fmt.Sprintf("uniqueCustomerEmail%d", 1002))
	assert.Equal(t, founded.Email, fmt.Sprintf("uniqueCustomerEmail%d", 1002))
	assert.Equal(t, founded.Birthdate.Format(consts.DefaultDateLayout), c.Birthdate.Format(consts.DefaultDateLayout))
}

func TestSearchCustomers(t *testing.T) {
	conn, err := InitConnDbForTest()
	if err != nil {
		t.Fatalf("db connection error %#v", err)
	}
	repo := adapters.NewCustomerRepository(conn)

	c, err := repo.GetCustomerByID(1)
	assert.Nil(t, err)

	founded, err := repo.SearchCustomers(c.FirstName, c.LastName, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(founded))
	assert.NotNil(t, founded)
	assert.Equal(t, founded[0].FirstName, c.FirstName)
	assert.Equal(t, founded[0].LastName, c.LastName)
}

func TestSearchCustomersByEmail(t *testing.T) {
	conn, err := InitConnDbForTest()
	if err != nil {
		t.Fatalf("db connection error %#v", err)
	}
	repo := adapters.NewCustomerRepository(conn)

	c := &customer.Customer{
		FirstName: fmt.Sprintf("uniqueCustomerName%d", 1003),
		LastName:  fmt.Sprintf("uniqueCustomerLast%d", 1003),
		Gender:    "Male",
		Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 1003, "@gmail.com"),
		Address:   fmt.Sprintf("uniqueCustomerAddress%d", 1003),
		Birthdate: time.Date(1990, time.Month(10), 01, 0, 0, 0, 0, time.UTC),
	}
	created, err := repo.CreateCustomer(c)
	assert.Nil(t, err)

	founded, err := repo.SearchCustomersByEmail(created.Email)
	assert.Nil(t, err)
	assert.NotNil(t, founded)
	assert.NotNil(t, founded)
	assert.Equal(t, founded.FirstName, c.FirstName)
	assert.Equal(t, founded.LastName, c.LastName)
	assert.Equal(t, founded.Email, c.Email)
}

func TestDeleteCustomer(t *testing.T) {
	conn, err := InitConnDbForTest()
	if err != nil {
		t.Fatalf("db connection error %#v", err)
	}
	repo := adapters.NewCustomerRepository(conn)

	c, err := repo.GetCustomerByID(1)
	assert.Nil(t, err)

	err = repo.DeleteCustomerByID(c.ID)
	assert.Nil(t, err)

	founded, err := repo.GetCustomerByID(c.ID)
	assert.Nil(t, err)
	assert.Nil(t, founded)
}
*/
