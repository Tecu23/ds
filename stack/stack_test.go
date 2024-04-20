package stack

import "testing"

func TestStackSimple(t *testing.T) {
	s := New()

	for i := 0; i < 100; i++ {
		s.Push(i)
	}
	for i := 99; i >= 0; i-- {
		if s.Peek().(int) != i {
			t.Error("peek", i, "had value", s.Peek())
		}
		x := s.Pop()
		if x != i {
			t.Error("remove", i, "had value", x)
		}
	}
}

//	func TestStackWrapping(t *testing.T) {
//		s := New()
//
//		for i := 0; i < 100; i++ {
//			s.Push(i)
//		}
//		for i := 0; i < 3; i++ {
//			s.Pop()
//			s.Push(100 + i)
//		}
//
//		for i := 0; i < 100; i++ {
//			if s.Peek().(int) != i+3 {
//				t.Error("peek", i, "had value", s.Peek())
//			}
//			s.Pop()
//		}
//	}
func TestStackLen(t *testing.T) {
	s := New()

	if s.Len() != 0 {
		t.Error("empty sueue length not 0")
	}

	for i := 0; i < 1000; i++ {
		s.Push(i)
		if s.Len() != i+1 {
			t.Error("adding: sueue with", i, "elements has length", s.Len())
		}
	}
	for i := 0; i < 1000; i++ {
		s.Pop()
		if s.Len() != 1000-i-1 {
			t.Error("removing: sueue with", 1000-i-i, "elements has length", s.Len())
		}
	}
}

// func TestStackGetNegative(t *testing.T) {
// 	s := New()
//
// 	for i := 0; i < 1000; i++ {
// 		s.Push(i)
// 		for j := 1; j <= s.Len(); j++ {
// 			if s.Get(-j).(int) != q.Len()-j {
// 				t.Errorf("index %d doesn't contain %d", -j, s.Len()-j)
// 			}
// 		}
// 	}
// }

// func TestStackGetOutOfRangePanics(t *testing.T) {
// 	s := New()
//
// 	s.Push(1)
// 	s.Push(2)
// 	s.Push(3)
//
// 	assertPanics(t, "should panic when negative index", func() {
// 		s.Get(-4)
// 	})
//
// 	assertPanics(t, "should panic when index greater than length", func() {
// 		s.Get(4)
// 	})
// }

// func TestStackPeekOutOfRangePanics(t *testing.T) {
// 	s := New()
//
// 	assertPanics(t, "should panic when peeking empty sueue", func() {
// 		s.Peek()
// 	})
//
// 	s.Push(1)
// 	s.Pop()
//
// 	assertPanics(t, "should panic when peeking emptied sueue", func() {
// 		s.Peek()
// 	})
// }
//
// func TestStackDesueueOutOfRangePanics(t *testing.T) {
// 	s := New()
//
// 	assertPanics(t, "should panic when removing empty sueue", func() {
// 		s.Pop()
// 	})
//
// 	s.Push(1)
// 	s.Pop()
//
// 	assertPanics(t, "should panic when removing emptied sueue", func() {
// 		s.Pop()
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

func BenchmarkStackSerial(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.Push(nil)
	}
	for i := 0; i < b.N; i++ {
		s.Peek()
		s.Pop()
	}
}

func BenchmarkArrStackSerial(b *testing.B) {
	s := ArrNew()
	for i := 0; i < b.N; i++ {
		s.ArrPush(nil)
	}
	for i := 0; i < b.N; i++ {
		s.ArrPeek()
		s.ArrPop()
	}
}

// func BenchmarkStackGet(b *testing.B) {
// 	s := New()
// 	for i := 0; i < b.N; i++ {
// 		s.Push(i)
// 	}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		s.Get(i)
// 	}
// }

func BenchmarkStackTickTock(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.Push(nil)
		s.Peek()
		s.Pop()
	}
}

func BenchmarkArrStackTickTock(b *testing.B) {
	s := ArrNew()
	for i := 0; i < b.N; i++ {
		s.ArrPush(nil)
		s.ArrPeek()
		s.ArrPop()
	}
}
