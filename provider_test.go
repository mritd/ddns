package main

//func TestNameCom_Create(t *testing.T) {
//	initConf()
//	p, err := NewNameCom()
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = p.Create("1.1.1.1")
//	if err != nil {
//		t.Fatal(err)
//	}
//}

//func TestNameCom_Update(t *testing.T) {
//	initConf()
//	p, err := NewNameCom()
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = p.Update("2.2.2.2")
//	if err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestNameCom_Query(t *testing.T) {
//	initConf()
//	p, err := NewNameCom()
//	if err != nil {
//		t.Fatal(err)
//	}
//	r, err := p.Query()
//	if err != nil {
//		t.Fatal(err)
//	} else {
//		t.Log(r)
//	}
//}

//func TestGoDaddy_Create(t *testing.T) {
//	initConf()
//	p, err := NewGoDaddy()
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = p.Create("1.1.1.1")
//	if err != nil {
//		t.Fatal(err)
//	}
//}

//func TestGoDaddy_Update(t *testing.T) {
//	initConf()
//	p, err := NewGoDaddy()
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = p.Update("2.2.2.2")
//	if err != nil {
//		t.Fatal(err)
//	}
//}

//func TestGoDaddy_Query(t *testing.T) {
//	initConf()
//	p, err := NewGoDaddy()
//	if err != nil {
//		t.Fatal(err)
//	}
//	r, err := p.Query()
//	if err != nil {
//		t.Fatal(err)
//	} else {
//		t.Log(r)
//	}
//}

//func TestGandi_Create(t *testing.T) {
//	initConf()
//	p, err := NewGandi()
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = p.Create("1.1.1.1")
//	if err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestGandi_Update(t *testing.T) {
//	initConf()
//	p, err := NewGandi()
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = p.Update("2.2.2.2")
//	if err != nil {
//		t.Fatal(err)
//	}
//}

func preTest() {
	debug = true
	initLog()
	initConf()
}
