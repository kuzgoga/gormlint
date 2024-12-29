package null_safety

type Order5 struct {
	Id          uint `gorm:"primaryKey"`
	Description string
	// not nullable - nullable
	CustomerId uint `gorm:"null;foreignKey:CustomerId;"` // want "Null safety error in \"Order5\" model, field \"CustomerId\": column nullable policy doesn't match to tag nullable policy"
}

type Order6 struct {
	Id          uint `gorm:"primaryKey"`
	Description string
	// nullable - not nullable
	CustomerId *uint `gorm:"not null;foreignKey:CustomerId;"` // want "Null safety error in \"Order6\" model, field \"CustomerId\": column nullable policy doesn't match to tag nullable policy"
}

type Order7 struct {
	Id          uint `gorm:"primaryKey"`
	Description string
	// not nullable - not nullable, nullable
	CustomerId uint `gorm:"not null;foreignKey:CustomerId;null;"` // want "Null safety error: tags \"null\" and \"not null\" are specified at one field"
}

type Order8 struct {
	Id          uint `gorm:"primaryKey"`
	Description string
	// nullable - not nullable, nullable
	CustomerId *uint `gorm:"not null;foreignKey:CustomerId;null;"` // want "Null safety error: tags \"null\" and \"not null\" are specified at one field"
}
