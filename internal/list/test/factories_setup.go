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

func fooCollection(v ...Foo) FooList {
	list := FooList{}
	list = append(list, v...)
	return list
}

func thingCollection(v ...Thing) ThingList {
	list := ThingList{}
	list = append(list, v...)
	return list
}

func ip(n int) *Num2 {
	v := Num2(n)
	return &v
}

