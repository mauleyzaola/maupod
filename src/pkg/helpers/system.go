package helpers

import "os/exec"

func ProgramExists(programName string) bool {
	_, err := exec.LookPath(programName)
	return err == nil
}
