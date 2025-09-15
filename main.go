package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var maxGoroutines int

func main() {
	var (
		cpuProfile = flag.String("cpuprofile", "", "write CPU profile to file")
		goroutines = flag.Int("goroutines", 0, "num goroutines for parallel solutions (default NumCPU)")
	)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage: go-r1bc [-cpuprofile=PROFILE] INPUTFILE\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	maxGoroutines = *goroutines
	if maxGoroutines == 0 {
		maxGoroutines = runtime.NumCPU()
	}

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
	}
	inputPath := args[0]

	st, err := os.Stat(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	size := st.Size()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	start := time.Now()
	output := bufio.NewWriter(os.Stdout)

	err = version1(inputPath, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	output.Flush()
	elapsed := time.Since(start)
	fmt.Fprintf(os.Stderr, "Processed %.1fMB in %s\n",
		float64(size)/(1024*1024), elapsed)
}
