# **Вычислятор**

Приложение принимает на вход два типа инструкций:

1. `calc`: вычислить арифметическую операцию (умножение, сложение, вычитание) над двумя сущностями и сохранить результат в переменную. Сущность может быть либо литерал типа int64, либо имя переменной. В одну и ту же переменную записывать результат можно только один раз.
2. `print`: вывести значение указанной переменной, например, `x = 5`
Считаем, что каждая арифметическая операция из инструкций вычисляется долго, например, 50ms. Требуется наиболее быстро вычислять результаты списка инструкций, т.е. выводить требуемые переменные.

**Примеры запросов и ответов**

input
```json
[
  { "type": "calc", "op": "+", "var": "x", "left": 1,  "right": 2 },
  { "type": "print", "var": "x" }
]
```

output
```json
{ "items": [
    { "var": "x","value": 3}
  ]
}
```

input
```json
[
  { "type": "calc", "op": "+", "var": "x",   "left": 10,  "right": 2  },
  { "type": "print",             "var": "x"                     },
  { "type": "calc", "op": "-", "var": "y",   "left": "x",  "right": 3  },
  { "type": "calc", "op": "*", "var": "z",   "left": "x",  "right": "y" },
  { "type": "print",             "var": "w"                     },
  { "type": "calc", "op": "*", "var": "w",   "left": "z",  "right": 0  }
]
```

output
```json
{ "items": [
    {"var": "x","value": 12},
    {"var": "w","value": 0}
  ]
}
```

input
```json
[
  { "type": "calc", "op": "+", "var": "x",        "left": 10,   "right": 2    },
  { "type": "calc", "op": "*", "var": "y",        "left": "x",  "right": 5    },
  { "type": "calc", "op": "-", "var": "q",        "left": "y",  "right": 20   },
  { "type": "calc", "op": "+", "var": "unusedA",  "left": "y",  "right": 100  },
  { "type": "calc", "op": "*", "var": "unusedB",  "left": "unusedA", "right": 2 },
  { "type": "print",             "var": "q"                        },
  { "type": "calc", "op": "-", "var": "z",        "left": "x",  "right": 15   },
  { "type": "print",             "var": "z"                        },
  { "type": "calc", "op": "+", "var": "ignoreC",  "left": "z",  "right": "y"  },
  { "type": "print",             "var": "x"                        }
]
```

output
```json
{ "items": [
    {"var": "q","value": 40},
    {"var": "z","value": -3},
    {"var": "x","value": 12}
  ]
}
```

# План реализации

## Ядро

1. Создаём репозиторий на [Github](https://github.com/)
2. Копируем ссылку на репозиторий, вводим в терминал команду `git clone <link>`, скачиваем репозиторий локально
3. Переходим в корень проекта и вводим команду `go mod init github.com/<username>/<project_name>` в терминале, чтобы проект стал считаться модулем
4. По пути <project_name> /internal/model/ создаём файл instruction.go, в котором описываем формат входящих инструкций и формат ответа
5. По пути <project_name> /internal/core/ создаём файл parser.go, который будет переводить инструкции из локального json-файла в инструкции, который будет понимать язык Go
6. По пути <project_name> /internal/core/ создаём файл executor.go, где содержится вся бизнес-логика проекта
7. В корне проекта создаём файл main.local.go и проверяем работоспособность кода

## Сервера

1. По пути <project_name>/cmd/server/ создаём файл main.go, который будет содержать в себе код для запуска http и grpc серверов
2. По пути <project_name>/internal/api/ создаём файл http.go, в котором описываем http-сервер
3. По пути <project_name>/internal/api/ создаём файл grpc.go, в котором описываем grpc-сервер
4. По пути <project_name>/proto/ создаём файл <project_name>.proto, в котором описываем интерфейс grpc-сервера
5. В корне проекта вводим команды `go get google.golang.org/grpc` и `go get google.golang.org/protobuf`
6. Из корня проекта выполняем команду `protoc --go_out=. --go-grpc_out=. proto/<project_name>.proto` и генерируем Go-код для реализации интерфейса сервера
7. Добавляем grpc-сервер в main.go и запускаем сервера командой `go run cmd/server/main.go`

## Документация

1. Вводим команду `go install [github.com/swaggo/swag/cmd/swag@latest](mailto:github.com/swaggo/swag/cmd/swag@latest)` в терминале
2. По пути <project_name>/internal/api/http.go добавляем к основной функции комментарии, по которым будет сгенерирована документация
3. Вводим команду `swag init -g internal/api/http.go` в терминале и генерируем документацию
4. Переходим на http://localhost:8080/swagger/index.html и убеждаемся, что всё работает

## Тесты

1. По пути <project_name>/internal/core/ создаём файл executor_test.go и пишем тесты для нашего проекта
2. Вводим команду `go test ./internal/core -cover` в терминал и убеждаемся, что покрытие тестами составляет более 20 процентов

## Docker

1. В корне создаём и описываем Dockerfile для запуска приложения в контейнере
2. Там же создаём файл docker-compose.yml для сохранения возможности масштабирования и запуска проекта командой `docker compose up`
