package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const PhotoApi = "https://api.pexels.com/v1"

type Client struct {
	Token          string
	Hc             http.Client
	RemainingTimes int32
}

func (c *Client) SearchPhotos(query string, perpage int, page int) (*SearchResult, error) {
	url := fmt.Sprintf(PhotoApi+"/search?query=%s&per_page=%d&page=%d", query, perpage, page)
	response, err := c.requestWithAuth("GET", url)
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var result SearchResult
	err = json.Unmarshal(data, &result)
	return &result, err
}

func (c *Client) requestWithAuth(method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	//This is how you authorise
	req.Header.Add("Authorization", c.Token)
	response, err := c.Hc.Do(req)
	if err != nil {
		return nil, err
	}
	times, err := strconv.Atoi(response.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return response, nil
	} else {
		c.RemainingTimes = int32(times)
	}
	return response, nil
}

func (c *Client) CuratedPhotos(perpage int, page int) (*CuratedResult, error) {
	url := fmt.Sprintf(PhotoApi+"/curated?per_page=%d&page=%d", perpage, page)
	response, err := c.requestWithAuth("GET", url)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result CuratedResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetPhoto(id int32) (*Photo, error) {
	url := fmt.Sprintf(PhotoApi+"/curated?per_page=%d&page=%d", perpage, page)
}
