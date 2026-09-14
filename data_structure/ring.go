package data_structure

import "container/ring"

type SlideWindow struct {
	n    int
	ring *ring.Ring
}

// type Ring struct {
// 	next, prev *Ring
// 	Value      any // for use by client; untouched by this library
// }

func NewSlideWindow(n int) *SlideWindow {
	return &SlideWindow{
		n:    n,
		ring: ring.New(n),
	}
}

func (w *SlideWindow) Push(data float64) {
	w.ring.Value = data
	w.ring = w.ring.Next()
}

func (w *SlideWindow) Mean() float64 {
	var sum, count float64
	w.ring.Do(func(a any) {
		if a != nil {
			count += 1.0
			sum += a.(float64)
		}
	})
	return sum / count
}
