# Inter
Проект представляет собой реализацию сервиса по размещению задачей, подсчету значения арифметической прогрессии в них, и получения сведений о всех поставленных "задачах"

# Инструкция по запуску
```shell
go run cmd/main.go 5
#Запуск сервиса с установлением количества единовремено выполняемых задач 5
```
# Запросы обрабатываемые сервисом
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