package e2e

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

var (
	binName = "counter-test"
)

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", binName, "../..")

	buf := &bytes.Buffer{}
	cmd.Stderr = buf

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to build binary:a ", err, buf.String())
		os.Exit(1)
	}

	result := m.Run()

	os.Remove(binName)
	os.Exit(result)
}
