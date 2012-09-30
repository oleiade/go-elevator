package main

import (
    "fmt"
    "bytes"
    "github.com/ugorji/go-msgpack"
    zmq "github.com/alecthomas/gozmq"
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

type Elevator struct {
    Socket      zmq.Socket
}

func NewElevator(endpoint string) (*Elevator) {
    elevator := new(Elevator)
    
    context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.XREQ)
    socket.Connect("tcp://127.0.0.1:4141")

    elevator.Socket = socket

    return elevator
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

func newMessage(r *Request) ([][]byte) {
    var preq *bytes.Buffer = packRequest(r)
    var parts = [][]byte{preq.Bytes()}

    return parts
}

func (e Elevator) send(r *Request) (*Response) {
    msg := newMessage(r)
    err := e.Socket.SendMultipart(msg, 0)
    
    if err != nil {
        fmt.Println(err.Error())
    }

    parts, _ := e.Socket.RecvMultipart(0)
    response, err := unpackResponse(parts)
    if err != nil {
        fmt.Println(err.Error())
    }

    return response
}

func main() {
    req := &Request{
        Db: "a9032698-12f8-45a6-ab3e-0e00915ed700",
        Command: "GET",
        Args: []string{"1"},
    }

    elevator := NewElevator("tcp://127.0.0.1:4141")
    response := elevator.send(req)
    fmt.Println(response)
    resp2 := elevator.send(req)
    fmt.Println(resp2)
    // response := send(socket, req)
    // fmt.Println(response)
}