package host

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Host int

const (
	Null Host = iota
	RPi
	BBB
	CubieTruck
	Galileo
)

func execOutput(name string, arg ...string) (output string, err error) {
	var out []byte
	if out, err = exec.Command(name, arg...).Output(); err != nil {
		return
	}
	output = string(out)
	return
}

func nodeName() (string, error) {
	return execOutput("uname", "-n")
}

func parseVersion(str string) (major, minor, patch int, err error) {
	parts := strings.Split(str, ".")
	len := len(parts)

	if major, err = strconv.Atoi(parts[0]); err != nil {
		return 0, 0, 0, err
	}
	if minor, err = strconv.Atoi(parts[1]); err != nil {
		return 0, 0, 0, err
	}
	if len > 2 {
		part := parts[2]
		part = strings.TrimSuffix(part, "+")
		if patch, err = strconv.Atoi(part); err != nil {
			return 0, 0, 0, err
		}
	}

	return major, minor, patch, err
}

func kernelVersion() (major, minor, patch int, err error) {
	output, err := execOutput("uname", "-r")
	if err != nil {
		return 0, 0, 0, err
	}

	return parseVersion(output)
}

func Detect() (Host, int, error) {
	major, minor, patch, err := kernelVersion()
	if err != nil {
		return Null, 0, err
	}

	if major < 3 || (major == 3 && minor < 8) {
		err = fmt.Errorf("embd: linux kernel versions lower than 3.8 are not supported. you have %v.%v.%v", major, minor, patch)
		return Null, 0, err
	}

	node, err := nodeName()
	if err != nil {
		return Null, 0, err
	}

	var host Host
	var rev int

	switch node {
	case "raspberrypi":
		host = RPi
	case "beaglebone":
		host = BBB
	default:
		err = fmt.Errorf("embd: your host %q is not supported at this moment. please request support at https://github.com/kidoman/embd/issues", node)
		return Null, 0, err
	}

	return host, rev, nil
}
