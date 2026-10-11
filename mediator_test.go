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

// 中介者是一种行为设计模式，让程序组件通过特殊的中介者对象进行间接沟通，达到减少组件之间依赖关系的目的。
// 当一些对象和其他对象紧密耦合以致难以对其进行修改时，可使用中介者模式。
// 当组件因过于依赖其他组件而无法在不同应用中复用时，可使用中介者模式。
// 如果为了能在不同情景下复用一些基本行为，导致你需要被迫创建大量组件子类时，可使用中介者模式。
package design_pattern_test

import (
	"fmt"
	"testing"
)

/*
列车调度是中介者模式的典型场景：`PassengerTrain`（客运列车）与 `FreightTrain`（货运列车）之间本会因抢占站台而互相依赖，
这里它们不再直接通信，而是统一通过 `StationManager`（站台管理调度）这个中介者协调。
列车到站先向 `CanArrive` 询问是否有空站台，被拒则进入等待队列；
离开时通过 `NotifyAboutDeparture` 通知中介者，由中介者按队列顺序唤醒下一列等待的列车 `PermitArrival`。
这样把多对多的交互收敛为各组件与中介者之间的一对一通信，列车之间彼此解耦、可独立复用与修改。
*/

func TestMediator(t *testing.T) {
	stationManager := NewStationManger()

	var passengerTrain Train = &PassengerTrain{
		mediator: stationManager,
	}
	var freightTrain Train = &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}

// ===

type Train interface {
	Arrive()
	Depart()
	PermitArrival()
}

// =

type PassengerTrain struct {
	mediator Mediator
}

func (g *PassengerTrain) Arrive() {
	if !g.mediator.CanArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.NotifyAboutDeparture()
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.Arrive()
}

// =

type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) Arrive() {
	if !g.mediator.CanArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.NotifyAboutDeparture()
}

func (g *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.Arrive()
}

// =

type Mediator interface {
	CanArrive(Train) bool
	NotifyAboutDeparture()
}

// =

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func NewStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) CanArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) NotifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}
