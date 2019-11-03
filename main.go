package main

import (
	"fmt"
	"html"
	"os"

	lyric_api "github.com/meysampg/lyrik/lyric-api"
)

func main() {
	if len(os.Args) < 3 {
		panic(fmt.Errorf("use in %s artist track format", os.Args[0]))
	}

	var artist = html.EscapeString(os.Args[1])
	var track = html.EscapeString(os.Args[2])

	lyric, err := lyric_api.Fetch(artist, track)
	if err != nil {
		panic(fmt.Errorf("there is an error on fetching lyric, %s", err))
	}

	fmt.Println(lyric)
}
