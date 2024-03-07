package binance

import (
	"context"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// PingService ping server
type PingService struct {
	c *Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ping",
	}
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// ServerTimeService get server time
type ServerTimeService struct {
	c *Client
}

// Do send request
func (s *ServerTimeService) Do(ctx context.Context, opts ...RequestOption) (serverTime int64, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/time",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return 0, err
	}
	j, err := common.GetJSON(data)
	if err != nil {
		return 0, err
	}

	ts, err := j.Get("serverTime").Int64()
	if err == nil {
		return ts, nil
	}
	return j.Int64()
}

// SetServerTimeService set server time
type SetServerTimeService struct {
	c *Client
}

// Do send request
func (s *SetServerTimeService) Do(ctx context.Context, opts ...RequestOption) (timeOffset int64, err error) {
	serverTime, err := s.c.NewServerTimeService().Do(ctx)
	if err != nil {
		return 0, err
	}
	timeOffset = currentTimestamp() - serverTime
	s.c.TimeOffset = timeOffset
	return timeOffset, nil
}
