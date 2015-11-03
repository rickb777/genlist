package plumbing

const Plumbing = `
//-------------------------------------------------------------------------------------------------

// This plumbing suite provides standard functions for piping data between goroutines.
// All these functions run until the input channel is closed (or all input channels are closed, if
// multiple). They then close their output channel(s).
// None of these functions create a goroutine - this must be done at the call site.

//-------------------------------------------------------------------------------------------------

// {{.TName}}Generator produces a stream of {{.PName}} based on a supplied generator function.
// The function is invoked N times with the integers from 0 to N-1. Each result is sent out.
// Finally, the output channel is closed and the generator terminates.
func {{.TName}}Generator(out chan<- {{.PName}}, iterations int, fn func(int) {{.PName}}) {
	{{.TName}}Generator3(out, 0, iterations-1, 1, fn)
}

// {{.TName}}Generator produces a stream of {{.PName}} based on a supplied generator function.
// The function is invoked *(|to - from|) / |stride|* times with the integers in the range specified by
// *from*, *to* and *stride*. If *stride* is negative, *from* should be greater than *to*.
// For each iteration, the computed function result is sent out.
// If *stride* is zero, the loop never terminates. Otherwise, after the generator has reached the
// loop end, the output channel is closed and the generator terminates.
func {{.TName}}Generator3(out chan<- {{.PName}}, from, to, stride int, fn func(int) {{.PName}}) {
	if (from > to && stride > 0) || (from < to && stride < 0) {
		panic("Loop conditions are divergent.")
	}
	if (from > to && stride < 0) {
		for i := from; i >= to; i += stride {
			out <- fn(i)
		}
	} else {
		for i := from; i <= to; i += stride {
			out <- fn(i)
		}
	}
	close(out)
}

// {{.TName}}Delta duplicates a stream of {{.PName}} to two output channels.
// When the sender closes the input channel, both output channels are closed then the function terminates.
func {{.TName}}Delta(in <-chan {{.PName}}, out1, out2 chan<- {{.PName}}) {
	for v := range in {
		select {
		case out1 <- v: out2 <- v
		case out2 <- v: out1 <- v
		}
	}
	close(out1)
	close(out2)
}

// {{.TName}}Zip2 interleaves two streams of {{.PName}}. Each input channel is used in turn, alternating between them.
// The function terminates when *both* input channels have been closed by their senders. The output channel is
// then closed also.
func {{.TName}}Zip2(in1, in2 <-chan {{.PName}}, out chan<- {{.PName}}) {
	closed2 := false
	for v := range in1 {
		out <- v
		v, ok := <- in2
		if ok {
			out <- v
		} else {
			closed2 = true
		}
	}
	// need to drain in2 as well?
	if !closed2 {
		for _ = range in2 {
		}
	}
	close(out)
}

// {{.TName}}Mux2 multiplexes two streams of {{.PName}}. Each input channel is used as soon as it is ready.
// The function terminates when *both* input channels have been closed by their senders. The output channel is
// then closed also.
//func {{.TName}}Mux2(in1, in2 <-chan {{.PName}}, out chan<- {{.PName}}) {
//	open1 := true -- TODO detect closed channels
//	open2 := true
//	for open1 || open2 {
//		select {
//		case v := <- in1
//			out <- v
//		case v := <- in2
//			out <- v
//		}
//	}
//	close(out)
//}

// {{.TName}}BlackHole silently consumes a stream of {{.PName}}. It terminates when the sender closes the channel.
func {{.TName}}BlackHole(in <-chan {{.PName}}) {
	for _ = range in {
		// om nom nom
	}
}

// {{.TName}}Filter filters a stream of {{.PName}}, silently dropping elements that do not match the predicate p.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.TName}}Filter(in <-chan {{.PName}}, out chan<- {{.PName}}, p func({{.PName}}) bool) {
	for v := range in {
		if p(v) {
			out <- v
		}
	}
	close(out)
}

// {{.TName}}Partition filters a stream of {{.PName}} into two output streams using a predicate p.
// When the sender closes the input channel, both output channels are closed then the function terminates.
func {{.TName}}Partition(in <-chan {{.PName}}, matching, others chan<- {{.PName}}, p func({{.PName}}) bool) {
	for v := range in {
		if p(v) {
			matching <- v
		} else {
			others <- v
		}
	}
	close(matching)
	close(others)
}

// {{.TName}}Map transforms a stream of {{.PName}} by applying a function to each item in the stream.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.TName}}Map(in <-chan {{.PName}}, out chan<- {{.PName}}, fn func({{.PName}}) {{.PName}}) {
	for v := range in {
		out <- fn(v)
	}
	close(out)
}

// {{.TName}}FlatMap transforms a stream of {{.PName}} by applying a function to each item in the stream that
// gives zero or more results, all of which are sent out.
// When the sender closes the input channel, the output channel is closed then the function terminates.
func {{.TName}}FlatMap(in <-chan {{.PName}}, out chan<- {{.PName}}, fn func({{.PName}}) {{.TName}}Collection) {
	for vi := range in {
		c := fn(vi)
		if c.NonEmpty() {
			for vo := range c.Send() {
				out <- vo
			}
		}
	}
	close(out)
}

`
