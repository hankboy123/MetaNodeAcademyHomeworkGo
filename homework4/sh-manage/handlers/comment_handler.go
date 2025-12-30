package handlers

import "sh-manage/services"

type CommentHandler struct {
	CommentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{
		CommentService: commentService,
	}
}
