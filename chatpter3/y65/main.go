package main

type Integer int

type Rect struct {
	x, y          float32
	width, height float64
}

type IFile interface {
	Read()(i int,err error)
	Writer()(i int ,err error)
	Seek()(i int,err error)
}

type SFile struct {

}

func(sf *SFile)Read()(i int,err error){
	println("sfile::Read()")
	return i,err
}

func(sf *SFile)Writer()(i int ,err error){
	return i,err
}

func(sf *SFile)Seek()(i int,err error){
	return i,err
}

func main() {


	var iff IFile= new(SFile)
	iff.Read()

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
