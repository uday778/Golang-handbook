package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalUrl  string    `json:"originalurl"`
	ShortUrl     string    `json:"shorturl"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortUrl(OriginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl))
	fmt.Println("hasher:", hasher)
	data := hasher.Sum(nil)
	fmt.Println("hasher data : ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("encoded to string :", hash)
	fmt.Println("final string:", hash[:8])
	return hash[:8]
}

func createURL(originalUrl string) string {
	shortURL := generateShortUrl(originalUrl)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OriginalUrl:  originalUrl,
		ShortUrl:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL

}

func getUrl(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid req body", http.StatusBadRequest)
		return
	}
	shortURL_ := createURL(data.URL)
	// fmt.Fprintf(w, shortURL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
	

}


func rediectURLHandler(w http.ResponseWriter, r *http.Request)  {
	id:= r.URL.Path[len("/redirect/"):]
	url,err := getUrl(id)
	if err != nil {
		http.Error(w,"invalid Request",http.StatusNotFound)
	}
	http.Redirect(w,r,url.OriginalUrl,http.StatusFound)
}

func main() {
	// fmt.Println("url-shortner")
	// OriginalUrl:= "https://github.com/uday-778/"
	// generateShortUrl(OriginalUrl)

	//register the handler function
	http.HandleFunc("/", Handler)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/{id}", rediectURLHandler)

	//starting the server
	fmt.Println("starting the server...")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("error while starting the server")
	}

}
