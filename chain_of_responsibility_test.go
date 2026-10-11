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

// 责任链是一种行为设计模式，允许你将请求沿着处理者链进行发送，直至其中一个处理者对其进行处理。
// 当程序需要使用不同方式处理不同种类请求，而且请求类型和顺序预先未知时，可以使用责任链模式。
// 当必须按顺序执行多个处理者时，可以使用该模式。
// 如果所需处理者及其顺序必须在运行时进行改变，可以使用责任链模式。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
病人就诊流程是责任链模式的典型场景：`Reception`、`Doctor`、`Medical`、`Cashier` 各是一个处理者，都实现 `Department` 接口。
每个处理者只负责自己这一步（挂号、问诊、取药、缴费），处理完后通过 `next` 把请求 `Patient` 交给下一个处理者，
从而串成一条「接待 → 医生 → 药房 → 收费」的责任链。
请求发送者（客户端）只调用链头的 `Execute`，无需关心整个流程由谁、按何顺序完成；
处理者与顺序都通过 `SetNext` 在运行时动态组装，扩展或调整环节时不必改动现有处理者的代码。
*/

func TestChainOfResponsibility(t *testing.T) {
	var cashier Department = &Cashier{}

	// Set next for medical department
	var medical Department = &Medical{}
	medical.SetNext(cashier)

	// Set next for doctor department
	var doctor Department = &Doctor{}
	doctor.SetNext(medical)

	// Set next for reception department
	var reception Department = &Reception{}
	reception.SetNext(doctor)

	patient := &Patient{Name: "abc"}
	// Patient visiting
	reception.Execute(patient)
}

// ===

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

// =

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}

// =

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}

// =

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}

// =

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
