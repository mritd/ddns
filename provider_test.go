package main

import (
	"testing"
)

func TestNameCom_Create(t *testing.T) {
	initConf()
	err := NewNameCom().Create("1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNameCom_Update(t *testing.T) {
	initConf()
	err := NewNameCom().Update("2.2.2.2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNameCom_Query(t *testing.T) {
	initConf()
	r, err := NewNameCom().Query()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}

func TestGodaddy_Create(t *testing.T) {
	initConf()
	err := NewGodaddy().Create("1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGodaddy_Update(t *testing.T) {
	initConf()
	err := NewGodaddy().Update("2.2.2.2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGodaddy_Query(t *testing.T) {
	initConf()
	r, err := NewGodaddy().Query()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
