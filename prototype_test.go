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

// 原型模式是一种创建型设计模式，使你能够复制已有对象，而又无需使代码依赖它们所属的类。
// 如果你需要复制一些对象，同时又希望代码独立于这些对象所属的具体类，可以使用原型模式。
// 如果子类的区别仅在于其对象的初始化方式，那么你可以使用该模式来减少子类的数量。别人创建这些子类的目的可能是为了创建特定类型的对象。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
文件系统层级结构的复制是原型模式的典型场景：`Inode` 接口声明了 `Clone()`，`File_`（叶子）与 `Folder_`（容器）各自实现克隆逻辑，
其中 `Folder_.Clone` 会递归地克隆其所有子节点，从而深拷贝整棵树。
客户端只面向 `Inode` 接口调用 `Clone`，无需知道被复制的具体类型——即使新增了新的节点类型，
只要实现 `Clone` 即可被复制，实现「复制已有对象而代码不依赖其具体类」的目标。
*/

func TestPrototype(t *testing.T) {
	var file1 Inode = &File_{name: "File1"}
	var file2 Inode = &File_{name: "File2"}
	var file3 Inode = &File_{name: "File3"}

	var folder1 Inode = &Folder_{
		children: []Inode{file1},
		name:     "Folder1",
	}

	var folder2 Inode = &Folder_{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("Printing hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("Printing hierarchy for Clone Folder")
	cloneFolder.Print("  ")
}

// ===

type Inode interface {
	Print(string)
	Clone() Inode
}

// =

type File_ struct {
	name string
}

func (f *File_) Print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File_) Clone() Inode {
	return &File_{name: f.name + "_clone"}
}

// =

type Folder_ struct {
	children []Inode
	name     string
}

func (f *Folder_) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.Print(indentation + indentation)
	}
}

func (f *Folder_) Clone() Inode {
	cloneFolder := &Folder_{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copied := i.Clone()
		tempChildren = append(tempChildren, copied)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
