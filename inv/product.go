package inv

import "github.com/satori/go.uuid"

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Barcode1    string
	Barcode2    string
}
