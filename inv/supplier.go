package inv

import "github.com/satori/go.uuid"

type Supplier struct {
	ID      uuid.UUID
	Name    string
	Address string
}
