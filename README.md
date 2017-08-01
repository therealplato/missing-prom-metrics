```
go get -u github.com/therealplato/missing-prom-metrics && cd $GOPATH/src/github.com/therealplato/missing-prom-metrics/service
go build
./service
```

Hit http://localhost:9009/metrics to get the promhttp metrics page.

Send a GET request to http://localhost:9000 to trigger a metric.

Observe that the metric is not shown on the metrics page.

Changing the vendoring strategy to `dep` has solved the problem but I don't know
why, see the `dep` branch.
