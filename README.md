# ObjectManager
 
### Управление объектами.
 
 ##### Использование fiber, jwt, telegram, postgre
 ##### Просмотр, удаление, создание объекта со следующими полями: Дата (создания), Название объекта, Стоимость.
 ##### Работа по API через Vue 3 (https://github.com/Vladicorn/ObjectManagerFront), также работа через телеграмм.
 
 ##### По HTTP авторизация с использованием JWT, в телеграмме проверка только на логин.
 
### API:
- Post("/api/register") - регистрация
    {
        "Name" : "Проверка",
        "Email" : "10000@mail.ru",
        "Password": "1234",
        "PasswordConfirm": "1234"
    }

- Post("/api/login") - логин
    {
        "Email" : "10000@mail.ru",
        "Password": "1234",
    }

- Get("/api/logout") - разлогин

- Get("/api/users") - список всех пользователей

- Get("/api/objects") - список всех объектов
- Get("/api/object/:id") - список отдельного объекта
- Post("/api/object") - создание объекта
    {
        "Name" : "Проверка",
        "PriceSum": "5",
    }

- Put("/api/object/:id") - обновление объекта
    {
        "Name" : "Проверка",
        "PriceSum": "5",
    }

- Delete("/api/object/:id") - удаление объекта

### Телеграмм команды
- /create 123@ya.ru - авторизация в тг
- /all - список всех объектов
- /add test 123 - добавить объект test - название объекта, 123 - стоимость
- /del test - удалить объект test - название объекта

