# Inter
Проект представляет собой реализацию сервиса по размещению задач, подсчету значения арифметической прогрессии в них и получения сведений о всех поставленных "задачах"

# Инструкция по запуску
```shell
go run cmd/main.go 5
#Запуск сервиса с установкой количества единовременно выполняемых задач 5

make testing
#Запуск тестов через http сервиса через запросы. Происходит запуск сервиса с ограниченным колчесетвом
#параллельных процессов и отправка запросов к сервису. После тестов все процессы останавливаются
```
# Запросы, обрабатываемые сервисом
`POST localhost:3000/create` - размещение новой задачи<br/>
`GET localhost:3000/get` - возвращает все имеющиеся задачи

# Формат подаваемых задач
Задачи подаются в теле запроса в формате JSON:
```shell
{
    "n": 1,
    "d": 3,
    "n1": 10.23,
    "l": 10,
    "TTL": 15
}
```
