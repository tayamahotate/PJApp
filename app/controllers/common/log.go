package common

import (
	"time"
	"fmt"
)

func WriteLog(screenName, button, action string) {
	fmt.Println(time.Now().String() + screenName + "." + button + "->" + action)
}

func WriteLogStr(screenName, button, action, str string) {
	fmt.Println(time.Now().String() + screenName + "." + button + "->" + action + " " + str)
}
