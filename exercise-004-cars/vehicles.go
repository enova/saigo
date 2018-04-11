package main

var VehicleMap = make(map[string]*Vehicle, 0)

type Vehicle struct {
  Name  string
  Count int
}

type Vehicles struct {
  List []*Vehicle
}

type View struct {
  Username string
  Vehicles Vehicles
}

func ViewVehicles(username string) View {
  return View{
    Username: username,
    Vehicles: AsVehicles(),
  }
}

func AsVehicles() Vehicles {
  vlist := make([]*Vehicle, len(VehicleMap))
  i := 0
  for _, vehicle := range VehicleMap {
    if vehicle == nil {
      continue
    }
    vlist[i] = vehicle
    i++
  }
  return Vehicles{List: vlist}
}

func AddVehicle(model string, speed string) {
  name := model + ": " + speed
  if v, e := VehicleMap[name]; !e {
    veh := new(Vehicle)
    veh.Name = name
    veh.Count = 1
    VehicleMap[name] = veh
  } else {
    v.Count++
    VehicleMap[name] = v
  }
}