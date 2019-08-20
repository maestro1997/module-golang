package main

import (
   "fmt"
   "sync"
   "time"
)

var wg sync.WaitGroup
var cutting_time = 5 // time of hair cutting
var walk_time = 3    // time of walk of client that can't seat on chair
var check_time = 1   // time for what barber checks waiting room





type Barber struct {
    client chan Client
    waiters chan Client
    state int   // 0 -sleeps, 1 - cuts, 2 - checks room
}

type Client struct {
    id int
}

func CreateClient(id int) Client {
    
    return Client{id}
}

func CreateBarber (room_size int) Barber {
    barber := Barber{make (chan Client,1), make (chan Client,room_size),0}
    return barber
}

func (b Barber) run_barber () {
    fmt.Println("Barber is sleeping\n")
    for {
         client := <-b.client
	 b.haircut(client)
	 b.check_room()
	 select  {		 
             case c:= <-b.waiters:
       		 b.client <- c
		 fmt.Printf("Barber takes client %d\n",c.id)
	     default:
		 fmt.Printf("Barber is going to sleep\n")
	 }
    }
}

func (b Barber) check_room() {
    time.Sleep(time.Second)
}

func (b Barber) haircut (client Client) {
   fmt.Printf("Barber serves client %d\n",client.id)
   time.Sleep(time.Second)
   <-b.client
   fmt.Printf("Barber cuts client client %d\n",client.id)
   wg.Done()
}

func (c Client) run_client (b *Barber, flag int) {
    var way string
    if flag == 0 {
        way = "came"
    } else {
        way = "return"
    }	
    fmt.Printf("Client %d %s to barbershop\n",c.id,way)
    for {
         select {
	 case <-b.client:
	     b.client <- c
	     fmt.Printf("Client %d wokes up the barber\n",c.id)
	 default : 
	     select {
	     case  b.waiters <- c:
	         fmt.Printf("Client %d seats in chair\n",c.id)
		 break
	     default:
	         fmt.Printf("Client %d went for walk from barbershop\n",c.id)
		 time.Sleep(time.Second)
		 c.run_client(b,1)
	     }
	 }
    }
}

func main() {
    room_size := 1
    client_count:= 1
    barber := CreateBarber(room_size)
    go barber.run_barber()
    time.Sleep(time.Second)
    wg.Add(client_count)
    for i:= 0; i < client_count; i++ {
        go CreateClient(i).run_client(&barber,0)
    }
    wg.Wait()
 }
