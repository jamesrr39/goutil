package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jamesrr39/goutil/bolt-tools/boltviz"
	"github.com/jamesrr39/goutil/must"
)

func main() {
	flag.Parse()
	filePath := flag.Arg(0)

	db, err := bolt.Open(filePath, 0600, nil)
	must.Must(err)
	defer db.Close()

	h, err := boltviz.NewHandlerFunc(db, boltviz.TemplateMap{
		PrintKey: func(pair boltviz.KVPairDisplay) string {
			return string(pair.Key)
		},
		PrintValue: func(pair boltviz.KVPairDisplay) string {
			return string(pair.Value)
		},
	}, time.Now().Format(time.Kitchen))

	must.Must(err)

	http.ListenAndServe("localhost:8090", h)
}
