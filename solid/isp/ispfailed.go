package solid

import "time"

// This interface violating ISP
type MobileOperation interface {
	SetAlarm(time time.Duration)
	DisableAlarm(time time.Duration)
	ConnectCall(number int)
	DisconnectCall(number int)
}

type Mobile struct {
}

func (m *Mobile) SetAlarm(time time.Duration) {
	// implementation here
}
func (m *Mobile) DisableAlarm(time time.Duration) {
	// implementation here
}
func (m *Mobile) ConnectCall(number int) {
	// implementation here
}
func (m *Mobile) DisconnectCall(number int) {
	// implementation here
}
