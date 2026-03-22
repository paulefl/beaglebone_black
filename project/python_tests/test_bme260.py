import subprocess

try:
    # Binary ausführen
    result = subprocess.run(
        ["./bme260_main"],      # Pfad zum Binary
        capture_output=True,    # stdout und stderr erfassen
        text=True,              # als String, nicht Bytes
        check=True              # wirft CalledProcessError bei Fehlercode != 0
    )

    print("=== Output ===")
    print(result.stdout)       # stdout des Binaries
    print("=== Errors ===")
    print(result.stderr)       # stderr des Binaries (falls vorhanden)

except subprocess.CalledProcessError as e:
    print(f"Binary returned non-zero exit code {e.returncode}")
    print("Output:", e.output)
    print("Errors:", e.stderr)

except FileNotFoundError:
    print("Binary './bme260_main' not found. War es gebaut?")

except Exception as e:
    print("Unexpected error:", str(e))