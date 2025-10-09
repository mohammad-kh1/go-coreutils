package helpers

import (
	"syscall"
	"os"
	"runtime"
)

func getTotalMemory() uint64 {
	var info syscall.Sysinfo_t
	err := syscall.Sysinfo(&info)
	if err != nil {
		return 0
	}
	return info.Totalram * uint64(info.Unit)
}

func SetBufferSize(path string) int {
	const (
		KB = 1024
		MB = 1024 * KB
	)

	//Default buffer size
	buffSize := 4 * KB

	// Check for file size
	info , err := os.Stat(path)
	if err == nil {
		size := info.Size()
		if size > 100 * MB {
			buffSize = 	32 * KB
		}
	}

	// Check system memory
	mem := getTotalMemory()
	if mem > 0 && mem < 256 * MB {
		buffSize = 2 * KB
	}

	//Check Arch
	if runtime.GOARCH == "arm" || runtime.GOARCH == "arm64"{
		buffSize = 2 * KB
	}

	return buffSize

}