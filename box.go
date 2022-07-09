package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity != 0 {
		b.shapes = append(b.shapes, shape)
		b.shapesCapacity -= 1
		return nil
	}

	return fmt.Errorf("not enough room")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	l := len(b.shapes)
	if i >= 0 && i < l {
		return b.shapes[i], nil
	}

	return nil, fmt.Errorf("index doesn't exist or went out of the range")
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if shape, err := b.GetByIndex(i); err == nil {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return shape, nil
	} else {
		return nil, err
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if s, err := b.GetByIndex(i); err == nil {
		b.shapes[i] = shape
		return s, nil
	} else {
		return nil, err
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64

	for i := range b.shapes {
		sum += b.shapes[i].CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64

	for i := range b.shapes {
		sum += b.shapes[i].CalcArea()
	}

	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var exist bool

	for i := 0; i < len(b.shapes); i++ {
		if _, ok := b.shapes[i].(*Circle); ok {
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			i -= 1
			exist = true
		}
	}

	if !exist {
		return fmt.Errorf("circles are not exist in the list")
	}

	return nil
}
