package model

type Category struct {
	// 数据库映射
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	CategoryName string `gorm:"unique_index ，not_null" json:"category_name"`
	CategoryLevel uint32 `json:"category_level"`
	CategoryParent int64 `json:"category_parent"`
	CategoryImage string `json:"category_image"`
	CategoryDescription string `json:"category_description"`
}
