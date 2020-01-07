package main

type Integer int

type Rect struct {
	x, y          float32
	width, height float64
}

func main() {
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
}

func (a Integer) Add(b Integer) {
	a += b
}

func (rt *Rect) Area() float64 {
	return rt.height * rt.width
}
