package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		Deposit(200)
		fmt.Printf("阿A的余额：%d\n", Balance())
		ch <- struct{}{}
	}()

	go func() {
		Deposit(100)
		fmt.Printf("阿B的余额：%d\n", Balance())
		ch <- struct{}{}
	}()

	go func() {
		fmt.Printf("取钱：%v\n", Withdraw(100))
		ch <- struct{}{}
	}()
	<-ch
	<-ch
	<-ch
	fmt.Println("done")
}
