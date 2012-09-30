package elevator

import (
	"bytes"
	"fmt"
	"github.com/ugorji/go-msgpack"
)

type Request struct {
	Db      string   `msgpack:"DB_UID"`
	Command string   `msgpack:"COMMAND"`
	Args    []string `msgpack:"ARGS"`
}

type Response struct {
	Status int      `msgpack:"STATUS"`
	Datas  []string `msgpack:"DATAS"`
}

func NewRequest(command string, args []string) *Request {
	return &Request{
		Command: command,
		Args:    args,
	}
}

func packRequest(r *Request) *bytes.Buffer {
	buffer := &bytes.Buffer{}
	enc := msgpack.NewEncoder(buffer)
	enc.Encode(r)

	return buffer
}

func unpackResponse(parts [][]byte) (*Response, error) {
	var err error
	response := new(Response)
	msg := parts[0]
	dec := msgpack.NewDecoder(bytes.NewBuffer(msg), nil)
	dec.Decode(response) // Ignore msgpack decoder errors

	if response.Status == FAILURE_STATUS {
		var code, msg = response.Datas[0], response.Datas[1]

		err = ElevatorError{
			Code: code,
			Msg:  msg,
		}

		return nil, err
	}

	return response, err
}

func main() {
	elevator := NewElevator("tcp://127.0.0.1:4141")

	value, err := elevator.Range("1", "3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}
