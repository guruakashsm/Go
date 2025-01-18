package main

type HandlerFunc func(string)

func (f HandlerFunc) ServeHTTP(s string) {
	f(s)
}

type Handler interface {
	ServeHTTP(string)
}

func main() {
	handler := func(s string) {
		println(s + " world")
	}

	result := HandlerFunc(handler)

	TestHandler(result)
}

func TestHandler(h Handler) {
	h.ServeHTTP("hello")
}
