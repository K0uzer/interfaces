package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	// "math"
)

// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }
// type Circle struct {
// 	Radius float64
// }
// func (c Circle) Area() float64 {
// 	return math.Pi * c.Radius * c.Radius	
// }
// func (c Circle) Perimeter() float64 {
// 	return 2 * math.Pi * c.Radius
// }
// type Rectangle struct {
// 	Width, Height float64
// }
// func (r Rectangle) Area() float64 {
// 	return r.Width * r.Height
// }
// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.Width + r.Height)
// }
// type Triangle struct {
// 	Base, Height float64
// }
// func (t Triangle) Area() float64 {
// 	return t.Base * t.Height / 2
// }
// func (t Triangle) Perimeter() float64 {
// 	side := math.Sqrt(t.Height*t.Height + (t.Base/2)*(t.Base/2))
//  return t.Base + 2*side
// }
// func TotalArea(shapes []Shape) float64 {
// 	var total float64
// 	for _, shape := range shapes {
// 		total += shape.Area()
// 	}
// 	return total
// }

// func DescribeShape(s Shape) string {
// 	switch v := s.(type) {
// 		case Circle:
// 			return fmt.Sprintf("круг радиусом %.1f", v.Radius)
// 		case Rectangle:
// 			return fmt.Sprintf("прямоугольник (высота - %.1f, длина - %.1f)", v.Height, v.Width)
// 		case Triangle:
// 			return fmt.Sprintf("триугольник (основание: %.1f высота: %.1f)", v.Base, v.Height)
// 		default:
// 			return fmt.Sprintf("фигура (площадь=%.2f)", s.Area())

// 	}
// }

// 2

// func LineCount(r io.Reader) (int, error) {
// 	count := 0
// 	scanner := bufio.NewScanner(r)

// 	for scanner.Scan() {
// 		count++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return 0, fmt.Errorf("ошибка сканера")
// 	}

// 	return count, nil
// }

// type LimitReader struct {
//     R         io.Reader // внутренний reader
//     Remaining int       // сколько байт ещё можно прочитать
// }

// func (lr *LimitReader) Read(p []byte) (int, error) {
// 	if lr.Remaining <= 0 {
// 		return 0, io.EOF
// 	}
// 	if len(p) > lr.Remaining {
// 		p = p[:lr.Remaining]
// 	}
// 	n, err := lr.R.Read(p)
// 	lr.Remaining -= n
// 	return n, err
// }

func LineCount(s io.Reader) (int, error) {
	scanner := bufio.NewScanner(s)
	scanner.Split(bufio.ScanWords)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count, scanner.Err()
}

func main() {
	// shapes := []Shape{
	// 	Circle{Radius: 5},
	// 	Rectangle{Width: 3, Height: 4},
	// 	Circle{Radius: 2},
	// 	Triangle{Base: 2, Height: 4},
	// }
	// fmt.Printf("Общая площадь: %.2f\n", TotalArea(shapes))

	// for _, shape := range shapes {
	// 	fmt.Println(DescribeShape(shape))
	// }

	// Проектирование
	// 1
	// n, _ := LineCount(strings.NewReader("one\ntwo\nthree"))
	// fmt.Println(n)
	
	// f, _ := os.Open("testdata/poem.txt")
	// defer f.Close()
	// n, _ = LineCount(f)
	// fmt.Println(n)

	// // 2
	// r := strings.NewReader("Hello, World!")
	// lr := &LimitReader{R: r, Remaining: 5}
	
	// data, _ := io.ReadAll(lr)
	// fmt.Println(string(data)) // "Hello"

	// 4
	file, err := os.Open("testdata/poem.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	count, err := LineCount(strings.NewReader("hello world"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Количество строк:", count)
	countF, err := LineCount(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Количество строк:", countF)
}