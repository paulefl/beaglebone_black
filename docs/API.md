# REST API Dokumentation

Base URL: `http://192.168.7.2:5000`

## Health

```
GET /health
```
```json
{
  "status":  "ok",
  "backend": "auto",
  "driver":  "Auto Driver (C→Rust)",
  "arch":    "armv7"
}
```

## BME280

```
GET /api/v1/bme280
GET /api/v1/bme280/stream    (SSE)
```
```json
{
  "temperature": 23.45,
  "humidity":    48.72,
  "pressure":    1013.25,
  "altitude":    89.3,
  "timestamp":   "2026-03-17T10:00:00Z",
  "backend":     "c"
}
```

## GPIO

```
GET  /api/v1/gpio
GET  /api/v1/gpio/{pin}
POST /api/v1/gpio/{pin}      Body: {"value": 0|1}
POST /api/v1/gpio/{pin}/export
POST /api/v1/gpio/{pin}/direction   Body: {"direction": "in"|"out"}
```

## UART

```
POST /api/v1/uart/config     Body: {"port": "/dev/ttyO1", "baud": 115200}
POST /api/v1/uart/send       Body: {"data": [bytes]}
GET  /api/v1/uart/receive?timeout=1000
```

## SPI

```
POST /api/v1/spi/config      Body: {"device": "/dev/spidev0.0", "speed": 1000000, "mode": 0}
POST /api/v1/spi/transfer    Body: {"device": "...", "speed": ..., "tx": [bytes]}
```
```json
{"rx_buf": [1, 2, 3], "length": 3, "backend": "rust"}
```

## Backend

```
POST /api/v1/backend         Body: {"backend": "c"|"rust"|"auto"}
```

## Fehlercodes

| HTTP | Bedeutung |
|---|---|
| 200 | OK |
| 500 | Hardware-Fehler (I2C, GPIO, etc.) |
