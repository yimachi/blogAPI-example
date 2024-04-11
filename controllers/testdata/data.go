package testdata

import "github.com/yourname/reponame/models"

var articleTestData = []models.Article{
	models.Article{
		ID:          1,
		Title:       "firstpost",
		Contents:    "first blog",
		UserName:    "hoge",
		NiceNum:     2,
		CommentList: commentTestData,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog",
		UserName: "hoge",
		NiceNum:  4,
	},
}

var commentTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "1st comment",
	},
	models.Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "welcome",
	},
}
