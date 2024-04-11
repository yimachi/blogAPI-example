package services

import (
	"github.com/yourname/reponame/apperrors"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

// 引数の情報をもとに新しいコメントを作り、結果を返却します
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComemnt(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
