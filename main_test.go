package main

import (
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestDockerRunning(t *testing.T) {
	cmd := exec.Command("docker", "ps")
	cmd.Stderr = os.Stderr
	cmd.Start()

	timer := time.AfterFunc(1*time.Second, func() {
		err := cmd.Process.Kill()
		if err != nil {
			panic(err) // panic as can't kill a process.
		}
	})
	err := cmd.Wait()
	timer.Stop()

	if err != nil {
		t.Errorf("Docker not running. try `boot2docker up`?")
	}
}

func TestFail(t *testing.T) {
	t.Errorf("Fail")
}
