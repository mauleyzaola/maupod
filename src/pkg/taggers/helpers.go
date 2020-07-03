package taggers

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
)

func run(program, filename string, params ...string) error {
	if filename == "" {
		return errors.New("missing parameter filename")
	}
	if _, err := os.Stat(filename); err != nil {
		return err
	}
	if len(params) == 0 {
		return errors.New("mising parameter: params")
	}
	if !helpers.ProgramExists(program) {
		return fmt.Errorf("could not find program: %s", program)
	}

	var pars = []string{}
	pars = append(pars, params...)
	pars = append(pars, filename)
	cmd := exec.Command(program, pars...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
