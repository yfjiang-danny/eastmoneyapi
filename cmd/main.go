package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	math_rand "math/rand"
	"time"

	"github.com/yfjiang-danny/eastmoneyapi/client"
	"github.com/yfjiang-danny/eastmoneyapi/config"
)

var configPath string

func init() {
	math_rand.Seed(time.Now().Unix())
	log.SetFlags(log.Lshortfile)
}
func init() {
	flag.StringVar(&configPath, "config", "", "")
	flag.Parse()
	if configPath != "" {
		config.SetConfigPath(configPath)
	}
}

func main() {
	// z := service.NewZ513050Svc()
	// z.Start()

	c := client.NewEastMoneyClient(config.GetConfig().EastMoneyClientConfig)

	new, err := c.GetCanBuyNewStockList()
	if err != nil {
		panic(err)
	}

	res, err := c.SubmitBatTrade(new.GetSubmitBatTradeParams())
	if err != nil {
		panic(err)
	}
	str, _ := json.Marshal(res)
	fmt.Println(string(str))
}
