package rpcutility

import (
	"net/rpc"
	"errors"
	"sync"
	"time"
	"math/rand"
	"fmt"
)

const HelloServiceName = "path/to/pkg.HelloService"
const DividedServiceName = "path/to/pkg.DividedService"
const WatchServiceName = "path/to/pkg.WatchService"

func init(){
	RegisterHelloServer(new(HelloServer))
	RegisterDividedServer(new(DividedServer))
	RegisterWatchServer(NewKVStoreService())
}

type Number struct{
	A int32
	B int32
}

type KVStoreService struct{
	m 		map[string]string
	filter	map[string]func(key string)
	mu 		sync.Mutex
}

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

type DividedServiceInterface interface{
	Divide(num Number, result *Number) error
}

type WatchServiceInterface interface{
	Watch(timeoutSecond int, keyChanged *string) error
}

func NewKVStoreService() *KVStoreService{
	return &KVStoreService{
		m:		make(map[string]string),
		filter:	make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error{
	p.mu.Lock()
	defer p.mu.Unlock()
	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}
	return errors.New("key not found")
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error{
	p.mu.Lock()
	defer p.mu.Unlock()
	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value{
		for _, fn := range p.filter{
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}

func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error{
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)

	p.mu.Lock()
	p.filter[id] = func(key string) {ch <- key}
	p.mu.Unlock()

	select{
	case <- time.After(time.Duration(timeoutSecond) * time.Second):
		return errors.New("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
	return nil
}

func RegisterHelloServer(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func RegisterDividedServer(svc DividedServiceInterface) error {
	return rpc.RegisterName(DividedServiceName, svc)
}

func RegisterWatchServer(svc WatchServiceInterface) error {
	return rpc.RegisterName(WatchServiceName, svc)
}

type HelloServer struct{}

type DividedServer struct{}

func (h *HelloServer) Hello(request string, reply *string) error {
	*reply = "Hello : " + request
	return nil
}

func (d *DividedServer) Divide(num Number, result *Number) error {
	if num.B == 0 {
		return errors.New("Invalid Divisor")
	}
	result.A = num.A / num.B
	result.B = num.A % num.B
	return nil
}