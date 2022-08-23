package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	gocurrency "golang.org/x/text/currency"
	"time"
)

type CurrencyPair struct {
	CurrencyFrom string     `json:"currency_from" db:"currency_from"`
	CurrencyTo   string     `json:"currency_to" db:"currency_to"`
	Well         float64    `json:"well" db:"well"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}

type CurrencyPairs []*CurrencyPair

func (pairs CurrencyPairs) MapByCurrency() CurrencyExchangeMap {
	currencyMap := CurrencyExchangeMap{}
	for _, pair := range pairs {
		currencyTo, ok := currencyMap[pair.CurrencyFrom]
		if !ok {
			var to []string
			currencyMap[pair.CurrencyFrom] = append(to, pair.CurrencyTo)
			continue
		}

		currencyMap[pair.CurrencyFrom] = append(currencyTo, pair.CurrencyTo)
	}

	return currencyMap
}

type CurrencyExchangeMap map[string][]string

type CurrencyPairCreateInput struct {
	CurrencyFrom string `json:"currencyFrom"`
	CurrencyTo   string `json:"currencyTo"`
}

func (i *CurrencyPairCreateInput) Validate() error {
	err := validation.ValidateStruct(
		i,
		validation.Field(&i.CurrencyFrom, validation.Required),
		validation.Field(&i.CurrencyTo, validation.Required),
	)
	if err != nil {
		return err
	}

	_, err = gocurrency.ParseISO(i.CurrencyFrom)
	if err != nil {
		return err
	}

	_, err = gocurrency.ParseISO(i.CurrencyTo)
	if err != nil {
		return err
	}

	if err != nil {
		// TODO: Add input error
		return err
	}
	return nil
}

type CurrencyPairExchangeInput struct {
	CurrencyFrom string  `json:"currencyFrom"`
	CurrencyTo   string  `json:"currencyTo"`
	Value        float64 `json:"value"`
}
