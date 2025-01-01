package relations_check

// Many-to-many

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

type Publisher struct {
	Id      uint      `gorm:"primaryKey"`
	Writers []*Writer `gorm:"many2many:publisher_books;"`
}

type Writer struct {
	Id         uint        `gorm:"primaryKey"`
	Publishers []Publisher `gorm:"many2many:publisher_books;"`
}
