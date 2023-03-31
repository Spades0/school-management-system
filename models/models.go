package models

type Student struct {
	ID      uint     `gorm:"primary_key"`
	Name    string   `gorm:"not null"`
	Email   string   `gorm:"not null;unique"`
	Courses []Course `gorm:"many2many:student_courses"`
}

type Course struct {
	ID       uint      `gorm:"primary_key"`
	Name     string    `gorm:"not null"`
	Code     string    `gorm:"not null;unique"`
	Students []Student `gorm:"many2many:student_courses"`
}
