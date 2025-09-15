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
		min, max, total float64
		count           int64
	}
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	stationStats := make(map[string]LocationStats)
	/*
		-- bufio to save on system calls
		-- essentially, we read large chunks
		-- of memory at once to save on calls
	*/
	parser := bufio.NewScanner(file)

	for parser.Scan() {
		/*
			-- process each row in the file
			-- check for existence of semicolon
				-- ensures only clean data is used
		*/
		row := parser.Text()
		station, tempValue, Semicolon := strings.Cut(row, ";")
		if !Semicolon {
			continue
		}

		/* string conversion */
		temperature, err := strconv.ParseFloat(tempValue, 64)
		if err != nil {
			return err
		}

		s, ok := stats[station]
		if !ok {
			s.min = temperature
			s.max = temperature
			s.total = temperature
			s.count = 1
		} else {
			s.min = min(s.min, temperature)
			s.max = max(s.max, temperature)
			s.total += temperature
			s.count += 1
		}
		stationStats[station] = s
	}

	stations := make([]string, 0, len(stationStats))
	for station := range stationStats {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	fmt.Fprint(output, "{")
	for i, station := range stations {
		if i > 0 {
			fmt.Fprint(output, ", ")
		}
		s := stationStats[station]
		mean := s.total / float64(s.count)
		fmt.Fprintf(output, "%s=%.1f/%.1f/%.1f", station, s.min, mean, s.max)
	}
	fmt.Fprint(output, "}\n")
	return nil
}
