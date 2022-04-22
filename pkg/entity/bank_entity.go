package entity

type Bank struct {
	ID             uint    `gorm:"primaryKey;autoIncrement"`
	Name           string  `gorm:"type:varchar(100);not null"`
	InterestRate   float64 `gorm:"not null"`
	MaxLoan        float64 `gorm:"not null"`
	MinDownPayment float64 `gorm:"not null"`
	LoanTermMonths int     `gorm:"not null"`
}
