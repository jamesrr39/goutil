package profile

import (
	"bytes"
	"sync"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Profiler(t *testing.T) {
	timeIncrement := int64(0)
	byteBuffer := bytes.NewBuffer(nil)
	profiler := Profiler{
		byteBuffer,
		func() time.Time { timeIncrement++; return time.Unix(0, timeIncrement*2000) },
		sync.Mutex{},
	}

	testRun := profiler.NewRun("add ints and strings")

	profiler.Mark(testRun, "adding 1 and 2")
	// some work here

	profiler.Mark(testRun, "building a string")
	// some work here

	err := profiler.StopAndRecord(testRun, "finished successfully")
	require.NoError(t, err)

	t.Run("recorded info correct", func(t *testing.T) {
		var run Run
		err := proto.Unmarshal(byteBuffer.Bytes(), &run)
		require.NoError(t, err)

		assert.Equal(t, "add ints and strings", run.Name)
		assert.Equal(t, "finished successfully", run.Summary)
		assert.Equal(t, int64(0), run.StartTimeNanos)
		assert.Equal(t, int64(6000), run.EndTimeNanos)
		require.Len(t, run.Events, 2)
		expectedEvents := []*Event{
			{Name: "adding 1 and 2", TimeNanos: 2000},
			{Name: "building a string", TimeNanos: 4000},
		}
		assert.Equal(t, expectedEvents, run.Events)
	})
}

// func Test_a(t *testing.T) {
// 	fp := filepath.Join("/tmp", time.Now().Format(time.RFC3339Nano)+".pbf")

// 	f, err := os.Create(fp)
// 	require.NoError(t, err)
// 	defer f.Close()

// 	p := NewProfiler(streamtostorage.NewWriter(f))
// 	run := p.NewRun("test--")

// 	time.Sleep(800 * time.Millisecond)

// 	p.Mark(run, "step 1 -- do some work")

// 	time.Sleep(1500 * time.Millisecond)

// 	p.Mark(run, "step 2 -- do some more work")

// 	time.Sleep(2300 * time.Millisecond)

// 	p.Mark(run, "step 3 -- do some more work 3")

// 	err = p.StopAndRecord(run, "finished successfully")
// 	require.NoError(t, err)
// }
