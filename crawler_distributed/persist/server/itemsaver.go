package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/1800g/crawler_distributed/crawler/config"
	"github.com/1800g/crawler_distributed/crawler_distributed/persist"
	"github.com/1800g/crawler_distributed/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
