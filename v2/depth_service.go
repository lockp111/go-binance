package binance

import (
	"context"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// DepthService show depth info
type DepthService struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *DepthService) Symbol(symbol string) *DepthService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *DepthService) Limit(limit int) *DepthService {
	s.limit = &limit
	return s
}

// Do send request
func (s *DepthService) Do(ctx context.Context, opts ...RequestOption) (res *DepthResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/depth",
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	j, err := common.GetJSON(data)
	if err != nil {
		return nil, err
	}
	res = new(DepthResponse)
	res.LastUpdateID, _ = j.Get("lastUpdateId").Int64()
	bidsArr, _ := j.Get("bids").Array()
	bidsLen := len(bidsArr)
	res.Bids = make([]Bid, bidsLen)
	for i := 0; i < bidsLen; i++ {
		item := j.Get("bids").Index(i)
		price, _ := item.Index(0).String()
		quantity, _ := item.Index(1).String()
		res.Bids[i] = Bid{
			Price:    price,
			Quantity: quantity,
		}
	}
	asksArr, _ := j.Get("asks").Array()
	asksLen := len(asksArr)
	res.Asks = make([]Ask, asksLen)
	for i := 0; i < asksLen; i++ {
		item := j.Get("asks").Index(i)
		price, _ := item.Index(0).String()
		quantity, _ := item.Index(1).String()
		res.Asks[i] = Ask{
			Price:    price,
			Quantity: quantity,
		}
	}
	return res, nil
}

// DepthResponse define depth info with bids and asks
type DepthResponse struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
}

// Ask is a type alias for PriceLevel.
type Ask = common.PriceLevel

// Bid is a type alias for PriceLevel.
type Bid = common.PriceLevel
