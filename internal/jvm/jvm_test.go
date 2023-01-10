package jvm

import (
	"fmt"
	"testing"
)

func TestJvm_List(t *testing.T) {
	jvm := NewJvm()
	list, err := jvm.ListProcesses()
	fmt.Print(err)
	fmt.Print(list)
}

func TestJvm_Version(t *testing.T) {
	jvm := NewJvm()
	version, err := jvm.Version()
	fmt.Print(err)
	fmt.Print(version)
}
