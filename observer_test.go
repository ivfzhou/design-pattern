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

// 观察者是一种行为设计模式，允许一个对象将其状态的改变通知其他对象。
// 当一个对象状态的改变需要改变其他对象，或实际对象是事先未知的或动态变化的时，可使用观察者模式。
// 当应用中的一些对象必须观察其他对象时，可使用该模式。但仅能在有限时间内或特定情况下使用。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
商品补货通知是观察者模式的典型场景：`Item`（被观察者/主题）维护一份观察者列表，`Customer`（观察者）实现 `Observer` 接口。
消费者通过 `Register`/`Deregister` 动态订阅或退订某个商品，当商品状态变化时（`UpdateAvailability`），
`Item` 通过 `NotifyAll` 遍历并逐一通知所有订阅者，使它们无需主动轮询即可在变化发生时收到更新，实现了主题与观察者之间的松耦合。
*/

func TestObserver(t *testing.T) {
	shirtItem := NewItem("Nike Shirt")

	var observerFirst Observer = &Customer{id: "abc@gmail.com"}
	var observerSecond Observer = &Customer{id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()

	shirtItem.Deregister(observerSecond)
	shirtItem.UpdateAvailability()
}

// ===

type Subject interface {
	Register(Observer)
	Deregister(Observer)
	NotifyAll()
}

// =

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.NotifyAll()
}

func (i *Item) Register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) Deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.observerList {
		observer.Update(i.name)
	}
}

// =

type Observer interface {
	Update(string)
	GetID() string
}

// =

type Customer struct {
	id string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) GetID() string {
	return c.id
}

// =

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = nil, observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
