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

// 工厂方法模式是一种创建型设计模式，其在父类中提供一个创建对象的方法，允许子类决定实例化对象的类型。
// 当你在编写代码的过程中，如果无法预知对象确切类别及其依赖关系时，可使用工厂方法。
// 如果你希望用户能扩展你软件库或框架的内部组件，可使用工厂方法。
// 如果你希望复用现有对象来节省系统资源，而不是每次都重新创建对象，可使用工厂方法。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
枪械创建是工厂方法模式的典型场景：所有枪都实现统一的 `IGun` 接口（`Ak47`、`Musket` 通过内嵌 `Gun` 复用公共字段与方法），
而具体实例化逻辑被集中在工厂函数 `GetGun` 中——它根据传入的类型决定创建哪种枪并返回 `IGun`，
客户端只面向接口编程，不直接 `new` 具体类型。
这样当需要新增一种枪时，只需扩展工厂分支，调用方代码保持不变，从而将「对象的创建」与「对象的使用」解耦。
*/

func TestFactoryMethod(t *testing.T) {
	ak47, _ := GetGun("ak47")
	musket, _ := GetGun("Musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}

// ===

type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "Musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

// =

type Gun struct {
	name  string
	power int
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func (g *Gun) GetPower() int {
	return g.power
}

// =

type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		name:  "AK47 gun",
		power: 4,
	}
}

// =

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		name:  "Musket gun",
		power: 1,
	}
}
