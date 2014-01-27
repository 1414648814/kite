package kite

import (
	"fmt"
	"testing"
)

func TestInfo(t *testing.T) {
	i, err := systemInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	fmt.Printf("info: %+v\n", i)

	fmt.Printf("MemoryTotal: %dM\n", i.MemoryTotal/1024/1024)
	fmt.Printf("MemoryUsage: %dM\n", i.MemoryUsage/1024/1024)
	fmt.Printf("DiskTotal: %dG\n", i.DiskTotal/1024/1024)
	fmt.Printf("DiskUsage: %dG\n", i.DiskUsage/1024/1024)

	if i.MemoryTotal == 0 {
		t.Errorf("unexpected memory total", i.MemoryTotal)
	}

	if i.MemoryUsage == 0 || i.MemoryUsage > i.MemoryTotal {
		t.Errorf("unexpected memory usage", i.MemoryUsage)
	}
}
