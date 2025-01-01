package relations_check

type Student struct {
	Id      uint     `gorm:"primaryKey"`
	Courses []Course `gorm:"many2many:student_courses;"`
}

type Course struct {
	Id       uint     `gorm:"primaryKey"`
	Students []Course `gorm:"many2many:student_courses;"` // want "Invalid type `Course` in M2M relation \\(use \\[\\]\\*Student or self-reference\\)"
}

type Author struct {
	Id       uint      `gorm:"primaryKey"`
	Articles []Article `gorm:"many2many:author_articles;"`
}

type Article struct {
	Id      uint   `gorm:"primaryKey"`
	Authors Author `gorm:"many2many:author_articles;"` // want "M2M relation `author_articles` with bad type `Author` \\(should be a slice\\)"
}
