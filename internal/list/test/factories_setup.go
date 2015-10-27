package main

func num1Collection(v ...Num1) Num1List {
	list := Num1List{}
	list = append(list, v...)
	return list
}

func num2Collection(v ...*Num2) Num2List {
	list := make([]*Num2, len(v))
	for i, n := range v {
		list[i] = n
	}
	return list
}
