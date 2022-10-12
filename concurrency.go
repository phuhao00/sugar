package sugar

import "sync"

type synchronize struct {
	locker sync.Locker
}

func (s *synchronize) Do(cb func() error) {
	s.locker.Lock()
	Try(cb)
	s.locker.Unlock()
}

func Synchronize(opt ...sync.Locker) synchronize {
	if len(opt) > 1 {
		panic("unexpected arguments")
	} else if len(opt) == 0 {
		opt = append(opt, &sync.Mutex{})
	}

	return synchronize{locker: opt[0]}
}

func Async[A any](f func() A) chan A {
	ch := make(chan A)

	go func() {
		ch <- f()
	}()

	return ch
}

// FanIn ...
func FanIn(in ...chan any) <-chan any {
	out := make(chan any)
	for i := range in {
		tmp := in[i]
		go func() {
			out <- tmp
		}()
	}
	return out
}

func FanOut(ch <-chan any, n int) []chan any {
	cs := make([]chan any, 0, n)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan any))
	}
	distributeToChannels := func(ch <-chan any, cs []chan any) {
		defer func(cs []chan any) {
			for _, c := range cs {
				close(c)
			}
		}(cs)

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}

					c <- val
				}
			}
		}
	}
	go distributeToChannels(ch, cs)

	return cs
}
