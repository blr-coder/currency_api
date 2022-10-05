package exchange_rates

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const abstractAPIFormat = "https://exchange-rates.abstractapi.com/v1/live/?api_key=%s&base=%s&target=%s"

type Client struct {
	http     *http.Client
	apiToken string
}

func (c *Client) Close() error {
	return nil
}

func New(apiToken string) *Client {
	return &Client{
		http:     &http.Client{},
		apiToken: apiToken,
	}
}

type ExchangeRatesInfo struct {
	Base          string             `json:"base"`
	ExchangeRates map[string]float64 `json:"exchange_rates"`
}

func (c *Client) GetRates(ctx context.Context, from string, currenciesTo []string) (*ExchangeRatesInfo, error) {

	logrus.Info("GetRates for CURRENCY_FROM:", from)
	logrus.Info("GetRates for CURRENCIES_TO:", currenciesTo)

	url := fmt.Sprintf(
		abstractAPIFormat,
		c.apiToken,
		from,
		strings.Join(currenciesTo, ","),
	)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	exchangeInfo := &ExchangeRatesInfo{}
	err = json.Unmarshal(body, exchangeInfo)
	if err != nil {
		return nil, err
	}

	return exchangeInfo, nil
}
