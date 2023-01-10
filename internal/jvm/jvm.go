package jvm

import (
	"bytes"
	"os/exec"
	"sort"
	"strings"
)

type Jvm struct {
}

type Process struct {
	ID   string
	Name string
}

func FromCmdOutput(out string) []Process {
	var jvmProcesses []Process
	for _, p := range strings.Split(out, "\n") {
		pParts := strings.Split(p, " ")
		var name string
		if len(pParts) > 1 {
			name = pParts[1]
		}
		jvmProcesses = append(jvmProcesses, Process{
			ID:   pParts[0],
			Name: name,
		})
	}
	sort.Slice(jvmProcesses, func(i, j int) bool {
		return jvmProcesses[i].Name > jvmProcesses[j].Name
	})
	return jvmProcesses
}

func NewJvm() *Jvm {
	return &Jvm{}
}

func (j *Jvm) ListProcesses() ([]Process, error) {
	var out bytes.Buffer
	cmd := exec.Command("jps")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return FromCmdOutput(out.String()), nil
}

func (j *Jvm) Version() (string, error) {
	args := []string{"-version"}
	cmd := exec.Command("java", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	s := string(output)
	versionParts := strings.Split(s, "\"")
	return versionParts[1], nil
}

func (j *Jvm) OpenInJConsole(pid string) error {
	cmd := exec.Command("jconsole", pid)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return cmd.Err
}

func (j *Jvm) KillProcess(pid string) error {
	cmd := exec.Command("kill", "-9", pid)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return cmd.Err
}
