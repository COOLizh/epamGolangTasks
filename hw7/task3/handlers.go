package main

import (
	"fmt"
	"net/http"
	"time"
)

func addCookie(w http.ResponseWriter, name string, value string) {
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		address := r.FormValue("address")
		addCookie(w, name, address)
		http.ServeFile(w, r, "form.html")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
