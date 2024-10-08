// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAccount = "accounts"

// Account mapped from table <accounts>
type Account struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Owner     string    `gorm:"column:owner;not null" json:"owner"`
	Balance   int64     `gorm:"column:balance;not null" json:"balance"`
	Currency  string    `gorm:"column:currency;not null" json:"currency"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}

// TableName Account's table name
func (*Account) TableName() string {
	return TableNameAccount
}
