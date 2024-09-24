package balancer

type Request struct {
	Fn func() int // The operation to perform.
	C  chan int   // The channel to return the result.
}
