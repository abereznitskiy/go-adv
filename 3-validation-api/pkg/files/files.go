package files

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonDb struct {
	filename string
	data     map[string]string
}

func NewJsonDb(name string) *JsonDb {
	db := &JsonDb{
		filename: name,
		data:     make(map[string]string),
	}

	err := db.load()
	if err != nil {
		fmt.Println("Error data loading:", err)
	}

	return db
}

func (db *JsonDb) Set(key, value string) error {
	db.data[key] = value
	return db.save()
}

func (db *JsonDb) Get(key string) (string, bool) {
	value, exists := db.data[key]
	return value, exists
}

func (db *JsonDb) Delete(key string) error {
	delete(db.data, key)
	return db.save()
}

func (db *JsonDb) load() error {
	file, err := os.ReadFile(db.filename)
	if err != nil {
		if os.IsNotExist(err) {
			db.data = make(map[string]string)
			return nil
		}
		return err
	}

	if len(file) == 0 {
		db.data = make(map[string]string)
		return nil
	}

	return json.Unmarshal(file, &db.data)
}

func (db *JsonDb) save() error {
	data, err := json.Marshal(db.data)
	if err != nil {
		return err
	}

	return os.WriteFile(db.filename, data, 0644)
}
