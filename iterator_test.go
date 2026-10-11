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

// 迭代器是一种行为设计模式，让你能在不暴露复杂数据结构内部细节的情况下遍历其中所有的元素。
// 当集合背后为复杂的数据结构，且你希望对客户端隐藏其复杂性时（出于使用便利性或安全性的考虑），可以使用迭代器模式。
// 使用该模式可以减少程序中重复的遍历代码。
// 如果你希望代码能够遍历不同的甚至是无法预知的数据结构，可以使用迭代器模式。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
用户集合的遍历是迭代器模式的典型场景：`UserCollection` 作为集合，负责通过 `CreateIterator` 生成对应的 `UserIterator`；
`Iterator` 接口抽象出统一的遍历方式（`HasNext`/`GetNext`），
而遍历所需的游标 `index` 及底层 `users` 切片都被封装在迭代器内部。
客户端只面向 `Iterator` 接口遍历，无需了解集合的内部存储结构；
即使将来集合改用其他数据结构，只要换一个迭代器实现，遍历代码保持不变，从而实现「集合」与「遍历」的解耦。
*/

func TestIterator(t *testing.T) {
	user1 := &User{
		Name: "a",
		Age:  30,
	}
	user2 := &User{
		Name: "b",
		Age:  20,
	}

	var userCollection Collection = &UserCollection{
		users: []*User{user1, user2},
	}

	var iterator Iterator = userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}

// ===

type Collection interface {
	CreateIterator() Iterator
}

type Iterator interface {
	HasNext() bool
	GetNext() *User
}

type User struct {
	Name string
	Age  int
}

// =

type UserCollection struct {
	users []*User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

// =

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) HasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false

}

func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
