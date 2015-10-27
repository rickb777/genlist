package main

func num1Collection(v ...Num1) Num1List {
	if len(v) == 0 {
		return Num1List{}
	}
	return Num1List(v)
}
