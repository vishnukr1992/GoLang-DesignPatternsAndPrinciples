package solid

import (
	"fmt"
	"time"
)

// Here we granulated the operations to different interfaces

type AlarmOperation interface {
	SetAlarm(time time.Duration)
	DisableAlarm(time time.Duration)
}

type CallOperation interface {
	ConnectCall(number int)
	DisconnectCall(number int)
}

type GameConsoleOperation interface {
	PlayGame(name string)
}

type Alarm struct {
}

type Phone struct {
}

type GameConsole struct {
}

func (a *Alarm) SetAlarm(time time.Duration) {
	fmt.Println("Alarm setting to ", time)
}
func (a *Alarm) DisableAlarm(time time.Duration) {
	fmt.Println("Alarm disabled at ", time)
}
func (p *Phone) ConnectCall(number int) {
	fmt.Println("calling to number ", number)
}
func (p *Phone) DisconnectCall(number int) {
	fmt.Println("Disconnect the call:", number)
}

func (g *GameConsole) PlayGame(name string) {
	fmt.Println("Opening game:", name)
}

type MobileIsp struct {
	Alarm
	Phone
	GameConsole
}
