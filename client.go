package main

import (
    "fmt"
    "bytes"
    "github.com/ugorji/go-msgpack"
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
    enc := msgpack.NewEncoder(bbuffer)
    enc.Encode(r)

    return buffer
}

func main() {
}