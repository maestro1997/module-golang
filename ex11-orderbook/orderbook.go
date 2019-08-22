//package orderbook
package main

type Trade struct {
	Bid    *Order
	Ask    *Order
	Volume uint64
	Price  uint64
}

type Orderbook struct {
	Bids []MyOrder
	Asks []MyOrder
}

type MyOrder struct {
	Id    int
	Kind  byte     // 2 - LIMIT, 1 - MARKET
	Count uint64
	Price uint64
}

type Order struct {
	Id     int
	Side   byte
	Kind   byte
	Count  uint64
	Price  uint64
}

func New() *Orderbook {
	Bids := make([]MyOrder,0)
	Asks := make([]MyOrder,0)
	Orderbook := Orderbook{Bids,Asks}
	return &Orderbook
}

func handle_bid(orderbook *Orderbook, order *MyOrder) ([]*Trade, *Order) {
	return nil, nil
}

func handle_ask(orderbook *Orderbook, order *MyOrder) ([]*Trade, *Order) {
	return nil, nil
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	MyOrder := MyOrder{order.Id, order.Kind, order.Count, order.Price}
	if order.Side == 1 {
		return handle_bid(orderbook, &MyOrder)
	}
	return handle_ask(orderbook, &MyOrder)
}

func main() {

}
