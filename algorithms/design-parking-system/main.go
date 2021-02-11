package main

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big, medium, small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 && this.big > 1 {
		this.big--
		return true
	} else if carType == 2 && this.medium > 1 {
		this.medium--
		return true
	} else if carType == 3 && this.small > 1 {
		this.small--
		return true
	}

	return false
}
