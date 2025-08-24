package pokeapi

import (
	"errors"
	"io"
	"net/http"
)

func getData(url string) ([]byte, error) {
	var tounmarshal []byte

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("404 Not Found")
	}
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	tounmarshal = dat
	return tounmarshal, nil

}
