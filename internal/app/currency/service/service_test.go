package service

import (
	"context"
	"currency_api/internal/app/currency/models"
	"currency_api/internal/app/currency/repository"
	"currency_api/internal/app/currency/repository/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/jaswdr/faker"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

func NewCurrencyPair(t *testing.T) *models.CurrencyPair {
	t.Helper()
	f := faker.New()
	updatedAt := f.Time().Time(time.Now()).UTC()
	return &models.CurrencyPair{
		CurrencyFrom: f.Currency().Code(),
		CurrencyTo:   f.Currency().Code(),
		Well:         float64(f.RandomDigit()),
		UpdatedAt:    &updatedAt,
	}
}

func NewCurrencyPairList(t *testing.T) models.CurrencyPairs {
	t.Helper()
	f := faker.New()

	var currencyPairs models.CurrencyPairs

	for i := 1; i <= rand.Intn(20); i++ {
		updatedAt := f.Time().Time(time.Now()).UTC()
		currencyPairs = append(currencyPairs, &models.CurrencyPair{
			CurrencyFrom: f.Currency().Code(),
			CurrencyTo:   f.Currency().Code(),
			Well:         float64(f.RandomDigit()),
			UpdatedAt:    &updatedAt,
		})
	}

	return currencyPairs
}

// CurrencyPairTestsSuite - набор тестов валютной пары :)
type CurrencyPairTestsSuite struct {
	suite.Suite

	currencyPairMock *mock.CurrencyPairMock
	service          *Service
}

func (ts *CurrencyPairTestsSuite) SetupTest() {
	mockController := minimock.NewController(ts.T())
	ts.currencyPairMock = mock.NewCurrencyPairMock(mockController)
	ts.service = New(&repository.Repository{
		Pair: ts.currencyPairMock,
	}, "")
}

func (ts *CurrencyPairTestsSuite) clear() {
	ts.currencyPairMock.MinimockCreateDone()
	ts.currencyPairMock.MinimockFinish()
}

func TestCurrencyPairs(t *testing.T) {
	suite.Run(t, new(CurrencyPairTestsSuite))
}

func (ts *CurrencyPairTestsSuite) TestCreateSuccess() {
	defer ts.clear()

	ctx := context.Background()
	randomCurrencyPair := NewCurrencyPair(ts.T())
	want := randomCurrencyPair

	ts.currencyPairMock.CreateMock.Return(want, nil)

	actual, err := ts.service.Create(ctx, &models.CurrencyPairCreateInput{
		CurrencyFrom: randomCurrencyPair.CurrencyFrom,
		CurrencyTo:   randomCurrencyPair.CurrencyTo,
	})
	ts.Require().NoError(err)
	ts.Require().NotNil(actual)
	ts.Require().Equal(*want, *actual)
}

func (ts *CurrencyPairTestsSuite) TestCreateError() {
	defer ts.clear()

	ctx := context.Background()

	ts.currencyPairMock.CreateMock.Return(nil, errors.New("some err"))

	actual, err := ts.service.Create(ctx, &models.CurrencyPairCreateInput{})
	ts.Require().Error(err)
	ts.Require().Nil(actual)
}

func (ts *CurrencyPairTestsSuite) TestGetSuccess() {
	defer ts.clear()

	ctx := context.Background()
	randomCurrencyPair := NewCurrencyPair(ts.T())

	want := randomCurrencyPair

	ts.currencyPairMock.GetMock.Return(want, nil)

	actual, err := ts.service.Get(ctx, randomCurrencyPair.CurrencyFrom, randomCurrencyPair.CurrencyTo)
	ts.Require().NoError(err)
	ts.Require().NotNil(actual)
	ts.Require().Equal(*want, *actual)
}

/*func (ts *CurrencyPairTestsSuite) TestGetError() {
	defer ts.clear()

	ctx := context.Background()

	ts.currencyPairMock.GetMock.Return(nil, errors.New("some err"))

	actual, err := ts.service.Get(ctx, "", "")
	ts.Require().Error(err)
	ts.Require().Nil(actual)
}*/

func (ts *CurrencyPairTestsSuite) TestListSuccess() {
	defer ts.clear()

	ctx := context.Background()
	randomCurrencyPairList := NewCurrencyPairList(ts.T())

	want := randomCurrencyPairList

	ts.currencyPairMock.ListMock.Return(want, nil)

	actual, err := ts.service.List(ctx)
	ts.Require().NoError(err)
	ts.Require().NotNil(actual)
	ts.Require().Equal(want, actual)
}

func (ts *CurrencyPairTestsSuite) TestGetError() {
	defer ts.clear()

	ctx := context.Background()
	wanrErr := errors.New("that is bad")

	ts.currencyPairMock.GetMock.When(ctx, "bad", "bad").Then(nil, wanrErr)
	/*ts.currencyPairMock.GetMock.When(ctx, "g", "g").Then(nil, errors.New("that is bad"))
	ts.currencyPairMock.GetMock.When(ctx, "z", "hb").Then(nil, errors.New("that is bad"))
	ts.currencyPairMock.GetMock.When(ctx, "fdss", "reg").Then(nil, errors.New("that is bad"))
	ts.currencyPairMock.GetMock.When(ctx, "rebw", "bgerw").Then(nil, errors.New("that is bad"))*/

	tests := []struct {
		name    string
		args    []string
		wantErr error
		want    *models.CurrencyPair
	}{
		{
			name: "simple",
			args: []string{
				"bad", "bad",
			},
			wantErr: wanrErr,
			want:    nil,
		},
	}

	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actual, err := ts.service.Get(ctx, tt.args[0], tt.args[1])
			ts.Require().Equal(tt.wantErr, err)
			ts.Require().Equal(tt.want, actual)
		})
	}
}
