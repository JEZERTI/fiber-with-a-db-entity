package services

import (
	"errors"
	"fitness_club/models"
	"time"
)

var clients = []models.Client{
	{ID: 1, Name: "Айжан Сагындыкова", Phone: "87071234501", Email: "a.sagyndikova@mail.kz", BirthDate: time.Date(1996, 3, 14, 0, 0, 0, 0, time.UTC)},
	{ID: 2, Name: "Алина Кузнецова", Phone: "87071234502", Email: "a.kuznetsova@mail.ru", BirthDate: time.Date(1995, 7, 22, 0, 0, 0, 0, time.UTC)},
	{ID: 3, Name: "Мадина Абдрахманова", Phone: "87071234503", Email: "m.abdrakhmanova@mail.kz", BirthDate: time.Date(1998, 1, 9, 0, 0, 0, 0, time.UTC)},
	{ID: 4, Name: "Дарья Иванова", Phone: "87071234504", Email: "d.ivanova@mail.ru", BirthDate: time.Date(1994, 11, 30, 0, 0, 0, 0, time.UTC)},
	{ID: 5, Name: "Айгерим Нуртаева", Phone: "87071234505", Email: "a.nurtayeva@mail.kz", BirthDate: time.Date(1997, 5, 18, 0, 0, 0, 0, time.UTC)},
	{ID: 6, Name: "Екатерина Смирнова", Phone: "87071234506", Email: "e.smirnova@mail.ru", BirthDate: time.Date(1993, 2, 27, 0, 0, 0, 0, time.UTC)},
	{ID: 7, Name: "Аружан Тлеубаева", Phone: "87071234507", Email: "a.tleubayeva@mail.kz", BirthDate: time.Date(1999, 9, 4, 0, 0, 0, 0, time.UTC)},
	{ID: 8, Name: "Полина Морозова", Phone: "87071234508", Email: "p.morozova@mail.ru", BirthDate: time.Date(1996, 12, 12, 0, 0, 0, 0, time.UTC)},
	{ID: 9, Name: "Динара Касымова", Phone: "87071234509", Email: "d.kassymova@mail.kz", BirthDate: time.Date(1995, 6, 21, 0, 0, 0, 0, time.UTC)},
	{ID: 10, Name: "Мария Петрова", Phone: "87071234510", Email: "m.petrova@mail.ru", BirthDate: time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC)},
	{ID: 11, Name: "Жансая Алимбекова", Phone: "87071234511", Email: "zh.alimbekova@mail.kz", BirthDate: time.Date(2000, 1, 3, 0, 0, 0, 0, time.UTC)},
	{ID: 12, Name: "Анна Волкова", Phone: "87071234512", Email: "a.volkova@mail.ru", BirthDate: time.Date(1997, 4, 19, 0, 0, 0, 0, time.UTC)},
	{ID: 13, Name: "Салтанат Омарова", Phone: "87071234513", Email: "s.omarova@mail.kz", BirthDate: time.Date(1994, 10, 8, 0, 0, 0, 0, time.UTC)},
	{ID: 14, Name: "Ольга Федорова", Phone: "87071234514", Email: "o.fedorova@mail.ru", BirthDate: time.Date(1991, 6, 25, 0, 0, 0, 0, time.UTC)},
	{ID: 15, Name: "Меруерт Бекенова", Phone: "87071234515", Email: "m.bekenova@mail.kz", BirthDate: time.Date(1998, 2, 17, 0, 0, 0, 0, time.UTC)},
	{ID: 16, Name: "Наталья Орлова", Phone: "87071234516", Email: "n.orlova@mail.ru", BirthDate: time.Date(1993, 9, 29, 0, 0, 0, 0, time.UTC)},
	{ID: 17, Name: "Айдана Муратова", Phone: "87071234517", Email: "a.muratova@mail.kz", BirthDate: time.Date(1999, 12, 6, 0, 0, 0, 0, time.UTC)},
	{ID: 18, Name: "Юлия Соколова", Phone: "87071234518", Email: "y.sokolova@mail.ru", BirthDate: time.Date(1995, 5, 11, 0, 0, 0, 0, time.UTC)},
	{ID: 19, Name: "Гульмира Ахметова", Phone: "87071234519", Email: "g.akhmetova@mail.kz", BirthDate: time.Date(1996, 7, 2, 0, 0, 0, 0, time.UTC)},
	{ID: 20, Name: "Ксения Лебедева", Phone: "87071234520", Email: "k.lebedeva@mail.ru", BirthDate: time.Date(1994, 3, 23, 0, 0, 0, 0, time.UTC)},
}

func UpdateClient(id int, data map[string]interface{}) (models.Client, error) {
	for i, client := range clients {
		if client.ID == id {

			if name, ok := data["name"].(string); ok {
				client.Name = name
			}
			if phone, ok := data["phone"].(string); ok {
				client.Phone = phone
			}
			if email, ok := data["email"].(string); ok {
				client.Email = email
			}
			if birhDate, ok := data["birth_date"].(time.Time); ok {
				client.BirthDate = birhDate
			}
			clients[i] = client
			return client, nil
		}
	}
	return models.Client{}, errors.New("client not found")
}

func GetClients() []models.Client {
	return clients
}

func GetClientByID(id int) (*models.Client, error) {
	for _, client := range clients {
		if client.ID == id {
			return &client, nil
		}
	}
	return nil, errors.New("client not found")
}

func CreateClient(client models.Client) models.Client {
	client.ID = len(clients) + 1
	clients = append(clients, client)
	return client
}

func DeleteClient(id int) error {
	for i, client := range clients {
		if client.ID == id {
			clients = append(clients[:1], clients[i+1:]...)
			return nil
		}
	}
	return errors.New("client not found")
}

func UpdateClient2(id int, updated models.Client) (models.Client, error) {
	for i, client := range clients {
		if client.ID == id {
			updated.ID = id
			clients[i] = updated
			return updated, nil
		}
	}
	return models.Client{}, errors.New("client not found")
}
