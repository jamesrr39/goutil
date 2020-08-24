package profile

import (
	"context"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/jamesrr39/goutil/errorsx"
)

var (
	runCtxKey      = struct{}{}
	profilerCtxKey = struct{}{}
)

func Middleware(profiler *Profiler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			run := profiler.NewRun(uuid.New().String())
			r = r.WithContext(
				context.WithValue(r.Context(), runCtxKey, run),
				context.WithValue(r.Context(), profilerCtxKey, profiler),
			)

			next.ServeHTTP(w, r)

			duration := time.Now().Sub(now)
			err := profiler.StopAndRecord(run, "finished in %s", duration.String())
			if err != nil {
				log.Printf("ERROR: profiler: could not StopAndRecord. Error: %q\n", err)
			}
		}

		return http.HandlerFunc(fn)
	}
}

func MarkOnCtx(ctx context.Context, eventName string) errorsx.Error {
	run := ctx.Value(runCtxKey)
	if nil {
		return errorsx.Errorf("Profile: MarkOnCtx: no profile run found on context")
	}
	profiler := ctx.Value(profilerCtxKey)
	if nil {
		return errorsx.Errorf("Profile: MarkOnCtx: no profile found on context")
	}

	p, ok := profiler.(*Profiler)
	if !ok {
		return errorsx.Errorf("Profile: MarkOnCtx: profiler type was not *Profiler (was %T)", profiler)
	}
	r, ok := run.(*Run)
	if !ok {
		return errorsx.Errorf("Profile: MarkOnCtx: run type was not *Run (was %T)", run)
	}

	return p.Mark(r, eventName)
}

type Profiler struct {
	writer  io.Writer
	nowFunc func() time.Time
	writeMu sync.Mutex
}

func NewProfiler(writer io.Writer) *Profiler {
	return &Profiler{writer, time.Now, sync.Mutex{}}
}

func (profiler *Profiler) NewRun(runName string) *Run {
	return &Run{
		Name:           runName,
		StartTimeNanos: profiler.nowFunc().UnixNano(),
		Events:         []*Event{},
	}
}

func (profiler *Profiler) Mark(run *Run, eventName string) {
	now := profiler.nowFunc()
	run.Events = append(run.Events, &Event{
		Name:      eventName,
		TimeNanos: now.UnixNano(),
	})
}

func (profiler *Profiler) StopAndRecord(run *Run, summaryMessage string) errorsx.Error {
	endTime := profiler.nowFunc()

	run.EndTimeNanos = endTime.UnixNano()
	run.Summary = summaryMessage

	b, err := proto.Marshal(run)
	if err != nil {
		return errorsx.Wrap(err)
	}

	profiler.writeMu.Lock()
	defer profiler.writeMu.Unlock()

	_, err = profiler.writer.Write(b)
	if err != nil {
		return errorsx.Wrap(err)
	}

	return nil
}
