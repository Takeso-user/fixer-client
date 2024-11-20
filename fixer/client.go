package fixer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	API_URL = "http://data.fixer.io/api/"
)

type Client struct {
	client *http.Client
}

func (c Client) GetRate() (ResponseBody, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file with key")
	}
	var accessKey = os.Getenv("ACCESS_KEY")
	url := API_URL + "latest?access_key=" + accessKey
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ResponseBody{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return ResponseBody{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	var response ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&response)

	return response, err

}

func (c Client) CovertCcy(from, to string, amount float64) (ResponseBodyConversation, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file with key")
	}
	var accessKey = os.Getenv("ACCESS_KEY")
	url := fmt.Sprintf("%sconvert?access_key=%s&from=%s&to=%s&amount=%f", API_URL, accessKey, from, to, amount)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ResponseBodyConversation{}, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return ResponseBodyConversation{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}

	}(resp.Body)
	var response ResponseBodyConversation
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err

}

func (c Client) GetAllSymbols() (ResponseBodySymbols, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file with key")
	}
	var accessKey = os.Getenv("ACCESS_KEY")
	url := API_URL + "symbols?access_key=" + accessKey
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ResponseBodySymbols{}, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return ResponseBodySymbols{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	var response ResponseBodySymbols
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout cant be zero")
	}
	return &Client{
		client: &http.Client{
			Transport: &LogRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
			Timeout: timeout,
		},
	}, nil
}
