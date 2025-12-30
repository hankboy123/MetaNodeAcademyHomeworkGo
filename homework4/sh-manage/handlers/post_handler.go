package handlers

import "sh-manage/services"

type PostHandler struct {
	PostService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{
		PostService: postService,
	}
}
