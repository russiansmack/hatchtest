package hatchtest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const base = "https://api.thecatapi.com/v1"

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient(apiKey string) *Client {
	c := &Client{}
	c.apiKey = apiKey

	return c
}

//create request
func (c *Client) makeRequest(endpoint string, method string) (*http.Request, error) {

	url := base + endpoint

	req, err := c.httpClient.NewRequest(method, url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	return req, err
}

//running the request
func (c *Client) runRequest(req *http.Request, v interface{}) (*http.Response, error) {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(v)

	return response, err
}

//Fetching cat images
func (c *Client) GetImageSearch( /*, opts *GetImagesOptions*/ ) {

	req, err := makeRequest("/images/search", http.MethodGet)
	if err != nil {
		return nil, err
	}

	response, err := runRequest(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func main() {
	client := NewClient("7a1768b0-1600-4c55-9769-83721284ab92")
	images, _ := client.GetImageSearch()
	fmt.Println(images)
}

/*
type GetImagesOptions struct {
  Query   string `url:"q"`
  ShowAll bool   `url:"all"`
  Page    int    `url:"page"`
}
opt := GetImagesOptions{ "foo", true, 2 }
v, _ := query.Values(opt)
fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"

*/
