package simian

import (
	"fmt"
	"time"
)

type Engine struct {
	Name      string
	startTime time.Time
	endTime   time.Time
	minDelay  uint
}

func NewEngine(name string) *Engine {

	engine := &Engine{
		Name:      name,
		startTime: time.Time{},
		endTime:   time.Time{},
		minDelay:  1,
	}
	fmt.Printf("Engine: %+v\n", engine)
	return engine
}

func PrintEngine(e Engine) {
	fmt.Printf("\n")
}

func (engine Engine) Run() {
	engine.startTime = time.Now()
}
