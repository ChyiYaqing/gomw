package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

// Go语言中，并不需要显式声明实现那个接口，只需要直接实现该接口对应的方法即可.
func (stu *Student) getName() string {
	return stu.name
}

func main() {
	// 没有显性赋值的变量被赋予默认值
	stu := &Student{
		name: "Tom",
	}
	msg := stu.hello("Jack")
	fmt.Println(msg)

	stu2 := new(Student)
	fmt.Println(stu2.hello("Alice"))

	// 实例化 Student后，强制类型转换为接口类型Person.
	var p Person = &Student{
		name: "Tom",
		age:  18,
	}
	fmt.Println(p.getName())

	// 检测某个类型实现了某个接口的所有方法
	var _ Person = (*Student)(nil)
	var _ Person = (*Worker)(nil)

	// 实例可以强制类型转换为接口，接口可以强制类型转换为实例
	stu = p.(*Student)
	fmt.Println(stu.hello(""))

	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 18
	m["scores"] = [3]int{98, 99, 85}
	fmt.Println(m)
}

// 接口
// -- 接口定义一组方法的集合，接口不能被实例化，一个类型可以实现多个接口

type Person interface {
	getName() string
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

// 定义一个没有任何方法的空接口，那么这个接口可以表示任意类型
