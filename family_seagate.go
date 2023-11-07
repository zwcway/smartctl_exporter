package main

import "strings"

type Seagate struct {
}

func (s *Seagate) DeviceAttributeRaw(id string, name string, val float64) float64 {
	switch id {
	case "1", "7":
		v := int64(val)
		v >>= 47
		v &= 0b111111
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
