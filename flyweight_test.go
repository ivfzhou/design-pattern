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

// 享元模式是一种结构型设计模式，它摒弃了在每个对象中保存所有数据的方式，通过共享多个对象所共有的相同状态，让你能在有限的内存容量中载入更多对象。
// 仅在程序必须支持大量对象且没有足够的内存容量时使用享元模式。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
游戏中大量玩家共享外观是享元模式的典型场景：玩家对象把状态划分为
「内在状态」（如衣服颜色、阵营外观，由 `TerroristDress`/`CounterTerroristDress` 表示，可被多个玩家共享）
与「外在状态」（如玩家的位置 `lat`/`long`，随个体变化）。
`DressFactory` 作为享元工厂，按类型缓存并复用衣服对象——同类型的玩家拿到的是同一个 `Dress` 实例，而不是各自创建一份，
从而在大规模玩家场景下显著节省内存；而每个玩家独有的位置信息则仍保存在各自的 `Player` 中。
*/

func TestFlyweight(t *testing.T) {
	game := NewGame()

	// Add Terrorist
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)

	// Add CounterTerrorist
	game.AddCounterTerrorist(CounterTerroristDressType)
	game.AddCounterTerrorist(CounterTerroristDressType)
	game.AddCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}

// ===

const (
	// TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	// CounterTerroristDressType terrorist dress type
	CounterTerroristDressType = "ctDress"
)

type Dress interface {
	GetColor() string
}

// =

type DressFactory struct {
	dressMap map[string]Dress
}

var dressFactorySingleInstance = &DressFactory{
	dressMap: make(map[string]Dress),
}

func GetDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

func (d *DressFactory) GetDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = NewTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = NewCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}

// =

type TerroristDress struct {
	color string
}

func NewTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

func (t *TerroristDress) GetColor() string {
	return t.color
}

// =

type CounterTerroristDress struct {
	color string
}

func NewCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}

func (c *CounterTerroristDress) GetColor() string {
	return c.color
}

// =

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func NewPlayer(playerType, dressType string) *Player {
	dress, _ := GetDressFactorySingleInstance().GetDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

// =

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func NewGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (c *Game) AddTerrorist(dressType string) {
	player := NewPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
}

func (c *Game) AddCounterTerrorist(dressType string) {
	player := NewPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
}
