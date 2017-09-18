package debounce

import "time"

type Debouncer struct {
	wait     time.Duration
	callback *func()
}

func NewDebouncer(wait time.Duration) *Debouncer {
	return &Debouncer{wait, nil}
}

func (d *Debouncer) Run(callback func()) {
	d.callback = &callback
	go func() {
		time.Sleep(d.wait)
		if d.callback == &callback {
			callback()
		}
	}()
}
