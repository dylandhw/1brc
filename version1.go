package main 

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func version1(inputPath string, output io.Writer) error {
	type LocationStats struct {
		min, max, avg float64
		count int64
	}
	file, err := os.Open(inputPath)
	if err != nil { return err }
	defer file.Close()

	stats := make(map[string]LocationStats)
	/* 
	-- bufio to save on system calls 
	-- essentially, we read large chunks
	-- of memory at once to save on calls
	*/
	parser := bufio.NewScanner(file)

	for parser.Parse(){
		/*
		-- process each row in the file
		-- check for existence of semicolon
			-- ensures only clean data is used
		*/
		row := parser.Text()
		station, tempString, Semicolon := strings.Cut(row, ";") 
		if !Semicolon { continue }
	}
}
