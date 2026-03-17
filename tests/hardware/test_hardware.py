import requests, os, pytest

HOST    = os.getenv("BEAGLE_HOST","192.168.7.2")
BACKEND = os.getenv("HW_BACKEND","auto")
API     = f"http://{HOST}:5000"

@pytest.fixture(autouse=True)
def set_backend():
    requests.post(f"{API}/api/v1/backend",json={"backend":BACKEND})

class TestBME280:
    def test_erreichbar(self):
        r = requests.get(f"{API}/api/v1/bme280",timeout=5)
        assert r.status_code == 200
    def test_temperatur(self):
        d = requests.get(f"{API}/api/v1/bme280").json()
        assert -40 <= d["temperature"] <= 85
    def test_luftfeuchte(self):
        d = requests.get(f"{API}/api/v1/bme280").json()
        assert 0 <= d["humidity"] <= 100
    def test_luftdruck(self):
        d = requests.get(f"{API}/api/v1/bme280").json()
        assert 300 <= d["pressure"] <= 1100
    def test_hoehe(self):
        d = requests.get(f"{API}/api/v1/bme280").json()
        assert -500 <= d["altitude"] <= 9000

class TestGPIO:
    def test_high(self):
        r = requests.post(f"{API}/api/v1/gpio/60",json={"value":1})
        assert r.status_code == 200
    def test_low(self):
        r = requests.post(f"{API}/api/v1/gpio/60",json={"value":0})
        assert r.status_code == 200
    def test_lesen(self):
        requests.post(f"{API}/api/v1/gpio/60",json={"value":1})
        d = requests.get(f"{API}/api/v1/gpio/60").json()
        assert d["value"] == 1

class TestFallback:
    def test_c_zu_rust(self):
        requests.post(f"{API}/api/v1/backend",json={"backend":"c"})
        r1 = requests.get(f"{API}/api/v1/bme280").json()
        assert r1["backend"] == "c"
        requests.post(f"{API}/api/v1/backend",json={"backend":"rust"})
        r2 = requests.get(f"{API}/api/v1/bme280").json()
        assert r2["backend"] == "rust"
