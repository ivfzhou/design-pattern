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

// 单例模式是一种创建型设计模式，让你能够保证一个类只有一个实例，并提供一个访问该实例的全局节点。
// 如果程序中的某个类对于所有客户端只有一个可用的实例，可以使用单例模式。
// 如果你需要更加严格地控制全局变量，可以使用单例模式。
package design_pattern_test

import (
	"fmt"
	"sync"
	"testing"
)

/*
本例通过 `GetInstance` 演示单例模式：它保证 `Single` 全局只有一个实例，并提供统一的全局访问点。
为避免并发场景下重复创建，这里采用了「双重检查锁定」——先在外层判空快速返回，避免每次访问都加锁；
仅当实例尚未创建时才加锁，并在锁内再次判空，从而兼顾线程安全与访问性能。
*/

func TestSingleton(t *testing.T) {
	for range 30 {
		go GetInstance()
	}

	fmt.Scanln()
}

// ===

var lock = &sync.Mutex{}

type Single struct{}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating Single instance now.")
			singleInstance = &Single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
