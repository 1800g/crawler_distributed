package main

import (
	"github.com/1800g/crawler_distributed/crawler/config"
	"github.com/1800g/crawler_distributed/crawler/engine"
	"github.com/1800g/crawler_distributed/crawler/persist"
	"github.com/1800g/crawler_distributed/crawler/scheduler"
	"github.com/1800g/crawler_distributed/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.starter.url.here",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
