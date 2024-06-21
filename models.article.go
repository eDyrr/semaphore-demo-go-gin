package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "article 1", Content: "article 1 body"},
	article{ID: 2, Title: "article 2", Content: "article 2 body"},
}

// return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, v := range articleList {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, errors.New("Article not found")
}
