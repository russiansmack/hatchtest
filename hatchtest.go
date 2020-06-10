package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const base = "https://api.thecatapi.com/v1"

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient(apiKey string) *Client {
	c := &Client{}
	c.httpClient = &http.Client{}
	c.apiKey = apiKey

	return c
}

//create request
func (c *Client) makeRequest(endpoint string, method string) (*http.Request, error) {

	//Debug endpoint
	//fmt.Println(endpoint)

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
func (c *Client) runRequest(req *http.Request, results interface{}) error {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	//Debug json
	//body, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))

	err = json.NewDecoder(response.Body).Decode(results)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

type ImageSearchOptions struct {
	Size      string `url:"size"`
	MimeTypes string `url:"mime_types"`
	Order     string `url:"order"`
	Page      int    `url:"page"`
	Limit     int    `url:"limit"`
}

//Fetching cat images
func (c *Client) GetImageSearch(options ImageSearchOptions) (*ImageResults, error) {

	//Default settings for GET parameters
	if options == (ImageSearchOptions{}) {
		options = ImageSearchOptions{"med", "jpg,gif,png", "RANDOM", 1, 1}
	}

	//GET params
	params, _ := query.Values(options)

	req, err := c.makeRequest("/images/search?"+params.Encode(), http.MethodGet)
	if err != nil {
		return nil, err
	}

	v := &ImageResults{}

	err = c.runRequest(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

//Fetching categories
func (c *Client) GetCategories() (*CategoryResults, error) {

	req, err := c.makeRequest("/categories", http.MethodGet)
	if err != nil {
		return nil, err
	}

	v := &CategoryResults{}

	err = c.runRequest(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func main() {
	c := NewClient("7a1768b0-1600-4c55-9769-83721284ab92")

	//*
	options := ImageSearchOptions{}
	options.Limit = 5
	options.Order = "ASC"
	options.Page = 0

	images, err := c.GetImageSearch(options)
	if err != nil {
		fmt.Println("badjuju")
	}
	fmt.Println(images)
	/*/
	categories, err := c.GetCategories()
	if err != nil {
		fmt.Println("badjuju")
	}
	fmt.Println(categories)*/
}

type CategoryResults []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ImageResults []struct {
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
