package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/mcandre/goxcart"
)

var flagRemove = flag.Bool("remove", true, "Automatically remove Docker containers upon termination")
var flagImage = flag.String("image", "", "Docker image name, e.g. mcandre/docker-gox")
var flagOutput = flag.String("output", "", "output directory, e.g. bin")
var flagRepository = flag.String("repo", "", "Repository namespace, e.g. github.com/mcandre/go-hextime")
var flagBanner = flag.String("banner", "", "artifact label (required), e.g. hextime-0.0.1")
var flagCommands = flag.String("commands", "", "command paths (required), e.g. ./cmd/...")
var flagVerbose = flag.Bool("verbose", false, "Enable additional logging")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

var signalErrorPattern = regexp.MustCompile("^signal: .+$")

func killContainersAndExit() {
	if err := goxcart.KillContainers(); err != nil {
		panic(err)
	}

	os.Exit(1)
}

func main() {
	flag.Parse()

	config := goxcart.NewPortConfig()

	switch {
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(1)
	case *flagVersion:
		fmt.Println(goxcart.Version)
		os.Exit(0)
	}

	config.RemoveContainer = *flagRemove
	config.Verbose = *flagVerbose

	if *flagImage != "" {
		config.Image = *flagImage
	}

	if *flagOutput != "" {
		config.OutputDirectory = *flagOutput
	}

	if *flagRepository != "" {
		config.Repository = *flagRepository
	}

	if *flagBanner != "" {
		config.Banner = *flagBanner
	} else {
		flag.PrintDefaults()
	}

	if *flagCommands != "" {
		config.CommandPaths = *flagCommands
	} else {
		flag.PrintDefaults()
	}

	if config.Verbose {
		log.Printf("PortConfig: %v", config)
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		<-signalCh
		killContainersAndExit()
	}()

	if err := config.Port(); err != nil {
		if !signalErrorPattern.MatchString(err.Error()) {
			panic(err)
		}
	}
}
