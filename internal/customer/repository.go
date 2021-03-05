package customer

type Repository interface {
	CreateCustomer(*DbCustomer) (*DbCustomer, error)
	UpdateCustomer(*DbCustomer) error
	GetAllCustomers(limit, offset *int64) ([]*DbCustomer, error)
	GetCustomerByID(id int64) (*DbCustomer, error)
	SearchCustomers(firstName, lastName string, limit, offset *int64) ([]*DbCustomer, error)
	SearchCustomersByEmail(email string) (*DbCustomer, error)
	DeleteCustomerByID(id int64) error
}
