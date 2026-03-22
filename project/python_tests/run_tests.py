# file: run_tests.py
import subprocess
import sys
import os

# Liste der Binaries, die getestet werden sollen
BINARIES = [
    {"name": "bme260_main", "path": "/output/bme260_main"},
    {"name": "rust_app", "path": "/output/rust_app"},
    {"name": "go_app", "path": "/output/go_app"}
]

def run_binary(binary):
    """Führt ein Binary aus und prüft Ausgabe und Exitcode."""
    path = binary["path"]
    name = binary["name"]
    
    if not os.path.isfile(path):
        print(f"[ERROR] Binary '{name}' nicht gefunden unter {path}")
        return False

    try:
        result = subprocess.run(
            [path],
            capture_output=True,
            text=True,
            check=True
        )
        print(f"[OK] {name} ausgeführt:")
        print("=== STDOUT ===")
        print(result.stdout.strip())
        if result.stderr.strip():
            print("=== STDERR ===")
            print(result.stderr.strip())
        return True

    except subprocess.CalledProcessError as e:
        print(f"[FAIL] {name} exit code {e.returncode}")
        print("=== STDOUT ===")
        print(e.stdout.strip())
        print("=== STDERR ===")
        print(e.stderr.strip())
        return False
    except Exception as e:
        print(f"[ERROR] {name} konnte nicht ausgeführt werden: {str(e)}")
        return False

def main():
    all_passed = True
    for binary in BINARIES:
        print(f"\n--- Testing {binary['name']} ---")
        passed = run_binary(binary)
        if not passed:
            print(f"\n--- Testing {binary['name']} failed ---")
            all_passed = False

    if all_passed:
        print("\n=== Alle Tests bestanden ✅ ===")
        sys.exit(0)
    else:
        print("\n=== Einige Tests fehlgeschlagen ❌ ===")
        sys.exit(1)

if __name__ == "__main__":
    main()