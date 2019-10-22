package plants

import (
	"database/sql"
	"time"
)

type Plant struct {
	Id                uint    `json:"id"`
	SoilMoistureLevel float32 `json:"soilMoistureLevel"`
	SoilDataTimestamp string  `json:"soilDataTimestamp"`
}

type GreenHouse struct {
	Db *sql.DB
}

func NewGreenHouse(db *sql.DB) *GreenHouse {
	return &GreenHouse{Db: db}
}

func (s *GreenHouse) ListCriticalPlants() ([]*Plant, error) {
	rows, err := s.Db.Query("SELECT * FROM plants WHERE soilmoisturelevel<=0.2")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Plant
	for rows.Next() {
		var c Plant
		if err := rows.Scan(&c.Id, &c.SoilMoistureLevel, &c.SoilDataTimestamp); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Plant, 0)
	}
	return res, nil
}

func (s *GreenHouse) AddMoistureLevel(id uint, soilMoistureLevel float32, soilDataTimestamp string) error {
	curentTime := time.Now().Format(time.RFC3339)
	_, err1 := s.Db.Exec("DELETE FROM plants WHERE id=$1", id)
	if err1 != nil {
		return err1
	}
	_, err2 := s.Db.Exec("INSERT INTO plants VALUES ($1, $2, $3)", id, soilMoistureLevel, curentTime)
	return err2
}
