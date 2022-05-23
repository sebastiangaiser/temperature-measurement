# Simple room temperature and humidity measurement

This small tool uses a [library](github.com/MichaelS11/go-dht) to measure the room temperature and humidity.
The measured values are displayed on the console and also exported for Prometheus available under path `/metrics`.

## Used hardware

- Raspberry Pi 2b
- DHT22 sensor

## Used GPIO pins

[//]: # (TODO add picture)

## Compile and run

```shell
GOOS=linux GOARCH=arm GOARM=7 go build -o out/temp_hum
```

Before starting the binary you have to provide a location for the Prometheus exporter.
The location is used as label for the exported metrics.

```shell
export LOCATION=test ./temp_hum
```

### (Optional) Add systemd service

```shell
scp temp_hum.service pi@temp.local:/home/pi/temp_hum.service
ssh pi@raspi.local "sudo mv temp_hum.service /etc/systemd/system/"
ssh pi@raspi.local "sudo systemctl daemon-reload"
ssh pi@raspi.local "sudo systemctl enable temp_hum.service"
ssh pi@raspi.local "sudo systemctl start temp_hum.service"
```
