package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sesi6/webserver/repositories"
	"time"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("platform")

	users, err := repositories.GetUsers()

	if err != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if userAgent == "WEB" || userAgent == "" {
		// process as html
		tpl, err := template.ParseFiles("webserver/views/static/index.html", "webserver/views/static/header.html")
		if err != nil {
			writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		tpl.Execute(w, users)
		return
	} else if userAgent != "API" {
		// handle userAgent error
		writeJsonResponse(w, http.StatusForbidden, map[string]interface{}{
			"error": fmt.Sprintf("agent %s not allowed", userAgent),
		})
		return
	}

	// process as API
	writeJsonResponse(w, http.StatusOK, map[string]interface{}{
		"payload": users,
	})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != http.MethodPost {
		writeJsonResponse(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"error": fmt.Sprintf("method %s not allowed !", method),
		})
		return
	}

	var req repositories.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	users, _ := repositories.GetUsers()
	req.ID = len(users) + 1
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	err = repositories.CreateUser(&req)
	if err != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	writeJsonResponse(w, http.StatusOK, map[string]interface{}{
		"payload": req,
	})
}

func writeJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
