package openexchangerates

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client is an HTTP client capable of communicating with OpenExchangeRates API
type Client struct {
	httpClient *http.Client
	appID      string
	baseURL    string
}

func NewClient(httpClient *http.Client, appID string, baseURL string) *Client {
	return &Client{
		httpClient: httpClient,
		appID:      appID,
		baseURL:    baseURL,
	}
}

func (c *Client) buildURL(path string) (*url.URL, error) {
	return url.Parse(c.baseURL + path)
}

func (c *Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	url, err := c.buildURL(path)
	if err != nil {
		return nil, err
	}

	query := url.Query()
	query.Set("app_id", c.appID)

	url.RawQuery = query.Encode()

	return http.NewRequest(method, url.String(), body)
}

func (c *Client) newGetRequest(path string) (*http.Request, error) {
	return c.newRequest(http.MethodGet, path, nil)
}

type GetCurrenciesResponse map[string]string

func (c *Client) GetCurrencies() (GetCurrenciesResponse, error) {
	request, err := c.newGetRequest("/currencies.json")
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list currencies")
	}

	var body GetCurrenciesResponse

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

type GetLatestExchangeRatesResponse struct {
	// Base is the code of the base currency used to calculate the exchange rates
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

// GetLatestExchangeRates returns the relative value of a base currency in terms of a set of different currencies
func (c *Client) GetLatestExchangeRates() (*GetLatestExchangeRatesResponse, error) {
	request, err := c.newGetRequest("/latest.json")
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list exchange rates")
	}

	var body GetLatestExchangeRatesResponse

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
