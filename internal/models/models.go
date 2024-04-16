package models

type Student struct {
	StudentId uint `gorm:"primaryKey"`
	GradeId   int
	Mail      string
	Age       int
}

type Grade struct {
	GradeId uint `gorm:"primaryKey"`
	Type    string
}
