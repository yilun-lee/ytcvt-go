package utils

import (
	"os/exec"
	"strings"
)

type Subprocess interface {
	// run
	Run() (string, error)
	// post
	parseOutput() ([]string, error)
	// get cmd
	GetCmd()
}

// ---------------------------------------------------------------------------------------------------
type basicCmd struct {
	out []byte
	err error
}

func (T *basicCmd) parseOutput() (string, error) {
	myString := string(T.out[:])
	return myString, T.err
}

// ---------------------------------------------------------------------------------------------------
type OneLineCmd struct {
	basicCmd
	Cmd string
}

func NewOneLineCmd(cmd string, args map[string]string) *OneLineCmd {
	var T OneLineCmd
	T.Cmd = StringReplacer(cmd, args)
	return &T
}

func (T *OneLineCmd) Run() (string, error) {

	T.out, T.err = exec.Command("bash", "-c", T.Cmd).Output()
	/*if T.err != nil {
		log.Fatal(T.err)
	}
	*/
	return T.parseOutput()
}

func (T *OneLineCmd) GetCmd() string {
	return T.Cmd
}

// ---------------------------------------------------------------------------------------------------
type MultiLineCmd struct {
	OneLineCmd
	rawCmd []string
}

func NewMultiLineCmd(cmd []string, args map[string]string) *MultiLineCmd {
	var T MultiLineCmd
	T.rawCmd = make([]string, len(cmd))
	for cc, line := range cmd {
		T.rawCmd[cc] = StringReplacer(line, args)
	}
	T.Cmd = strings.Join(T.rawCmd, "; ")
	return &T
}

func (T *MultiLineCmd) GetCmd() []string {
	return T.rawCmd
}
