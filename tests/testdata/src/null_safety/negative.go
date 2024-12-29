package null_safety

type Order1 struct {
	Id          uint `gorm:"primaryKey"`
	Description string
	// not nullable - not nullable
	CustomerId uint `gorm:"not null;foreignKey:CustomerId;"`
}

type Order2 struct {
	Id          uint `gorm:"primaryKey"`
	Description string
	// nullable - nullable
	CustomerId *uint `gorm:"null;foreignKey:CustomerId;"`
}

type Order3 struct {
	Id uint `gorm:"primaryKey"`
	// nullable - unspecified
	Status      *string
	Description string
}

type Order4 struct {
	Id uint `gorm:"primaryKey"`
	// not nullable - unspecified
	Status      *string
	Description string
}
