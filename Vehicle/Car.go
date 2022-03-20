package Vehicle

type Car struct{
  Wheels int
  Window string
  Steering string
}

func carBuilder()*Car{
   return &Car{}
}

func (c *Car)setWheels(){
   c.Wheels=4
}

func (c *Car)setWindow(){
   c.Window="Window Set"
}

func (c *Car)setSteering(){
   c.Steering="Steering set"
}

func (c *Car)getVehicle()model{
  return model{Wheels:c.Wheels,Window:c.Window,Steering:c.Steering}
}

