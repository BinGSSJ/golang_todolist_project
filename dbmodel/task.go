package dbmodel

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index:not null"` // title and use index
	Status    int    `gorm:"default:'0'"`    // memo status, 0 is unfinished
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64 `gorm:"default:0"`
}
