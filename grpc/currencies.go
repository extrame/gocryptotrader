package grpc

import (
	"time"

	ic "github.com/influxdata/influxdb/client/v2"
)

type InfluxContext interface {
	GetBpCfg() ic.BatchPointsConfig
	Write(ic.BatchPoints) error
}

func (cs *AllEnabledExchangeCurrencies) StoreInInflux(ctx InfluxContext) error {
	at := time.Now() //有延时，需要更正，至少应该放在server侧
	// write transactions
	currentBatchPoints, err := ic.NewBatchPoints(ctx.GetBpCfg())

	for _, c := range cs.ExchangeCurrencies {

		for _, p := range c.ExchangeValues {
			tags := make(map[string]string)
			tags["exchange"] = c.ExchangeName
			tags["pair_first"] = p.Pair.FirstCurrency
			tags["pair_second"] = p.Pair.SecondCurrency
			fields := map[string]interface{}{
				"price": p.Last,
			}
			var tmp *ic.Point
			if tmp, err = ic.NewPoint("price", tags, fields, at); err == nil {
				currentBatchPoints.AddPoint(tmp)
			}
		}
	}
	err = ctx.Write(currentBatchPoints)
	return err
}

func (rep *UpdateTickerReport) StoreInInflux(ctx InfluxContext) error {
	at := time.Now() //有延时，需要更正，至少应该放在server侧
	// write transactions
	currentBatchPoints, err := ic.NewBatchPoints(ctx.GetBpCfg())

	tags := make(map[string]string)
	tags["exchange"] = rep.ExchangeName
	tags["pair_first"] = rep.Price.Pair.FirstCurrency
	tags["pair_second"] = rep.Price.Pair.SecondCurrency
	fields := map[string]interface{}{
		"price": rep.Price.Last,
	}

	var tmp *ic.Point
	if tmp, err = ic.NewPoint("price", tags, fields, at); err == nil {
		currentBatchPoints.AddPoint(tmp)
	}

	err = ctx.Write(currentBatchPoints)
	return err
}
