// +build windows

package node

import (
	"github.com/shirou/gopsutil/process"
	"github.com/duanhf2012/origin/log"
)
func KillProcess(processId int){
	p, err := process.NewProcess(int32(processId)) // Specify process id of parent
	if err != nil {
		log.Error("process.NewProcess error:%s", err.Error())
		return
	}

	children, _ := p.Children()
	for _, v := range children {
		err := v.Kill()  // Kill each child
		if err != nil {
			log.Error("kill child faild:%s", err.Error())
		}
	}

	err = p.Kill() // Kill the parent process
	if err != nil {
		log.Error("kill error:%s", err.Error())
	}
}