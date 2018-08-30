package backends

import (
	"fmt"
	"testing"

	"github.com/SmartMeshFoundation/Spectrum/core"
)

func TestBackend(t *testing.T) {
	sb := NewSimulatedBackend(core.GenesisAlloc{})
	fmt.Println(sb)
}
