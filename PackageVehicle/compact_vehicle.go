package packagevehicle

import "fmt"

type CompactVehicle struct {
	Vehicle
	Color     string
	sitNumber int
}

func NewCompactVehicle(make string, speed float32, owner string, color string, sitNumber int) (*CompactVehicle, error) {
	// 驗證 sitNumber 範圍
	if sitNumber < 0 || sitNumber > 4 {
		return nil, fmt.Errorf("sitNumber must be between 0 and 4")
	}

	return &CompactVehicle{
		Vehicle:   *NewVehicle(make, speed, owner),
		Color:     color,
		sitNumber: sitNumber,
	}, nil
}

func (cv *CompactVehicle) GetSitNumber() int {
	return cv.sitNumber
}

func (cv *CompactVehicle) Paint(newColor string) {
	cv.Color = newColor
}
