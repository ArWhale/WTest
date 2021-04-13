package adapters

import (
	"database/sql"
	"github.com/ArWhale/WTest/internal/customer"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type CustomerRepository struct {
	db     *sql.DB
	logger logrus.FieldLogger
}

func NewCustomerRepository(db *sql.DB, logger logrus.FieldLogger) *CustomerRepository {
	return &CustomerRepository{
		db:     db,
		logger: logger,
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
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error CreateCustomer - tx.Begin")
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	err = cr.db.QueryRow(sqlstr, c.FirstName, c.LastName, c.Birthdate, c.Gender, c.Email, c.Address).Scan(&c.ID)
	if err != nil {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error CreateCustomer - queryRow")
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error CreateCustomer - tx.Commit")
		return nil, err
	}

	return c, nil
}

func (cr *CustomerRepository) UpdateCustomer(c *customer.Customer) error {
	var err error

	const sqlstr = `UPDATE public.customers SET
					first_name = $1, last_name = $2, birthdate = $3, gender = $4, e_mail = $5, address = $6
					WHERE id = $7`

	tx, err := cr.db.Begin()
	if err != nil {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error UpdateCustomer - tx.Begin")
		return err
	}
	defer func() { _ = tx.Rollback() }()

	_, err = cr.db.Exec(sqlstr, c.FirstName, c.LastName, c.Birthdate, c.Gender, c.Email, c.Address, c.ID)
	if err != nil {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error UpdateCustomer - exec")
		return err
	}

	err = tx.Commit()
	if err != nil {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error UpdateCustomer - tx.Commit")
		return err
	}

	return nil
}

func (cr *CustomerRepository) GetAllCustomers(limit, offset *int64) ([]*customer.Customer, error) {
	var err error

	const sqlstr = `SELECT
		id, first_name, last_name, birthdate, gender, e_mail, address
		FROM public.customers
		LIMIT $1 :: NUMERIC
		OFFSET $2 :: NUMERIC`

	q, err := cr.db.Query(sqlstr, limit, offset)
	if err != nil {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error query GetAllCustomers - query")
		return nil, err
	}
	defer q.Close()

	var res []*customer.Customer
	for q.Next() {
		c := customer.Customer{}
		err = q.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Birthdate, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			cr.logger.WithFields(logrus.Fields{
				"resource": "db",
				"err":      err,
			}).Error("Error query GetAllCustomers - scan")
			return nil, err
		}

		res = append(res, &c)
	}

	if err := q.Err(); err != nil {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error query GetAllCustomers - query row")
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
	err = cr.db.QueryRow(sqlstr, id).Scan(&c.ID, &c.FirstName, &c.LastName, &c.Birthdate, &c.Gender, &c.Email, &c.Address)
	if err != nil && err == sql.ErrNoRows {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error query GetCustomerByID - query row")
		return nil, err
	}

	return &c, nil
}

func (cr *CustomerRepository) SearchCustomers(firstName, lastName *string, limit, offset *int64) ([]*customer.Customer, error) {
	var err error

	var sqlstr = `SELECT id, first_name, last_name, birthdate, gender, e_mail, address 
				FROM public.customers 
				WHERE ( $1 :: TEXT IS NULL or first_name like  $1 :: TEXT) 
				AND ($2 :: TEXT IS NULL or last_name like  $2 :: TEXT) 
				LIMIT $3 :: NUMERIC 
				OFFSET $4 :: NUMERIC`

	q, err := cr.db.Query(sqlstr, firstName, lastName, limit, offset)
	if err != nil {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error SearchCustomers - query")
		return nil, err
	}
	defer q.Close()

	var res []*customer.Customer
	for q.Next() {
		c := customer.Customer{}

		err = q.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Birthdate, &c.Gender, &c.Email, &c.Address)
		if err != nil {
			cr.logger.WithFields(logrus.Fields{
				"resource": "db",
				"err":      err,
			}).Error("Error SearchCustomers - scan ")
			return nil, err
		}

		res = append(res, &c)
	}

	if err := q.Err(); err != nil {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error SearchCustomers - query row")
	}

	return res, nil
}

func (cr *CustomerRepository) SearchCustomersByEmail(email string) (*customer.Customer, error) {
	var err error

	const sqlstr = `SELECT
		id, first_name, last_name, birthdate, gender, e_mail, address
		FROM public.customers
		WHERE e_mail = $1`

	c := customer.Customer{}
	err = cr.db.QueryRow(sqlstr, email).Scan(&c.ID, &c.FirstName, &c.LastName, &c.Birthdate, &c.Gender, &c.Email, &c.Address)
	if err != nil && err == sql.ErrNoRows {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error SearchCustomersByEmail - query")
		return nil, nil
	}

	return &c, nil
}

func (cr *CustomerRepository) DeleteCustomerByID(id int64) error {
	var err error

	const sqlstr = `DELETE FROM public.customers WHERE id = $1`
	_, err = cr.db.Exec(sqlstr, id)
	if err != nil {
		cr.logger.WithFields(logrus.Fields{
			"resource": "db",
			"err":      err,
		}).Error("Error DeleteCustomerByID - exec")
		return err
	}

	return nil
}
