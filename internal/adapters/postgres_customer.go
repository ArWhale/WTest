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

func (cr *CustomerRepository) CreateCustomer(c *customer.DbCustomer) (*customer.DbCustomer, error) {
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

	err = cr.db.QueryRow(sqlstr, c.FirstName, c.LastName, c.Birthdate, c.Gender, c.Email, c.Address).Scan(&c.ID)
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

func (cr *CustomerRepository) UpdateCustomer(c *customer.DbCustomer) error {
	var err error

	const sqlstr = `UPDATE public.customers SET
					first_name = $1, last_name = $2, birthdate = $3, gender = $4, e_mail = $5, address = $6
					WHERE id = $7`

	tx, err := cr.db.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Error update customer  - ")
	}
	defer func() { _ = tx.Rollback() }()

	_, err = cr.db.Exec(sqlstr, c.FirstName, c.LastName, c.Birthdate, c.Gender, c.Email, c.Address, c.ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"resource": "db",
			"info":     "commit transaction failed",
			"err":      err,
		}).Error("Error update customer- ")
	}

	return nil
}

func (cr *CustomerRepository) GetAllCustomers(limit, offset *int64) ([]*customer.DbCustomer, error) {
	var err error

	const sqlstr = `SELECT
		id, first_name, last_name, birthdate, gender, e_mail, address
		FROM public.customers
		LIMIT $1 :: NUMERIC
		OFFSET $2 :: NUMERIC`

	q, err := cr.db.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	var res []*customer.DbCustomer
	for q.Next() {
		c := customer.DbCustomer{}
		err = q.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Birthdate, &c.Gender, &c.Email, &c.Address)
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

func (cr *CustomerRepository) GetCustomerByID(id int64) (*customer.DbCustomer, error) {
	var err error

	const sqlstr = `SELECT
		id, first_name, last_name, birthdate, gender, e_mail, address
		FROM public.customers
		WHERE id = $1`

	c := customer.DbCustomer{}
	err = cr.db.QueryRow(sqlstr, id).Scan(&c.ID, &c.FirstName, &c.LastName, &c.Birthdate, &c.Gender, &c.Email, &c.Address)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (cr *CustomerRepository) SearchCustomers(firstName, lastName *string, limit, offset *int64) ([]*customer.DbCustomer, error) {
	var err error

	var sqlstr = `SELECT id, first_name, last_name, birthdate, gender, e_mail, address 
				FROM public.customers 
				WHERE ( $1 :: TEXT IS NULL or first_name like  $1 :: TEXT) 
				AND ($2 :: TEXT IS NULL or last_name like  $2 :: TEXT) 
				LIMIT $3 :: NUMERIC 
				OFFSET $4 :: NUMERIC`

	q, err := cr.db.Query(sqlstr, firstName, lastName, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	var res []*customer.DbCustomer
	for q.Next() {
		c := customer.DbCustomer{}

		err = q.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Birthdate, &c.Gender, &c.Email, &c.Address)
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

func (cr *CustomerRepository) SearchCustomersByEmail(email string) (*customer.DbCustomer, error) {
	var err error

	const sqlstr = `SELECT
		id, first_name, last_name, birthdate, gender, e_mail, address
		FROM public.customers
		WHERE e_mail = $1`

	c := customer.DbCustomer{}
	err = cr.db.QueryRow(sqlstr, email).Scan(&c.ID, &c.FirstName, &c.LastName, &c.Birthdate, &c.Gender, &c.Email, &c.Address)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (cr *CustomerRepository) DeleteCustomerByID(id int64) error {
	var err error

	const sqlstr = `DELETE FROM public.customers WHERE id = $1`
	_, err = cr.db.Exec(sqlstr, id)
	if err != nil {
		return err
	}

	return nil
}
