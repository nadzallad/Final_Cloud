package entity

type City struct {
	CityID    int     `gorm:"column:city_id;primaryKey"`
	Name      string  `gorm:"column:name"`
	Latitude  float64 `gorm:"column:latitude"`
	Longitude float64 `gorm:"column:longitude"`
}

func (City) TableName() string {
	return "cities"
}