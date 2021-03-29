package dvras

// Service TODO
type Service interface {
    Start(command *StartCommand) error
    Stop(command *StopCommand) error
}

// StartCommand TODO
type StartCommand struct {
    // DeviceID   UUID
    Annotation string `json:"annotation"`
}

// StopCommand TODO
type StopCommand struct {
    // DeviceID   UUID
    Annotation string `json:"annotation"`
}
