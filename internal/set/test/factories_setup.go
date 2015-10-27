package main

func num1Collection(v ...Num1) Num1Set {
	return NewNum1Set(v...)
}

func fooCollection(v ...Foo) FooSet {
	return NewFooSet(v...)
}

func thingCollection(v ...Thing) ThingSet {
	return NewThingSet(v...)
}
