package repositories_test

import (
	"testing"

	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want coment of articleID %d but got ID %d\n", articleID, comment.ArticleID)
		}
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "CommentInsertTest",
	}

	exptectedCommentID := 3
	newComment, err := repositories.InsertComemnt(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}
	if newComment.CommentID != exptectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", exptectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?
		`
		testDB.Exec(sqlStr, comment.Message)
	})
}
