package repository

import (
	"gorm.io/gorm"
	"mortgage-calculator/pkg/entity"
)

type BankRepository interface {
	Get(id int) (entity.Bank, error)
	GetAll() ([]entity.Bank, error)
	Save(bank entity.Bank) error
	Update(bank entity.Bank) error
	Delete(id int) error
}

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepository{db: db}
}

func (bankRepo *bankRepository) Get(id int) (entity.Bank, error) {
	var bank entity.Bank
	err := bankRepo.db.First(&bank, id).Error
	return bank, err
}

func (bankRepo *bankRepository) GetAll() ([]entity.Bank, error) {
	var banks []entity.Bank
	err := bankRepo.db.Find(&banks).Error
	return banks, err
}

func (bankRepo *bankRepository) Save(bank entity.Bank) error {
	return bankRepo.db.Create(&bank).Error
}

func (bankRepo *bankRepository) Update(bank entity.Bank) error {
	return bankRepo.db.Save(&bank).Error
}

func (bankRepo *bankRepository) Delete(id int) error {
	return bankRepo.db.Delete(&entity.Bank{}, id).Error
}
