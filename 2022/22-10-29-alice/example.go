package main

type Handler interface {
	ServeTest()
}

type HandlerFunc func()

func (f HandlerFunc) ServeTest() {
	f()
}

type Constructor func(Handler) Handler

func test1Middleware() Constructor {
	return func(h Handler) Handler {
		return HandlerFunc(func() {
			println("test1")
			// 关键代码
			h.ServeTest()
		})
	}
}
func test2Middleware() Constructor {
	return func(h Handler) Handler {
		return HandlerFunc(func() {
			println("test2")
			// 关键代码
			h.ServeTest()
		})
	}
}

func Then(h Handler) Handler {
	array := []Constructor{test1Middleware(), test2Middleware()}
	for i := range array {
		h = array[len(array)-1-i](h)
	}
	return h
}

func main() {
	fn := HandlerFunc(func() {
		println("fn")
	})
	t := Then(fn)
	t.ServeTest()
}
