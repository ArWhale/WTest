package customer

type Repository interface {
	CreateCustomer(*Customer) (*Customer, error)
	UpdateCustomer() error
	GetAllCustomers() ([]*Customer, error)
	GetCustomerByID(id int64) (*Customer, error)
	SearchCustomers(limit, offset int64) ([]*Customer, error)
}
