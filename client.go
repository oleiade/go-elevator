package main

import (
    "bytes"
    "github.com/ugorji/go-msgpack"
    "github.com/alecthomas/gozmq"
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
func main() {
}