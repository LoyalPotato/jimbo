package emulator

import (
	"github.com/LoyalPotato/jimbo/pkg/cpu"
	"github.com/LoyalPotato/jimbo/pkg/memory"
)

type Ctx struct {
	Running bool
	CPU     cpu.CPU
	Memory  memory.Memory
}

var GlobalCtx *Ctx

func Init() {
	GlobalCtx = &Ctx{
		Running: true,
		CPU:     cpu.CPU{},
		Memory:  memory.Memory{},
	}

	GlobalCtx.Memory.Init()

	GlobalCtx.CPU.Init()
	GlobalCtx.CPU.Boot()

	for GlobalCtx.Running {
		if !GlobalCtx.Running {
			break
		}
	}
}
