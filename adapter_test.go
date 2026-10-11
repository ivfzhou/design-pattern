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

// 适配器模式是一种结构型设计模式，它能使接口不兼容的对象能够相互合作。
// 当你希望使用某个类，但是其接口与其他代码不兼容时，可以使用适配器类。
// 如果您需要复用这样一些类，他们处于同一个继承体系，并且他们又有了额外的一些共同的方法，
// 但是这些共同的方法不是所有在这一继承体系中的子类所具有的共性。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
电脑外接接口是适配器模式的典型场景：客户端期望的接口是 `Computer_`（含 `InsertIntoLightningPort`），
`Mac_` 天然兼容；而 `Windows_` 只提供 `InsertIntoUSBPort`，接口不兼容。
`WindowsAdapter` 作为适配器，实现 `Computer_` 接口并内部持有一个 `Windows_`，把客户端的 Lightning 调用转换为 Windows 的 USB 调用，
从而让原本无法直接配合的 `Client` 与 `Windows_` 能够协同工作，而无需修改任何一方的原有代码。
*/

func TestAdapter(t *testing.T) {
	client := &Client{}

	var mac Computer_ = &Mac_{}
	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows_{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

// ===

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer_) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

// =

type Computer_ interface {
	InsertIntoLightningPort()
}

// =

type Mac_ struct{}

func (m *Mac_) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// =

type WindowsAdapter struct {
	windowMachine *Windows_
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.InsertIntoUSBPort()
}

// =

type Windows_ struct{}

func (w *Windows_) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
