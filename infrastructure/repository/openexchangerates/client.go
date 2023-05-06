package openexchangerates

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
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

func (c *Client) do(method string, url *url.URL, body io.Reader) (*http.Response, error) {
	query := url.Query()
	query.Set("app_id", c.appID)
	url.RawQuery = query.Encode()

	request, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	return c.httpClient.Do(request)
}

func (c *Client) get(url *url.URL) (*http.Response, error) {
	return c.do(http.MethodGet, url, nil)
}

type GetCurrenciesResponse map[string]string

func (c *Client) GetCurrencies() (GetCurrenciesResponse, error) {
	url, err := c.buildURL("/currencies.json")
	if err != nil {
		return nil, err
	}

	response, err := c.get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

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
	url, err := c.buildURL("/latest.json")
	if err != nil {
		return nil, err
	}

	response, err := c.get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var body ResponseError

		if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
			return nil, err
		}

		return nil, errors.WithMessage(body, "failed to list exchange rates")
	}

	var body GetLatestExchangeRatesResponse

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
