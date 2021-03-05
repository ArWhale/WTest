package test

import (
	"database/sql"
	"fmt"
	"github.com/SArtemJ/WTest/internal/adapters"
	"github.com/SArtemJ/WTest/internal/consts"
	"github.com/SArtemJ/WTest/internal/customer"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

//TODO run docker-compose up postgres-test before run tests
func InitConnDbForTest() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:postgres@postgres:15432/customers_test?sslmode=disable&log_statement=all")
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection, Err: %+v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database, Err: %+v", err)
	}

	_, err = db.Exec(`DROP TABLE IF EXISTS customers;`)
	if err != nil {
		return nil, fmt.Errorf("failed to drop test database, Err: %+v", err)
	}

	_, err = db.Exec(sqlFake)
	if err != nil {
		return nil, fmt.Errorf("failed to inset test data, Err: %+v", err)
	}

	return db, nil
}

const sqlFake = `
	CREATE TABLE IF NOT EXISTS customers(
   id SERIAL PRIMARY KEY,
   first_name text NOT NULL,
   last_name text NOT NULL,
   birthdate date NOT NULL,
   gender text NOT NULL,
   e_mail text NOT NULL,
   address text NOT NULL,
   UNIQUE(e_mail)
);

INSERT INTO customers(first_name, last_name, e_mail, gender, birthdate, address) VALUES
('Dalila', 'Coddrington', 'dcoddrington0@reddit.com', 'Male', '1997-04-23', '9332 Menomonie Plaza'),
('Cordie', 'Sangar', 'csangar1@booking.com', 'Male', '2020-10-18', '8700 International Circle'),
('Janeta', 'Sweynson', 'jsweynson2@myspace.com', 'Male', '2008-03-10', '968 Ilene Road'),
('Mercedes', 'Tummasutti', 'mtummasutti3@hexun.com', 'Female', '2016-05-25', '0 Novick Crossing'),
('Bradney', 'Fist', 'bfist4@usa.gov', 'Male', '2012-09-11', '135 Vahlen Plaza'),
('Arlen', 'Stryde', 'astryde5@usgs.gov', 'Female', '2011-03-14', '59 Granby Lane'),
('Farah', 'Hassan', 'fhassan6@slate.com', 'Female', '2002-06-01', '4 Westridge Center'),
('Maurine', 'Sommerling', 'msommerling7@nature.com', 'Male', '2017-06-13', '41 Homewood Court'),
('Lorelle', 'Spino', 'lspino8@ibm.com', 'Female', '2004-10-01', '726 Maple Point'),
('Marcos', 'Nuccitelli', 'mnuccitelli9@godaddy.com', 'Female', '1994-03-22', '8 Eastwood Plaza'),
('Meredith', 'Shorey', 'mshoreya@about.me', 'Male', '2013-03-14', '5154 Eagle Crest Terrace'),
('Tabbi', 'Ubsdall', 'tubsdallb@theglobeandmail.com', 'Male', '1997-04-30', '8618 Golf View Crossing'),
('Wylie', 'Borman', 'wbormanc@psu.edu', 'Male', '2013-07-20', '3 Barby Alley'),
('Falkner', 'Kearns', 'fkearnsd@posterous.com', 'Male', '2018-04-12', '3 Harper Parkway'),
('Valerie', 'Totterdell', 'vtotterdelle@businessweek.com', 'Female', '2016-11-24', '73524 Hudson Parkway')ON CONFLICT DO NOTHING;`

func TestCreateCustomerHappy(t *testing.T) {
	conn, err := InitConnDbForTest()
	if err != nil {
		t.Fatalf("db connection error %#v", err)
	}
	repo := adapters.NewCustomerRepository(conn)

	c := &customer.DbCustomer{
		FirstName: fmt.Sprintf("uniqueCustomerName%d", 1000),
		LastName:  fmt.Sprintf("uniqueCustomerLast%d", 1000),
		Gender:    "Male",
		Email:     fmt.Sprintf("uniqueCustomerEmail%d%s", 1000, "@gmail.com"),
		Address:   fmt.Sprintf("uniqueCustomerAddress%d", 1000),
		Birthdate: time.Date(1989, time.Month(10), 01, 0, 0, 0, 0, time.UTC),
	}
	created, err := repo.CreateCustomer(c)
	assert.Nil(t, err)
	assert.NotNil(t, created.ID)
	assert.NotEqual(t, "", created.ID)

	founded, err := repo.GetCustomerByID(created.ID)
	assert.Nil(t, err)
	assert.NotNil(t, founded)
	assert.Equal(t, founded.Birthdate.Format(consts.DefaultDateLayout), created.Birthdate.Format(consts.DefaultDateLayout))
}

func TestCreateCustomerError(t *testing.T) {
	conn, err := InitConnDbForTest()
	if err != nil {
		t.Fatalf("db connection error %#v", err)
	}

	c := &customer.DbCustomer{
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

	c := &customer.DbCustomer{
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
