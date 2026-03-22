import subprocess

def test_binary_runs():
    result = subprocess.run(["./bme260_main"], capture_output=True, text=True)
    assert result.returncode == 0