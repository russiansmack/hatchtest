package main

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
	c.httpClient = http.DefaultClient
	c.apiKey = apiKey

	return c
}

//create request
func (c *Client) makeRequest(endpoint string, method string) (*http.Request, error) {

	url := base + endpoint
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	return req, err
}

//running the request
func (c *Client) runRequest(req *http.Request, v interface{}) (ImageResult, error) {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	//body, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))

	var results ImageResult

	err = json.NewDecoder(response.Body).Decode(&results)
	if err != nil {
		fmt.Println(err)
	}

	return results, nil
}

//Fetching cat images
func (c *Client) GetImageSearch() (ImageResult, error) {

	req, err := c.makeRequest("/images/search", http.MethodGet)
	if err != nil {
		return nil, err
	}

	//var v ImageResult
	var v interface{}

	response, err := c.runRequest(req, v)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func main() {
	c := NewClient("7a1768b0-1600-4c55-9769-83721284ab92")
	images, err := c.GetImageSearch()
	if err != nil {
		fmt.Println("badjuju")
	}
	fmt.Println(images[0].URL)
}

type ImageResult []struct {
	Breeds []struct {
		Weight struct {
			Imperial string `json:"imperial"`
			Metric   string `json:"metric"`
		} `json:"weight"`
		ID               string `json:"id"`
		Name             string `json:"name"`
		CfaURL           string `json:"cfa_url"`
		VetstreetURL     string `json:"vetstreet_url"`
		VcahospitalsURL  string `json:"vcahospitals_url"`
		Temperament      string `json:"temperament"`
		Origin           string `json:"origin"`
		CountryCodes     string `json:"country_codes"`
		CountryCode      string `json:"country_code"`
		Description      string `json:"description"`
		LifeSpan         string `json:"life_span"`
		Indoor           int    `json:"indoor"`
		AltNames         string `json:"alt_names"`
		Adaptability     int    `json:"adaptability"`
		AffectionLevel   int    `json:"affection_level"`
		ChildFriendly    int    `json:"child_friendly"`
		DogFriendly      int    `json:"dog_friendly"`
		EnergyLevel      int    `json:"energy_level"`
		Grooming         int    `json:"grooming"`
		HealthIssues     int    `json:"health_issues"`
		Intelligence     int    `json:"intelligence"`
		SheddingLevel    int    `json:"shedding_level"`
		SocialNeeds      int    `json:"social_needs"`
		StrangerFriendly int    `json:"stranger_friendly"`
		Vocalisation     int    `json:"vocalisation"`
		Experimental     int    `json:"experimental"`
		Hairless         int    `json:"hairless"`
		Natural          int    `json:"natural"`
		Rare             int    `json:"rare"`
		Rex              int    `json:"rex"`
		SuppressedTail   int    `json:"suppressed_tail"`
		ShortLegs        int    `json:"short_legs"`
		WikipediaURL     string `json:"wikipedia_url"`
		Hypoallergenic   int    `json:"hypoallergenic"`
	} `json:"breeds"`
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
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
