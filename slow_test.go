package slow

import (
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		i := 0
		fn := func() {
			i++
		}
		debounced := Debounce(fn, time.Millisecond*50, nil)
		debounced()
		time.Sleep(time.Millisecond * 60)
		if i != 1 {
			t.FailNow()
		}
		debounced()
		if i != 1 {
			t.FailNow()
		}
		time.Sleep(time.Millisecond * 20)
		debounced()
		time.Sleep(time.Millisecond * 20)
		debounced()
		time.Sleep(time.Millisecond * 20)
		debounced()
		time.Sleep(time.Millisecond * 20)
		debounced()
		time.Sleep(time.Millisecond * 20)
		debounced()
		if i != 1 {
			t.FailNow()
		}
		time.Sleep(time.Millisecond * 60)
		if i != 2 {
			t.FailNow()
		}
		time.Sleep(time.Millisecond * 60)
		if i != 2 {
			t.FailNow()
		}
	})
}

func TestThrottle(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		i := 0
		fn := func() {
			i++
		}
		throttled := Throttle(fn, time.Millisecond*50, nil)
		throttled()
		time.Sleep(time.Millisecond * 20)
		if i != 1 {
			t.FailNow()
		}
		time.Sleep(time.Millisecond * 60)
		if i != 2 {
			t.FailNow()
		}
		throttled()
		throttled()
		throttled()
		throttled()
		time.Sleep(time.Millisecond * 60)
		throttled()
		time.Sleep(time.Millisecond * 20)
		if i != 5 {
			t.FailNow()
		}
		throttled()
		time.Sleep(time.Millisecond * 20)
		if i != 5 {
			t.FailNow()
		}
		throttled()
		time.Sleep(time.Millisecond * 20)
		if i != 6 {
			t.FailNow()
		}
		throttled()
		time.Sleep(time.Millisecond * 20)
		throttled()
		time.Sleep(time.Millisecond * 20)
		throttled()
		time.Sleep(time.Millisecond * 20)
		if i != 7 {
			t.FailNow()
		}
		time.Sleep(time.Millisecond * 60)
		if i != 9 {
			t.FailNow()
		}
	})
}
