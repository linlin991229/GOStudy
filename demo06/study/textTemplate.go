package study

import (
	"fmt"
	"os"
	"text/template"
)

func TestTextTemplate() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Values {{.}}\n")

	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Values {{.}}\n"))

	t1.Execute(os.Stdout, "Hello World")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"Hello", "World"})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// struct和map使用{{.fileName}}访问字段
	t2 := Create("t2", "Name {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"John Doe"})
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})
	// if-else模板
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")
	// range模板
	t4 := Create("t4", "Range: {{range .}} {{.}} {{end}}\n")

	t4.Execute(os.Stdout, []string{"a", "b", "c"})
}
func TestShadowing() {
	i := 10
	// 多赋值时，变量可以重复声明（重复使用相同的变量名），
	i, j := test()
	{
		i := 11
		fmt.Printf("i=%d j=%d\n", i, j)
	}
	fmt.Println(i, j)
}
func test() (int, int) {
	return 10, 20
}
