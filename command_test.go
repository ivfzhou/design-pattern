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

// 命令是一种行为设计模式，它可将请求或简单操作转换为一个对象。
// 如果你需要通过操作来参数化对象，可使用命令模式。
// 如果你想要将操作放入队列中、操作的执行或者远程执行操作，可使用命令模式。
// 如果你想要实现操作回滚功能，可使用命令模式。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
遥控器按钮是命令模式的典型场景：`TV` 是接收者（真正执行开关机操作），`OnCommand`、`OffCommand`把「开机」「关机」这些操作封装成独立的命令对象，
`Button` 则是调用者，它只持有一个 `Command` 接口并调用 `Execute`，完全不知道具体执行的是哪个设备的什么操作。
这样就把「发起请求的按钮」与「实际执行的设备」解耦——同一个按钮可以绑定不同命令，命令也可被替换、排队或用于后续回滚扩展。
*/

func TestCommand(t *testing.T) {
	var tv Device = &TV{}

	var onCommand Command = &OnCommand{
		device: tv,
	}

	var offCommand Command = &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.Press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.Press()
}

// ===

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

// =

type Command interface {
	Execute()
}

// =

type OnCommand struct {
	device Device
}

func (c *OnCommand) Execute() {
	c.device.On()
}

// =

type OffCommand struct {
	device Device
}

func (c *OffCommand) Execute() {
	c.device.Off()
}

// =

type Device interface {
	On()
	Off()
}

// =

type TV struct {
	isRunning bool
}

func (t *TV) On() {
	t.isRunning = true
	fmt.Println("Turning tv On")
}

func (t *TV) Off() {
	t.isRunning = false
	fmt.Println("Turning tv Off")
}
