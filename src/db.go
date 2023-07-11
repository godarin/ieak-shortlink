package api

import (
	"database/sql"
	"fmt"
)

type row struct {
	Key   string
	Value string
}
type Table struct {
	Name string
	Data []row
}

func InitDB() *Table {
	loadedData, err := loadFromSQLite("data.db")
	if err != nil {
		fmt.Println("Ошибка при загрузке данных:", err)
		return &Table{}
	}
	return &Table{
		Data: loadedData,
	}
}

func (t *Table) AddKey(key string, value string) error {
	if link, _, err := t.GetValue(key); err == nil {
		return fmt.Errorf("ссылка уже существует: %s ", link)
	} else {
		t.Data = append(t.Data, row{
			key,
			value,
		})
		err := saveToSQLite("data.db", t.Data)
		if err != nil {
			return nil
		}
		return nil
	}
}

func (t *Table) DeleteKey(key string) error {
	if _, i, err := t.GetValue(key); err == nil {
		t.Data = append(t.Data[:i], t.Data[i+1:]...)
		err := saveToSQLite("data.db", t.Data)
		return err

	} else {
		return err
	}

}

func (t *Table) GetValue(key string) (string, int, error) {
	loadedData, err := loadFromSQLite("data.db")
	if err != nil {
		fmt.Println("Ошибка при загрузке данных:", err)
	}
	t.Data = loadedData
	for i, datum := range t.Data {
		if datum.Key == key {
			return datum.Value, i, nil
		}
	}
	return "", 0, fmt.Errorf("ссылки не существует")
}

func saveToSQLite(databaseName string, data []row) error {
	db, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		return err
	}
	defer db.Close()

	// Создаем таблицу, если она не существует
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS rows (key TEXT PRIMARY KEY, value TEXT)")
	if err != nil {
		return err
	}

	// Очищаем таблицу перед сохранением новых данных
	_, err = db.Exec("DELETE FROM rows")
	if err != nil {
		return err
	}

	// Вставляем каждую запись в таблицу
	stmt, err := db.Prepare("INSERT INTO rows(key, value) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, r := range data {
		_, err := stmt.Exec(r.Key, r.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadFromSQLite(databaseName string) ([]row, error) {
	db, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT key, value FROM rows")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []row
	for rows.Next() {
		var r row
		err := rows.Scan(&r.Key, &r.Value)
		if err != nil {
			return nil, err
		}
		data = append(data, r)
	}

	return data, nil
}
