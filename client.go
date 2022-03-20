package main
import "fmt"
import "DesignPattern/Vehicle"

func main(){
   carbuilder:=Vehicle.VehicleBuilder("car")
   director:=Vehicle.NewDirector(carbuilder)
   carmodel:=director.BuildVehicle()
   fmt.Println("Car window is :",carmodel.Window)
   fmt.Println("Car wheels is :",carmodel.Wheels)
   fmt.Println("Car steering is :",carmodel.Steering)
}