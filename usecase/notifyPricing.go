package usecase

import (
	"context"
	"fmt"
	"net/http"
	"surge/config"
	natsBroker "surge/pkg/nats"
	redisWrapper "surge/pkg/redis"
)

type NotifyInterface interface {
}

type Notifier struct {
}

type districtCoefficient struct {
	coe float64
}

type pricingNotifyRequest struct {
	districtId string
	coe        float64
}

// pricing service should reply Acknowledge if code is 200
// surge service considers it as a success
type pricingNotifyResponse struct {
	code    int64
	message string
}

// NotifyPricing notifies pricing service when request coe is not equal with current coe in pricing service
// it doesn't handle errors cause it runs on every request so if it fails in one request it will be called
// with the next request. it saves pricing last coe for every district and if current coe in pricing is equal to
// request coe it will not send any request to pricing service for network efficiency reason.
func NotifyPricing(ctx context.Context, coe float64, districtId string) {
	key := fmt.Sprintf("%s:%s", config.GetCnf().ServiceName, districtId)
	var currentCoe *districtCoefficient
	err := redisWrapper.Get(ctx, key, currentCoe)
	if err == nil && currentCoe.coe == coe {
		return
	}

	prReq := &pricingNotifyRequest{
		districtId: districtId,
		coe:        coe,
	}
	var prRes *pricingNotifyResponse
	cnf := config.GetCnf()
	for i := 0; i < config.GetCnf().NotifyPricingRetryCount; i++ {
		err = natsBroker.Request(ctx, cnf.PricingSubject, uint16(cnf.NotifyPricingTimeout), prReq, prRes)
		if err == nil {
			break
		}
	}

	if err == nil {
		// pricing service was informed successfully ,so we set current coe in redis for district id
		if prRes.code == http.StatusOK {
			currentCoe = &districtCoefficient{
				coe: coe,
			}
			err = redisWrapper.Set(ctx, key, currentCoe, redisWrapper.NoTTLTime)
		}
	}
}
