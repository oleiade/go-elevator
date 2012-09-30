package elevator

import "fmt"

type ElevatorError struct {
	Msg 	string
}

func (e ElevatorError) Error() string {
	return fmt.Sprintf("[Error] Server says: %v", e.Msg)
}