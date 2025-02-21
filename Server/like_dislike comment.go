package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	structs "forum/Data"
	database "forum/Database"
)

func LikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Errors(w, structs.Error{Code: http.StatusMethodNotAllowed, Message: "Method not allowed", Page: "Home", Path: "/"})
		return
	}
	ids := strings.Split(r.URL.Path[len("/like_comment/"):], "/")
	if len(ids) != 2 {
		Errors(w, structs.Error{Code: http.StatusBadRequest, Message: "Invalid ID", Page: "Home", Path: "/"})
		return
	}
	id_post, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		Errors(w, structs.Error{Code: http.StatusBadRequest, Message: "Invalid post ID", Page: "Home", Path: "/"})
		return
	}
	id_comment, err := strconv.ParseInt(ids[1], 10, 64)
	if err != nil {
		Errors(w, structs.Error{Code: http.StatusBadRequest, Message: "Invalid comment ID", Page: "Post", Path: "/post/" + ids[0]})
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Error Adding Like", Page: "Post", Path: "/post/" + ids[0]})
		return
	}
	user := database.GetUserConnected(cookie.Value)
	if user == nil {
		http.SetCookie(w, &http.Cookie{Name: "session", Value: "", MaxAge: -1})
		Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Please log to Adding Like", Page: "Post", Path: "/post/" + ids[0]})
		return
	} else if !database.CheckLikeComment(user.UserID, id_post, id_comment) {
		if database.AddLikeComment(user.UserID, id_post, id_comment) != nil {
			Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Error Adding Like", Page: "Post", Path: "/post/" + ids[0]})
			return
		}
	} else if database.DeleteLikeComment(user.UserID, id_post, id_comment) != nil {
		Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Error Deleting Like", Page: "Post", Path: "/post/" + ids[0]})
		return
	}
	cookie.Expires = time.Now().Add(5 * time.Minute)
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	updatedLikes, errLikesComment := database.CountLikesComment(id_post, id_comment)
	if errLikesComment != nil {
		Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Error counting Like", Page: "Post", Path: "/post/" + ids[0]})
		return
	}
	updatedDislikes, errDislikesComment := database.CountDislikesComment(id_post, id_comment)
	if errDislikesComment != nil {
		Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Error counting Dislike", Page: "Post", Path: "/post/" + ids[0]})
		return
	}
	response := map[string]interface{}{
		"updatedLikes":    updatedLikes,
		"updatedDislikes": updatedDislikes,
		"isLiked":         database.CheckLikeComment(user.UserID, id_post, id_comment),
		"isDisliked":      database.CheckDislikeComment(user.UserID, id_post, id_comment),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DislikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Errors(w, structs.Error{Code: http.StatusMethodNotAllowed, Message: "Method not allowed", Page: "Home", Path: "/"})
		return
	}
	ids := strings.Split(r.URL.Path[len("/dislike_comment/"):], "/")
	if len(ids) != 2 {
		Errors(w, structs.Error{Code: http.StatusBadRequest, Message: "Invalid ID", Page: "Home", Path: "/"})
		return
	}
	id_post, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		Errors(w, structs.Error{Code: http.StatusBadRequest, Message: "Invalid post ID", Page: "Home", Path: "/"})
		return
	}
	id_comment, err := strconv.ParseInt(ids[1], 10, 64)
	if err != nil {
		Errors(w, structs.Error{Code: http.StatusBadRequest, Message: "Invalid comment ID", Page: "Home", Path: "/"})
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Error Adding Dislike", Page: "Post", Path: "/post/" + ids[0]})
		return
	}
	user := database.GetUserConnected(cookie.Value)
	if user == nil {
		http.SetCookie(w, &http.Cookie{Name: "session", Value: "", MaxAge: -1})
		Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Please log to Adding Dislike", Page: "Post", Path: "/post/" + ids[0]})
		return
	} else if !database.CheckDislikeComment(user.UserID, id_post, id_comment) {
		if database.AddDislikeComment(user.UserID, id_post, id_comment) != nil {
			Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Error Adding Dislike", Page: "Post", Path: "/post/" + ids[0]})
			return
		}
	} else if database.DeleteDislikeComment(user.UserID, id_post, id_comment) != nil {
		Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Error Deleting Dislike", Page: "Post", Path: "/post/" + ids[0]})
		return
	}
	cookie.Expires = time.Now().Add(5 * time.Minute)
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	updatedLikes, errLikesComment := database.CountLikesComment(id_post, id_comment)
	if errLikesComment != nil {
		Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Error counting Like", Page: "Post", Path: "/post/" + ids[0]})
		return
	}
	updatedDislikes, errDislikesComment := database.CountDislikesComment(id_post, id_comment)
	if errDislikesComment != nil {
		Errors(w, structs.Error{Code: http.StatusInternalServerError, Message: "Error counting Dislike", Page: "Post", Path: "/post/" + ids[0]})
		return
	}
	response := map[string]interface{}{
		"updatedLikes":    updatedLikes,
		"updatedDislikes": updatedDislikes,
		"isLiked":         database.CheckLikeComment(user.UserID, id_post, id_comment),
		"isDisliked":      database.CheckDislikeComment(user.UserID, id_post, id_comment),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
