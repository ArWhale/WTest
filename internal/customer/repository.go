package customer

type Repository interface {
	CreateCustomer(*Customer) (*Customer, error)
	UpdateCustomer(*Customer) error
	GetAllCustomers(limit, offset *int64) ([]*Customer, error)
	GetCustomerByID(id int64) (*Customer, error)
	SearchCustomers(firstName, lastName *string, limit, offset *int64) ([]*Customer, error)
	SearchCustomersByEmail(email string) (*Customer, error)
	DeleteCustomerByID(id int64) error
}
