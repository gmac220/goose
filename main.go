package main

import (
	"os/exec"
)

func main() {
	// path := os.Getenv("PATH")
	// fmt.Println(path)

	// hostname, err := os.Hostname()
	// if err != nil {
	// 	os.Exit(0)
	// }
	// fmt.Println(hostname)

	// cmd := exec.Command("ls", "-l")
	// output, stderr := cmd.Output()
	// if stderr != nil {
	// 	fmt.Println(stderr)
	// }
	// fmt.Println(string(output))

	dockerstop := exec.Command("docker", "stop", "exampledb")
	dockerstop.Run()
	dockerdelete := exec.Command("docker", "container", "rm", "exampledb")
	dockerdelete.Run()
	dockerrm := exec.Command("docker", "rmi", "exampledb")
	dockerrm.Run()
	dockerbuild := exec.Command("docker", "build", "-t", "exampledb", ".")
	dockerbuild.Run()
	dockerrun := exec.Command("docker", "run", "--name=exampledb", "-p=5432:5432", "-d", "exampledb")
	dockerrun.Run()
	// doutput, _ := dockercmd.Output()
	// fmt.Println(string(doutput))
}
