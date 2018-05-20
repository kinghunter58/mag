package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/urfave/cli"
)

var word = regexp.MustCompile(`\W`)

var (
	errNotInMag    = errors.New("curent directory doesn't contain magconfig.json OR angular/angular.json")
	errTooManyArgs = errors.New("This function has too many arguments")
)

func buildAction(c *cli.Context) error {
	if c.NArg() != 0 {
		return errTooManyArgs
	}
	root := filepath.Dir(os.Args[0])
	// fmt.Println("root", root)
	_, err := os.Stat(filepath.ToSlash(root) + "/magconfig.json")
	if err != nil {
		return errNotInMag
	}
	_, err = os.Stat(filepath.ToSlash(root) + "/angular/angular.json")
	if err != nil {
		return errNotInMag
	}
	conf, err := getConfig("magconfig.json")
	if err != nil {
		return err
	}
	conf.Dir = word.ReplaceAllString(conf.Dir, "")
	BC := make(chan error)
	go ngBuild(conf.Dir, BC)
	go goBuild(BC)
	if err := <-BC; err != nil {
		return err
	}
	if err := <-BC; err != nil {
		return err
	}
	return nil
}

func ngBuild(path string, c chan error) {
	cmd := exec.Command("ng", "build", "--output-path", "../"+path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "angular"
	err := cmd.Run()
	if err != nil {
		c <- fmt.Errorf("An error by building angular")
	}
	c <- nil
}

func goBuild(c chan error) {
	cmd := exec.Command("go", "build")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		c <- fmt.Errorf("An error by building go")
	}
	c <- nil
}
