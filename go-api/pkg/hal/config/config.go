package config

import (
	"flag"
	"os"
	"strings"
	"time"
)

type Config struct {
	Backend    string
	I2CBus     string
	BME280Addr uint8
	UARTPort   string
	UARTBaud   uint32
	SPIDevice  string
	SPISpeed   uint32
	Timeout    time.Duration
}

func LoadConfig() *Config {
	backend := flag.String("backend", "auto", "c | rust | auto")
	uartPort := flag.String("uart", "/dev/ttyO1", "UART Port")
	i2cBus := flag.String("i2c", "/dev/i2c-1", "I2C Bus")
	flag.Parse()
	if env := os.Getenv("HW_BACKEND"); env != "" {
		*backend = env
	}
	return &Config{
		Backend:    strings.ToLower(*backend),
		I2CBus:     *i2cBus,
		BME280Addr: 0x76,
		UARTPort:   *uartPort,
		UARTBaud:   115200,
		SPIDevice:  "/dev/spidev0.0",
		SPISpeed:   1000000,
		Timeout:    10 * time.Second,
	}
}
