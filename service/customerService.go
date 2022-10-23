package service

import (
	"github.com/pilseong/banking/domain"
	"github.com/pilseong/banking/dto"
	"github.com/pilseong/banking/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	customersResponse := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		customersResponse = append(customersResponse, *c.ToDto())
	}

	return customersResponse, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()

	return response, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
