package testdata

import "github.com/yourname/reponame/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "first blog",
		UserName: "hoge",
		NiceNum:  2,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog",
		UserName: "hoge",
		NiceNum:  4,
	},
}
