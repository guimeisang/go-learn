package cgss

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-learn/cgss/src/cg"
	"go-learn/cgss/src/ipc"
)

var centerClient *cg.CenterClient

func startCenterService() error {
	server := ipc.NewIpcClient(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{client}
	return nil
}

func Help(args []string) int{
	fmt.Println(`
		Commands:
			login<username><level><exp>
			login<username>
			send<message>
			listplayer
			quit(q)
			help(h)
	`)
	return 0
}

func Quit(args []string) int {
	return 1
}

func Logout(args []string) int{
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
	}
	return 0
}


