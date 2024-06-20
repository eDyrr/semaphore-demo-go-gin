package main

import "testing"

// test the fucntion that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// check taht the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Fail()
			break
		}

	}
}
