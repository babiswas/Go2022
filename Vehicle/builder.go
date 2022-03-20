package Vehicle

type Vehicle interface{
   setWindow()
   setWheels()
   setSteering()
   getVehicle() model
}

func VehicleBuilder(modelType string)Vehicle{
    if modelType=="car"{
      return &Car{}
    }else if modelType=="truck"{
      return &Truck{}
    }else{
      return nil
    }
}