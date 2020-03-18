package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type Alipay struct {
}

func (a *Alipay) CanUseFaceID() {

}

type Cash struct {
}

func (c *Cash) Stolen() {

}

type CantainCanUseFaceID interface {
	CanUseFaceID()
}
type ContainStolen interface {
	Stolen()
}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case CantainCanUseFaceID:
		log.Printf("%T can use faceid\n", payMethod)
	case ContainStolen:
		log.Printf("%T canuse stolen\n", payMethod)
	}
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, errors.New("math:square root of negative number")
	}
	return math.Sqrt(f), nil
}

type selfErrors struct {
	Num     float64
	problem string
}

func (se *selfErrors) Error() string {
	return fmt.Sprintf("Wrong!!!,because \"%f\" is a negative number ", se.Num)
}

func sSqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, &selfErrors{f, "error"}
	}
	return math.Sqrt(f), nil
}

func main() {
	f, err := sSqrt(-3.2)
	log.Println(f, err)

	//print(new(Alipay))
	//print(new(Cash))
	//printType(1024)
	//printType(33.3)
	//printType([]byte("1234"))
	//printType("lsy")
	//log.Println("hello")
}

func printType(v interface{}) {
	switch v.(type) {
	case int:
		log.Println(v, "is int")
	case float32:
		log.Println(v, "is float32")
	case float64:
		log.Println(v, "is float64")
	case bool:
		log.Println(v, "is bool")
	case byte:
		log.Println(v, "is byte")
	case string:
		log.Println(v, "is string")
	case []byte:
		log.Println(v, "is []byte")
	default:
		log.Println(v, "is no type")
	}
}
