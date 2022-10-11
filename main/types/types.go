package types

type Console interface {
	SayHello()
}

type consoleImpl struct {
}

func (c *consoleImpl) SayHello() {
	println("Hello from main module!")
}
