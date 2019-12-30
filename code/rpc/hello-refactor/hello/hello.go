package hello

import "net/rpc"

// ServiceName is the name of the hello service.
const ServiceName = "hello/HelloService"

// ServiceInterface is the interface for the hello service.
type ServiceInterface interface {
	Hello(request string, reply *string) error
}

// RegisterService is used to register a hello service to RPC.
func RegisterService(svc ServiceInterface) error {
	return rpc.RegisterName(ServiceName, svc)
}

// ServiceClient is the client of the hello service.
type ServiceClient struct {
	*rpc.Client
}

// DialService is used to connect a hello service for a client.
func DialService(network, address string) (*ServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &ServiceClient{Client: client}, nil
}

// Hello is the RPC function of the hello service.
func (p *ServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(ServiceName+".Hello", request, reply)
}
