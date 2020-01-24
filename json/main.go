package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`        // json name.
	Color  bool `json:"color,omitempty"` // no JSON output should be produced if the field has the zero value for its type.
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey", "Ingrid"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json.Marshal %s\n", err)
	}
	// fmt.Println(data) // bytes
	fmt.Printf("%s\n", data)
}
