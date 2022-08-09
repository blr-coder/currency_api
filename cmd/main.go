package main

import (
	"context"
	"currency_api/config"
	"currency_api/internal/models"
	"currency_api/internal/repository"
	"currency_api/internal/repository/repository_interfaces"
	"currency_api/internal/storage/postgres"
	"currency_api/internal/transport/rest"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {

	appConfig := config.NewConfig()

	psqlDB, err := postgres.NewDB(
		appConfig.Postgres.Host,
		appConfig.Postgres.Port,
		appConfig.Postgres.Name,
		appConfig.Postgres.User,
		appConfig.Postgres.Password,
	)
	if err != nil {
		panic(fmt.Errorf("failed to init postgres: %v", err.Error()))
	}

	logger := logrus.StandardLogger()

	r := repository.NewCurrencyPairRepository(psqlDB, logger)

	checkManager := NewCheckManager(r)

	go checkManager.checkRates()

	h := rest.NewCurrencyHandler(r)

	app := rest.NewCurrencyApp(h)

	err = app.Run(appConfig.Port)
	if err != nil {
		panic(fmt.Errorf("failed to run app: %v", err.Error()))
	}
}

type CheckManager struct { // Вероятно тоже должен быть свой интерфейс
	sync.WaitGroup
	currencyRepo repository_interfaces.CurrencyPairRepositoryI
}

func NewCheckManager(currencyRepo repository_interfaces.CurrencyPairRepositoryI) *CheckManager {
	return &CheckManager{currencyRepo: currencyRepo}
}

func (m *CheckManager) checkRates() {

	// ???
	ctx := context.TODO()

	listCurrencyPairs, err := m.currencyRepo.List(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	currencyMap := listCurrencyPairs.MapByCurrency()

	fmt.Println("cMap:", currencyMap)

	//makeRequest("USD", []string{"EUR", "RUB"})
	for tick := range time.Tick(20 * time.Second) {
		fmt.Println("Tick", tick.UTC().Format(time.RFC3339))

		// TODO: "Free plan is limited to 1 request per second."
		var secondsForSleep int64

		//wg := &sync.WaitGroup{}
		for f, t := range currencyMap {

			// Костыльная задержка для отправки запросов т.к. API банит больше 1 запроса в секунду
			secondsForSleep++

			m.Add(1)
			go m.makeRequest(f, t, time.Duration(secondsForSleep))
			m.Wait()
		}
	}
}

func (m *CheckManager) makeRequest(currencyFrom string, currenciesTo []string, s time.Duration) {

	//"Free plan is limited to 1 request per second."
	time.Sleep(s * time.Second)

	// TODO: Ключь от API естессно надо вынести в конфиг
	url := fmt.Sprintf(
		"https://exchange-rates.abstractapi.com/v1/live/?api_key=%s&base=%s&target=%s",
		"f0685cd6c6744cc686d60fa9dc6477c0",
		currencyFrom,
		strings.Join(currenciesTo, ","),
	)
	//resp, err := http.Get("https://exchange-rates.abstractapi.com/v1/live/?api_key=f0685cd6c6744cc686d60fa9dc6477c0&base=USD&target=EUR,RUB")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	aResp := &models.CurrencyExchangeInfo{}
	err = json.Unmarshal(body, aResp)
	if err != nil {
		log.Fatalln(err)
	}

	spew.Dump(aResp)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	// makeRequest выполняется в отдельной горутине в цикле... запросы в базу в цикле это фигово наверное...
	err = m.currencyRepo.UpdateCurrencyWell(context.TODO(), aResp)
	if err != nil {
		log.Fatalln(err)
	}
	m.Done()
}
