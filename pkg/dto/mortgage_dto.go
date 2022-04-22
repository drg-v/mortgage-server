package dto

type MortgageDto struct {
	BankID      int     `json:"bankID"`
	InitialLoan float64 `json:"initialLoan"`
	DownPayment float64 `json:"downPayment"`
}
