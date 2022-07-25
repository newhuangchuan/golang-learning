package main

import (
    "log"
    "sync"
    "time"
    )

var HcMutex sync.Mutex

func runMutex(id int) {
    log.Printf("[task id : %d , go to get mutex ]", id)
    HcMutex.Lock()
    log.Printf("[task id : %d , have get mutex ]", id)
    time.Sleep(time.Second * 2)
    HcMutex.Unlock()
    log.Printf("[task id : %d , open this mutex]", id)
}

func runLock() {
    runMutex(1)
    runMutex(2)
    runMutex(3)
    runMutex(4)
    runMutex(5)
    runMutex(6)

}

func main() {

    runLock()
    time.Sleep(3 * time.Second)
    
}
