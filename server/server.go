package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/EMC-Dojo/go-url-preview/title"
)

func GetTitle(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	url := r.URL.Query().Get("url")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Fprint(w, "error getting from URL")
	}
	if titleStr, ok := title.GetHtmlTitle(resp.Body); ok {
		fmt.Println(titleStr)
		fmt.Fprintf(w, `{"title": "%s"}`, titleStr)
	} else {
		fmt.Println("Fail to get HTML title")
		fmt.Fprint(w, "error getting from URL")
	}
}

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.HandleFunc("/getTitle", GetTitle)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}
