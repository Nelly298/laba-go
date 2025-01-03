package main // Определяем пакет main, который является точкой входа в программу

import (
	"errors" // Импортируем пакет для работы с ошибками
	"fmt"    // Импортируем пакет для форматированного ввода-вывода
)

// Функция formatIP принимает массив из 4 байтов и возвращает строку с IP-адресом
func formatIP(ip [4]byte) string {
	// Используем fmt.Sprintf для форматирования строки в виде "x.x.x.x"
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Функция listEven принимает диапазон и возвращает срез четных чисел и ошибку
func listEven(start, end int) ([]int, error) {
	// Проверяем, если левая граница больше правой
	if start > end {
		// Если это так, возвращаем nil и ошибку с сообщением
		return nil, errors.New("левая граница больше правой")
	}

	// Объявляем срез для хранения четных чисел
	var evens []int
	// Проходим по всем числам в заданном диапазоне
	for i := start; i <= end; i++ {
		// Проверяем, является ли число четным
		if i%2 == 0 {
			// Если четное, добавляем его в срез evens
			evens = append(evens, i)
		}
	}
	// Возвращаем срез четных чисел и nil как значение ошибки
	return evens, nil
}

func main() {
	// Пример использования функции formatIP
	// Создаем массив из 4 байтов, представляющий IP-адрес
	ip := [4]byte{127, 0, 0, 1}
	// Выводим отформатированный IP-адрес на экран
	fmt.Println("IP-адрес:", formatIP(ip))

	// Пример использования функции listEven
	// Вызываем функцию listEven с диапазоном от 1 до 10
	evens, err := listEven(1, 10)
	// Проверяем, произошла ли ошибка
	if err != nil {
		// Если ошибка есть, выводим сообщение об ошибке
		fmt.Println("Ошибка:", err)
	} else {
		// Если ошибки нет, выводим срез четных чисел
		fmt.Println("Чётные числа:", evens)
	}

	// Пример с ошибкой
	// Вызываем функцию listEven с неправильным диапазоном (10, 1)
	evens, err = listEven(10, 1)
	// Проверяем, произошла ли ошибка
	if err != nil {
		// Если ошибка есть, выводим сообщение об ошибке
		fmt.Println("Ошибка:", err)
	}
}
