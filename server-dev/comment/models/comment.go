package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Content   string    `gorm:"type:text" json:"content"`
	Category  string    `gorm:"type:varchar(255);index" json:"category"` // 非唯一索引
	CreatedAt time.Time `json:"created_at"`
}
