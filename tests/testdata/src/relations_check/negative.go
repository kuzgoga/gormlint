package relations_check

type Library struct {
	Id    uint    `gorm:"primaryKey"`
	Books []*Book `gorm:"many2many:library_book;"`
}

type Book struct {
	Id        uint       `gorm:"primaryKey"`
	Libraries []*Library `gorm:"many2many:library_book;"`
}

type Employee struct {
	Id           uint        `gorm:"primaryKey"`
	Subordinates []*Employee `gorm:"many2many:employee_subordinates;"` // self-reference
}
