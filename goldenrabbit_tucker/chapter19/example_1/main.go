package main

import (
	"fmt"
)

type MethodStruct struct {
	msg string
}

func (ms *MethodStruct) PointReceiverPrint() {
	fmt.Println("pointer: \t", ms.msg)
}


func (ms MethodStruct) ValueReceiverPrint() {
	fmt.Println("value: \t\t", ms.msg)
}


func (ms *MethodStruct) PointerReceiverSetMsg(msg string) {
	ms.msg = msg
}

func (ms MethodStruct) ValueReceiverSetMsg(msg string) {
	ms.msg = msg
}

func main() {
	pms := new(MethodStruct)
	pms.msg = "hello world!!!"
	fmt.Println("===pointer struct pms===")
	pms.PointReceiverPrint()
	pms.ValueReceiverPrint()

	vms := MethodStruct{
		msg: "hello world!!!",
	}
	fmt.Println("====value struct vms===")
	vms.PointReceiverPrint()
	vms.ValueReceiverPrint()

	pms.PointerReceiverSetMsg("bye world!!!")
	fmt.Println("===pointer struct pms===")
	pms.PointReceiverPrint()
	pms.ValueReceiverPrint()

	vms.ValueReceiverSetMsg("bye world!!!")
	fmt.Println("====value struct vms===")
	vms.PointReceiverPrint()
	vms.ValueReceiverPrint()

	pms.msg = "hello world!!!"
	pms.ValueReceiverSetMsg("bye world!!!")
	fmt.Println("===pointer struct pms===")
	pms.PointReceiverPrint()
	pms.ValueReceiverPrint()

	vms.msg = "hello world!!!"
	vms.PointerReceiverSetMsg("bye world!!!")
	fmt.Println("====value struct vms===")
	vms.PointReceiverPrint()
	vms.ValueReceiverPrint()
}
