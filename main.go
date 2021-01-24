package cda

import (
"fmt"
"github.com/shopspring/decimal"
"sort"
)

type DecimalPriceOrderBook struct {
	IsSell         bool
	PriceVolumeMap map[float64]int
}

func NewDecimalPriceOrderBook(isSell bool) *DecimalPriceOrderBook {
	return &DecimalPriceOrderBook{
		IsSell:         isSell,
		PriceVolumeMap: make(map[float64]int),
	}
}

func (book *DecimalPriceOrderBook) Add(price decimal.Decimal, volume int) {
	if volume == 0{
		return
	}
	priceFloat, _ := price.Float64()
	if _, ok := book.PriceVolumeMap[priceFloat]; ok {
		book.PriceVolumeMap[priceFloat] += volume
		if book.PriceVolumeMap[priceFloat] == 0 {
			delete(book.PriceVolumeMap, priceFloat)
		}
	} else {
		book.PriceVolumeMap[priceFloat] = volume
	}
}
//
//type DecimalSlice []decimal.Decimal
//
//func (p DecimalSlice) Len() int           { return len(p) }
//func (p DecimalSlice) Less(i, j int) bool { return p[i].LessThan(p[j]) }
//func (p DecimalSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (book *DecimalPriceOrderBook) String() string {
	result := "{"
	indexes := make([]float64, 0)
	for price, volume := range book.PriceVolumeMap {
		if price == 0 {
			result += fmt.Sprintf("\"%s\":%d,", decimal.NewFromFloat(price), volume)
		} else {
			indexes = append(indexes, price)
		}
	}
	if book.IsSell {
		sort.Float64s(indexes)
	} else {
		sort.Sort(sort.Reverse(sort.Float64Slice(indexes)))
	}
	for _, key := range indexes {
		result += fmt.Sprintf("\"%s\":%d,", decimal.NewFromFloat(key), book.PriceVolumeMap[key])
	}
	if len(book.PriceVolumeMap) > 0 {
		result = result[:len(result)-1]
	}
	return  result + "}"
}

func (book *DecimalPriceOrderBook) GetBestOrder() (decimal.Decimal, int) {
	if len(book.PriceVolumeMap) == 0 {
		return decimal.Zero, 0
	}
	indexes := make([]float64, 0)
	for price, volume := range book.PriceVolumeMap {
		if price == 0 {
			return decimal.NewFromFloat(price), volume
		} else {
			indexes = append(indexes, price)
		}
	}
	if book.IsSell {
		sort.Float64s(indexes)
	} else {
		sort.Sort(sort.Reverse(sort.Float64Slice(indexes)))
	}
	return decimal.NewFromFloat(indexes[0]), book.PriceVolumeMap[indexes[0]]
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
