package lyric_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
)

const url = "http://lyric-api.herokuapp.com/api/find/%s/%s"

type lyric struct {
	Lyric string
	Err   string
}

func Fetch(artist, track string) (string, error) {
	var url = fmt.Sprintf(url, artist, track)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	var body = make([]byte, response.ContentLength)
	_, err = response.Body.Read(body)
	defer response.Body.Close()
	if err != nil && err != io.EOF {
		return "", err
	}

	data := &lyric{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return "", err
	}

	if data.Err != "none" {
		return "", errors.New(data.Err)
	}

	return data.Lyric, nil
}
