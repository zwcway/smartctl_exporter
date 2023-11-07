package main

import "strings"

type Seagate struct {
}

func (s *Seagate) DeviceAttributeRaw(id string, name string, val float64) float64 {
	switch id {
	case "1", "7":
		v := int64(val)
		// raw48:54
		// 5 is 47-40bit
		// 4 is 39-32bit
		// 54 is Little-Endian
		v >>= 32
		return float64(v)
	}

	return val
}

func SeagateDetector(family, serial, model string) SMARTDeviceFamily {
	if strings.HasPrefix(family, "Seagate") {
		return &Seagate{}
	}
	return nil
}
