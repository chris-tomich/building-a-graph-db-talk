package main

import (
	"fmt"
	"time"
)

func main() {
	sharedMap := NewChannelSharedMap()

	go sharedMap.WriteToMap("Emma", "A message from Liam to Emma")
	go sharedMap.WriteToMap("Liam", "A message from Emma to Liam")
	go sharedMap.WriteToMap("Olivia", "A message from Noah to Olivia")
	go sharedMap.WriteToMap("Noah", "A message from Olivia to Noah")
	go sharedMap.WriteToMap("Ava", "A message from Isabella to Ava")
	go sharedMap.WriteToMap("Lucas", "A message from Mason to Lucas")
	go sharedMap.WriteToMap("Isabella", "A message from Ava to Isabella")
	go sharedMap.WriteToMap("Mason", "A message from Lucas to Mason")
	go sharedMap.WriteToMap("Sophia", "A message from Logan to Sophia")
	go sharedMap.WriteToMap("Logan", "A message from Sophia to Logan")

	go sharedMap.ReadFromMap("Emma")
	go sharedMap.ReadFromMap("Liam")
	go sharedMap.ReadFromMap("Olivia")
	go sharedMap.ReadFromMap("Noah")
	go sharedMap.ReadFromMap("Ava")
	go sharedMap.ReadFromMap("Lucas")
	go sharedMap.ReadFromMap("Isabella")
	go sharedMap.ReadFromMap("Mason")
	go sharedMap.ReadFromMap("Sophia")
	go sharedMap.ReadFromMap("Logan")

	<-sharedMap.Done
}

func NewChannelSharedMap() *ChannelSharedMap {
	m := &ChannelSharedMap{
		sharedMap: make(map[string]string),
		write:     make(chan *Message),
		read:      make(chan string),
		Done:      make(chan bool, 1),
	}

	go m.Loop()

	return m
}

type Message struct {
	Recipient string
	Message   string
}

type ChannelSharedMap struct {
	sharedMap map[string]string
	write     chan *Message
	read      chan string
	Done      chan bool
}

func (m *ChannelSharedMap) Loop() {
	ticker := time.NewTicker(time.Second * 5)
	defer func () {
		ticker.Stop()
	}()

	isInactive := false

	for {
		select {
		case <-ticker.C:
			if isInactive {
				fmt.Println("inactive for 10 secs, shutting down")
				m.Done <- true
				return
			} else {
				fmt.Println("inactive for 5 secs, will wait another 5 secs")
				isInactive = true
				continue
			}
		case msg := <-m.write:
			if msg != nil {
				m.sharedMap[msg.Recipient] = msg.Message
			}
		case recipient := <-m.read:
			message := m.sharedMap[recipient]

			if message == "" {
				go m.ReadFromMap(recipient)
				continue
			}

			delete(m.sharedMap, recipient)
			fmt.Println(message)
		}

		isInactive = false
	}
}

func (m *ChannelSharedMap) WriteToMap(recipient string, message string) {
	m.write <- &Message{Recipient: recipient, Message: message}
}

func (m *ChannelSharedMap) ReadFromMap(recipient string) {
	m.read <- recipient
}
