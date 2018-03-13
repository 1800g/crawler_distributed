package client

import (
	"net/rpc"

	"github.com/1800g/crawler_distributed/crawler/engine"
	"github.com/1800g/crawler_distributed/crawler_distributed/config"
	"github.com/1800g/crawler_distributed/crawler_distributed/worker"
)

func CreateProcessor(
	clientChan chan *rpc.Client) engine.Processor {

	return func(
		req engine.Request) (
		engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc,
			sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult),
			nil
	}
}
