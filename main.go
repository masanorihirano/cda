package cda

import (
"fmt"
"github.com/shopspring/decimal"
"sort"
)

type DecimalPriceOrderBook struct {
	isSell         bool
	priceVolumeMap map[decimal.Decimal]int
}

func NewDecimalPriceOrderBook(isSell bool) *DecimalPriceOrderBook {
	return &DecimalPriceOrderBook{
		isSell:         isSell,
		priceVolumeMap: make(map[decimal.Decimal]int),
	}
}

func (book *DecimalPriceOrderBook) Add(price decimal.Decimal, volume int) {
	if v, ok := book.priceVolumeMap[price]; ok {
		v += volume
		if v == 0 {
			delete(book.priceVolumeMap, price)
		}
	} else {
		book.priceVolumeMap[price] = volume
	}
}

type DecimalSlice []decimal.Decimal

func (p DecimalSlice) Len() int           { return len(p) }
func (p DecimalSlice) Less(i, j int) bool { return p[i].LessThan(p[j]) }
func (p DecimalSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (book *DecimalPriceOrderBook) String() string {
	result := "{"
	indexes := make(DecimalSlice, 0)
	for price, volume := range book.priceVolumeMap {
		if price.Equal(decimal.New(0, 0)) {
			result += fmt.Sprintf("%s:%d,", price, volume)
		} else {
			indexes = append(indexes, price)
		}
	}
	if book.isSell {
		sort.Sort(indexes)
	} else {
		sort.Sort(sort.Reverse(indexes))
	}
	for _, key := range indexes {
		result += fmt.Sprintf("%s:%d,", key, book.priceVolumeMap[key])
	}
	return result[:len(result)-1] + "}"
}

func (book *DecimalPriceOrderBook) GetBestOrder() (decimal.Decimal, int) {
	indexes := make(DecimalSlice, 0)
	for price, volume := range book.priceVolumeMap {
		if price.Equal(decimal.New(0, 0)) {
			return price, volume
		} else {
			indexes = append(indexes, price)
		}
	}
	if book.isSell {
		sort.Sort(indexes)
	} else {
		sort.Sort(sort.Reverse(indexes))
	}
	return indexes[0], book.priceVolumeMap[indexes[0]]
}

type DecimalPriceCdaMarket struct {
	SellOrderBook *DecimalPriceOrderBook
	BuyOrderBook  *DecimalPriceOrderBook
}

func NewDecimalPriceCdaMarket() *DecimalPriceCdaMarket {
	return &DecimalPriceCdaMarket{
		SellOrderBook: NewDecimalPriceOrderBook(true),
		BuyOrderBook:  NewDecimalPriceOrderBook(false),
	}
}

func (market *DecimalPriceCdaMarket) AddOrder(price decimal.Decimal, volume int, isSell bool) {
	if isSell {
		market.SellOrderBook.Add(price, volume)
	} else {
		market.BuyOrderBook.Add(price, volume)
	}
}

func (market *DecimalPriceCdaMarket) Execution(){
	// ToDo To be implemented
}
