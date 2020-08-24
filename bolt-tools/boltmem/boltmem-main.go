package main

import (
	"encoding/binary"
	"flag"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jamesrr39/goutil/must"
)

func main() {
	iterationsPerTx := flag.Int("iterations-per-tx", 100000, "iterations in one transaction")
	numTransactions := flag.Int("num-transactions", 10, "number of transactions to perform")
	flag.Parse()

	err := os.MkdirAll("/tmp/bolt", 0700)
	must.Must(err)

	filePath := "/tmp/bolt/" + time.Now().Format(time.RFC3339Nano)

	go func() {
		var memstats runtime.MemStats
		for {
			runtime.ReadMemStats(&memstats)

			log.Printf("mem total alloc: %d\nHeap alloc: %d\nalloc: %d\n", memstats.TotalAlloc, memstats.HeapAlloc, memstats.Alloc)

			time.Sleep(time.Millisecond * 500)
		}
	}()

	log.Printf("creating db at %q\n", filePath)

	db, err := bolt.Open(filePath, 0600, nil)
	must.Must(err)

	for iteration := 0; iteration < *numTransactions; iteration++ {
		iterationStartTime := time.Now()
		tx, err := db.Begin(true)
		must.Must(err)

		bucket, err := tx.CreateBucketIfNotExists([]byte("test"))
		must.Must(err)

		for i := 0; i < *iterationsPerTx; i++ {
			if (i % 1000) == 0 {
				log.Printf("inserting %d\n", i)
			}

			vbinary := make([]byte, 8)
			binary.BigEndian.PutUint64(vbinary, uint64(i))
			err = bucket.Put(vbinary, vbinary)
			must.Must(err)
		}

		commitStartTime := time.Now()
		err = tx.Commit()
		must.Must(err)

		endTime := time.Now()

		log.Printf("iteration %d: commit time: %s, total iteration time %s\n",
			iteration,
			endTime.Sub(commitStartTime).String(),
			endTime.Sub(iterationStartTime).String(),
		)
	}

	must.Must(err)
}
