package relations_check

type Student struct {
	Id      uint     `gorm:"primaryKey"`
	Courses []Course `gorm:"many2many:student_courses;constraint:OnDelete:CASCADE;"`
}

type Course struct {
	Id       uint     `gorm:"primaryKey"`
	Students []Course `gorm:"many2many:student_courses;constraint:OnDelete:CASCADE"` // want "Invalid type `Course` in M2M relation \\(use \\[\\]\\*Student or self-reference\\)"
}

type Author struct {
	Id       uint      `gorm:"primaryKey"`
	Articles []Article `gorm:"many2many:author_articles;constraint:OnDelete:CASCADES;"`
}

type Article struct {
	Id      uint   `gorm:"primaryKey"`
	Authors Author `gorm:"many2many:author_articles;constraint:OnDelete:CASCADE;"` // want "M2M relation `author_articles` with bad type `Author` \\(should be a slice\\)"
}

type Kuzbass struct {
	Id     uint   `gorm:"primaryKey"`
	Cities []City // want "Expected foreignKey `KuzbassId` in model `City` for 1:M relation with model `Kuzbass`"
}

type City struct {
	Id      uint    `gorm:"primaryKey"`
	Kuzbass Kuzbass // want "Invalid relation in field `Kuzbass`"
}

type Federation struct { // want "Id field should be presented model \"Federation\""
	Lands []Land `gorm:"constraint:OnDelete:CASCADE;"`
}

type Land struct {
	Id           uint `gorm:"primaryKey"`
	FederationId uint
}

// Belongs to

type Owner struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	CompanyId int
	Company   Company `gorm:"constraint:OnDelete:CASCADE;"`
}

type Company struct {
	Id   int
	Name string
}
