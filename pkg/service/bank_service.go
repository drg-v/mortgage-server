package service

import (
	"errors"
	"math"
	"mortgage-calculator/pkg/dto"
	"mortgage-calculator/pkg/entity"
	"mortgage-calculator/pkg/repository"
)

type BankService interface {
	Get(id int) (dto.BankDto, error)
	GetAll() ([]dto.BankDto, error)
	Create(bank dto.BankDto) error
	Update(bank dto.BankDto) error
	Delete(id int) error
	CalculateMortgage(mortgage dto.MortgageDto) (float64, error)
}

type bankService struct {
	bankRepository repository.BankRepository
}

func NewBankService(bankRepo repository.BankRepository) BankService {
	return &bankService{bankRepository: bankRepo}
}

func (bankService *bankService) Get(id int) (dto.BankDto, error) {
	bankEntity, err := bankService.bankRepository.Get(id)
	if err != nil {
		return dto.BankDto{}, errors.New("bank service - unable to find the bank")
	}
	bankDto := dto.BankDto{
		ID:             bankEntity.ID,
		Name:           bankEntity.Name,
		InterestRate:   bankEntity.InterestRate,
		MaxLoan:        bankEntity.MaxLoan,
		MinDownPayment: bankEntity.MinDownPayment,
		LoanTermMonths: bankEntity.LoanTermMonths,
	}
	return bankDto, nil
}

func (bankService *bankService) GetAll() ([]dto.BankDto, error) {
	bankEntities, err := bankService.bankRepository.GetAll()
	if err != nil {
		return []dto.BankDto{}, errors.New("bank service - unable to find all banks")
	}
	bankDtoSlice := make([]dto.BankDto, 0, len(bankEntities))
	for _, val := range bankEntities {
		bankDtoSlice = append(bankDtoSlice, dto.BankDto{
			ID:             val.ID,
			Name:           val.Name,
			InterestRate:   val.InterestRate,
			MaxLoan:        val.MaxLoan,
			MinDownPayment: val.MinDownPayment,
			LoanTermMonths: val.LoanTermMonths,
		})
	}
	return bankDtoSlice, nil
}

func (bankService *bankService) Create(bank dto.BankDto) error {
	bankEntity := entity.Bank{
		ID:             bank.ID,
		Name:           bank.Name,
		InterestRate:   bank.InterestRate,
		MaxLoan:        bank.MaxLoan,
		MinDownPayment: bank.MinDownPayment,
		LoanTermMonths: bank.LoanTermMonths,
	}
	err := bankService.bankRepository.Save(bankEntity)
	if err != nil {
		return errors.New("bank service - error creating new bank")
	}
	return nil
}

func (bankService *bankService) Update(bank dto.BankDto) error {
	bankEntity := entity.Bank{
		ID:             bank.ID,
		Name:           bank.Name,
		InterestRate:   bank.InterestRate,
		MaxLoan:        bank.MaxLoan,
		MinDownPayment: bank.MinDownPayment,
		LoanTermMonths: bank.LoanTermMonths,
	}
	err := bankService.bankRepository.Update(bankEntity)
	if err != nil {
		return errors.New("bank service - error updating bank")
	}
	return nil
}

func (bankService *bankService) Delete(id int) error {
	err := bankService.bankRepository.Delete(id)
	if err != nil {
		return errors.New("bank service - error deleting bank")
	}
	return nil
}

func (bankService *bankService) CalculateMortgage(mortgage dto.MortgageDto) (float64, error) {
	bank, err := bankService.bankRepository.Get(mortgage.BankID)
	if err != nil {
		return 0, errors.New("bank service - unable to find the bank")
	}
	if mortgage.InitialLoan > bank.MaxLoan {
		return 0, errors.New("bank service - loan limit exceeded")
	}
	if mortgage.DownPayment < mortgage.InitialLoan*bank.MinDownPayment/100.0 {
		return 0, errors.New("bank service - down payment is insufficient")
	}
	amount := mortgage.InitialLoan - mortgage.DownPayment
	monthlyPayment := bankService.mortgageFormula(amount, bank.InterestRate, bank.LoanTermMonths)
	return monthlyPayment, nil
}

func (bankService *bankService) mortgageFormula(amount, interestRate float64, months int) float64 {
	interestRate = interestRate / 100.0 / 12
	numerator := amount * interestRate * math.Pow(1.0+interestRate, float64(months))
	denominator := math.Pow(1.0+interestRate, float64(months)) - 1.0
	return numerator / denominator
}
