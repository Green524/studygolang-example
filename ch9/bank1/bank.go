package bank

type WithdrawStruct struct {
	Amount int
	Msg    string
	Flag   chan bool
}

var (
	deposits  = make(chan int)
	balances  = make(chan int)
	withdraws = make(chan WithdrawStruct)
)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) WithdrawStruct {
	msg := new(WithdrawStruct)
	msg.Amount = amount
	msg.Flag = make(chan bool)
	if Balance() >= amount {
		msg.Msg = "取款成功"
	} else {
		msg.Msg = "余额不足"
	}
	withdraws <- *msg
	return *msg
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case msg := <-withdraws:
			if msg.Msg == "取款成功" {
				msg.Flag <- true
				balance -= msg.Amount
			} else {
				msg.Flag <- false
				msg.Amount = balance
			}

		}
	}
}
func init() {
	go teller()
}
