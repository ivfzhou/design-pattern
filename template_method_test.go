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

// 模版方法是一种行为设计模式，它在基类中定义了一个算法的框架，允许子类在不修改结构的情况下重写算法的特定步骤。
// 当你只希望客户端扩展某个特定算法步骤，而不是整个算法或其结构时，可使用模板方法模式。
// 当多个类的算法除一些细微不同之外几乎完全一样时，你可使用该模式。但其后果就是，只要算法发生变化，你就可能需要修改所有的类。
package design_pattern_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

/*
一次性口令（OTP）校验是模板方法模式的典型场景：校验流程由固定的算法骨架组成——生成随机码、保存到缓存、组织消息内容、发送通知、校验口令。
其中「生成、缓存、校验」这类通用步骤沉淀在 `OTPImpl` 中作为默认实现；
而「组织消息」与「发送通知」是随渠道变化的步骤，由 `Sms`、`Email` 各自实现。
`Sms`、`Email` 通过内嵌 `OTPImpl` 复用公共步骤，仅重写差异部分，从而在不改动整体流程结构的前提下扩展出不同的发送渠道。
*/

func TestTemplateMethod(t *testing.T) {
	var o OTP

	o = &Sms{}
	otp := o.GenRandomOTP(4)
	o.SaveOTPCache(otp)
	message := o.GetMessage(otp)
	t.Log(o.SendNotification(message))

	o = &Email{}
	otp = o.GenRandomOTP(4)
	o.SaveOTPCache(otp)
	message = o.GetMessage(otp)
	t.Log(o.SendNotification(message))
}

// ===

// OTP one-time password
type OTP interface {
	// GenRandomOTP 生成随机的 n 位数字
	GenRandomOTP(int) string

	// SaveOTPCache 在缓存中保存这组数字以便进行后续验证
	SaveOTPCache(string)

	// GetMessage 准备内容
	GetMessage(string) string

	// SendNotification 发送通知
	SendNotification(string) error

	ValidateOTP(string) bool
}

// =

type OTPImpl struct {
	OTP string
}

func (o *OTPImpl) GenRandomOTP(l int) string {
	sb := strings.Builder{}
	sb.Grow(l)
	for i := 0; i < l; i++ {
		sb.WriteString(strconv.Itoa(rand.Intn(10)))
	}
	return sb.String()
}

func (o *OTPImpl) SaveOTPCache(otp string) {
	o.OTP = otp
}

func (o *OTPImpl) ValidateOTP(otp string) bool {
	return o.OTP == otp
}

// =

type Sms struct {
	OTPImpl
}

func (s *Sms) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) SendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

// =

type Email struct {
	OTPImpl
}

func (s *Email) GetMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *Email) SendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}
