package models

type Student struct {
	StudentId uint `gorm:"primaryKey"`
	GradeId   int
	Name      string
	Lastname  string
	Mail      string
	Age       int
}

type Grade struct {
	GradeId uint `gorm:"primaryKey"`
	Type    string
}
