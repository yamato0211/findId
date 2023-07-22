package a

func f() {
	var Id int  // want "NG: Id, OK: ID"
	println(Id) // want "NG: Id, OK: ID"
}
