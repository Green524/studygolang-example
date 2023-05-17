package bank

import "sync"

var (
	//mu      sync.Mutex
	mu      sync.RWMutex //多读单写锁
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

//死锁（对已经锁上的goroutine 再次上锁）
//func Withdraw(amount int) bool {
//	mu.Lock()
//	defer mu.Unlock()
//	Deposit(-amount)
//	if Balance() < 0 {
//		Deposit(amount)
//		return false // insufficient funds
//	}
//	return true
//}
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

//此方法需要加锁
func deposit(amount int) {
	balance += amount
}
