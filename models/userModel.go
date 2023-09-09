package models

import (
	"time"

	"gorm.io/gorm"
)

type UserStatus string
type UserRole string

const (
	Pending  UserStatus = "pending"
	Active   UserStatus = "active"
	Inactive UserStatus = "inactive"
)

const (
	Admin     UserRole = "admin"
	Moderator UserRole = "moderator"
	Normal    UserRole = "normal"
)

// Gorm model will automatically create id, created_at, updated_at and deleted_at for you
// Name will be automatically converted to snake_case
// See https://gorm.io/docs/models.html for reference of field tags
// See https://dev.mysql.com/doc/refman/8.0/en/data-types.html for data types if your use 'type:'
type User struct {
	gorm.Model
	Username     string     `gorm:"uniqueIndex;type:varchar(20);not null"`
	Email        string     `gorm:"uniqueIndex;type:varchar(256);not null"`
	Password     string     `gorm:"type:varchar(256);not null"`
	DisplayName  string     `gorm:"uniqueIndex;type:varchar(30);not null"`
	Firstname    string     `gorm:"index;type:varchar(100);not null"`
	Middlename   string     `gorm:"type:varchar(100);default:null"`
	Lastname     string     `gorm:"index;type:varchar(100);not null"`
	ProfileImage string     `gorm:"serializer:json;default:null"`
	Role         UserRole   `gorm:"type:enum('admin', 'moderator', 'normal');not null;default:'normal'"`
	Status       UserStatus `gorm:"type:enum('pending', 'active', 'inactive');not null;default:'pending'"`
	VerifiedAt   time.Time  `gorm:"serializer:unixtime;type:time;default:null"`
}
