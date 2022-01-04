package main

import "fmt"

type RecordNotFoundErr struct {
	Record *Record
}

func (e RecordNotFoundErr) Error() string {
	return fmt.Sprintf("record %s not found", e.Record)
}

func NewRecordNotFoundErr(r *Record) RecordNotFoundErr {
	return RecordNotFoundErr{
		Record: r,
	}
}
