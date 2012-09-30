package elevator


import (
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


func (e *Elevator) send(r *Request) (*Response, error) {
    // Insert elevator connector db_uid in Request
    r.Db = e.Db

    msg := newMessage(r)
    if err := e.Socket.SendMultipart(msg, 0); err != nil {
        return nil, err
    }

    parts, _ := e.Socket.RecvMultipart(0)
    response, err := unpackResponse(parts)
    if err != nil {
        return nil, err
    }

    return response, nil
}


func (e *Elevator) Connect(db_name string) (error) {
    req := NewRequest("DBCONNECT", []string{db_name})
    response, err := e.send(req)
    if err != nil {
        return err
    }

    e.Db = response.Datas[0]
    return nil 
}


func (e *Elevator) CreateDb(db_name string) (error) {
    req := NewRequest("DBCREATE", []string{db_name})
    _, err := e.send(req)

    return err
}


func (e *Elevator) ListDb() ([]string, error) {
    req := NewRequest("DBLIST", []string{})
    response, err := e.send(req)
    if err != nil {
        return nil, err
    }
    value := response.Datas

    return value, nil
}


func (e *Elevator) Get(key string) (string, error) {
    req := NewRequest("GET", []string{key})
    response, err := e.send(req)
    if err != nil {
        return "", err
    }
    value := response.Datas[0]

    return value, nil
}


func (e *Elevator) Put(key string, value string) (error) {
    req := NewRequest("PUT", []string{key, value})
    _, err := e.send(req)

    return err
}


func (e *Elevator) Delete(key string) (error) {
    req := NewRequest("DELETE", []string{key})
    _, err := e.send(req)

    return err
}