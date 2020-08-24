package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jamesrr39/goutil/must"
)

func main() {
	iterationsPerTx := flag.Int("iterations-per-tx", 10, "iterations in one transaction")
	numGoroutines := flag.Int("num-goroutines", 10, "number of goroutines batching transactions")
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

	doneChan := make(chan bool, *numGoroutines)

	for goRoutineID := 0; goRoutineID < *numGoroutines; goRoutineID++ {
		go addItem(db, doneChan, *iterationsPerTx, goRoutineID)
		time.Sleep(time.Second)
	}

	<-doneChan
}

func addItem(db *bolt.DB, doneChan chan bool, iterationsPerTx, goRoutineID int) {
	var commitStartTime time.Time

	iterationStartTime := time.Now()

	err := db.Batch(func(tx *bolt.Tx) error {

		if goRoutineID == 5 {
			return errors.New("test error")
		}

		bucket, err := tx.CreateBucketIfNotExists([]byte("test"))
		must.Must(err)

		for i := 0; i < iterationsPerTx; i++ {
			b := []byte(fmt.Sprintf("%d_%d", goRoutineID, i))

			err = bucket.Put(b, b)
			must.Must(err)

			log.Printf("inserted from goroutine %d\n", goRoutineID)

			// time.Sleep(time.Second)
		}

		commitStartTime = time.Now()

		return nil
	})

	must.Must(err)

	endTime := time.Now()

	log.Printf("inserts: %d, commit time: %s, total iteration time: %s\n",
		iterationsPerTx,
		endTime.Sub(commitStartTime).String(),
		endTime.Sub(iterationStartTime).String(),
	)

	doneChan <- true
}
