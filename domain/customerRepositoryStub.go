package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {

	customers := []Customer{
		{
			Id:          "1001",
			Name:        "Pilseong",
			City:        "SeongName",
			Zipcode:     "12345",
			DateOfBirth: "1979-10-18",
			Status:      "1",
		},
		{
			Id:          "1002",
			Name:        "Sangmi",
			City:        "SeongName",
			Zipcode:     "12345",
			DateOfBirth: "1978-09-11",
			Status:      "1",
		},
	}

	return CustomerRepositoryStub{customers}
}
