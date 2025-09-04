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
}
