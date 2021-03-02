package adapters

import (
	"database/sql"
	"github.com/SArtemJ/WTest/internal/customer"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (cr *CustomerRepository) CreateCustomer(c *customer.Customer) (*customer.Customer, error) {
	var err error

	const sqlstr = `INSERT INTO public.customers (
		first_name, last_name, birthdate, gender, e_mail, address
		) VALUES (
		$1, $2, $3, $4, $5, $6
		) RETURNING id`

	tx, err := cr.db.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Error create customer  - ")
	}

	defer func() { _ = tx.Rollback() }()

	err = cr.db.QueryRow(sqlstr, c.FirstName, c.LastName, c.BirthDate, c.Gender, c.Email, c.Address).Scan(&c.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"resource": "db",
			"info":     "commit transaction failed",
			"err":      err,
		}).Error("Error create customer- ")
	}

	return c, nil
}

func (cr *CustomerRepository) UpdateCustomer() error {
	return nil
}

func (cr *CustomerRepository) GetAllCustomers() ([]*customer.Customer, error) {
	var err error

	const sqlstr = `SELECT
		id, first_name, last_name, birthdate, gender, e_mail, address
		FROM public.customers`

	q, err := cr.db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	res := []*customer.Customer{}
	for q.Next() {
		c := customer.Customer{}
		err = q.Scan(&c.ID, &c.FirstName, &c.LastName, &c.BirthDate, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	if err := q.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (cr *CustomerRepository) GetCustomerByID(id int64) (*customer.Customer, error) {
	var err error

	const sqlstr = `SELECT
		id, first_name, last_name, birthdate, gender, e_mail, address
		FROM public.customers
		WHERE id = $1`

	c := customer.Customer{}

	err = cr.db.QueryRow(sqlstr, id).Scan(&c.ID, &c.FirstName, &c.LastName, &c.BirthDate, &c.Gender, &c.Address)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (cr *CustomerRepository) SearchCustomers(limit, offset int64) ([]*customer.Customer, error) {
	return nil, nil
}
