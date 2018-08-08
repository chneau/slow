package slow

import (
	"time"
)

// const debounce = (func, delay) => {
// 	let inDebounce
// 	return function() {
// 	  const context = this
// 	  const args = arguments
// 	  clearTimeout(inDebounce)
// 	  inDebounce = setTimeout(() => func.apply(context, args), delay)
// 	}
//   }

// Options ...
type Options struct {
	Leading  bool
	MaxWait  time.Duration
	Trailing bool
}

// Debounce ...
//
// leading := false
//
// maxWait := time.Duration(0)
//
// trailing := true
//
func Debounce(fn func(), wait time.Duration, options *Options) func() {
	lastCallTime := time.Time{}
	lastInvokeTime := time.Time{}
	leading := false
	maxWait := time.Duration(0)
	trailing := true
	maxing := false
	var timer *time.Timer
	var timerExpired func()
	var trailingEdge func(time.Time)
	var remainingWait func(time.Time) time.Duration
	if options != nil {
		leading = options.Leading
		maxing = options.MaxWait != 0
		if maxing == true {
			if wait > options.MaxWait {
				maxWait = wait
			} else {
				maxWait = options.MaxWait
			}
		}
		trailing = options.Trailing
	}

	invokeFunc := func(t time.Time) {
		lastInvokeTime = t
		go fn()
	}

	leadingEdge := func(t time.Time) {
		lastInvokeTime = t
		timer = time.AfterFunc(remainingWait(t), timerExpired)
		if leading {
			invokeFunc(t)
		}
	}

	remainingWait = func(t time.Time) time.Duration {
		timeSinceLastCall := t.Sub(lastCallTime)
		timeSinceLastInvoke := t.Sub(lastInvokeTime)
		timeWaiting := wait - timeSinceLastCall
		if maxing {
			if timeWaiting < maxWait-timeSinceLastInvoke {
				return timeWaiting
			}
			return maxWait - timeSinceLastInvoke
		}
		return timeWaiting
	}

	shouldInvoke := func(t time.Time) bool {
		timeSinceLastCall := t.Sub(lastCallTime)
		timeSinceLastInvoke := t.Sub(lastInvokeTime)
		return lastCallTime.IsZero() || timeSinceLastCall > wait ||
			timeSinceLastCall < 0 || (maxing && timeSinceLastInvoke >= maxWait)
	}

	timerExpired = func() {
		t := time.Now()
		if shouldInvoke(t) {
			trailingEdge(t)
			return
		}
		timer = time.AfterFunc(remainingWait(t), timerExpired)
	}

	trailingEdge = func(t time.Time) {
		timer = nil
		if trailing {
			invokeFunc(t)
			return
		}
	}

	debounced := func() {
		t := time.Now()
		isInvoking := shouldInvoke(t)
		lastCallTime = t
		if isInvoking {
			if timer == nil {
				leadingEdge(lastCallTime)
				return
			}
			if maxing {
				timer = time.AfterFunc(wait, timerExpired)
				invokeFunc(lastCallTime)
				return
			}
		}
		if timer == nil {
			timer = time.AfterFunc(wait, timerExpired)
		}
	}
	return debounced
}

// Throttle ...
func Throttle(fn func(), wait time.Duration, options *Options) func() {
	leading := true
	trailing := true
	if options != nil {
		leading = options.Leading
		trailing = options.Trailing
	}
	return Debounce(fn, wait, &Options{
		Leading:  leading,
		MaxWait:  wait,
		Trailing: trailing,
	})
}
