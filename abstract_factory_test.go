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

// 抽象工厂模式是一种创建型设计模式，它能创建一系列相关的对象，而无需指定其具体类。
// 如果代码需要与多个不同系列的相关产品交互，但是由于无法提前获取相关信息，或者出于对未来扩展性的考虑，
// 你不希望代码基于产品的具体类进行构建，在这种情况下，你可以使用抽象工厂。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
运动品牌产品族是抽象工厂模式的典型场景：`ISportsFactory` 是抽象工厂，声明了 `MakeShoe`、`MakeShirt` 两个创建方法；
`Adidas`、`Nike` 是具体工厂，各自返回同一品牌下「相互配套」的鞋和衣服（产品族）。
客户端通过 `GetSportsFactory` 拿到某个品牌的工厂后，只面向 `ISportsFactory`、`IShoe`、`IShirt` 等抽象接口编程，
完全不依赖具体产品类。这样更换整个产品族只需替换工厂实例，且能保证同一工厂产出的鞋与衣服品牌一致。
*/

func TestAbstractFactory(t *testing.T) {
	adidasFactory, _ := GetSportsFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()
	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	nikeFactory, _ := GetSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()
	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Printf("Size: %d", s.GetSize())
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Printf("Size: %d", s.GetSize())
}

// ===

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}

// =

type Adidas struct{}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		logo: "adidas",
		size: 14,
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		logo: "adidas",
		size: 14,
	}
}

// =

type Nike struct{}

func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		logo: "nike",
		size: 14,
	}
}

func (n *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		logo: "nike",
		size: 14,
	}
}

// =

type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}

func (s *Shoe) GetSize() int {
	return s.size
}

type AdidasShoe struct {
	Shoe
}

type NikeShoe struct {
	Shoe
}

// =

type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) GetLogo() string {
	return s.logo
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}

func (s *Shirt) GetSize() int {
	return s.size
}

type AdidasShirt struct {
	Shirt
}

type NikeShirt struct {
	Shirt
}
