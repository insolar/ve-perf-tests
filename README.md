### ve-perf-tests
VE performance tests

### Local run
Setup network
```
scripts/insolard/launchnet.sh -g
```

Run single test for N nodes network
```
TARGET=http://127.0.0.1:32304 REPORT_CSV_FILE=scaling.csv WALLETS=100 NODES=5 go run cmd/test/main.go
```

Generate report
```
REPORT_CSV_FILE=scaling.csv REPORT_PNG_FILE=report.png go run cmd/report/main.go
```
