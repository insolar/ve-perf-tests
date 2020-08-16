module github.com/insolar/ve-perf-tests

go 1.14

require (
	github.com/gojuno/minimock/v3 v3.0.8 // indirect
	github.com/insolar/assured-ledger/ledger-core v0.0.0-20200811095133-eb75ba92a497
	github.com/insolar/consensus-reports v0.0.0-20200515131339-fea7a784f1d6
	github.com/insolar/insconfig v0.0.0-20200513150834-977022bc1445
	github.com/insolar/loaderbot v0.0.22
	github.com/json-iterator/go v1.1.9
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.6.1
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/sys v0.0.0-20200810151505-1b9f1253b3ed // indirect
	golang.org/x/tools v0.0.0-20200811172722-d77521d07411 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
