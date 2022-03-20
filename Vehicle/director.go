package Vehicle

type Director struct {
	vehicle Vehicle
}

func NewDirector(v Vehicle) *Director {
	return &Director{
		vehicle: v,
	}
}

func (d *Director) SetVehicle(v Vehicle) {
	d.vehicle = v
}

func (d *Director) BuildVehicle() model {
	d.vehicle.setWheels()
	d.vehicle.setWindow()
	d.vehicle.setSteering()
	return d.vehicle.getVehicle()
}
