package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Pedro", City: "Salvador", ZipCode: "110011", DateOfBirth: "2023-05-06", Status: "1"},
		{Id: "1001", Name: "Lucas", City: "Feira de Santana", ZipCode: "220022", DateOfBirth: "2023-08-10", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
