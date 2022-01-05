package utils

import (
	"fmt"
	"os/exec"
)

// git client clone
func OsExecClone(workspace, url string) error {
	fmt.Println("Try again with git client clone")
	cmd := exec.Command("git", "clone", url, workspace)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	return err
}

// git client pull
func OsExecPull(workspace, url, referenceName, refType string) error {
	if refType == "branch" {
		fmt.Println(fmt.Sprintf("git client pull code by  branch: %s", referenceName))
	} else if refType == "tag" {
		fmt.Println(fmt.Sprintf("git client pull code by  tag: %s", referenceName))
	} else if refType == "commit" {
		fmt.Println(fmt.Sprintf("git client pull code by  commit id: %s", referenceName))
	}
	cmd := exec.Command("git", "pull", url, referenceName)
	cmd.Dir = workspace
	out, err := cmd.CombinedOutput()
	fmt.Println("git client pull code . " + string(out))
	return err
}
