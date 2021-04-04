package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/cole-ratner/docker-crawler/internal/docker"
)

func main() {

	// setting flags
	host := flag.String("h", "hub.docker.com", "The container registry that you would like to search.")
	var worker string; flag.StringVar(&worker, "w", "", "The name of the worker app to use for enumerating the container. (Required)")
	var searchTerm string; flag.StringVar(&searchTerm, "searchterm", "", "The term that you would like to search within the container registry. (Required)")
	flag.Parse()
	
	// validating non-empty values
	if *host == "" || worker == "" || searchTerm == "" {
		flag.Usage()
		os.Exit(1)
	}

	// enumerating images in the public registry per searchterm
	imageList, err := docker.ListImages(host, searchTerm)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// instantiating a waitgroup and waiting for all spawned worker procs to complete
	var wg sync.WaitGroup	
	docker.ScheduleWorkers(imageList, worker, &wg)
	wg.Wait()
}
