package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../models"
)

type PostsPageData struct {
	Friends []models.UserList
	Posts   []models.PostsList
}

type Follow struct {
	UserId     int
	FollowerId int
}

func Posts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		cookieToken, err := r.Cookie("token")
		if err != nil || cookieToken.Value == "" {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}

		userId, err := db.t.GetUserIdFromToken(cookieToken.Value)
		if err != nil {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}

		postsData := PostsPageData{
			Friends: db.l.GetFollowerSuggestions(userId),
			Posts:   db.l.GetFollowerPosts(userId, &db.up),
		}
		log.Println("Posts=======", postsData.Posts)

		t, _ := template.ParseFiles("./views/html/posts.html")
		t.Execute(w, postsData)

	} else if r.Method == "POST" {
		type Status struct {
			Status string
		}
		var statusMessage Status

		userCookie, err := r.Cookie("user_id")
		if err != nil || userCookie.Value == "" {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &statusMessage)
		user, err := strconv.Atoi(userCookie.Value)
		if err != nil {
			panic(err)
		}
		text := statusMessage.Status

		db.up.AddPost(user, text)
		log.Println("db.up===", db.up)
		ReturnAPIResponse(w, r, 200, "Status shared successfully!!", make(map[string]string))
	}
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userCookie, err := r.Cookie("user_id")
		if err != nil || userCookie.Value == "" {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}
		var followsData Follow
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &followsData)
		user, err := strconv.Atoi(userCookie.Value)
		if err != nil {
			panic(err)
		}
		follower := followsData.FollowerId
		db.l.FollowUser(user, follower)
		log.Println("db.l===", db.l)
		ReturnAPIResponse(w, r, 200, "User Followed successfully!!", make(map[string]string))
	}
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userCookie, err := r.Cookie("user_id")
		if err != nil || userCookie.Value == "" {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}
		var followsData Follow
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &followsData)
		user, err := strconv.Atoi(userCookie.Value)
		if err != nil {
			panic(err)
		}
		follower := followsData.FollowerId
		db.l.UnfollowUser(user, follower)
		log.Println("db.l===", db.l)
		ReturnAPIResponse(w, r, 200, "User UnFollowed successfully!!", make(map[string]string))
	}
}
