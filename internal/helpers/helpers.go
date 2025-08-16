package helpers

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/aybangueco/dotsync/internal/config"
)

func ExpandPath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[1:]), nil
	}
	return path, nil
}

func ValidateConfig(conf []config.DotSyncConfig) error {
	for _, c := range conf {
		if c.Target == "" {
			return fmt.Errorf("%s's target path must not be empty", c.Source)
		}

		if c.Source == "" {
			return errors.New("one of the sources are empty, either remove it or add some source")
		}
	}

	return nil
}

func RemoveFromTarget(c config.DotSyncConfig, target string) error {
	if c.IsDirectory {
		rmDir := exec.Command("rm", "-rf", fmt.Sprintf("%s/%s", target, c.Source))
		output, err := rmDir.CombinedOutput()
		if err != nil {
			return fmt.Errorf("Error removing directory: %v\nOutput: %s", err, string(output))
		}
	} else {
		rm := exec.Command("rm", fmt.Sprintf("%s/%s", target, c.Source))
		output, err := rm.CombinedOutput()
		if err != nil {
			return fmt.Errorf("Error removing file: %v\nOutput: %s", err, string(output))
		}
	}

	return nil
}

func RemoveFromSource(c config.DotSyncConfig) error {
	if c.IsDirectory {
		rmDir := exec.Command("rm", "-rf", c.Source)
		output, err := rmDir.CombinedOutput()
		if err != nil {
			return fmt.Errorf("Error removing directory: %v\nOutput: %s", err, string(output))
		}
	} else {
		rm := exec.Command("rm", c.Source)
		output, err := rm.CombinedOutput()
		if err != nil {
			return fmt.Errorf("Error removing file: %v\nOutput: %s", err, string(output))
		}
	}

	return nil
}
