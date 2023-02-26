package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/exec"
	"syscall"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "run command",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("invalid args")
		}
		cmd := ctx.Args().Get(0)
		tty := ctx.Bool("ti")
		Run(cmd, tty)
		return nil
	},
}

func NewParentProcess(cmd string, tty bool) *exec.Cmd {
	args := []string{"init", cmd}
	process := exec.Command("/proc/self/exe", args...)
	process.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,
	}
	if tty {
		process.Stdin = os.Stdin
		process.Stdout = os.Stdout
		process.Stderr = os.Stderr
	}
	return process
}

func Run(cmd string, tty bool) {
	parent := NewParentProcess(cmd, tty)
	if err := parent.Start(); err != nil {
		log.Fatalf("Start Parent error+%v", err)
	}
	parent.Wait()
	os.Exit(-1)
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "init process",
	Action: func(ctx *cli.Context) error {
		cmd := ctx.Args().Get(0)
		return RunInitProcess(cmd, nil)
	},
}

func RunInitProcess(cmd string, args []string) error {
	log.Printf("run init porocess,cmd=%s", cmd)
	defaultMountArgs := syscall.MS_NOSUID | syscall.MS_NOEXEC | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountArgs), "")
	argv := []string{cmd}
	if err := syscall.Exec(cmd, argv, os.Environ()); err != nil {
		log.Fatalf("exec cmmand fail +%v", err)
	}
	return nil
}
