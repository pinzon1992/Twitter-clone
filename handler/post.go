package handler

import (
	"../models"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

type PostsPageData struct {
	Friends []models.UserList
}

type Follow struct {
	UserId     int
	FollowerId int
}

func Posts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		postsData := PostsPageData{
			Friends: db.l.GetFollowerSuggestions(1),
		}

		t, _ := template.ParseFiles("./views/html/posts.html")
		t.Execute(w, postsData)
	} else {
		r.ParseForm()

		//used to post data
	}

}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var followsData Follow
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &followsData)
		user := followsData.UserId
		follower := followsData.FollowerId
		db.l.FollowUser(user, follower)

		ReturnAPIResponse(w, r, 200, "User Followed successfully!!")
	}
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var followsData Follow
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &followsData)
		user := followsData.UserId
		follower := followsData.FollowerId
		db.l.UnfollowUser(user, follower)

		ReturnAPIResponse(w, r, 200, "User UnFollowed successfully!!")
	}
}
