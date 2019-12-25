package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	IsBuy  bool
}

//Value ... Calculates value of trade
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.IsBuy {
		value = -value
	}
	return value
}

func main() {
	t1 := Trade{"MSFT", 100, 99.96, true}
	fmt.Printf("T1 : %+v\n", t1)
	fmt.Println(t1.Value())

	t2 := Trade{
		Symbol: "APPL",
		Price:  1000,
		Volume: 100,
		IsBuy:  true,
	}
	fmt.Println(t2.Symbol)

	t3 := Trade{}
	fmt.Println(t3)
}
