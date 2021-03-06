// The MIT License (MIT)
//
// Copyright (c) 2016,2017  aerth <aerth@riseup.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package spawn a process like a salmon
package spawn

import (
	"os"
	"os/exec"
	"strconv"
)

const (
	// SPAWNTIME is an environmental variable that your application may use to count spawn depth.
	// SPAWNTIME of 1 means first spawn, SPAWNTIME=2 means was spawned from a spawn, etc
	SPAWNTIME = "SPAWNTIME"
)

// Spawn better than a salmon!
func Spawn() error {

	// Increment SPAWNTIME count
	i, _ := strconv.Atoi(os.Getenv(SPAWNTIME))
	i++
	os.Setenv(SPAWNTIME, strconv.Itoa(i))

	// Spawned process has new environmental variable: SPAWNED=true
	os.Setenv("SPAWNED", "true")
	me, medir, args := Exe()

	cmd := exec.Command(me, args...)
	cmd.Dir = medir

	return cmd.Start()
}

// Destroy is the same as os.Exit(0) for now.
func Destroy() {
	os.Exit(0)
}
