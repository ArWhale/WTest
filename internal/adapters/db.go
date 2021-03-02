package adapters

import (
	"database/sql"
	"fmt"
)

type Repositories struct {
	CustomerRepository *CustomerRepository
	db                 *sql.DB
}

func NewRepositories(pgURL string) (*Repositories, error) {
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database connection. Err: %+v", err)
	}

	return &Repositories{
		CustomerRepository: NewCustomerRepository(db),
		db:                 db,
	}, nil
}

func (s *Repositories) Close() error {
	return s.db.Close()
}
