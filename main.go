package main

import "github.com/mdirkse/i3ipc"
import "fmt"

var class string = "st-256color"
var scratchpad string = "__i3_scratch"

func main() {
	ipcsocket, _ := i3ipc.GetIPCSocket()
	tree, _ := ipcsocket.GetTree()
	leaves := tree.Leaves()
	for _, leaf := range leaves {
		if leaf.Window_Properties.Class != class {
			continue
		}
		workspace := leaf.Parent.Parent.Name
		focused := leaf.Focused
		if workspace == scratchpad {
			ipcsocket.Command(fmt.Sprintf("[class=%s] focus", class))
			ipcsocket.Command("move window to output HDMI-1")
		} else if focused {
			ipcsocket.Command("move scratchpad")
		} else {
			ipcsocket.Command(fmt.Sprintf("[class=%s] focus", class))
		}
	}
}
