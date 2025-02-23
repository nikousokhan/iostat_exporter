package connector

import (
	"bytes"
	"os/exec"
)

func RunIostat() (string, error) {
	cmd := exec.Command("iostat", "-xd", "5")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}
