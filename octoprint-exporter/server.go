package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	yaml "gopkg.in/yaml.v2"
)

// Structure for our config.yaml
type Conf struct {
	Listen    string
	Apikey    string
	Probetime int
	Octopi    string
}

var c Conf

var configpath = flag.String("config", "config.yaml", "Choose your config File")

func main() {
	// Parse Configs
	flag.Parse()
	c.getConf()
	// Register Metrics
	jobinfo := newJobCollector()
	prometheus.MustRegister(jobinfo)

	// Start the Prometheus HTTP Server
	log.Println("Started HTTP Server: " + c.Listen)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(c.Listen, nil))
}

func (c *Conf) getConf() *Conf {

	yamlFile, err := ioutil.ReadFile(*configpath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		os.Exit(255)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
