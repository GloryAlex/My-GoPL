package bank

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func TestBalance(t *testing.T) {
	type test struct {
		value    int
		routines int
	}
	rand.Seed(time.Now().Unix())
	name := "bank"
	tests := []test{}
	for i := 0; i < 10; i++ {
		tests = append(tests, test{rand.Int() % 1000, rand.Int()%5 + 1})
	}

	total := 0
	for _, tt := range tests {
		total += tt.value * tt.routines
		t.Run(name, func(t *testing.T) {
			isEnd := make(chan struct{}, tt.routines)
			for i := 0; i < tt.routines; i++ {
				go func() {
					Deposit(tt.value)
					isEnd <- struct{}{}
				}()
			}
			for i := 0; i < tt.routines; i++ {
				<-isEnd
			}
			got := Balance()
			info := fmt.Sprintf("Balance() = %v, want %v", got, total)
			if got != total {
				t.Errorf(info)
			} else {
				t.Logf(info)
			}
		})
	}
}

func TestWithdraw(t *testing.T) {
	type test struct {
		value    int
		routines int
	}
	rand.Seed(time.Now().Unix())
	name := "bank"
	tests := []test{}
	for i := 0; i < 10; i++ {
		tests = append(tests, test{rand.Int() % 1000, rand.Int()%10 + 1})
	}

	var total int64 = rand.Int63() % 10000
	Deposit(int(total))
	for _, tt := range tests {
		t.Run(name, func(t *testing.T) {
			isEnd := make(chan struct{}, tt.routines)
			for i := 0; i < tt.routines; i++ {
				go func() {
					if Withdraw(tt.value) {
						atomic.AddInt64(&total, int64(-tt.value))
					}
					isEnd <- struct{}{}
				}()
			}
			for i := 0; i < tt.routines; i++ {
				<-isEnd
			}
			got := Balance()
			info := fmt.Sprintf("Balance() = %v, want %v", got, total)
			if got != int(total) {
				t.Errorf(info)
			} else {
				t.Logf(info)
			}
		})
	}
}
