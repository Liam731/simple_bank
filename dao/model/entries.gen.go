// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEntry = "entries"

// Entry mapped from table <entries>
type Entry struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AccountID int64     `gorm:"column:account_id;not null" json:"account_id"`
	Amount    int64     `gorm:"column:amount;not null" json:"amount"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}

// TableName Entry's table name
func (*Entry) TableName() string {
	return TableNameEntry
}
