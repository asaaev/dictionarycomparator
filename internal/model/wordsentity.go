package model

type Word struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	WordBody string `gorm:"uniqueIndex;not null"`
}
