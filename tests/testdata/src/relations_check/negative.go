package relations_check

// Many-to-many

type Library struct {
	Id    uint    `gorm:"primaryKey"`
	Books []*Book `gorm:"many2many:library_book;constraint:OnDelete:CASCADE;"`
}

type Book struct {
	Id        uint       `gorm:"primaryKey"`
	Libraries []*Library `gorm:"many2many:library_book;constraint:OnDelete:CASCADE;"`
}

type Employee struct {
	Id           uint        `gorm:"primaryKey"`
	Subordinates []*Employee `gorm:"many2many:employee_subordinates;constraint:OnDelete:CASCADE;"` // self-reference
}

type Publisher struct {
	Id      uint      `gorm:"primaryKey"`
	Writers []*Writer `gorm:"many2many:publisher_books;constraint:OnDelete:CASCADE;"`
}

type Writer struct {
	Id         uint        `gorm:"primaryKey"`
	Publishers []Publisher `gorm:"many2many:publisher_books;constraint:OnDelete:CASCADE;"`
}

// One-to-many
type Comment struct {
	Id            uint `gorm:"primaryKey"`
	CommentatorId uint
	Commentator   Commentator
}

type Commentator struct {
	Id       uint      `gorm:"primaryKey"`
	Comments []Comment `gorm:"foreignKey:CommentatorId;references:Id;constraint:OnDelete:CASCADE;"`
}

type Post struct {
	Id    uint    `gorm:"primaryKey"`
	Files []*File `gorm:"constraint:OnDelete:CASCADE;"`
}

type File struct {
	Id     uint `gorm:"primaryKey"`
	PostId uint
	Post   Post
}
