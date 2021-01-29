package main

func defangIPaddr(address string) string {
	new := ""
	for _, i2 := range address {
		if i2 == 46 {
			new += "[.]"
		} else {
			new += string(i2)
		}
	}
	return new
}