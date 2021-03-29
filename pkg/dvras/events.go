package dvras

import "time"

// Event TODO
type Event interface {
    Apply(*Device)
}

// StartedEvent TODO
type StartedEvent struct {
    DeviceID   DeviceID `json:"aggregate_id"`
    Annotation string   `json:"annotation"`
}

// Apply TODO
func (event *StartedEvent) Apply(device *Device) {
    device.state = On

    return
}

// StoppedEvent TODO
type StoppedEvent struct {
    DeviceID   DeviceID `json:"aggregate_id"`
    Annotation string   `json:"annotation"`
}

// Apply TODO
func (event *StoppedEvent) Apply(device *Device) {
    device.state = Off

    return
}

// DataPointAcquiredEvent TODO
type DataPointAcquiredEvent struct {
    DeviceID  DeviceID  `json:"aggregate_id"`
    Timestamp time.Time `json:"timestamp"`
    Ch1       []int16   `json:"ch1"`
    Ch2       []int16   `json:"ch2"`
    PPS       []int16   `json:"pps"`
}

// Apply TODO
func (event *DataPointAcquiredEvent) Apply(device *Device) {
    return
}
