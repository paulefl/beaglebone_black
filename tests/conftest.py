import os
import socket
import pytest

HOST = os.getenv("BEAGLE_HOST", "192.168.7.2")
_HOST_EXPLICIT = "BEAGLE_HOST" in os.environ


def _reachable(host: str, port: int = 5000, timeout: float = 2.0) -> bool:
    try:
        with socket.create_connection((host, port), timeout=timeout):
            return True
    except OSError:
        return False


def pytest_collection_modifyitems(items):
    if _reachable(HOST):
        return
    if _HOST_EXPLICIT:
        pytest.exit(f"BEAGLE_HOST={HOST} gesetzt aber nicht erreichbar (Port 5000)", returncode=1)
    skip = pytest.mark.skip(reason=f"BeagleBone nicht erreichbar ({HOST}:5000) — BEAGLE_HOST setzen für Hardware-Tests")
    for item in items:
        item.add_marker(skip)
