package main

import (
	"flag"
	"fmt"
	"os/exec"
	"os"
	"path"
	"strings"
)

const (
	baseOutputPath = "/tmp/output"
)


func main() {
	var image string 
	flag.StringVar(&image, "image", "", "The name of an image to scan.")
	flag.Parse()

	if image == "" {
		flag.Usage()
		os.Exit(1)
	}

	var pathPrefix 	string
	var imageName 	string = image
	if strings.Contains(imageName, "/") {
		imageParts := strings.Split(imageName, "/")
		pathPrefix = imageParts[0]
		imageName = imageParts[1]
	}
	
	fileName := fmt.Sprintf("%s.txt", imageName)
	outPath := path.Join(baseOutputPath, pathPrefix)
	fullPath := path.Join(outPath, fileName)
	err := os.Mkdir(outPath, 0777)
	if err != nil && !os.IsExist(err) {
		fmt.Println(err)
	}

	f, err := os.Create(fullPath)
	defer f.Close()

	out, err := exec.Command("docker", "history", image).CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Now enumerating: %s\n", image)
	f.Write(out)
}