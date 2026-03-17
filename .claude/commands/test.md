Run the Go unit tests for the HAL layer using the mock driver (no hardware required).

```bash
cd go-api && go test ./pkg/hal/... -v -race -count=3
```

After the tests complete:
- Report how many tests passed/failed
- Show any race conditions detected
- If tests fail, analyze the failure and suggest a fix
- Remind the user that integration tests require a real BeagleBone: `BEAGLE_HOST=192.168.7.2 pytest tests/hardware/ -v`
