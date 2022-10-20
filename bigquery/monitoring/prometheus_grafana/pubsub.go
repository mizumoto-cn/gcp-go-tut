package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/mizumoto-cn/gcp-go-tut/bigquery/monitoring/prometheus_grafana/constexpr"
	"github.com/mizumoto-cn/gcp-go-tut/bigquery/monitoring/prometheus_grafana/parser"
)

//

// Gauge metrics for Prometheus
var (
	requestCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_request_count",
		Help: "total request count",
	})
	operatesGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_operates_gauge",
		Help: "total operates gauge",
	})
)

// Pulling messages from pub/sub
func pullMessages() {
	log.Println("STARTED PULLING MESSAGES")

	ctx := context.Background()

	// Set your $PROJECT_ID
	client, err := pubsub.NewClient(ctx, "thisol-bps-apigee-project1")
	if err != nil {
		log.Fatal(err)
	}

	// Set your $SUBSCRIPTION
	subID := "xu-prometheus-puller"
	var mu sync.Mutex

	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()

		log.Print("Got message: " + string(msg.Data))

		if !containsInitPayload(string(msg.Data)) {
			t := parser.Payload{Data: msg.Data}
			r := parser.RawLogHandler{}
			p, err := r.Parse(t)
			if err != nil {
				log.Fatal(err)
			}
			operate := p.MethodName
			operatesGauge.Set(constexpr.GetOps().Get())
		}

		msg.Ack()
	})
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	cancel()
}

// fucntion to check if the message is the init payload
func containsInitPayload(payload string) bool {
	// if strings.Contains(payload, "wubba lubba dub dub") {
	// 	return true
	// }
	// fill in your own init payload if you want
	return false
}

func main() {
	go pullMessages()

	log.Println("STARTED PROMETHEUS")

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
