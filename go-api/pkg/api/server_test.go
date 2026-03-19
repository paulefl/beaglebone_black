// [SDOC_LINK: SWR-006]
package api_test

import (
	"net/http"
	"testing"
	"time"

	"myproject/pkg/api"
)

// Regression Test für Issue #21:
// NewHTTPServer muss einen http.Server mit konfigurierten Timeouts zurückgeben.
func TestNewHTTPServer_HasTimeouts(t *testing.T) {
	srv := api.NewHTTPServer(":5000", http.NewServeMux())

	if srv.ReadTimeout == 0 {
		t.Error("ReadTimeout ist 0 — muss > 0 sein")
	}
	if srv.WriteTimeout == 0 {
		t.Error("WriteTimeout ist 0 — muss > 0 sein")
	}
	if srv.IdleTimeout == 0 {
		t.Error("IdleTimeout ist 0 — muss > 0 sein")
	}
}

func TestNewHTTPServer_TimeoutValues(t *testing.T) {
	srv := api.NewHTTPServer(":5000", http.NewServeMux())

	if srv.ReadTimeout < 5*time.Second {
		t.Errorf("ReadTimeout zu kurz: %v (min 5s)", srv.ReadTimeout)
	}
	if srv.WriteTimeout < 10*time.Second {
		t.Errorf("WriteTimeout zu kurz: %v (min 10s)", srv.WriteTimeout)
	}
	if srv.IdleTimeout < 30*time.Second {
		t.Errorf("IdleTimeout zu kurz: %v (min 30s)", srv.IdleTimeout)
	}
}

func TestNewHTTPServer_AddrAndHandler(t *testing.T) {
	mux := http.NewServeMux()
	srv := api.NewHTTPServer(":5000", mux)

	if srv.Addr != ":5000" {
		t.Errorf("Addr: got %q, want %q", srv.Addr, ":5000")
	}
	if srv.Handler != mux {
		t.Error("Handler nicht korrekt gesetzt")
	}
}
