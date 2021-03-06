// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package util contains shared functions for several AoC days.
package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ReadLines returns the contents of a text file as a slice of strings representing the lines. The
// newline separators are not kept. The last line need not have a newline character at the end.
func ReadLines(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading lines: %v", err)
	}
	lines := strings.Split(string(data), "\n")
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines, nil
}

// ReadChunks returns the contents of a text file as a slice of strings representing all paragraphs,
// as defined by text separated by a blank line (two consecutive newlines).
func ReadChunks(path string) (chunks []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(ScanChunks)
	for s.Scan() {
		chunks = append(chunks, s.Text())
	}
	return chunks, nil
}

// ReadIntRows parses a text file formatted as one integer per line.
func ReadIntRows(path string) ([]int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	var ints []int
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("parsing ints from %s: %v", path, err)
		}
		ints = append(ints, i)
	}
	return ints, nil
}

// P represents a two-dimensional integer-valued coordinate.
type P struct {
	X, Y int
}

// ScanChunks implements a bufio.SplitFunc for scanning paragraphs delimited by a blank line
// (i.e., two consecutive '\n' bytes).
func ScanChunks(data []byte, atEOF bool) (advance int, token []byte, err error) {
	delim := []byte{'\n', '\n'}

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, delim); i >= 0 {
		return i + 2, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func (p P) Neigh() [4]P {
	return [4]P{{p.X, p.Y - 1}, {p.X, p.Y + 1}, {p.X - 1, p.Y}, {p.X + 1, p.Y}}
}
