package docker

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/shirou/gopsutil/v3/cpu"
)

// ListImages queries a public container registry for images per search term
func ListImages(host *string, searchTerm string) ([]string, error) {
	out, err := exec.Command("docker", "search", searchTerm, "--format", "{{.Name}}").Output()
	if err != nil {
		return nil, err
	}

	imageList := strings.Split(string(out), "\n")
	fmt.Printf("Found %d images under your requested search term.\n", len(imageList))
	return imageList, nil
}

func getCPUInfo() (int, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return 0, err
	}
	return len(cpuStat), nil
}

// ScheduleWorkers schedules workers in go routines
func ScheduleWorkers(images []string, workerImage string, wg *sync.WaitGroup) error {
	fmt.Printf("Now Scheduling Workers.\n")
	cores, err := getCPUInfo()
	if err != nil {
		return err
	}

	runtime.GOMAXPROCS(cores)
	fmt.Printf("Found %d cores. Limiting max processes to %d...\n\n", cores, cores)
	
	c := make(chan string)
	for _, i := range images {
		wg.Add(1)
		wArgs := []string{"run", "-v", "/home/data:/tmp", "-v", "/var/run/docker.sock:/var/run/docker.sock", workerImage, "-image", i}
		w := NewWorker("/bin/docker", wArgs, c)
		go w.Run(wg)
	}
	go Collect(c)
	return nil
}
