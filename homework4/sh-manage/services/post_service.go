package services

import "gorm.io/gorm"

type PostService struct {
	// 这里可以添加数据库连接等依赖
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}
