package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi"
	"github.com/jamesrr39/goutil/must"
	"github.com/jamesrr39/goutil/profile"
	"github.com/jamesrr39/goutil/streamtostorage"
)

func main() {

	// set up server
	addr := "localhost:9001"
	apiPath := "/api/v1/myroute"

	tempdir, err := ioutil.TempDir("", "")
	must.NoError(err)

	profilePath := filepath.Join(tempdir, "profile.pbf")
	log.Printf("writing profile to %s\n", profilePath)

	f, err := os.Create(profilePath)
	must.NoError(err)
	defer f.Close()

	profileWriter, err := streamtostorage.NewWriter(f, streamtostorage.MessageSizeBufferLenDefault)
	must.NoError(err)

	profiler := profile.NewProfiler(profileWriter)

	router := chi.NewRouter()
	router.Use(profile.Middleware(profiler))
	router.Get(apiPath, func(w http.ResponseWriter, r *http.Request) {
		// handler. This won't do anything, just mark some events, and sleep between them so we can see the timeline in the profileviz summary
		ctx := r.Context()

		err := profile.MarkOnCtx(ctx, "step 1")
		must.NoError(err)

		time.Sleep(time.Second)

		err = profile.MarkOnCtx(ctx, "step 2")
		must.NoError(err)

		time.Sleep(time.Millisecond * 500)

		err = profile.MarkOnCtx(ctx, "step 3")
		must.NoError(err)
	})

	errChan := make(chan error)
	go func() {
		// start server
		err = http.ListenAndServe(addr, router)
		if err != nil {
			errChan <- err
		}
	}()

	doneChan := make(chan struct{})
	go func() {
		const maxAttempts = 20
		for i := 0; i < maxAttempts; i++ {
			time.Sleep(time.Second)
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s%s", addr, apiPath), nil)
			if err != nil {
				errChan <- err
				return
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				errChan <- err
				return
			}

			if resp.StatusCode == 200 {
				doneChan <- struct{}{}
				return
			}
		}
		errChan <- fmt.Errorf("reached max attempts: %d", maxAttempts)
	}()

	select {
	case err := <-errChan:
		log.Fatalln(err)
	case <-doneChan:
		log.Println("finished successfully")
	}
}
