#!/usr/bin/env bash
# =============================================================================
# setup-dev.sh — Entwicklungsumgebung für BeagleBone Black vorbereiten
#
# Installiert: Go 1.23, ARM Cross-Compiler, Rust-Toolchain, Python-Test-Deps
# Führt aus:   go mod tidy, cargo fetch
# =============================================================================
set -euo pipefail

GREEN='\033[0;32m'; BLUE='\033[0;34m'; YELLOW='\033[1;33m'; RED='\033[0;31m'; NC='\033[0m'
log()     { echo -e "${BLUE}[INFO]${NC}  $*"; }
success() { echo -e "${GREEN}[OK]${NC}    $*"; }
warn()    { echo -e "${YELLOW}[WARN]${NC}  $*"; }
error()   { echo -e "${RED}[ERROR]${NC} $*"; exit 1; }

GO_VERSION="1.23.8"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

# ── Go ────────────────────────────────────────────────────────────────────────
install_go() {
    if command -v go &>/dev/null; then
        local current
        current=$(go version | awk '{print $3}' | sed 's/go//')
        if [[ "$current" == "$GO_VERSION"* ]]; then
            success "Go $current bereits installiert"
            return
        fi
        warn "Go $current gefunden, benötigt $GO_VERSION"
    fi

    log "Installiere Go $GO_VERSION..."
    local arch
    case $(uname -m) in
        x86_64)  arch="amd64" ;;
        aarch64) arch="arm64" ;;
        armv7l)  arch="armv6l" ;;
        *) error "Architektur nicht unterstützt: $(uname -m)" ;;
    esac

    local tarball="go${GO_VERSION}.linux-${arch}.tar.gz"
    curl -fsSL "https://go.dev/dl/${tarball}" -o "/tmp/${tarball}"
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf "/tmp/${tarball}"
    rm -f "/tmp/${tarball}"

    # PATH für aktuelle Session
    export PATH="/usr/local/go/bin:$PATH"

    # Permanent in Shell-Config eintragen
    for rc in ~/.bashrc ~/.zshrc; do
        if [[ -f "$rc" ]] && ! grep -q '/usr/local/go/bin' "$rc"; then
            echo 'export PATH="/usr/local/go/bin:$PATH"' >> "$rc"
        fi
    done

    success "Go $(go version | awk '{print $3}') installiert"
}

# ── ARM Cross-Compiler ────────────────────────────────────────────────────────
install_arm_toolchain() {
    if command -v arm-linux-gnueabihf-gcc &>/dev/null; then
        success "ARM Cross-Compiler bereits installiert"
        return
    fi
    log "Installiere ARM Cross-Compiler..."
    sudo apt-get update -qq
    sudo apt-get install -y --no-install-recommends \
        gcc-arm-linux-gnueabihf \
        libc6-dev-armhf-cross \
        make
    success "ARM Cross-Compiler installiert"
}

# ── Rust + Cross ──────────────────────────────────────────────────────────────
install_rust() {
    if command -v cargo &>/dev/null; then
        success "Rust bereits installiert ($(cargo --version))"
    else
        log "Installiere Rust..."
        curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y --quiet
        source "$HOME/.cargo/env"
        success "Rust $(cargo --version) installiert"
    fi

    source "$HOME/.cargo/env" 2>/dev/null || true

    log "Füge ARM-Target hinzu..."
    rustup target add armv7-unknown-linux-musleabihf

    if ! command -v cross &>/dev/null; then
        log "Installiere cross..."
        cargo install cross --quiet
        success "cross installiert"
    else
        success "cross bereits installiert"
    fi

    if ! command -v cbindgen &>/dev/null; then
        log "Installiere cbindgen..."
        cargo install cbindgen --quiet
        success "cbindgen installiert"
    else
        success "cbindgen bereits installiert"
    fi
}

# ── Python Test-Dependencies ──────────────────────────────────────────────────
install_python_deps() {
    log "Installiere Python Test-Dependencies..."
    if command -v pip3 &>/dev/null; then
        pip3 install --quiet pytest requests pytest-json-report
        success "Python Deps installiert"
    else
        warn "pip3 nicht gefunden — Python Tests nicht verfügbar"
    fi
}

# ── Go Module ─────────────────────────────────────────────────────────────────
go_mod_tidy() {
    log "Führe go mod tidy aus..."
    (cd "$REPO_ROOT/go-api" && go mod tidy)
    success "go.sum aktualisiert"
}

# ── Verify ────────────────────────────────────────────────────────────────────
verify() {
    log "Verifiziere Setup..."
    (cd "$REPO_ROOT/go-api" && go vet ./...) && success "go vet OK"
    (cd "$REPO_ROOT/go-api" && go test ./pkg/hal/... -count=1 -timeout 30s) && success "Tests OK"
}

# ── Main ──────────────────────────────────────────────────────────────────────
main() {
    echo ""
    echo -e "${BLUE}╔══════════════════════════════════════════╗${NC}"
    echo -e "${BLUE}║   BeagleBone Black — Dev Setup           ║${NC}"
    echo -e "${BLUE}╚══════════════════════════════════════════╝${NC}"
    echo ""

    case "${1:-all}" in
        all)
            install_go
            install_arm_toolchain
            install_rust
            install_python_deps
            go_mod_tidy
            verify
            ;;
        go)           install_go; go_mod_tidy ;;
        arm)          install_arm_toolchain ;;
        rust)         install_rust ;;
        python)       install_python_deps ;;
        tidy)         go_mod_tidy ;;
        verify)       verify ;;
        *)
            echo "Verwendung: $0 [all|go|arm|rust|python|tidy|verify]"
            echo ""
            echo "  all     — Alles installieren (Standard)"
            echo "  go      — Go $GO_VERSION + go mod tidy"
            echo "  arm     — ARM Cross-Compiler"
            echo "  rust    — Rust + cross + cbindgen"
            echo "  python  — pytest + requests"
            echo "  tidy    — Nur go mod tidy"
            echo "  verify  — go vet + Tests"
            exit 1
            ;;
    esac

    echo ""
    success "Setup abgeschlossen!"
    echo ""
}

main "$@"
