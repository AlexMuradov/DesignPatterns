package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Session struct {
}

var singleInstance *Session

func initSess() *Session {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &Session{}
			fmt.Println("Instance created")
		} else {
			fmt.Println("Instance already created")
		}
	} else {
		fmt.Println("Instance already created")
	}
	return singleInstance
}
