package main

import (
	"fmt"
	"github.com/MichaelS11/go-dht"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	prefix              = "temp_hum_"
	temperatureGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: prefix + "temperature_celsius",
		Help: "Current temperature in celsius for labeled room",
	},
		[]string{"location"},
	)
	humidityGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: prefix + "humidity_percent",
		Help: "Current humidity in percent for labeled room",
	},
		[]string{"location"},
	)
)

func getTempHum(location string) {
	for {
		err := dht.HostInit()
		if err != nil {
			fmt.Println("HostInit error:", err)
			return
		}

		sensor, err := dht.NewDHT("GPIO2", dht.Celsius, "")
		if err != nil {
			fmt.Println("NewDHT error:", err)
			return
		}

		humidity, temperature, err := sensor.ReadRetry(11)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}

		temperatureGaugeVec.WithLabelValues(location).Set(temperature)
		humidityGaugeVec.WithLabelValues(location).Set(humidity)

		log.Printf("timestamp: %v\n", time.Now())
		log.Printf("humidity: %v\n", humidity)
		log.Printf("temperature: %v\n", temperature)

		time.Sleep(15 * time.Second)
		temperatureGaugeVec.Reset()
		humidityGaugeVec.Reset()
	}
}

func init() {
	prometheus.MustRegister(temperatureGaugeVec)
	prometheus.MustRegister(humidityGaugeVec)
}

func main() {
	var location string
	if locationEnv := os.Getenv("LOCATION"); locationEnv != "" {
		location = locationEnv
		log.Printf("Used location is: " + location)
	} else {
		log.Fatal("LOCATION environment variable is required")
	}

	http.Handle("/metrics", promhttp.Handler())
	go getTempHum(location)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
