package slideWindow

import (
	"sync"
	"time"
)

type SlideWindow struct {
	count []int
	last []int64
	sync.Mutex
	concurrent int
}

func GetSlideWindow(concurrent int) *SlideWindow{
	s := new(SlideWindow)
	s.count = make([]int, 60)
	s.last = make([]int64, 60)
	s.concurrent = concurrent
	return s
}

func (w *SlideWindow)Visit() bool{
	w.Lock()
	defer func() {
		w.Unlock()
	}()
	t := time.Now().Unix()
	index := t % 60
	if t - w.last[index] >= 60{
		w.last[index] = t
		w.count[index] = 1
		return true
	}else {
		if w.count[index] < w.concurrent {
			w.count[index] = w.count[index] + 1
			return true
		}
	}
	return false
}
