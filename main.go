package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/marpaia/graphite-golang"
	"github.com/taiyoh/go-cm160"
)

func main() {
	g, err := graphite.NewGraphite("localhost", 2003)
	if err != nil {
		panic(err)
	}

	cb := func(buf *cm160.Record) {
		fmt.Println(buf)

		var t time.Time
		if buf.IsLive {
			t = time.Now().UTC()
		} else {
			t = time.Date(buf.Year, time.Month(buf.Month), buf.Day, buf.Hour, buf.Minute, 0, 0, time.UTC)
		}

		m := graphite.Metric{Name: "electricity.amps", Value: strconv.FormatFloat(float64(buf.Amps), 'f', -1, 32), Timestamp: t.Unix()}
		err := g.SendMetric(m)
		if err != nil {
			panic(err)
		}
	}

	monitor := cm160.Open()
	monitor.Wait(cb)
}
