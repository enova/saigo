package main

var vmap = make(map[string]*Vehicle, 0)

// Vehicle holds a vehicle name and count
type Vehicle struct {
	Name  string
	Count int
}

// Vehicles holds a slice of Vehicle structs
type Vehicles struct {
	List []*Vehicle
}

// View holds a Vehicles and a Username for rendering
type View struct {
	Username string
	Vehicles Vehicles
}

// ViewVehicles generates a View for a given username
func ViewVehicles(username string) View {
	return View{
		Username: username,
		Vehicles: AsVehicles(),
	}
}

// AsVehicles returns the current map as Vehicles
func AsVehicles() Vehicles {
	vlist := make([]*Vehicle, len(vmap))
	i := 0
	for _, vehicle := range vmap {
		if vehicle == nil {
			continue
		}
		vlist[i] = vehicle
		i++
	}
	return Vehicles{List: vlist}
}

// AddVehicle adds or updates the given model and speed in the map
func AddVehicle(model string, speed string) {
	name := model + ": " + speed
	if v, e := vmap[name]; !e {
		veh := new(Vehicle)
		veh.Name = name
		veh.Count = 1
		vmap[name] = veh
	} else {
		v.Count++
		vmap[name] = v
	}
}
