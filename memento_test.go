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

// 备忘录是一种行为设计模式，允许生成对象状态的快照并在以后将其还原。
// 当你需要创建对象状态快照来恢复其之前的状态时，可以使用备忘录模式。
// 当直接访问对象的成员变量、获取器或设置器将导致封装被突破时，可以使用该模式。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
笔记的保存与回滚是备忘录模式的典型场景：`Originator`（发起人）持有当前状态并负责创建、恢复快照，
`Memento`（备忘录）封装某一时刻的状态快照，`Caretaker`（负责人）只负责保存和取出快照、但不修改其内容。
每次修改状态前调用 `CreateMemento` 把快照交给 `Caretaker`存起来，
之后便可随时通过 `RestoreMemento` 回到任意历史状态，从而在不暴露对象内部实现的前提下实现状态撤销与恢复。
*/

func TestMemento(t *testing.T) {
	caretaker := &Caretaker{
		mementoArray: make([]*Memento, 0),
	}

	originator := &Originator{
		state: "A",
	}

	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(caretaker.GetMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

	originator.RestoreMemento(caretaker.GetMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.GetState())
}

// ===

type Originator struct {
	state string
}

func (e *Originator) CreateMemento() *Memento {
	return &Memento{state: e.state}
}

func (e *Originator) RestoreMemento(m *Memento) {
	e.state = m.GetSavedState()
}

func (e *Originator) SetState(state string) {
	e.state = state
}

func (e *Originator) GetState() string {
	return e.state
}

// =

type Memento struct {
	state string
}

func (m *Memento) GetSavedState() string {
	return m.state
}

// =

type Caretaker struct {
	mementoArray []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	return c.mementoArray[index]
}
