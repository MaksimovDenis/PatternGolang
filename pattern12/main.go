package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Очередь сообщений.
type Queue struct {
	Messages []Message
	// Композиция (внедрение) интерфейса является
	// реализацией шаблона "декоратор".
	io.Writer
}

// Сообщение.
type Message struct {
	ID   int
	Time int64
	Msg  string
}

func main() {
	var mq Queue
	mq.Messages = []Message{
		{ID: 1, Msg: "Message"},
	}
	f, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	mq.Writer = f
	b, err := json.Marshal(mq.Messages)
	if err != nil {
		log.Fatal(err)
	}
	// Возможности внедренного объекта
	// доступны внешнему объекту.
	// Поэтому у очереди появляется
	// метод Write.
	mq.Write(b)
}
