package references_check

// TODO: add test with annotations on back-references

type User struct {
	Name      string
	CompanyID string
	Company   Company `gorm:"references:code"` // want "Related field \"code\" doesn't exist on model \"Company\""
}

type Company struct {
	ID   int
	Code string
	Name string
}

type Order struct {
	Id        uint   `gorm:"primaryKey"`
	CompanyID string `gorm:"references:Code"` // want "Related model \"string\" doesn't exist"
	Company   Company
}
