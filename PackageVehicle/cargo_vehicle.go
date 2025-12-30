package packagevehicle

type CargoVehicle struct {
	Vehicle
	Load float32
}

func NewCargoVehicle(make string, speed float32, owner string, load float32) *CargoVehicle {
	return &CargoVehicle{
		Vehicle: *NewVehicle(make, speed, owner),
		Load:    load,
	}
}

func (cv *CargoVehicle) LoadCargo(weight float32) {
	cv.Load += weight
}
func (cv *CargoVehicle) UnloadCargo(weight float32) {

	if cv.Load < weight {
		cv.Load = 0
	} else {
		cv.Load -= weight
	}
}
