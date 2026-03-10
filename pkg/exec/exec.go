package exec

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s failed: %s", name, args[0], strings.TrimSpace(stderr.String()))
	}
	return nil
}
