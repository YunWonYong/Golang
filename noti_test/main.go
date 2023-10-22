package main

import (
	"fmt"
)

type (
	NotiInfo struct {
		Type string
		CurrentTime string
	}

	NotiManager struct {
		NotiPushQueue chan *NotiInfo
		NotiPushErrorQueue chan error
	}
)

func (ni *NotiInfo) send() error {
	// [TODO] send logic

	fmt.Printf("%#+v\n", ni)
	return nil
}

func (nm *NotiManager) Init(queueSize int) *NotiManager {
	nm.NotiPushQueue = make(chan *NotiInfo, queueSize)
        nm.NotiPushErrorQueue = make(chan error, queueSize)

        go func () {
                for {
                        select {
			case err := <- nm.NotiPushErrorQueue:
                                        fmt.Println(err.Error())
			case noti := <- nm.NotiPushQueue:
				noti.send()
                        }
                }
        }()
	return nm
}

func (nm *NotiManager) NotiPush(ni *NotiInfo) {
	nm.NotiPushQueue <- ni
}

func (nm *NotiManager)Send() {
	for {
		noti := <- nm.NotiPushQueue 

		go func () {
			if err := noti.send(); err != nil {
				nm.NotiPushErrorQueue <- err
			}
		}()
	}
}

const QUEQUE_SIZE = 100

var manager *NotiManager = new(NotiManager).Init(QUEQUE_SIZE)
func main() {

	notiList := append([]*NotiInfo{}, new(NotiInfo), new(NotiInfo), new(NotiInfo), new(NotiInfo), new(NotiInfo), new(NotiInfo), new(NotiInfo))

	go func() {
		for _, notiInfo := range notiList {
			manager.NotiPush(notiInfo)
		}
	}()

	for {
	}

}
