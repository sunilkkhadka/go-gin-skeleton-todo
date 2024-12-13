// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID        uint32         `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	FullName  string         `gorm:"column:full_name;type:varchar(45);not null" json:"full_name"`
	Phone     string         `gorm:"column:phone;type:varchar(15);not null;uniqueIndex:UQ_user_phone,priority:1" json:"phone"`
	Gender    string         `gorm:"column:gender;type:varchar(15);not null" json:"gender"`
	Email     string         `gorm:"column:email;type:varchar(100);not null;uniqueIndex:UQ_user_email,priority:1" json:"email"`
	Password  string         `gorm:"column:password;type:varchar(100);not null" json:"password"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
