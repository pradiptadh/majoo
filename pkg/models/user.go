package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Merchant struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	User         *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	MerchantName string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Outlet struct {
	ID         int       `json:"id"`
	MerchantID int       `json:"merchant_id"`
	Merchant   *Merchant `json:"merchant,omitempty" gorm:"foreignKey:MerchantID"`
	OutletName string    `json:"outlet_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Transaction struct {
	ID         int       `json:"id"`
	MerchantID int       `json:"merchant_id"`
	Merchant   *Merchant `json:"merchant,omitempty" gorm:"foreignKey:MerchantID"`
	OutletID   int       `json:"outlet_id"`
	Outlet     *Outlet   `json:"outlet,omitempty" gorm:"foreignKey:OutletID"`
	BillTotal  float64   `json:"bill_total"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
