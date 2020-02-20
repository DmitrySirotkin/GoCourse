package main

import "fmt"

func main() {

	type passengerСar struct {
		mark        string
		year        int
		bootVolume  int
		startEngine bool
		openWindows bool
		bootLoad    int
	}

	var firstPassCar = passengerСar{mark: "Lada", year: 2020, bootVolume: 10, startEngine: true, openWindows: true, bootLoad: 10}
	var secondPassCar = passengerСar{mark: "Skoda", year: 2010, bootVolume: 20, startEngine: true, openWindows: false, bootLoad: 20}
	var thridPassCar = passengerСar{mark: "Opel", year: 2000, bootVolume: 30, startEngine: false, openWindows: false, bootLoad: 30}
	firstPassCar.bootLoad = 0
	secondPassCar.startEngine = false
	thridPassCar.openWindows = true

	fmt.Println(firstPassCar, secondPassCar, thridPassCar)

	type truck struct {
		mark        string
		year        int
		truckVolume int
		startEngine bool
		openWindows bool
		truckLoad   int
	}

	var firstTruckCar = truck{mark: "Yaz", year: 2020, truckVolume: 100, startEngine: true, openWindows: true, truckLoad: 100}
	var secondTruckCar = truck{mark: "Man", year: 2010, truckVolume: 200, startEngine: false, openWindows: true, truckLoad: 100}
	var thridTruckCar = truck{mark: "Yaz", year: 2015, truckVolume: 300, startEngine: false, openWindows: true, truckLoad: 100}
	firstTruckCar.truckLoad = 0
	secondTruckCar.truckLoad = 0
	thridTruckCar.truckLoad = 0

	fmt.Println(firstTruckCar, secondTruckCar, thridTruckCar)
}
