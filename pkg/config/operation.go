package config

type Operation = int

const (
	InvalidOperation Operation = iota
	Print
	Create
	Delete
	Edit
)
