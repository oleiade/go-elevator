package elevator

import (
	"fmt"
	"bytes"
    zmq 		"github.com/alecthomas/gozmq"
)

type Elevator struct {
    Socket      zmq.Socket
    Db          string
}

func NewElevator(endpoint string) (*Elevator) {
    elevator := new(Elevator)
    
    context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.XREQ)
    socket.Connect(endpoint)

    elevator.Socket = socket
    elevator.Connect("default")

    return elevator
}

func newMessage(r *Request) ([][]byte) {
    var preq *bytes.Buffer = packRequest(r)
    var parts = [][]byte{preq.Bytes()}

    return parts
}

func (e *Elevator) send(r *Request) (*Response) {
    // Insert elevator connector db_uid in Request
    r.Db = e.Db

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

func (e *Elevator) Connect(db_name string) (error) {
    req := NewRequest("DBCONNECT", []string{db_name})
    response := e.send(req)
    e.Db = response.Datas[0]
    return nil 
}

func (e *Elevator) CreateDb(db_name string) (error) {
    req := NewRequest("DBCREATE", []string{db_name})
    e.send(req)
    return nil
}

func (e *Elevator) ListDb() ([]string, error) {
    req := NewRequest("DBLIST", []string{})
    response := e.send(req)
    value := response.Datas

    return value, nil
}

func (e *Elevator) Get(key string) (string, error) {
    req := NewRequest("GET", []string{key})
    response := e.send(req)
    value := response.Datas[0]

    return value, nil
}

func (e *Elevator) Put(key string, value string) (error) {
    req := NewRequest("PUT", []string{key, value})
    e.send(req)

    return nil
}

func (e *Elevator) Delete(key string) (error) {
    req := NewRequest("DELETE", []string{key})
    e.send(req)

    return nil
}