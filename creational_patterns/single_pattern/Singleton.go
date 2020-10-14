package main

import (
	"fmt"
	"sync"
) //单例模式
// 1、单例类只能有一个实例。
// 2、单例类必须自己创建自己的唯一实例。
// 3、单例类必须给所有其他对象提供这一实例。

var (
	// once 保证只执行一次
	once sync.Once

	// instance 单例对象
	instance *Instance
)

// Instance 单例类型结构体
type Instance struct {
	// Title 标题
	Title string
}

// GetInstance 获取单例对象。 注：必须使用指针对象，否则会在返回时，拷贝一份，对象就不是同一个了。
func GetInstance(title string) *Instance {
	once.Do(func() {
		fmt.Println("执行获取单例对象")
		instance = &Instance{Title: title}
	})
	return instance
}

func main() {
	instance1 := GetInstance("标题1")
	fmt.Println(fmt.Sprintf("instance1 = %p, name=%s", instance1, instance1.Title))

	instance2 := GetInstance("标题2")
	fmt.Println(fmt.Sprintf("instance2 = %p, name=%s", instance2, instance2.Title))
}
