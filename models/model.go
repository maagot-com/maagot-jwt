// create database models
package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `json:"first_name" gorm:"type:varchar(100);not null" binding:"required,alpha"`
	LastName    string `json:"last_name" gorm:"type:varchar(100);not null" binding:"required,alpha"`
	Email       string `json:"email" gorm:"type:varchar(200)"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(14);uniqueIndex;unique;not null" binding:"required,e164,min=11,max=14"`
	DeviceToken string `json:"device_token" gorm:"type:varchar(200);not null" binding:"required,alpha"`
	Pin         string `json:"pin" gorm:"type:varchar(200);not null" binding:"required,max=4"`
	Quota       uint   `json:"quota" gorm:"type:bigint;default:0;not null"`
	Locked      bool   `json:"locked" gorm:"type:boolean;default:false"`
	Photo       string `json:"photo" gorm:"type:text"`
	IsActive    bool   `json:"is_active" gorm:"type:boolean;default:true;index"`
}

type Wallet struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	Balance   uint   `json:"balance" gorm:"type:bigint;default:0;not null" binding:"required,number,min=0,max=10000000"`
	Currency  string `json:"currency" gorm:"type:varchar(3);default:XOF;not null" binding:"alpha,oneof=XOF GHS XAF GNH EUR USD"`
	IsActive  bool   `json:"is_active" gorm:"type:boolean;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Payment struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	WalletID  uint   `json:"wallet_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	Balance   uint   `json:"balance" gorm:"type:bigint;default:0;not null" binding:"required,number,min=0,max=10000000"`
	Currency  string `json:"currency" gorm:"type:varchar(3);default:XOF;not null" binding:"alpha,oneof=XOF GHS XAF GNH EUR USD"`
	IsActive  bool   `json:"is_active" gorm:"type:boolean;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Transaction struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	PaymentID uint   `json:"payment_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	WalletID  uint   `json:"wallet_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	Balance   uint   `json:"balance" gorm:"type:bigint;default:0;not null" binding:"required,number,min=0,max=10000000"`
	Currency  string `json:"currency" gorm:"type:varchar(3);default:XOF;not null" binding:"alpha,oneof=XOF GHS XAF GNH EUR USD"`
	IsActive  bool   `json:"is_active" gorm:"type:boolean;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Transfer struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	FromID    uint   `json:"from_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	ToID      uint   `json:"to_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	Balance   uint   `json:"balance" gorm:"type:bigint;default:0;not null" binding:"required,number,min=0,max=10000000"`
	Currency  string `json:"currency" gorm:"type:varchar(3);default:XOF;not null" binding:"alpha,oneof=XOF GHS XAF GNH EUR USD"`
	IsActive  bool   `json:"is_active" gorm:"type:boolean;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Notification struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	FromID    uint   `json:"from_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	ToID      uint   `json:"to_id" gorm:"type:bigint;not null" binding:"required,number,gt=0"`
	Balance   uint   `json:"balance" gorm:"type:bigint;default:0;not null" binding:"required,number,min=0,max=10000000"`
	Currency  string `json:"currency" gorm:"type:varchar(3);default:XOF;not null" binding:"alpha,oneof=XOF GHS XAF GNH EUR USD"`
	IsActive  bool   `json:"is_active" gorm:"type:boolean;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OTP struct {
	ID          uint      `json:"id" gorm:"primaryKey;unique"`
	Code        string    `json:"code" gorm:"type:varchar(6)"`
	IsUsed      bool      `json:"is_used" gorm:"type:boolean;not null;default:false"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(14);not null" binding:"required,e164,min=11,max=14"`
	KeyUID      string    `json:"key_uid" gorm:"type:varchar(100);not null"`
	ExpiryAt    time.Time `json:"expiry_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
