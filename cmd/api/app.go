package main

import (
	"context"
	"currency_api/internal/app/config"
	"currency_api/internal/app/currency/repository"
	"currency_api/internal/app/currency/service"
	"currency_api/internal/app/currency/transport/rest"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	dbTimeout = 3 * time.Second
)

func runApp() error {

	appConfig, err := config.NewConfig("configs/dev_config.yaml")
	if err != nil {
		return err
	}

	db, err := sqlx.Open("postgres", appConfig.PostgresConnLink)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return err
	}

	r := repository.New(db)

	/*c := http.Client{}
	checkManager := NewCheckManager(r, &c)
	go checkManager.checkRates()*/

	s := service.New(r)

	api := fiber.New()

	h := rest.New(s, api)
	h.InitCurrencyRoutes()

	// TODO: Move to func
	/*exchangeClient := exchange_rates.New(appConfig.AbstractApiKey)
	listCurrencyPairs, err := r.Pair.List(ctx)
	if err != nil {
		if err != nil {
			return err
		}
	}
	currencyMap := listCurrencyPairs.MapByCurrency()

	for f, t := range currencyMap {
		rates, err := exchangeClient.GetRates(ctx, f, t)
		if err != nil {
			return err
		}

		err = r.Pair.UpdateCurrencyWell(ctx, rates)
		if err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}*/
	// TODO: Move to func

	return api.Listen(fmt.Sprintf(":%s", appConfig.Port))
}

// FIXME: Drop after testing

//type CheckManager struct { // Вероятно тоже должен быть свой интерфейс
//	sync.WaitGroup
//	repository *repository.Repository
//	httpClient *http.Client
//}
//
//func NewCheckManager(repository *repository.Repository, httpClient *http.Client) *CheckManager {
//	return &CheckManager{repository: repository, httpClient: httpClient}
//}
//
//func (m *CheckManager) checkRates() {
//
//	// FIXME: Как тут лучше поступить с контекстом?
//	ctx := context.TODO()
//
//	listCurrencyPairs, err := m.repository.Pair.List(ctx)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	currencyMap := listCurrencyPairs.MapByCurrency()
//
//	//makeRequest("USD", []string{"EUR", "RUB"})
//	for tick := range time.Tick(10 * time.Second) {
//		fmt.Println("Tick", tick.UTC().Format(time.RFC3339))
//
//		/*testM := map[string]float64{
//			"EUR": 0.009473,
//			"RUB": 0.009524,
//			"PLN": 0.045005,
//		}
//
//		testUPD := &models.CurrencyExchangeInfo{
//			Base:          "USD",
//			ExchangeRates: testM,
//		}
//
//		err := m.repository.Pair.UpdateCurrencyWell(ctx, testUPD)
//		if err != nil {
//			log.Fatalln(err)
//		}*/
//
//		// TODO: "Free plan is limited to 1 request per second."
//		var secondsForSleep int64
//
//		for f, t := range currencyMap {
//
//			// Костыльная задержка для отправки запросов т.к. API банит больше 1 запроса в секунду
//			secondsForSleep++
//
//			m.Add(1)
//			go m.makeRequest(f, t, time.Duration(secondsForSleep))
//			m.Wait()
//		}
//	}
//}

/*func (m *CheckManager) makeRequest(currencyFrom string, currenciesTo []string, s time.Duration) {

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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := m.httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	aResp := &exchange_rates.ExchangeRatesInfo{}
	err = json.Unmarshal(body, aResp)
	if err != nil {
		log.Fatalln(err)
	}

	spew.Dump(aResp)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	// makeRequest выполняется в отдельной горутине в цикле... запросы в базу в цикле это фигово наверное...
	err = m.repository.Pair.UpdateCurrencyWell(context.TODO(), aResp)
	if err != nil {
		log.Fatalln(err)
	}
	// FIXME: Верно ли тут используется WaitGroup?
	m.Done()
}*/
