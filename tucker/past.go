package main

import "time"

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Id    int
	Price int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(p *Product) *Parcel {
	return &Parcel{Pdt: p, ShippedTime: time.Now()}
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}
