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

// 组合模式是一种结构型设计模式，你可以使用它将对象组合成树状结构，并且能像使用独立对象一样使用它们。
// 如果你需要实现树状对象结构，可以使用组合模式。
// 如果你希望客户端代码以相同方式处理简单和复杂元素，可以使用该模式。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
文件系统是组合模式的典型场景：`File`（叶子节点）与 `Folder`（容器节点）都实现了同一个 `Component` 接口，
`Folder` 通过持有一组 `Component` 把自己和子节点组织成树状结构，
从而让客户端能用统一的方式对待单个文件与整个文件夹——对 `Folder` 调用`Search` 时会递归地把请求下发给它的每一个子节点。
*/

func TestComposite(t *testing.T) {
	var file1 Component = &File{name: "File1"}
	var file2 Component = &File{name: "File2"}
	var file3 Component = &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}

// ===

type Component interface {
	Search(string)
}

// =

type File struct {
	name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) GetName() string {
	return f.name
}

// =

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}
