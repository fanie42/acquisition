package dvras

// DeviceRepository TODO
type DeviceRepository interface {
    GetDeviceByID(DeviceID) (*Device, error)
    Save(*Device) error
}
