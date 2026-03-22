#!/bin/sh
echo "=== Container gestartet ==="
echo "Hostname: $(hostname)"
echo "IP-Adresse: $(hostname -i)"
echo "Interfaces:"
ip a || true
echo "==========================="

# Starte den eigentlichen CMD-Prozess
exec "$@"