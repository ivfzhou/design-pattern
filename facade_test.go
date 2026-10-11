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

// 外观模式是一种结构型设计模式，能为程序库、框架或其他复杂类提供一个简单的接口。
// 如果你需要一个指向复杂子系统的直接接口，且该接口的功能有限，则可以使用外观模式。
// 如果需要将子系统组织为多层结构，可以使用外观。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
电子钱包是外观模式的典型场景：一次存取款实际上要协调 `Account`（账户校验）、`SecurityCode`（安全码校验）、
`Wallet`（余额增减）、`Notification`（通知）和 `Ledger`（记账）等多个子系统的复杂交互。
`WalletFacade` 作为外观，把这些内部细节封装成 `AddMoneyToWallet`、`DeductMoneyFromWallet`这样简单的统一入口，
客户端只需调用一个方法即可完成整套流程，无需关心子系统的实现细节与调用顺序。
*/

func TestFacade(t *testing.T) {
	walletFacade := NewWalletFacade("abc", 1234)

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}
}

// ===

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacade := &WalletFacade{
		account:      NewAccount(accountID),
		securityCode: NewSecurityCode(code),
		wallet:       NewWallet(),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
	fmt.Println("Account created")
	return walletFacade
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting Add money to wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.CreditBalance(amount)
	w.notification.SendWalletCreditNotification()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.SendWalletDebitNotification()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

// =

type Account struct {
	name string
}

func NewAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) CheckAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

// =

type SecurityCode struct {
	code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) CheckCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("security code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

// =

type Wallet struct {
	balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) CreditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
}

func (w *Wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}

// =

type Ledger struct{}

func (s *Ledger) MakeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
}

// =

type Notification struct{}

func (n *Notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
