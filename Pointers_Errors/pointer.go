package Pointerserrors

import "fmt"

type Money int
type Bank struct {
	amount Money
}

func (b *Bank) Deposit(value Money) {
	b.amount += value
}
func (b *Bank) Display() Money {
	return b.amount
}
func (m Money) String() string {
	return fmt.Sprintf("%d BTC", m)
}
func (b *Bank) Withdraw(value Money) {
	b.amount -= value
}
