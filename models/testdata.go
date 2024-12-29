package models

type Order struct {
	Id            uint `gorm:"primaryKey"`
	Status        string
	ProductTypeId uint
	ProductType   ProductType
	ProductAmount uint
	Description   string
	CustomerId    uint `gorm:"null;foreignKey:CustomerId;"`
	Customer      Customer
	CreatedAt     int64 `gorm:"autoCreateTime"`
	DeadlineDate  int64
}

type ProductType struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}

type Customer struct {
	Id      uint `gorm:"primaryKey"`
	Title   string
	Contact string
	Orders  []Order `gorm:"foreignKey:CustomerId"`
}
