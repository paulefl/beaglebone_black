Switch the active HAL backend on the BeagleBone Black.

Available backends: `c`, `rust`, `auto`

The user wants to switch to: $ARGUMENTS

Run on the BeagleBone:
```bash
ssh debian@192.168.7.2 "sudo systemctl set-environment HW_BACKEND=$ARGUMENTS && systemctl restart embedded-sw"
```

Then verify the backend is active:
```bash
curl http://192.168.7.2:5000/health
```

The `/health` response includes the active `backend` field. Confirm it matches the requested backend.

Backend descriptions:
- `c` — uses the C shared library via CGO (highest performance)
- `rust` — uses the Rust shared library via FFI (memory safety guarantees)
- `auto` — tries C first, falls back to Rust on error (default for production)
