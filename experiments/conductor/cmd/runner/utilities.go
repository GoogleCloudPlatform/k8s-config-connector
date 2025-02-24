// Copyright 2025 Google LLC
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

package runner

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type exitBash func()

func startBash() (io.WriteCloser, io.ReadCloser, exitBash, error) {
	cmd := exec.Command("bash")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		errBuffer := make([]byte, 1000)
		_, err = stderr.Read(errBuffer)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("BASH ERR %s\r\n", string(errBuffer))
	}()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	exit := func() {
		log.Printf("COMMAND: exit\r\n")
		if _, err = stdin.Write([]byte("exit\n")); err != nil {
			log.Fatal(err)
		}
		err := cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("BASH DONE\r\n")
	}
	return stdin, stdout, exit, err
}

func cdRepoBranchDirBash(opts *RunnerOptions, subdir string, stdin io.WriteCloser, stdout io.ReadCloser) string {
	dir := opts.branchRepoDir
	if subdir != "" {
		dir = filepath.Join(dir, subdir)
	}
	log.Printf("COMMAND: cd %s and echo\r\n", dir)
	if _, err := stdin.Write([]byte(fmt.Sprintf("cd %s && echo done\n", dir))); err != nil {
		log.Fatal(err)
	}
	outBuffer := make([]byte, 1000)
	done := false
	var msg string
	for !done {
		length, err := stdout.Read(outBuffer)
		if err != nil {
			log.Fatal(err)
		}
		msg += string(outBuffer[:length])
		done = strings.HasSuffix(msg, "done\n")
	}
	log.Printf("CD OUT %s\r\n", msg)
	return msg
}

type closer func()

func setLoggingWriter(opts *RunnerOptions, branch Branch) closer {
	// Initially force out to stdout in case we hit an error we don't
	// want to pollute a different runs logs with our logs.
	// TODO: Return a log object so we can run in parrellel.
	log.SetOutput(os.Stdout)
	if opts.loggingDir == "" {
		log.Println("Logging dir not set")
		return noOp
	}
	logDir := filepath.Join(opts.loggingDir, branch.Name)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Printf("Error creating logging dir %s, :%v", logDir, err)
			return noOp
		}
	}

	var out *os.File
	var err error
	logFile := filepath.Join(logDir, "out.log")
	if out, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755); err != nil {
		log.Printf("Error opening logging file %s, :%v", logFile, err)
		return noOp
	}
	log.SetOutput(out)

	/*
		var errF *os.File
		errFile := filepath.Join(logDir, "err.log")
		if _, err := os.Stat(errFile); os.IsNotExist(err) {
			errF, err = os.Create(errFile)
			if err != nil {
				return os.Stdout, os.Stderr
			}
		} else {
			errF, err = os.OpenFile(errFile, os.O_APPEND, 0755)
			if err != nil {
				return os.Stdout, os.Stderr
			}
		}
	*/

	return func() {
		// Initially force out to stdout in case we hit an error we don't
		// want to pollute a different runs logs with our logs.
		// TODO: Return a log object so we can run in parrellel.
		log.SetOutput(os.Stdout)
		if err := out.Close(); err != nil {
			log.Printf("Failed to clode logging file %s, :%v", logFile, err)
		}
	}
}

func noOp() {
	return
}
