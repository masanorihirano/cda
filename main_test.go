package cda

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestDecimalPriceOrderBook(t *testing.T) {
	sellOrderBook := NewDecimalPriceOrderBook(true)
	buyOrderBook := NewDecimalPriceOrderBook(false)

	sellOrderBook.Add(decimal.Zero, 100)
	buyOrderBook.Add(decimal.Zero, 100)
	sellOrderBook.Add(decimal.Zero, 100)
	buyOrderBook.Add(decimal.Zero, 100)
	price, volume := sellOrderBook.GetBestOrder()
	if !price.Equal(decimal.Zero){
		t.Errorf("Error 1: expected: %s, got %s",decimal.Zero, price)
	}
	if volume != 200 {
		t.Errorf("Error 2: expected: %d, got %d", 200, volume)
	}
	buyOrderBook.Add(decimal.New(100, 0), 100)
	price, volume = buyOrderBook.GetBestOrder()
	if !price.Equal(decimal.Zero){
		t.Errorf("Error 3: expected: %s, got %s",decimal.Zero, price)
	}
	if volume != 200 {
		t.Errorf("Error 4: expected: %d, got %d", 200, volume)
	}

	sellOrderBook.Add(decimal.New(100,0), 100)
	buyOrderBook.Add(decimal.New(100,0), 100)
	sellOrderBook.Add(decimal.New(200,0), 100)
	buyOrderBook.Add(decimal.New(200,0), 100)
	price, volume = sellOrderBook.GetBestOrder()
	if !price.Equal(decimal.Zero){
		t.Errorf("Error 5: expected: %s, got %s",decimal.Zero, price)
	}
	if volume != 200 {
		t.Errorf("Error 6: expected: %d, got %d", 200, volume)
	}
	price, volume = buyOrderBook.GetBestOrder()
	if !price.Equal(decimal.Zero){
		t.Errorf("Error 7: expected: %s, got %s",decimal.Zero, price)
	}
	if volume != 200 {
		t.Errorf("Error 8: expected: %d, got %d", 200, volume)
	}

	sellString := sellOrderBook.String()
	buyString := buyOrderBook.String()
	if sellString != "{0:200,100:100,200:100}"{
		t.Errorf("Error 9: expected: %s, got %s", "{0:200,100:100,200:100}", sellString)
	}
	if buyString != "{0:200,200:100,100:200}"{
		t.Errorf("Error 10: expected: %s, got %s", "{0:200,200:100,100:100}", sellString)
	}

	sellOrderBook.Add(decimal.Zero, -200)
	buyOrderBook.Add(decimal.Zero, -200)
	price, volume = sellOrderBook.GetBestOrder()
	if !price.Equal(decimal.New(100,0)){
		t.Errorf("Error 11: expected: %s, got %s",decimal.New(1000,0), price)
	}
	if volume != 100 {
		t.Errorf("Error 12: expected: %d, got %d", 100, volume)
	}
	price, volume = buyOrderBook.GetBestOrder()
	if !price.Equal(decimal.New(200,0)){
		t.Errorf("Error 13: expected: %s, got %s",decimal.New(200,0), price)
	}
	if volume != 100 {
		t.Errorf("Error 14: expected: %d, got %d", 100, volume)
	}

	sellOrderBook.Add(decimal.New(100,0), -100)
	buyOrderBook.Add(decimal.New(200,0), -100)

	price, volume = sellOrderBook.GetBestOrder()
	if !price.Equal(decimal.New(200,0)){
		t.Errorf("Error 15: expected: %s, got %s",decimal.New(200,0), price)
	}
	if volume != 100 {
		t.Errorf("Error 16: expected: %d, got %d", 100, volume)
	}
	price, volume = buyOrderBook.GetBestOrder()
	if !price.Equal(decimal.New(100,0)){
		t.Errorf("Error 17: expected: %s, got %s",decimal.New(100,0), price)
	}
	if volume != 200 {
		t.Errorf("Error 18: expected: %d, got %d", 200, volume)
	}

	sellOrderBook.Add(decimal.New(200,0), -100)
	buyOrderBook.Add(decimal.New(100,0), -100)

	price, volume = sellOrderBook.GetBestOrder()
	if !price.Equal(decimal.Zero){
		t.Errorf("Error 15: expected: %s, got %s",decimal.Zero, price)
	}
	if volume != 0 {
		t.Errorf("Error 16: expected: %d, got %d", 0, volume)
	}
	price, volume = buyOrderBook.GetBestOrder()
	if !price.Equal(decimal.New(100, 0)){
		t.Errorf("Error 17: expected: %s, got %s",decimal.New(100, 0), price)
	}
	if volume != 100 {
		t.Errorf("Error 18: expected: %d, got %d", 100, volume)
	}
	sellString = sellOrderBook.String()
	buyString = buyOrderBook.String()
	if sellString != "{}"{
		t.Errorf("Error 9: expected: %s, got %s", "{}", sellString)
	}
	if buyString != "{100:100}"{
		t.Errorf("Error 10: expected: %s, got %s", "{100:100}", sellString)
	}

}

