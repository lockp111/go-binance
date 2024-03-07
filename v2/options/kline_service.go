package options

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// KlinesService list klines
type KlinesService struct {
	c         *Client
	symbol    string
	interval  string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *KlinesService) Symbol(symbol string) *KlinesService {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *KlinesService) Interval(interval string) *KlinesService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *KlinesService) Limit(limit int) *KlinesService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *KlinesService) StartTime(startTime int64) *KlinesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *KlinesService) EndTime(endTime int64) *KlinesService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *KlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/klines",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Kline{}, err
	}
	j, err := common.GetJSON(data)
	if err != nil {
		return []*Kline{}, err
	}
	arr, _ := j.Array()
	num := len(arr)
	res = make([]*Kline, num)
	for i := 0; i < num; i++ {
		item := j.Index(i)
		m, _ := item.Map()
		if len(m) < 12 {
			err = fmt.Errorf("invalid kline response")
			return []*Kline{}, err
		}
		kline := &Kline{}
		kline.OpenTime, _ = item.Get("openTime").Int64()
		kline.Open, _ = item.Get("open").String()
		kline.High, _ = item.Get("high").String()
		kline.Low, _ = item.Get("low").String()
		kline.Close, _ = item.Get("close").String()
		kline.CloseTime, _ = item.Get("closeTime").Int64()
		kline.Amount, _ = item.Get("amount").String()
		kline.TakerAmount, _ = item.Get("takerAmount").String()
		kline.Volume, _ = item.Get("volume").String()
		kline.TakerVolume, _ = item.Get("takerVolume").String()
		kline.Interval, _ = item.Get("interval").String()
		kline.TradeCount, _ = item.Get("tradeCount").Int64()
		res[i] = kline
	}
	return res, nil
}

// Kline define kline info
type Kline struct {
	OpenTime    int64  `json:"openTime"`
	Open        string `json:"open"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Close       string `json:"close"`
	CloseTime   int64  `json:"closeTime"`
	Amount      string `json:"amount"`
	TakerAmount string `json:"takerAmount"`
	Volume      string `json:"volume"`
	TakerVolume string `json:"takerVolume"`
	Interval    string `json:"interval"`
	TradeCount  int64  `json:"tradeCount"`
}
