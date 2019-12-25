package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string
	Price  float64
	Volume int
	Buy    bool
}

func NewTrade(symbol string, price float64, volume int, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol cannot be empty")
	}

	if price < 0 {
		return nil, fmt.Errorf("Price cannot be < 0")
	}

	if volume < 0 {
		return nil, fmt.Errorf("Volume cannot be < 0")
	}

	trade := &Trade{symbol, price, volume, buy}
	return trade, nil
}

//Value ... Calculates value of trade
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {
	t, err := NewTrade("APPL", 1000.45, 10, true)
	if err != nil {
		fmt.Printf("Error has occured in creating trade %s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.Value())
}
