package config

import (
	"flag"
	"os"
	"testing"
)

func resetFlags() {
	// flag.Parse() kann nur einmal auf dem globalen FlagSet aufgerufen werden.
	// Für Tests ein frisches FlagSet setzen.
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}

func TestLoadConfig_Defaults(t *testing.T) {
	resetFlags()
	cfg := LoadConfig()
	if cfg.Backend != "auto" { t.Errorf("Backend default: %q", cfg.Backend) }
	if cfg.BME280Addr != 0x76 { t.Errorf("BME280Addr: %x", cfg.BME280Addr) }
	if cfg.UARTBaud != 115200 { t.Errorf("UARTBaud: %d", cfg.UARTBaud) }
	if cfg.SPISpeed != 1000000 { t.Errorf("SPISpeed: %d", cfg.SPISpeed) }
	if cfg.Timeout == 0 { t.Error("Timeout nicht gesetzt") }
}

func TestLoadConfig_EnvOverride(t *testing.T) {
	resetFlags()
	t.Setenv("HW_BACKEND", "rust")
	cfg := LoadConfig()
	if cfg.Backend != "rust" { t.Errorf("HW_BACKEND override: %q", cfg.Backend) }
}
