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
    status       int        `msgpack:"STATUS"`
    datas        []string   `msgpack:"DATAS"`
}

func packRequest(r *Request) (*bytes.Buffer) {
    buffer := &bytes.Buffer{}
    fmt.Println(buffer)
    enc := msgpack.NewEncoder(buffer)
    enc.Encode(r)

    return buffer
}

func unpackResponse(r []byte) {
    buffer := new(bytes.Buffer)
    dec := msgpack.NewDecoder(buffer, nil)
    err := dec.Decode(r[0])
    if err != nil {
        fmt.Println(err.Error())
    }
}

func newMessage(r *Request) ([][]byte) {
    var preq *bytes.Buffer = packRequest(r)
    var parts = [][]byte{preq.Bytes()}

    return parts
}

func send(s zmq.Socket, r *Request) ([][]byte) {
    msg := newMessage(r)
    err := s.SendMultipart(msg, 0)

    if err != nil {
        fmt.Println(err.Error())
    }

    parts, _ := s.RecvMultipart(0)

    return parts
}

func main() {
    req := &Request{
        Db: "a9032698-12f8-45a6-ab3e-0e00915ed700",
        Command: "GET",
        Args: []string{"1"},
    }

    context, _ := zmq.NewContext()
    defer context.Close()
    socket, _ := context.NewSocket(zmq.XREQ)
    socket.Connect("tcp://127.0.0.1:4141")

    send(socket, req)
}