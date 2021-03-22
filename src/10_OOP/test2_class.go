package main

import "fmt"

//如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	//如果说类的属性首字母大写, 表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name  string
	Ad    int
	level int
}

/*
func (t Hero) Show() {
	fmt.Println("Name = ", t.Name)
	fmt.Println("Ad = ", t.Ad)
	fmt.Println("Level = ", t.Level)
}

func (t Hero) GetName() string {
	return t.Name
}

func (t Hero) SetName(newName string) {
	//t 是调用该方法的对象的一个副本（拷贝）
	t.Name = newName
}
*/
func (t *Hero) Show() {
	fmt.Println("Name = ", t.Name)
	fmt.Println("Ad = ", t.Ad)
	fmt.Println("Level = ", t.level)
}

func (t *Hero) GetName() string {
	return t.Name
}

func (t *Hero) SetName(newName string) {
	//t 是调用该方法的对象的一个副本（拷贝）
	t.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100}

	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
