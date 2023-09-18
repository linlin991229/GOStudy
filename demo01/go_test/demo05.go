package go_test

import (
	"fmt"
	"image"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/reader"
)

func read() {
	r := strings.NewReader("hello,Reader!")

	buff := make([]byte, 8)
	for {
		n, err := r.Read(buff)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, buff)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
	}
}

type MyReader struct{}

func (myReader MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}
func readDemo() {
	reader.Validate(MyReader{})
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	fmt.Println(n)
	return n, err
}
func readDemo2() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
func imageDemo() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
