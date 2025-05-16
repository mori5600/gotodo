package todo

import "fmt"

type TodoStatus int

const (
	NotStarted TodoStatus = iota
	InProgress
	Done
	Unknown
)

var statusStrings = [...]string{
	"Not Started", // 0
	"In Progress", // 1
	"Done",        // 2
	"Unknown",     // 3
}

func (s TodoStatus) String() string {
	if int(s) < 0 || int(s) >= len(statusStrings) {
		return "Unknown"
	}
	return statusStrings[s]
}

func IntToStatus(num int) (TodoStatus, error) {
	if num < 0 || num >= int(Unknown) {
		return Unknown, fmt.Errorf("invalid status number: %d", num)
	}
	return TodoStatus(num), nil
}

func StatusToInt(status TodoStatus) int {
	return int(status)
}

func StatusToString(status TodoStatus) string {
	return status.String()
}
