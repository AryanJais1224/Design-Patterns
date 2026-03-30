type PaymentStrategy interface {
    Pay(amount int)
}

type CreditCard struct{}
func (c CreditCard) Pay(amount int) {}

type UPI struct{}
func (u UPI) Pay(amount int) {}

type PaymentContext struct {
    strategy PaymentStrategy
}

func (p *PaymentContext) Execute(amount int) {
    p.strategy.Pay(amount)
}
