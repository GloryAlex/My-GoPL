package bank

var deposit = make(chan int)
var balances = make(chan int)
var withdraw = make(chan int)
var withdrawResult = make(chan bool)

func Deposit(amount int) {
	deposit <- amount
}
func Balance() int {
	return <-balances
}
func Withdraw(amount int) bool {
	withdraw <- amount
	return <-withdrawResult
}
func teller() {
	var realBalance = 0
	for true {
		select {
		case amount := <-deposit:
			realBalance += amount
		case balances <- realBalance:
		case amount := <-withdraw:
			if amount > realBalance {
				withdrawResult <- false
			} else {
				realBalance -= amount
				withdrawResult <- true
			}
		}
	}
}

func init() {
	go teller()
}
