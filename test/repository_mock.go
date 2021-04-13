package test

import (
	"fmt"
	"github.com/ArWhale/WTest/internal/customer"
	"time"
)

type RepoMock struct {
	create  func(*customer.Customer) (*customer.Customer, error)
	update  func(*customer.Customer) error
	getAll  func(*int64, *int64) ([]*customer.Customer, error)
	getByID func(int64) (*customer.Customer, error)
	search  func(*string, *string, *int64, *int64) ([]*customer.Customer, error)
	delete  func(id int64) error
}

func NewRepoMock(
	create func(*customer.Customer) (*customer.Customer, error),
	update func(*customer.Customer) error,
	getAll func(*int64, *int64) ([]*customer.Customer, error),
	getByID func(int64) (*customer.Customer, error),
	search func(*string, *string, *int64, *int64) ([]*customer.Customer, error),
	delete func(id int64) error,
) *RepoMock {
	return &RepoMock{
		create:  create,
		update:  update,
		getAll:  getAll,
		getByID: getByID,
		search:  search,
		delete:  delete,
	}
}

func (r *RepoMock) CreateCustomer(c *customer.Customer) (*customer.Customer, error) {
	return r.create(c)
}

func (r *RepoMock) UpdateCustomer(c *customer.Customer) error {
	return r.update(c)
}

func (r *RepoMock) GetAllCustomers(limit, offset *int64) ([]*customer.Customer, error) {
	return r.getAll(limit, offset)
}

func (r *RepoMock) GetCustomerByID(id int64) (*customer.Customer, error) {
	return r.getByID(id)
}

func (r *RepoMock) SearchCustomers(firstName, lastName *string, limit, offset *int64) ([]*customer.Customer, error) {
	return r.search(firstName, lastName, limit, offset)
}

func (r *RepoMock) SearchCustomersByEmail(email string) (*customer.Customer, error) {
	return &customer.Customer{
		FirstName: fmt.Sprintf("uniqueCustomerName%d", 5),
		LastName:  fmt.Sprintf("uniqueCustomerLast%d", 5),
		Gender:    "Female",
		Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 5, "@gmail.com"),
		Address:   fmt.Sprintf("uniqueCustomerAddress%d", 5),
		Birthdate: time.Date(1999, time.Month(10), 05, 0, 0, 0, 0, time.UTC),
	}, nil
}

func (r *RepoMock) DeleteCustomerByID(id int64) error {
	return r.delete(id)
}
