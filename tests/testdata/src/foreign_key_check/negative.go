package foreign_key_check

type User struct {
	Name         string
	CompanyRefer uint
	Company      Company `gorm:"foreignKey:CompanyRefer"`
}

type Company struct {
	ID   int
	Name string
}
