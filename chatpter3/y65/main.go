package main

import "log"

type Integer int

type Rect struct {
	x, y          float32
	width, height float64
}

type Base struct {
	Name string
}

type Foo struct {
	*Base
}

type Job struct {
	Comman string
	*log.Logger
}

func (job *Job) Start() {

}

func (base *Base) Foo() {
	base.Name = "jinru base foo"
	println(base.Name)
}

func (base *Base) Bar() {

}

func (foo *Foo) Bar() {
	foo.Base.Bar()
}

func main() {
	var foo = Foo{new(Base)}
	foo.Foo()
	var b Integer = 2
	var a *Integer = &b
	a.Add(2)
	println("a=", a)

	var rect1 = Rect{}
	rect1.width = 2.34
	rect1.height = 4.22
	aa := rect1.Area()
	println("面积=", aa)
	//var rect2 = new(Rect)
	//var rect3 = &Rect{}
	var map1 map[int]string
	map1 = make(map[int]string)

	map1[2] = "1234"
	println("map1[2]=", map1[2])
	value, ok := map1[2]
	if ok {
		println("map1[2]=", value)
	}
}

func (a Integer) Add(b Integer) {
	a += b
}

func (rt *Rect) Area() float64 {
	return rt.height * rt.width
}
