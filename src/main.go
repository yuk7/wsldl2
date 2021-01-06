package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yuk7/wsldl/isregd"
	"github.com/yuk7/wsldl/lib/wslapi"
	"github.com/yuk7/wsldl/version"
)

func main() {
	efPath, _ := os.Executable()
	name := filepath.Base(efPath[:len(efPath)-len(filepath.Ext(efPath))])

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			version.Execute()

		case "isregd":
			isregd.Execute(name)

		default:
			fmt.Println("Invalid Arg.")
		}
	} else {

		fmt.Println(name)
		fmt.Println(wslapi.WslIsDistributionRegistered(name))
	}
}
