package config

import "github.com/tkanos/gonfig"

type Config struct {
	Port                  int
	AlphaVantageApiKey    string
	MarketDataProviderUrl string
	PgHostname            string
	PgPort                int
	PgDB                  string
	PgPassword            string
	PgUser                string
}

type Configuration struct {
	Config
}

var configuration *Configuration

func GetConfiguration() *Configuration {
	if configuration != nil {
		return configuration
	}
	config := Configuration{}
	err := gonfig.GetConf("configuration.json", &config)
	if err != nil {
		panic(err)
	}
	configuration = &config
	return configuration
}


