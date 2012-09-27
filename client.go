package main

import (
    "fmt"
    "log"
    "bytes"
    "github.com/ugorji/go-msgpack"
    zmq "github.com/alecthomas/gozmq"
)

type Request struct {
    db_uid      string
    command string
    args         []string
}

type Response struct {
    status       int
    content     []string
}

func packRequest(r *Request) (*bytes.Buffer) {
    buffer := new(bytes.Buffer)
    enc := msgpack.NewEncoder(buffer)
    enc.Encode(r)

    return buffer
}

func newMessage(r *Request) ([][]byte) {
    var preq *bytes.Buffer = packRequest(r)
    var breq byte

    breq, err := preq.ReadByte()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(breq)

    return [][]byte{[]byte{breq}}
}

func send(s zmq.Socket, r *Request) ([][]byte) {
    msg := newMessage(r)
    err := s.SendMultipart(msg, 0)
    
    fmt.Println(msg)
    if err != nil {
        log.Fatal(err)
    }

    parts, _ := s.RecvMultipart(0)

    return parts
}

func main() {
    req := &Request{
        db_uid: "a9032698-12f8-45a6-ab3e-0e00915ed700",
        command: "GET",
        args: []string{"1"},
    }

    context, _ := zmq.NewContext()
    defer context.Close()
    socket, _ := context.NewSocket(zmq.XREQ)
    socket.Connect("tcp://127.0.0.1:4141")

    res := send(socket, req)
    fmt.Println(res)
}