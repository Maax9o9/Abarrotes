package entities

import "time"

type Order struct {
	ID        int       
	OrderDate time.Time 
	Status    string
}
