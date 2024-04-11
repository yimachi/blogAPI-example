package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/yourname/reponame/apperrors"
	"github.com/yourname/reponame/controllers/services"
	"github.com/yourname/reponame/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment のハンドラ
// 記事にコメントを投稿する
// 投稿に成功したコメントの内容
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.BadParam.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
