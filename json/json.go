package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablance", Year: 1942, Color: false, Actors: []string{"Humphrey", "Ingrid"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve", "Jacqueline"}},
	}

	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		fmt.Println("JSON marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		fmt.Println("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
