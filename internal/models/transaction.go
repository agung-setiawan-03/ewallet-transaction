package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	ID                int       `json:"-"`
	UserID            int       `json:"user_id"`
	Amount            float64   `json:"amount" gorm:"column:amount;type:decimal(15,2);" valid:"required"`
	TransactionType   string    `json:"transaction_type" gorm:"column:transaction_type;type:enum('TOPUP','PURCHASE', 'REFUND');" valid:"required"`
	TransactionStatus string    `json:"transaction_status" gorm:"column:transaction_status;type:enum('PENDING','SUCCESS', 'FAILED', 'REVERSED');"`
	Reference         string    `json:"reference" gorm:"column:reference;type:varchar(255);"`
	Description       string    `json:"description" gorm:"column:description;type:varchar(255);" valid:"required"`
	AdditionalInfo    string    `json:"additional_info" gorm:"column:additional_info;type:text"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
}

func (l Transaction) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

func (*Transaction) TableName() string {
	return "transactions"
}

type CreateTransactionResponse struct {
	Reference          string `json:"reference"`
	TransactionStatus string `json:"transaction_status"`
}

type UpdateStatusTransaction struct {
	Reference         string `json:"reference" valid:"required"`
	TransactionStatus string `json:"transaction_status" valid:"required"`
	AdditionalInfo    string `json:"additional_info"`
}

func (l UpdateStatusTransaction) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
type RefundTransaction struct {
	Reference      string `json:"reference" valid:"required"`
	Description    string `json:"description" valid:"required"`
	AdditionalInfo string `json:"additional_info"`
}

func (l RefundTransaction) Validate() error {
	v := validator.New()
	return v.Struct(l)
}