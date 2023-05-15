package bank

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	//如果这里sema 的容量为0，那么这里往sema 输入struct之后会堵塞，因为容量为0的channel会等待另一个方接收struct，代码才会继续往下走
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
