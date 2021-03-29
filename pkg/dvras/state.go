package dvras

// State TODO
type State int

const (
    // On TODO
    On State = iota
    // Off TODO
    Off
)

// String TODO
func (state State) String() string {
    switch state {
    case On:
        return "on"
    case Off:
        return "off"
    }

    return ""
}
