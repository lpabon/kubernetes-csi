package mock

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	mock_driver "github.com/csi-volumes/kubernetes-csi/mock/driver"
	gomock "github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestPluginInfoResponse(t *testing.T) {

	// Setup mock
	m := gomock.NewController(t)
	defer m.Finish()
	driver := mock_driver.NewMockIdentityServer(m)

	// Setup input
	in := &csi.GetPluginInfoRequest{
		Version: &csi.Version{
			Major: 0,
			Minor: 1,
			Patch: 0,
		},
	}

	// Setup mock outout
	out := &csi.GetPluginInfoResponse{
		Reply: &csi.GetPluginInfoResponse_Result_{
			Result: &csi.GetPluginInfoResponse_Result{
				Name:          "mock",
				VendorVersion: "0.1.1",
				Manifest: map[string]string{
					"hello": "world",
				},
			},
		},
	}

	// Setup expectation
	driver.EXPECT().GetPluginInfo(nil, in).Return(out, nil).Times(1)

	// Actual call
	r, err := driver.GetPluginInfo(nil, in)
	name := r.GetResult().GetName()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if name != "mock" {
		t.Errorf("Unknown name: %s\n", name)
	}
}

func TestGRPCGetPluginInfoReponse(t *testing.T) {

	// Setup mock
	m := gomock.NewController(t)
	defer m.Finish()
	driver := mock_driver.NewMockIdentityServer(m)

	// Setup input
	in := &csi.GetPluginInfoRequest{
		Version: &csi.Version{
			Major: 0,
			Minor: 1,
			Patch: 0,
		},
	}

	// Setup mock outout
	out := &csi.GetPluginInfoResponse{
		Reply: &csi.GetPluginInfoResponse_Result_{
			Result: &csi.GetPluginInfoResponse_Result{
				Name:          "mock",
				VendorVersion: "0.1.1",
				Manifest: map[string]string{
					"hello": "world",
				},
			},
		},
	}

	// Setup expectation
	// !IMPORTANT!: Must set context expected value to gomock.Any() to match any value
	driver.EXPECT().GetPluginInfo(gomock.Any(), in).Return(out, nil).Times(1)

	// Create a new RPC
	server := mock_driver.NewMockCSIDriver()
	err := server.Start(driver)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer server.Stop()

	// Actual call
	conn, err := grpc.Dial(server.Address(), grpc.WithInsecure())
	if err != nil {
		t.Errorf(err.Error())
	}
	defer conn.Close()

	// Make call
	c := csi.NewIdentityClient(conn)
	r, err := c.GetPluginInfo(context.Background(), in)
	name := r.GetResult().GetName()
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if name != "mock" {
		t.Errorf("Unknown name: %s\n", name)
	}
}
