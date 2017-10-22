package driver

//go:generate mockgen -source=$GOPATH/src/github.com/container-storage-interface/spec/lib/go/csi/csi.pb.go -imports .=github.com/container-storage-interface/spec/lib/go/csi -package=driver -destination=driver.mock.go
import (
	"fmt"
	"net"
	"sync"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Check is To be removed
func Check() {
	fmt.Printf("Check")
}

type MockCSIDriverServer struct {
	Listener net.Listener
	server   *grpc.Server
	wg       sync.WaitGroup
}

func NewMockCSIDriver() *MockCSIDriverServer {
	return &MockCSIDriverServer{}
}

func (m *MockCSIDriverServer) goServe() {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		m.server.Serve(m.Listener)
	}()
}

func (m *MockCSIDriverServer) Address() string {
	return m.Listener.Addr().String()
}
func (m *MockCSIDriverServer) Start(s csi.IdentityServer) error {

	// Listen on a port assigned by the net package
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	m.Listener = l

	m.server = grpc.NewServer()
	csi.RegisterIdentityServer(m.server, s)
	reflection.Register(m.server)

	m.goServe()
	return nil
}

func (m *MockCSIDriverServer) Stop() error {
	m.server.Stop()
	m.wg.Wait()
	return nil
}
