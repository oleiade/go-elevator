go-elevator
===========

A go client for Elevator key/value store

## Installation
`go get github.com/oleiade/go-elevator`

## Usage

```go
    package main
    
    import (
        elevator "github.com/oleiade/go-elevator"
    )
    
    func main() {
      client := elevator.NewElevator("tcp://127.0.0.1")
    }
```