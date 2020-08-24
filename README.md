### ve-perf-tests

![ve performance check](https://github.com/insolar/ve-perf-tests/workflows/ve-performance%20check/badge.svg)

VE performance tests

### Local run
Setup network
```
scripts/insolard/launchnet.sh -g
```

Run single test for N nodes network
```
TARGET=http://52.168.125.103:34304 REPORT_CSV_FILE=scaling.csv NODES=5 go run cmd/test/main.go
```
The test will create NODES*1000 wallets to ensure good distribution

Generate report
```
REPORT_CSV_FILE=scaling.csv REPORT_PNG_FILE=report.png go run cmd/report/main.go
```
