package bank_test

import (
	"fmt"
	"studygolang-example/ch9/bank1"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	go func() {
		bank.Deposit(200)
		fmt.Printf("余额：%d\n", bank.Balance())
		withdraw := bank.Withdraw(200)
		<-withdraw.Flag
		fmt.Printf("%v，金额：%d\n", withdraw.Msg, withdraw.Amount)
		done <- struct{}{}
	}()

	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d want = %d", got, want)
	}
}
