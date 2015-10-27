package main

func num1Collection(v ...Num1) Num1Set {
	return NewNum1Set(v...)
}

func num2Collection(v ...*Num2) Num2Set {
	set := make(map[*Num2]struct{})
	for _, n := range v {
		set[n] = struct{}{}
	}
	return set
}
