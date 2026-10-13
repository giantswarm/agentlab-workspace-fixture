package main

import (
	"fmt"
	"os"

	"github.com/giantswarm/agentlab-workspace-fixture/greeting"
)

func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println(greeting.Hello(name))
}
