package Vehicle

type Truck struct{
  Wheels int
  Window string
  Steering string
}

func truckBuilder()*Truck{
   return &Truck{}
}

func (t *Truck)setWheels(){
   t.Wheels=4
}

func (t *Truck)setWindow(){
   t.Window="Window Set"
}

func (t *Truck)setSteering(){
   t.Steering="Steering"
}

func (t *Truck)getVehicle()model{
  return model{Wheels:t.Wheels,Window:t.Window,Steering:t.Steering}
}

