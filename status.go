package kite

import (
	"os/user"
	"runtime"
)

type Status struct{}

type Info struct {
	State       string `json:"state"`
	DiskUsage   uint64 `json:"diskUsage"`
	DiskTotal   uint64 `json:"diskTotal"`
	MemoryUsage uint64 `json:"memoryUsage"`
	MemoryTotal uint64 `json:"totalMemoryLimit"`
	HomeDir     string `json:"homeDir"`
	Uname       string `json:"uname"`
}

type memory struct {
	Usage uint64 `json:"memoryUsage"`
	Total uint64 `json:"memoryTotal"`
}

type disk struct {
	Usage uint64 `json:"diskUsage"`
	Total uint64 `json:"diskTotal"`
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}

	return usr.HomeDir
}

func systemInfo() (*Info, error) {
	disk, err := diskStats()
	if err != nil {
		return nil, err
	}

	mem, err := memoryStats()
	if err != nil {
		return nil, err
	}

	return &Info{
		State:       "RUNNING", // needed for client side compatibility
		DiskUsage:   disk.Usage,
		DiskTotal:   disk.Total,
		MemoryUsage: mem.Usage,
		MemoryTotal: mem.Total,
		HomeDir:     homeDir(),
		Uname:       runtime.GOOS,
	}, nil
}

func (Status) Info(r *Request) (interface{}, error) {
	return systemInfo()
}
