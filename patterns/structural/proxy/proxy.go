// Instead of calling the THING you want to call
// you call the THING which call the THING you want !!"
package proxy

import "fmt"

type service interface {
	DoSomething()
}

type realService struct {
	ServiceName string
}

func (s realService) DoSomething() {
	// do something
	fmt.Println("Doing something")
}

func newServiceProxy(serviceName string) service {
	return realServiceProxy{realService{serviceName}}
}

type realServiceProxy struct {
	realService
}

func (p realServiceProxy) DoSomething() {
	fmt.Println("Doing something with proxy")
	p.realService.DoSomething()
}

func RunProxy() {
	service := newServiceProxy("service")
	service.DoSomething()
}
