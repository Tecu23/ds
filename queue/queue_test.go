package queue

import "testing"

func TestQueueSimple(t *testing.T) {
	q := New()

	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 100; i++ {
		if q.Peek().(int) != i {
			t.Error("peek", i, "had value", q.Peek())
		}
		x := q.Dequeue()
		if x != i {
			t.Error("remove", i, "had value", x)
		}
	}
}

func TestQueueWrapping(t *testing.T) {
	q := New()

	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 3; i++ {
		q.Dequeue()
		q.Enqueue(100 + i)
	}

	for i := 0; i < 100; i++ {
		if q.Peek().(int) != i+3 {
			t.Error("peek", i, "had value", q.Peek())
		}
		q.Dequeue()
	}
}

func TestQueueLen(t *testing.T) {
	q := New()

	if q.Len() != 0 {
		t.Error("empty queue length not 0")
	}

	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
		if q.Len() != i+1 {
			t.Error("adding: queue with", i, "elements has length", q.Len())
		}
	}
	for i := 0; i < 1000; i++ {
		q.Dequeue()
		if q.Len() != 1000-i-1 {
			t.Error("removing: queue with", 1000-i-i, "elements has length", q.Len())
		}
	}
}

// func TestQueueGetNegative(t *testing.T) {
// 	q := New()
//
// 	for i := 0; i < 1000; i++ {
// 		q.Enqueue(i)
// 		for j := 1; j <= q.Len(); j++ {
// 			if q.Get(-j).(int) != q.Len()-j {
// 				t.Errorf("index %d doesn't contain %d", -j, q.Len()-j)
// 			}
// 		}
// 	}
// }

// func TestQueueGetOutOfRangePanics(t *testing.T) {
// 	q := New()
//
// 	q.Enqueue(1)
// 	q.Enqueue(2)
// 	q.Enqueue(3)
//
// 	assertPanics(t, "should panic when negative index", func() {
// 		q.Get(-4)
// 	})
//
// 	assertPanics(t, "should panic when index greater than length", func() {
// 		q.Get(4)
// 	})
// }

// func TestQueuePeekOutOfRangePanics(t *testing.T) {
// 	q := New()
//
// 	assertPanics(t, "should panic when peeking empty queue", func() {
// 		q.Peek()
// 	})
//
// 	q.Enqueue(1)
// 	q.Dequeue()
//
// 	assertPanics(t, "should panic when peeking emptied queue", func() {
// 		q.Peek()
// 	})
// }
//
// func TestQueueDequeueOutOfRangePanics(t *testing.T) {
// 	q := New()
//
// 	assertPanics(t, "should panic when removing empty queue", func() {
// 		q.Dequeue()
// 	})
//
// 	q.Enqueue(1)
// 	q.Dequeue()
//
// 	assertPanics(t, "should panic when removing emptied queue", func() {
// 		q.Dequeue()
// 	})
// }

func assertPanics(t *testing.T, name string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: didn't panic as expected", name)
		}
	}()

	f()
}

// WARNING: Go's benchmark utility (go test -bench .) increases the number of
// iterations until the benchmarks take a reasonable amount of time to run; memory usage
// is *NOT* considered. On a fast CPU, these benchmarks can fill hundreds of GB of memory
// (and then hang when they start to swap). You can manually control the number of iterations
// with the `-benchtime` argument. Passing `-benchtime 1000000x` seems to be about right.

func BenchmarkQueueSerial(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Enqueue(nil)
	}
	for i := 0; i < b.N; i++ {
		q.Peek()
		q.Dequeue()
	}
}

func BenchmarkArrQueueSerial(b *testing.B) {
	q := ArrNew()
	for i := 0; i < b.N; i++ {
		q.ArrEnqueue(nil)
	}
	for i := 0; i < b.N; i++ {
		q.ArrPeek()
		q.ArrDequeue()
	}
}

// func BenchmarkQueueGet(b *testing.B) {
// 	q := New()
// 	for i := 0; i < b.N; i++ {
// 		q.Enqueue(i)
// 	}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		q.Get(i)
// 	}
// }

func BenchmarkQueueTickTock(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Enqueue(nil)
		q.Peek()
		q.Dequeue()
	}
}

func BenchmarkArrQueueTickTock(b *testing.B) {
	q := ArrNew()
	for i := 0; i < b.N; i++ {
		q.ArrEnqueue(nil)
		q.ArrPeek()
		q.ArrDequeue()
	}
}
