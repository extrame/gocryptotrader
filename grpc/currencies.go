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
