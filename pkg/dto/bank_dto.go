package dto

type BankDto struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	InterestRate   float64 `json:"interestRate"`
	MaxLoan        float64 `json:"maxLoan"`
	MinDownPayment float64 `json:"minDownPayment"`
	LoanTermMonths int     `json:"loanTermMonths"`
}
