package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup // counter of clients

type Barber struct {
	sync.Mutex
	state int // 0 sleeping, 1 cutting, 2 checking 
}

func (b *Barber) String() string {
	if b.state == 0 {
		return "sleeping"
	}
	if b.state == 1 {
		return "cutting"
	}
	return "checking"
}

type Client struct {
	num int
}

func (c *Client) String() string {
	return fmt.Sprintf("Client N%d",c.num)
}

func NewBarber() (b *Barber) {
	return &Barber{
		state: 0,
	}
}

func barber(b *Barber, wr chan *Client, wakers chan *Client) {
	for {
		b.Lock()
		defer b.Unlock()
		b.state = 2
		fmt.Printf("Checking waiting room: %d clients waiting\n", len(wr))
		time.Sleep(time.Millisecond * 100)
		select {
		case c := <-wr:
			fmt.Printf("Barber takes %s\n",c)
			HairCut(c, b)
			b.Unlock()
		default:
			fmt.Printf("Barber going to sleep\n")
			b.state = 0
			b.Unlock()
			c := <-wakers
			b.Lock()
			fmt.Printf("Barber woken by %s\n", c)
			HairCut(c, b)
			b.Unlock()
		}
	}
}

func HairCut(c *Client, b *Barber) {
	b.state = 1
	b.Unlock()
	fmt.Printf("Cutting hair of %s\n", c)
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("Barber finished serving of %s\n",c)
	fmt.Printf("%s left barbershop\n",c)
	b.Lock()
	wg.Done()
}

func run_client(c *Client, b *Barber, wr chan<- *Client, wakers chan<- *Client,mode int) {
	time.Sleep(time.Millisecond * 50)
	if mode == 0 {
		fmt.Printf("%s came into barbershop\n",c)
	} else {
		fmt.Printf("%s returned to barbershop\n",c)
    }
	flag := 0
	if len(wr) > 0 {
		flag = 1
	}
	b.Lock()
	if flag == 0 {
		fmt.Printf("%s checks, barber %s | in room: %d\n",c, b , len(wr))
	}
	switch b.state {
	case 1:
		select {
		case wr <- c:
			fmt.Printf("%s seats in chair\n",c)
		default:
			fmt.Printf("No free chairs, %s will return soon\n",c)
			time.Sleep(2*time.Second)
			go run_client(c, b, wr, wakers, 1)
		}
	case 0:
		wakers <- c
	}
	b.Unlock()
}

func main() {
	b := NewBarber()
	room_size := 3
	cl_count := 6
	WaitingRoom := make(chan *Client, room_size)
	Wakers := make(chan *Client, 1)
	go barber(b, WaitingRoom, Wakers)
	time.Sleep(time.Second)
	wg = new(sync.WaitGroup)
	wg.Add(cl_count)
	for i := 0; i < cl_count; i++ {
		time.Sleep(time.Millisecond * 50)
		c := new(Client)
		c.num = i
		go run_client(c, b, WaitingRoom, Wakers,0)
	}
	wg.Wait()
	fmt.Println("No more clients for the day")
}
