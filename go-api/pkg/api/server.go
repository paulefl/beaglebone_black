package api

import (
	"net/http"
	"time"
)

// NewHTTPServer erstellt einen http.Server mit konfigurierten Timeouts (SWR-006).
// Verhindert Ressourcenerschöpfung durch hängende Verbindungen auf eingebetteten Systemen.
func NewHTTPServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
