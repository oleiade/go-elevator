package elevator

import (
    "fmt"
    "bytes"
    "github.com/ugorji/go-msgpack"
)

type Request struct {
    Db          string      `msgpack:"DB_UID"`
    Command     string      `msgpack:"COMMAND"`
    Args        []string    `msgpack:"ARGS"`
}

type Response struct {
    Status       int        `msgpack:"STATUS"`
    Datas        []string   `msgpack:"DATAS"`
}

func NewRequest(command string, args []string) (*Request) {
    return &Request{
        Command: command,
        Args: args,
    }
}

func packRequest(r *Request) (*bytes.Buffer) {
    buffer := &bytes.Buffer{}
    enc := msgpack.NewEncoder(buffer)
    enc.Encode(r)

    return buffer
}

func unpackResponse(parts [][]byte) (*Response, error) {
    response := new(Response)
    msg := parts[0]
    dec := msgpack.NewDecoder(bytes.NewBuffer(msg), nil)
    err := dec.Decode(response)

    return response, err
}


func main() {
    elevator := NewElevator("tcp://127.0.0.1:4141")
    elevator.Put("2", "b")
    val, _ := elevator.Get("2")
    fmt.Println(val)
    
    err := elevator.Delete("2")
    if err != nil {
        fmt.Println(err)
    }
    value, _ := elevator.ListDb()
    fmt.Println(value)
}