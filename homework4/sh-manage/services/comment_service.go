package services

import "gorm.io/gorm"

type CommentService struct {
	// 这里可以添加数据库连接等依赖
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db: db}
}
