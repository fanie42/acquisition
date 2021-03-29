package dvras

import (
    "time"

    "github.com/google/uuid"
)

// Device TODO
type Device struct {
    id    DeviceID
    state State
    // data  Data

    events []Event
}

// NewDevice TODO
func NewDevice() *Device {
    return &Device{
        id:     DeviceID(uuid.New()),
        state:  Off,
        events: make([]Event, 0),
    }
}

// ID TODO
func (device *Device) ID() DeviceID {
    return device.id
}

// State TODO
func (device *Device) State() State {
    return device.state
}

// Changes TODO
func (device *Device) Changes() []Event {
    return device.events
}

// Start TODO
func (device *Device) Start(
    annotation string,
) error {
    device.raise(&StartedEvent{
        DeviceID:   device.id,
        Annotation: annotation,
    })

    return nil
}

// Stop TODO
func (device *Device) Stop(
    annotation string,
) error {
    device.raise(&StoppedEvent{
        DeviceID:   device.id,
        Annotation: annotation,
    })

    return nil
}

// AcquireDataPoint TODO
func (device *Device) AcquireDataPoint(
    timestamp time.Time,
    ch1 []int16,
    ch2 []int16,
    pps []int16,
) error {
    if device.state == Off {
        err := device.Start("automatic start")
        if err != nil {
            return err
        }
    }
    device.raise(&DataPointAcquiredEvent{
        DeviceID:  device.id,
        Timestamp: timestamp,
        Ch1:       ch1,
        Ch2:       ch2,
        PPS:       pps,
    })

    return nil
}

func (device *Device) raise(event Event) {
    device.events = append(device.events, event)
    event.Apply(device)
}
