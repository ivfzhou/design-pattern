/*
 * Copyright (c) 2026 ivfzhou
 * design-pattern is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

// 访问者模式是一种行为设计模式，它能将算法与其所作用的对象隔离开来。
// 如果你需要对一个复杂对象结构（例如对象树）中的所有元素执行某些操作，可使用访问者模式。
// 可使用访问者模式来清理辅助行为的业务逻辑。
// 当某个行为仅在类层次结构中的一些类中有意义，而在其他类中没有意义时，可使用该模式。
package design_pattern_test

import (
	"fmt"
	"math"
	"testing"
)

/*
图形计算是访问者模式的典型场景：`Shape`（元素）实现 `Accept(Visitor)`，`Visitor` 接口声明 `Visit(Shape)`，
而 `AreaCalculator`、`MiddleCoordinate` 是具体的访问者，分别封装「计算面积」与「计算中心坐标」两种算法。
当需要为图形增加一种新操作时，不必改动各个图形类，只需新增一个访问者；
元素通过 `Accept` 把自身作为参数传给访问者，由访问者在 `Visit` 内按具体类型分发处理，从而将算法（操作）与元素结构解耦。
*/

func TestVisitor(t *testing.T) {
	var square Shape = &Square{Side: 2}
	var circle Shape = &Circle{Radius: 3}
	var rectangle Shape = &Rectangle{Length: 2, Width: 3}

	areaCalculator := &AreaCalculator{}
	square.Accept(areaCalculator)
	fmt.Println("AreaCalculator square", areaCalculator.GetArea())
	circle.Accept(areaCalculator)
	fmt.Println("AreaCalculator circle", areaCalculator.GetArea())
	rectangle.Accept(areaCalculator)
	fmt.Println("AreaCalculator rectangle", areaCalculator.GetArea())

	middleCoordinates := &MiddleCoordinate{}
	square.Accept(middleCoordinates)
	fmt.Println("MiddleCoordinate square", middleCoordinates.GetMiddleCoordinate())
	circle.Accept(middleCoordinates)
	fmt.Println("MiddleCoordinate circle", middleCoordinates.GetMiddleCoordinate())
	rectangle.Accept(middleCoordinates)
	fmt.Println("MiddleCoordinate rectangle", middleCoordinates.GetMiddleCoordinate())
}

// ===

type Shape interface {
	GetType() string
	Accept(Visitor)
}

// =

type Rectangle struct {
	Length int
	Width  int
}

func (t *Rectangle) Accept(v Visitor) {
	v.Visit(t)
}

func (t *Rectangle) GetType() string {
	return "Rectangle"
}

// =

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v Visitor) {
	v.Visit(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

// =

type Square struct {
	Side int
}

func (s *Square) Accept(v Visitor) {
	v.Visit(s)
}

func (s *Square) GetType() string {
	return "Square"
}

// =

type Visitor interface {
	Visit(Shape)
}

// =

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) Visit(s Shape) {
	switch v := s.(type) {
	case *Square:
		a.area = v.Side * v.Side
	case *Circle:
		a.area = int(float64(v.Radius*v.Radius) * math.Pi)
	case *Rectangle:
		a.area = v.Length * v.Width
	}
}

func (a *AreaCalculator) GetArea() int {
	return a.area
}

// =

type MiddleCoordinate struct {
	x int
	y int
}

func (m *MiddleCoordinate) Visit(s Shape) {
	switch v := s.(type) {
	case *Square:
		m.x = v.Side / 2
	case *Circle:
		m.x = v.Radius
		m.y = v.Radius
	case *Rectangle:
		m.x = v.Length
		m.y = v.Width
	}
}

func (m *MiddleCoordinate) GetMiddleCoordinate() [2]int {
	return [2]int{m.x, m.y}
}
