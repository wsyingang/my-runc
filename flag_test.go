package main

import (
	"fmt"
	"syscall"
	"testing"
)

func TestCloneFlag(t *testing.T) {
	// 2080505856
	flag := syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC | syscall.CLONE_FS
	fmt.Printf("flag=%v", flag)
}
