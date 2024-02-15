package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResponseData struct {
	Search       []Movie
	TotalResults string `json:"totalResults"`
	Response     string
}

type Movie struct {
	Title  string
	Year   string
	Type   string
	Poster string
	imdbID string
}

func main() {
	url := "https://www.omdbapi.com/?apikey=aa9e49f&s=spider"

	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var movieResponse ResponseData
	json.Unmarshal(data, &movieResponse)
	for _, m := range movieResponse.Search {
		fmt.Println(m.Title)
	}
	fmt.Printf("Total %v movies found\n", movieResponse.TotalResults)

}
