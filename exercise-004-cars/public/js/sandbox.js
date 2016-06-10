var VehicleSandbox = (function() {

  function VehicleSandbox(vehicles) {
    this.vehicles   = [];
    this.speeds     = [];
    this.positions  = [];

    // Add vehicles if vehicles given
    if (vehicles && vehicles.length > 0) {
      this.addVehicles(vehicles);
    }
  }

  //  Update canvas with background and vehicle positions
  VehicleSandbox.prototype.animate = function() {
    this.draw_background();
    this.draw_vehicles();
  };

  //  Go through each vehicle in vehicleList
  //  Determine the vehicle name+speed+count
  //  Add a new vehicle per type
  VehicleSandbox.prototype.addVehicles = function(vehicleList) {
    for(var x=0; x < vehicleList.length; x++) {
      var vehicle = vehicleList[x];

      var _name   = vehicle["Name"].replace(/\ /, '').split(':')[0];
      var _speed  = vehicle["Name"].replace(/\ /, '').split(':')[1];
      var _count  = vehicle["Count"];

      for(var y=0; y < _count; y++) {
        this.vehicles.push(_name);
        this.speeds.push(_speed);
        this.positions.push(0);
      }
    }
  };

  //  Determine how far to advance each vehicle based on type and speed
  VehicleSandbox.prototype.advance = function(vehicle, speed) {
    var base = 1;

    // Base Velocity
    switch (vehicle) {
      case "jeep":
        base = 50;
        break;
      case "bike":
        base = 15;
        break;
      case "boat":
        base = 30;
        break;
    }

    // Adjusted Velocity
    switch (speed) {
      case "slow":
        return base / 2;
      case "fast":
        return base * (0.1 + 0.9 * Math.random());
      case "rage":
        return base * 3 * (Math.random() * Math.random());
    }

    // Should Never Happen
    return 1;
  };

  //  Determine vehicle color
  VehicleSandbox.prototype.vehicle_color = function(vehicle) {
    switch (vehicle) {
      case "jeep":
        return '#FF0000';
      case "bike":
        return '#00FF00';
      case "boat":
        return '#0000FF';
    }

    return '#333333';
  };

  //  Determine vehicle size
  VehicleSandbox.prototype.vehicle_radius = function(vehicle) {
    switch (vehicle) {
      case "jeep":
        return 10;
      case "bike":
        return 3;
      case "boat":
        return 15;
    }
    return 1;
  };

  //  Determine Y limits
  VehicleSandbox.prototype.limit_vertical = function(y) {
    if (y > 300) {
      return 300 + (y - 300) * 0.2;
    }
    if (y < 100) {
      return 100 + (y - 100) * 0.2;
    }
    return y;
  };

  //  Determine updated vehicle position and render on canvas
  VehicleSandbox.prototype.draw_vehicles = function() {
    var canvas = document.getElementById("canvas");
    if (!canvas.getContext) { return; }

    var ctx = canvas.getContext("2d");

    for (var i = 0; i < this.vehicles.length; i++) {
      angle = (this.positions[i] * 2.0 * Math.PI) / 10000.0;

      var r = 200 - 10 * i;
      if (r <= 40) {
        r = 40;
      }

      var x = 300 + r * Math.cos(angle);
      var y = 200 + r * Math.sin(angle);

      y = this.limit_vertical(y);

      var vehicle = this.vehicles[i];
      var speed = this.speeds[i];

      var increment = this.advance(vehicle, speed);
      var radius = this.vehicle_radius(vehicle);
      var color = this.vehicle_color(vehicle);

      ctx.fillStyle = color;
      ctx.beginPath();
      ctx.arc(x, y, radius, 0, Math.PI*2, true);
      ctx.fill();

      this.positions[i] += increment;
    }
  };

  //  Setup background
  VehicleSandbox.prototype.draw_background = function() {
    var canvas = document.getElementById("canvas");
    if (!canvas.getContext) { return; }

    var ctx = canvas.getContext("2d");

    ctx.fillStyle = '#110011';
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    ctx.strokeStyle = '#333333';
    ctx.lineWidth = 1;

    for (var i = 0; i <= canvas.width; i += 50) {
      ctx.beginPath();
      ctx.moveTo(i, 0);
      ctx.lineTo(i, canvas.height);
      ctx.stroke();
    }

    for (var i = 0; i <= canvas.height; i += 50) {
      ctx.beginPath();
      ctx.moveTo(0, i);
      ctx.lineTo(canvas.width, i);
      ctx.stroke();
    }
  };

  return VehicleSandbox;
})();
