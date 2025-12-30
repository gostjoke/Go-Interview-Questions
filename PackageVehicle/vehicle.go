package packagevehicle

// Vehicle 結構體
type Vehicle struct {
	Make  string
	Speed float32
	owner string // private field
}

func (v *Vehicle) Accelerate() {
	v.Speed += 10
}

func (v *Vehicle) Decelerate() {
	if v.Speed >= 10 {
		v.Speed -= 10
	}
}

// 構造函數
func NewVehicle(make string, speed float32, owner string) *Vehicle {
	return &Vehicle{
		Make:  make,
		Speed: speed,
		owner: owner,
	}
}
