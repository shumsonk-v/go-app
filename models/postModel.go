package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `gorm:"type:varchar(100);default:null"`
	Body     string `gorm:"type:longtext;default:null"`
	Likes    int    `gorm:"type:INT UNSIGNED;default:0"`
	Draft    bool   `gorm:"not null;default:true"`
	AuthorId uint   `gorm:"user_id"`
}
