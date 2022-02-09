package database

import (
	"TaskManager/repo"
	"time"
)

//Выборка всех объектов
func (DB *Quer) GetAllObj() *[]repo.Object {
	rows, err := DB.DB.Query(`SELECT "Name","Start","Price","Id" FROM "Object" ORDER BY "Start"`)
	CheckError(err)

	var objects = []repo.Object{}
	defer rows.Close()

	for rows.Next() {
		object := repo.Object{}
		err = rows.Scan(&object.Name, &object.Start, &object.PriceSum, &object.Id)
		CheckError(err)
		objects = append(objects, object)
	}

	return &objects
}

//Выборка отдельного объекта
func (DB *Quer) GetObj(id int) []repo.Object {
	rows, err := DB.DB.Query(`SELECT "Name","Price","Id" FROM "Object"   WHERE "Object"."Id" = $1`, id)
	CheckError(err)

	var objects = []repo.Object{}
	defer rows.Close()

	for rows.Next() {
		object := repo.Object{}
		//ff := []repo.Task{}
		err = rows.Scan(&object.Name, &object.PriceSum, &object.Id)
		CheckError(err)
		objects = append(objects, object)
	}

	return objects
}

//Создать объект
func (DB *Quer) CreateObj(Object *repo.Object) error {
	t := time.Now()

	_, err := DB.DB.Exec(`INSERT INTO "Object"("Name","Price","Start") Values ($1,$2,$3)`, Object.Name, Object.PriceSum, t)
	CheckError(err)

	return err
}

//Обновить объект
func (DB *Quer) UpdateObj(Object *repo.Object, Id int) error {
	_, err := DB.DB.Exec(`UPDATE "Object" SET  "Price" = $1 WHERE "Id" = $2`, Object.PriceSum, Id)
	CheckError(err)

	return err
}

//Удалить объект
func (DB *Quer) DeleteObj(Id int) error {
	_, err := DB.DB.Exec(`DELETE FROM "Object"  WHERE "Id" = $1`, Id)
	CheckError(err)

	return err
}

/*
func (DB *Quer) GetObjWithCom(idO, idC int) []repo.Object {
	rows, err := DB.DB.Query(`SELECT "Tasks"."Name","Object"."Name","Object"."Price" FROM "Tasks"  JOIN "Object" ON "Object"."Id" = "Tasks"."FK_Object"  WHERE "Object"."Id" = $1 AND "Tasks"."Id" = $2`, idO, idC)
	CheckError(err)

	var objects = []repo.Object{}
	defer rows.Close()

	for rows.Next() {
		object := repo.Object{}
		//ff := []repo.Task{}
		err = rows.Scan(&object.Tasks.Name, &object.Name, &object.PriceSum)
		CheckError(err)
		objects = append(objects, object)
	}

	return objects
}
*/

//Выборка объектов для телеги
func (DB *Quer) GetAllObjTg(Id chan string) error {
	rows, err := DB.DB.Query(`SELECT "Name","Start","Price","Id" FROM "Object" ORDER BY "Start"`)
	CheckError(err)

	var objects = []repo.Object{}
	defer rows.Close()

	for rows.Next() {
		object := repo.Object{}
		err = rows.Scan(&object.Name, &object.Start, &object.PriceSum, &object.Id)
		CheckError(err)
		objects = append(objects, object)
	}
	Id <- objects[0].Name
	return nil
}

//Выборка отдельного объекта
func (DB *Quer) GetObjTg(name string) *repo.Object {
	rows, err := DB.DB.Query(`SELECT "Name","Price","Id" FROM "Object"   WHERE "Object"."Name" = $1`, name)
	CheckError(err)

	object := repo.Object{}
	defer rows.Close()

	for rows.Next() {

		//ff := []repo.Task{}
		err = rows.Scan(&object.Name, &object.PriceSum, &object.Id)
		CheckError(err)

	}

	return &object
}

//
//Получить логин по чатИД, используется для авторизации в тг (без пароля)
func (DB *Quer) GetLoginTg(chatId int64) *repo.User {

	rows, err := DB.DB.Query(`SELECT "Email","Password" FROM "Users"   WHERE "Telegram" = $1`, chatId)
	CheckError(err)

	user := repo.User{}
	defer rows.Close()

	for rows.Next() {

		//ff := []repo.Task{}
		err = rows.Scan(&user.Email, &user.Password)
		CheckError(err)

	}

	return &user
}

//Обновление в юзере чатИД
func (DB *Quer) UpdateUserTg(user *repo.User) error {

	_, err := DB.DB.Exec(`UPDATE "Users" SET "Telegram" = $1 WHERE "Email" = $2`, user.Telegram, user.Email)
	CheckError(err)

	return nil
}

//Удалить объект через ТГ
func (DB *Quer) DeleteObjTg(Name string) error {
	_, err := DB.DB.Exec(`DELETE FROM "Object"  WHERE "Name" = $1`, Name)
	CheckError(err)

	return err
}

//Получить всех юзеров
func (DB *Quer) GetAllUsers() *[]repo.User {
	rows, err := DB.DB.Query(`SELECT "Name","Email","Id" FROM "Users"`)
	CheckError(err)

	var users = []repo.User{}
	defer rows.Close()

	for rows.Next() {
		user := repo.User{}
		err = rows.Scan(&user.Name, &user.Email, &user.Id)
		CheckError(err)
		users = append(users, user)
	}

	return &users
}

//Создать пользователя
func (DB *Quer) CreateUser(User *repo.User) error {

	_, err := DB.DB.Exec(`INSERT INTO "Users"("Name","Email","Password") Values ($1,$2,$3)`, User.Name, User.Email, User.Password)
	CheckError(err)

	return err
}

//Получить пользователя по email, используется для авторизации через АПИ
func (DB *Quer) GetUser(email string) *repo.User {
	rows, err := DB.DB.Query(`SELECT "Name","Email","Password" FROM "Users"   WHERE "Email" = $1`, email)
	CheckError(err)

	user := repo.User{}
	defer rows.Close()

	for rows.Next() {

		//ff := []repo.Task{}
		err = rows.Scan(&user.Name, &user.Email, &user.Password)
		CheckError(err)

	}

	return &user
}
