package services

import (
	"sh-manage/dto"
	"sh-manage/models"
	"sh-manage/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostService struct {
	// 这里可以添加数据库连接等依赖
	db          *gorm.DB
	context     *gin.Context
	userService *UserService
}

func NewPostService(db *gorm.DB, userService *UserService, c *gin.Context) *PostService {
	return &PostService{db: db, userService: userService, context: c}
}

func (p *PostService) CreatePost(post *dto.PostDto) (*models.Post, *utils.AppError) {
	if post == nil {
		return nil, utils.NewAppError(500, "参数不能为空")
	}

	if err := post.Validate(); err != nil {
		return nil, err
	}

	postModel := &models.Post{
		Title:   post.Title,
		Content: post.Content,
		UserId:  utils.GetCurrentUserID(p.context),
	}

	if err := p.db.Create(&postModel).Error; err != nil {
		return nil, utils.NewAppError(500, "Failed to create post")
	}
	return postModel, nil

}
