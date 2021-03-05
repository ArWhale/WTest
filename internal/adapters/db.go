package adapters

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Repositories struct {
	CustomerRepository *CustomerRepository
	db                 *sql.DB
}

func NewRepositories(pgURL string, logger logrus.FieldLogger) (*Repositories, error) {
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database connection. Err: %+v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed ping database, Err: %+v", err)
	}

	return &Repositories{
		CustomerRepository: NewCustomerRepository(db, logger),
		db:                 db,
	}, nil
}

func (s *Repositories) Close() error {
	return s.db.Close()
}
