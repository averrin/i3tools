package main

import "github.com/mdirkse/i3ipc"
import "fmt"
import "os/exec"
import "os"

var termClass string = "st-256color"
var scratchpad string = "__i3_scratch"

func focusTerm() {
	ipcsocket, _ := i3ipc.GetIPCSocket()
	tree, _ := ipcsocket.GetTree()
	leaves := tree.Leaves()
	for _, leaf := range leaves {
		if leaf.Window_Properties.Class != termClass {
			continue
		}
		workspace := leaf.Parent.Parent.Name
		focused := leaf.Focused
		if workspace == scratchpad {
			ipcsocket.Command(fmt.Sprintf("[id=%v] focus", leaf.Window))
			ipcsocket.Command("move window to output HDMI-1")
			ipcsocket.Command(fmt.Sprintf("[id=%v] focus", leaf.Window))
		} else if focused {
			ipcsocket.Command("move scratchpad")
		} else {
			ipcsocket.Command(fmt.Sprintf("[id=%v] focus", leaf.Window))
		}
		break
	}
}

func ror(class string, cmd string) {
	ipcsocket, _ := i3ipc.GetIPCSocket()
	tree, _ := ipcsocket.GetTree()
	leaves := tree.Leaves()
	found := false
	for _, leaf := range leaves {
		if leaf.Window_Properties.Class != class || leaf.Focused {
			continue
		} else if !leaf.Focused {
			ipcsocket.Command(fmt.Sprintf("[id=%v] focus", leaf.Window))
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Running app…")
		cmd := exec.Command(cmd)
		cmd.Start()
	}
}

func main() {
	fmt.Println(os.Args)
	if len(os.Args) == 1 || os.Args[1] == "term" {
		focusTerm()
	} else if os.Args[1] == "ror" {
		ror(os.Args[2], os.Args[3])
	}
}
