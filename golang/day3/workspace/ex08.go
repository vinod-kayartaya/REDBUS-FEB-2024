package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Movie struct {
	Title  string
	Poster string
	Year   string
	Type   string
	imdbID string
}

func (m *Movie) Print() {
	fmt.Printf("Title: %v\n", m.Title)
	fmt.Printf("Year: %v\n", m.Year)
	fmt.Printf("Type: %v\n", m.Type)
	fmt.Printf("imdbID: %v\n", m.imdbID)
	fmt.Printf("Poster: %v\n", m.Poster)
}

func main() {
	file, err := os.Open("movie.json")

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}

	var movie Movie

	if err := json.Unmarshal(bytes, &movie); err != nil {
		panic(err)
	}

	movie.Print()
}
