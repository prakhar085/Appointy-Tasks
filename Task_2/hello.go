package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Article struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	SubTitle          string `json:"subtitle"`
	Content           string `json:"content"`
	CreationTimestamp string `json:"timestamp"`
}

type Articles []Article

type articleHandler struct {
	sync.Mutex
	articles Articles
}

func (ah *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ah.get(w, r)
	case "POST":
		ah.post(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid Method")
	}
}

func (ah *articleHandler) get(w http.ResponseWriter, r *http.Request) {
	defer ah.Unlock()
	ah.Lock()
	id, err := iDfromURL(r)
	if err != nil {
		respondWithJSON(w, http.StatusOK, ah.articles)
		return
	}
	if id >= len(ah.articles) || id < 0 {
		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}
	respondWithJSON(w, http.StatusOK, ah.articles[id])

}

func (ah *articleHandler) post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		respondWithError(w, http.StatusUnsupportedMediaType, "Content type 'application/json' required")
		return
	}
	var article Article
	err = json.Unmarshal(body, &article)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer ah.Unlock()
	ah.Lock()
	ah.articles = append(ah.articles, article)
	respondWithJSON(w, http.StatusCreated, article)

}

func iDfromURL(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		return 0, errors.New("Not Found")
	}
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, errors.New("Not Found")
	}
	return id, nil
}

func newarticleHandler() *articleHandler {
	t := time.Now()
	return &articleHandler{
		articles: Articles{
			Article{
				"1",
				"NA",
				"NA",
				"NA",
				t.Format("20060102150405"),
			},
			Article{
				"2",
				"NA",
				"NA",
				"NA",
				t.Format("20060102150405"),
			},
		},
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"ERROR": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	responce, _ := json.Marshal(data)
	w.Header().Add("content-type", "applications/jason")
	w.WriteHeader(code)
	w.Write(responce)

}

func main() {
	port := ":8080"
	ah := newarticleHandler()
	http.Handle("/articles", ah)
	http.Handle("/articles/", ah)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "_")
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
