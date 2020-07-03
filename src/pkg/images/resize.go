package images

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"

	"github.com/mauleyzaola/maupod/src/pkg/information"
)

func Size(r io.Reader) (x, y int, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := information.InfoString(scanner.Text())
		key, value := line.Split()
		switch key {
		case "Width":
			if val, err := strconv.Atoi(value); err == nil {
				x = val
			}
		case "Height":
			if val, err := strconv.Atoi(value); err == nil {
				y = val
			}
		default:
			continue
		}
		if x != 0 && y != 0 {
			break
		}
	}
	return
}

func ImageResize(source, target string, width, height int) error {
	const program = "convert"
	if !helpers.ProgramExists(program) {
		return fmt.Errorf("could not find program: %s in path", program)
	}
	if source == "" {
		return errors.New("missing parameter: source")
	}
	if target == "" {
		return errors.New("missing parameter: target")
	}

	var p = []string{
		source,
		fmt.Sprintf("-resize=%dx%d", width, height),
		target,
	}
	p = append(p, source)
	cmd := exec.Command(program, p...)
	output := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	cmd.Stdout = output
	cmd.Stderr = errOutput
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s %s : %v", output.String(), errOutput.String(), err)
	}
	return nil
}
