package model

type Money float32

func (m Money) Cents() int64 {
	return int64(m * 100)
}

func (m Money) Float32() float32 {
	return float32(m)
}

func NewMoney(money float32) Money {
	return Money(money)
}

func NewMoneyFromCents(cents int64) Money {
	return Money(float32(cents / 100))
}
