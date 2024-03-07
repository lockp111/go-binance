package binance

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
		endpoint: "/api/v3/klines",
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
	data, err := s.c.callAPI(ctx, r, opts...)
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
		arr, _ = item.Array()
		if len(arr) < 11 {
			err = fmt.Errorf("invalid kline response")
			return []*Kline{}, err
		}
		kline := &Kline{}
		kline.OpenTime, _ = item.Index(0).Int64()
		kline.Open, _ = item.Index(1).String()
		kline.High, _ = item.Index(2).String()
		kline.Low, _ = item.Index(3).String()
		kline.Close, _ = item.Index(4).String()
		kline.Volume, _ = item.Index(5).String()
		kline.CloseTime, _ = item.Index(6).Int64()
		kline.QuoteAssetVolume, _ = item.Index(7).String()
		kline.TradeNum, _ = item.Index(8).Int64()
		kline.TakerBuyBaseAssetVolume, _ = item.Index(9).String()
		kline.TakerBuyQuoteAssetVolume, _ = item.Index(10).String()
		res[i] = kline
	}
	return res, nil
}

// Kline define kline info
type Kline struct {
	OpenTime                 int64  `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	TradeNum                 int64  `json:"tradeNum"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}
