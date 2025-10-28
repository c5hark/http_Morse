# Morse Code Converter

Веб-сервер на Go, который позволяет загружать файлы через HTML-форму и автоматически определяет, содержат ли они текст или код Морзе.
После определения тип данных конвертируется в противоположный формат (текст ⇄ код Морзе), а результат сохраняется в новый файл.

# Функциональность

Загрузка файлов через HTML-форму

Автоматическое определение содержимого файла (текст или код Морзе)

Конвертация текста в код Морзе и наоборот

Сохранение результата в новый файл

Возврат результата пользователю

# Технологии
Golang, HTML, HTTP Server, Standard Library

# Ценность проекта

Проект демонстрирует навыки:

Создания и настройки веб-сервера на Go

Работы с HTML-формами и обработкой загрузок файлов

Организации многопакетной структуры проекта

Обработки ошибок и логирования

Использования стандартной библиотеки для реальных задач

# Запуск проекта

Клонируйте репозиторий:

`git clone https://github.com/c5hark/http_Morse.git`
`cd http_Morse`

Установите зависимости:

`go mod tidy`

Запустите сервер:

`go run main.go`

Откройте в браузере:

`http://localhost:8080`

----

A Go web server implementing a task scheduling service similar to a TODO list.
It allows creating, editing, deleting, and completing tasks, including recurring ones based on defined rules.

# Features

Add new tasks

Get task list

View task details

Edit tasks

Delete tasks

Mark tasks as completed with recurrence

# Technologies

Go (Golang), HTML, HTTP Server, Standard Library

# Value

This project demonstrates skills in:

Web server development with Go

Working with databases

CRUD operations implementation

REST API design

Writing and testing business logic

# Run the project

Clone the repository:

`git clone https://github.com/c5hark/http_Morse.git`
`cd http_Morse`

Install dependencies:

`go mod tidy`

Run the application:

`go run main.go`

Open in browser:

`http://localhost:8080`
