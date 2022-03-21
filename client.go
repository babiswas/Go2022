package main
import "fmt"
import "DesignPattern/Vehicle"

/*Used if multiple parameter is there in the constructor*/

func main(){
   carbuilder:=Vehicle.VehicleBuilder("car")
   director:=Vehicle.NewDirector(carbuilder)
   carmodel:=director.BuildVehicle()
   truckbuilder:=Vehicle.VehicleBuilder("truck")
   director1:=Vehicle.NewDirector(truckbuilder)
   truckmodel:=director1.BuildVehicle()
   fmt.Println("Car window is :",carmodel.Window)
   fmt.Println("Car wheels is :",carmodel.Wheels)
   fmt.Println("Car steering is :",carmodel.Steering)
   fmt.Println("Truck window is :",truckmodel.Window)
   fmt.Println("Truck wheels is :",truckmodel.Wheels)
   fmt.Println("Truck steering is :",truckmodel.Steering)
}