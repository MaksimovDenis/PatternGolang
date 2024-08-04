package main

import (
	"log"
	"os"
)

// Интерфейс определяет контракт на
// требуемое поведение объектов.
type Logger interface {
	Log(string) error
}

// Тип данных, выполняющий контракт Logger.
type FileLogger struct{}

func (*FileLogger) Log(s string) error {
	f, err := os.Create("./log.txt")
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(s))
	if err != nil {
		return err
	}
	return nil
}

// Тип данных, выполняющий контракт Logger.
type ConsoleLogger struct{}

func (*ConsoleLogger) Log(s string) error {
	log.Println(s)
	return nil
}

type Log struct {
	records []string
	log     Logger
}

func main() {
	l := Log{
		records: []string{"Record 1", "Record 2"},
	}
	//
	// Подменяем поведение при необходимости.
	// Используем разные "стратегии".
	//
	// Вывод сообщений в файл.
	l.log = new(FileLogger)
	err := l.log.Log(l.records[0])
	if err != nil {
		log.Fatal(err)
	}
	// Вывод сообщений на экран.
	l.log = new(ConsoleLogger)
	l.log.Log(l.records[1])
}
