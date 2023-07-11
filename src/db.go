package api

import "fmt"

type row struct {
	Key   string
	Value string
}

type Table struct {
	Name string
	Data []row
}

func InitDB() *Table {
	return &Table{}
}

func (t *Table) AddKey(key string, value string) error {
	if link, _, err := t.GetValue(key); err == nil {
		return fmt.Errorf("ссылка уже существует: %s ", link)
	} else {
		t.Data = append(t.Data, row{
			key,
			value,
		})
		return nil
	}
}

func (t *Table) DeleteKey(key string) error {
	if _, i, err := t.GetValue(key); err == nil {
		t.Data = append(t.Data[:i], t.Data[i+1:]...)
		return nil
	} else {
		return err
	}

}

func (t *Table) GetValue(key string) (string, int, error) {
	for i, datum := range t.Data {
		if datum.Key == key {
			return datum.Value, i, nil
		}
	}
	return "", 0, fmt.Errorf("ссылки не существует")
}
