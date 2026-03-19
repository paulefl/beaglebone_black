package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"myproject/pkg/hal"
	"myproject/pkg/hal/config"
	"myproject/pkg/hal/loader"
	"net/http"
	"sync"
	"time"
)

var (
	hw    hal.HardwareDriver
	hwMu  sync.Mutex
	hwCfg *config.Config
)

func main() {
	hwCfg = config.LoadConfig()
	var err error
	hw, err = loader.NewDriver(hwCfg)
	if err != nil {
		log.Fatalf("HAL init: %v", err)
	}
	defer hw.Close()
	log.Printf("Driver: %s Backend: %s", hw.Name(), hw.Backend())
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/api/v1/bme280", bme280Handler).Methods("GET")
	r.HandleFunc("/api/v1/bme280/stream", bme280StreamHandler).Methods("GET")
	r.HandleFunc("/api/v1/gpio", gpioAllHandler).Methods("GET")
	r.HandleFunc("/api/v1/gpio/{pin}", gpioReadHandler).Methods("GET")
	r.HandleFunc("/api/v1/gpio/{pin}", gpioWriteHandler).Methods("POST")
	r.HandleFunc("/api/v1/uart/config", uartConfigHandler).Methods("POST")
	r.HandleFunc("/api/v1/uart/send", uartSendHandler).Methods("POST")
	r.HandleFunc("/api/v1/uart/receive", uartReceiveHandler).Methods("GET")
	r.HandleFunc("/api/v1/spi/transfer", spiTransferHandler).Methods("POST")
	r.HandleFunc("/api/v1/backend", backendHandler).Methods("POST")
	log.Printf("API läuft auf :5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
func cors(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok", "backend": string(hw.Backend()), "driver": hw.Name(), "arch": "armv7"})
}
func bme280Handler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	data, err := hw.BME280Read()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func bme280StreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-r.Context().Done():
			return
		case <-ticker.C:
			m, err := hw.BME280Read()
			if err != nil {
				continue
			}
			data, _ := json.Marshal(m)
			fmt.Fprintf(w, "data: %s\n\n", data)
			w.(http.Flusher).Flush()
		}
	}
}
func gpioAllHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	pins := []uint32{60, 61, 66, 67}
	var results []interface{}
	for _, pin := range pins {
		data, err := hw.GPIORead(pin)
		if err == nil {
			results = append(results, data)
		}
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"pins": results})
}
func gpioReadHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	var pin uint32
	fmt.Sscanf(mux.Vars(r)["pin"], "%d", &pin)
	data, err := hw.GPIORead(pin)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func gpioWriteHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	var pin uint32
	fmt.Sscanf(mux.Vars(r)["pin"], "%d", &pin)
	var req struct {
		Value int `json:"value"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	if err := hw.GPIOWrite(pin, req.Value); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"pin": pin, "value": req.Value, "status": "ok"})
}
func uartConfigHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	var req struct {
		Port string `json:"port"`
		Baud uint32 `json:"baud"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	if err := hw.UARTOpen(req.Port, req.Baud); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
func uartSendHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	var req struct {
		Data []byte `json:"data"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	n, err := hw.UARTWrite(req.Data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"bytes_sent": n})
}
func uartReceiveHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	data, err := hw.UARTRead(1000)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func spiTransferHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	var req struct {
		Device string `json:"device"`
		Speed  uint32 `json:"speed"`
		TX     []byte `json:"tx"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	data, err := hw.SPITransfer(req.Device, req.Speed, req.TX)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func backendHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	var req struct {
		Backend string `json:"backend"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	cfg := &config.Config{}
	*cfg = *hwCfg
	cfg.Backend = req.Backend
	newHW, err := loader.NewDriver(cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hwMu.Lock()
	hw.Close()
	hw = newHW
	hwMu.Unlock()
	json.NewEncoder(w).Encode(map[string]string{"backend": string(newHW.Backend()), "status": "ok"})
}
