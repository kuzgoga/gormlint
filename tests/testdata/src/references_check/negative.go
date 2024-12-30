package references_check

type WorkArea struct {
	Id         uint     `gorm:"primaryKey"`
	Workshop   Workshop `gorm:"foreignKey:WorkshopId;references:Id;"`
	WorkshopId uint
}

type Workshop struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	WorkAreas []WorkArea `gorm:"constraint:OnDelete:CASCADE;"`
}

type TeamType struct {
	Code uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type TeamTask struct {
	Id         uint `gorm:"primaryKey"`
	TeamTypeId uint
	TeamType   TeamType `gorm:"references:Code;"`
}
