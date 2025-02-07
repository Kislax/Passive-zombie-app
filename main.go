package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Выводим сообщение в консоль
	fmt.Println("Приложение-эмулятор работает, можете запускать игру")
	fmt.Println("Для завершения работы нажмите Ctrl+C или закройте приложение")

	// Создаем канал для перехвата сигналов завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Запускаем бесконечный цикл, чтобы приложение не завершалось
	go func() {
		for {
			time.Sleep(time.Second) // Чтобы не нагружать процессор
		}
	}()

	// Ожидание сигнала завершения (блокирующий вызов)
	<-sigChan
	fmt.Println("\nПриложение завершено.")
}
