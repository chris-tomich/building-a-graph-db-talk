package main

import (
	"sync"
	"fmt"
)

func main() {
	sharedMap := NewMutexSharedMap()

	wg1 := &sync.WaitGroup{}

	go sharedMap.WriteToMap(wg1, "Emma", "A message from Liam to Emma")
	go sharedMap.WriteToMap(wg1, "Liam", "A message from Emma to Liam")
	go sharedMap.WriteToMap(wg1, "Olivia", "A message from Noah to Olivia")
	go sharedMap.WriteToMap(wg1, "Noah", "A message from Olivia to Noah")
	go sharedMap.WriteToMap(wg1, "Ava", "A message from Isabella to Ava")
	go sharedMap.WriteToMap(wg1, "Lucas", "A message from Mason to Lucas")
	go sharedMap.WriteToMap(wg1, "Isabella", "A message from Ava to Isabella")
	go sharedMap.WriteToMap(wg1, "Mason", "A message from Lucas to Mason")
	go sharedMap.WriteToMap(wg1, "Sophia", "A message from Logan to Sophia")
	go sharedMap.WriteToMap(wg1, "Logan", "A message from Sophia to Logan")

	wg1.Add(10)
	wg1.Wait()

	wg2 := &sync.WaitGroup{}

	go sharedMap.ReadFromMap(wg2, "Emma")
	go sharedMap.ReadFromMap(wg2, "Liam")
	go sharedMap.ReadFromMap(wg2, "Olivia")
	go sharedMap.ReadFromMap(wg2, "Noah")
	go sharedMap.ReadFromMap(wg2, "Ava")
	go sharedMap.ReadFromMap(wg2, "Lucas")
	go sharedMap.ReadFromMap(wg2, "Isabella")
	go sharedMap.ReadFromMap(wg2, "Mason")
	go sharedMap.ReadFromMap(wg2, "Sophia")
	go sharedMap.ReadFromMap(wg2, "Logan")

	wg2.Add(10)
	wg2.Wait()
}

func NewMutexSharedMap() *MutexSharedMap {
	m := &MutexSharedMap{
		sharedMap: make(map[string]string),
		RWMutex: &sync.RWMutex{},
	}

	return m
}

type MutexSharedMap struct {
	sharedMap map[string]string
	*sync.RWMutex
}

func (m *MutexSharedMap) WriteToMap(wg *sync.WaitGroup, recipient string, message string) {
	m.Lock()
	m.sharedMap[recipient] = message
	m.Unlock()
	wg.Done()
}

func (m *MutexSharedMap) ReadFromMap(wg *sync.WaitGroup, recipient string) {
	m.RLock()
	message := m.sharedMap[recipient]
	m.RUnlock()

	fmt.Println(message)
	wg.Done()
}
