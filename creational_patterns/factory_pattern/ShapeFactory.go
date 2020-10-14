package main

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct{}
type Square struct{}

//方法接受体是指针，不会发生值复制，且可以更改接受体的属性
func (rectangle *Rectangle) Draw() {
	fmt.Println("this is Rectangle!")
}

func (square *Square) Draw() {
	fmt.Println("this is Square!")
}

var ShapeMap = map[string]Shape{
	"Rectangle": &Rectangle{},
	"Square":    &Square{},
}

type Factory struct{}

func (f Factory) getShape(shapeType string) Shape {
	if p, ok := ShapeMap[shapeType]; ok {
		return p
	}
	return nil
}

//优点： 1、一个调用者想创建一个对象，只要知道其名称就可以了。
//2、扩展性高，如果想增加一个产品，只要扩展一个工厂类就可以。
//3、屏蔽产品的具体实现，调用者只关心产品的接口。
func main() {
	f := Factory{}
	f.getShape("Rectangle").Draw()
	f.getShape("Square").Draw()
}
