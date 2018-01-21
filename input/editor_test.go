package input

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

var exitCode = "0"

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", fmt.Sprintf("GO_WANT_EXIT_CODE=%s", exitCode)}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	if exitCode := os.Getenv("GO_WANT_EXIT_CODE"); exitCode != "0" {
		fmt.Printf("error code %s\n", exitCode)
		code, _ := strconv.Atoi(exitCode)
		os.Exit(code)
	}

	filename := os.Args[len(os.Args)-1]
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to open temp file for writing test post: %s", err.Error())
		os.Exit(1)
	}
	str := string(bytes)
	str = "this is the result" + str
	ioutil.WriteFile(filename, []byte(str), os.ModeAppend)

	os.Exit(0)
}

func TestEditor_Read(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	expectedResult := "this is the result"

	e, err := newEditor(nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := e.Read()
	if err != nil {
		t.Fatal(err)
	}

	if *res != expectedResult {
		t.Errorf("Output is different than expected ('%s' does not equal '%s')", *res, expectedResult)
	}
}

func TestEditor_Read_Fail(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	exitCode = "1"

	e, err := newEditor(nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = e.Read()
	if err == nil {
		t.Fatal(err)
	}
}
