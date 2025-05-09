package common

import (
	"fmt"

	"github.com/LoyalPotato/jimbo/pkg/state"
)

func Println(a ...any) {
	if state.DebugMode {
		fmt.Println(a...)
	}
}

func Printf(format string, a ...any) {
	if state.DebugMode {
		fmt.Printf(format, a...)
	}
}
