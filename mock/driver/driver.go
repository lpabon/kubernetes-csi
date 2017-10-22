package driver

//go:generate mockgen -source=$GOPATH/src/github.com/container-storage-interface/spec/lib/go/csi/csi.pb.go -imports .=github.com/container-storage-interface/spec/lib/go/csi -package=driver -destination=driver.mock.go
import (
	"fmt"
	//	_ "github.com/container-storage-interface/spec/lib/go/csi"
)

// Check is To be removed
func Check() {
	fmt.Printf("Check")
}
