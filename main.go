package main

import (
	"fmt"
	"sync"
	"time"

	memdb "github.com/wolvenspirit/memref/pkg"
)

type myData struct {
	Name string
	m    sync.Mutex
}

func (data *myData) mutate(value string) {
	data.m.Lock()
	data.Name = value
	data.m.Unlock()
}

var (
	Store = memdb.Storage[myData]{}
	stop  = make(chan int, 1)
)

func anotherContext() {
	time.Sleep(time.Second * 1)
	b := Store.Get("clearance-2")
	fmt.Printf("%+v", b[0].Value.Name)
	stop <- 1
}
func example() {

	someId := "sasfasf"
	someOtherId := "asfasfasd"
	v := memdb.Entity[myData]{Key: memdb.EntityKey{Surnames: []string{"top-secret", "clearance-1"}, Id: someId}, Value: myData{Name: "very important secret data"}}
	vv := memdb.Entity[myData]{Key: memdb.EntityKey{Surnames: []string{"top-secret", "clearance-2"}, Id: someOtherId}, Value: myData{Name: "confidential data"}}
	Store.Set(&v)
	Store.Set(&vv)
	l := Store.Get("top-secret")
	fmt.Printf("%+v %+v\n", l[0].Value.Name, l[1].Value.Name)
	go anotherContext()
	vv.Value.mutate("updated confidential data")
	<-stop
}

func main() {
	example()
}
