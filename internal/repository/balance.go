package repository

import (
	"github.com/jinzhu/gorm"
)

// BalanceRepository definition
type BalanceRepository struct {
	db *gorm.DB
}

// Balance : balance for given asset
type Balance struct {
	gorm.Model
	Asset string
	Free  float64
}

// Last : last know balance for giver asset
func (r *BalanceRepository) Last(asset string) *Balance {
	var lastBalance Balance
	r.db.Last(&lastBalance, "Asset = ?", asset)

	return &lastBalance
}

// NewBalanceRepository -
func NewBalanceRepository(db *gorm.DB) *BalanceRepository {
	return &BalanceRepository{db}
}
