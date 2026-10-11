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

// 装饰模式是一种结构型设计模式，允许你通过将对象放入包含行为的特殊封装对象中来为原对象绑定新的行为。
// 如果你希望在无需修改代码的情况下即可使用对象，且希望在运行时为对象新增额外的行为，可以使用装饰模式。
// 如果用继承来扩展对象行为的方案难以实现或者根本不可行，你可以使用该模式。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
披萨加配料是装饰模式的典型场景：`VeggieMania` 是被装饰的基础披萨，`CheeseTopping` 和`TomatoTopping` 是装饰器，
它们与被装饰对象实现同一个 `IPizza` 接口，并各自持有一个`IPizza`。
通过层层包裹（`TomatoTopping` 包 `CheeseTopping` 包 `VeggieMania`），
调用 `GetPrice` 时每一层都在内层价格上追加自己的配料费，从而在运行时动态地为原对象叠加行为，
而无需修改原有披萨的代码。
*/

func TestDecorator(t *testing.T) {
	var pizza IPizza = &VeggieMania{}

	// Add cheese topping
	var pizzaWithCheese IPizza = &CheeseTopping{
		pizza: pizza,
	}

	// Add tomato topping
	var pizzaWithCheeseAndTomato IPizza = &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}

// ===

type IPizza interface {
	GetPrice() int
}

// =

type VeggieMania struct{}

func (p *VeggieMania) GetPrice() int {
	return 15
}

// =

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.pizza.GetPrice()
	return pizzaPrice + 7
}

// =

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.pizza.GetPrice()
	return pizzaPrice + 10
}
