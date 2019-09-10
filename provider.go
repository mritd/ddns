package main

type Provider interface {
	Query() (string, error)
	Update(ip string) error
	Create(ip string) error
}
