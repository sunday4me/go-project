package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	name           string
	length, height float64
}

func (r *rectangle) area() float64 {
	return r.length * r.height
}

func (r *rectangle) perimeter() float64 {
	return 2*r.length + 2*r.height
}

type triangle struct {
	name    string
	a, b, c float64
}

func (t *triangle) area() float64 {
	return 0.5 * (t.a * t.b)
}

func (t *triangle) perimeter() float64 {
	return t.a + t.b + math.Sqrt((t.a*t.a)+(t.b*t.b))
}

func (t *triangle) String() string {
	return "I am an triangle"
	//return fmt.Sprintf("%s[sides: a=%.2f b=%.2f c=%.2f]", t.name, t.a, t.b, t.c)
}

func shapeInfo(s Shape) string {
	return fmt.Sprintf("Area = %.2f, Perimeter = %.2f", s.area(), s.perimeter())
}

func main() {
	r := &rectangle{"First Rectangle", 4, 4}
	fmt.Println(r, "=>", shapeInfo(r))

	t := &triangle{
		name: "First Triangle",
		a:    1,
		b:    2,
		c:    3,
	}
	fmt.Println(t)
	//First Triangle[sides: a=1.00 b=2.00 c=3.00]

	//&{First Rectangle 4 4} => Area = 16.00, Perimeter = 16.00
	//&{First Triangle 1 2 3} => Area = 1.00, Perimeter = 5.24

	//First Triangle[sides: a=1.00 b=2.00 c=3.00] => Area = 1.00, Perimeter = 5.24
} (edited) 
:white_check_mark:
2



Oladapo Abimbola
:thunder_icon-01:  7:54 PM
package main

import (
	"fmt"
	"io"
	"os"
)

//var i io.Reader
//var w io.Writer

type Counter int

func (c *Counter) Write(b []byte) (int, error) {
	l := len(b)
	*c += Counter(l) //casting l to an integer so it becomes the right type
	return l, nil
}

func main() {
	var c Counter
	f1, _ := os.Open("input.txt")
	//f2, _ := os.Create("output.txt")
	newWriter := &c
	n, _ := io.Copy(newWriter, f1)

	fmt.Println("Copied", n, "bytes")
	//fmt.Println(c)

}
