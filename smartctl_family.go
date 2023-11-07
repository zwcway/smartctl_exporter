package main

import "sync"

var (
	families = []func(family, serial, model string) SMARTDeviceFamily{
		SeagateDetector,
	}
	lock sync.Mutex
)

func RegisterFamily(defect func(family, serial, model string) SMARTDeviceFamily) {
	lock.Lock()
	families = append(families, defect)
	lock.Unlock()
}

func smartctl_family(family, serial, model string) SMARTDeviceFamily {
	for _, f := range families {
		fm := f(family, serial, model)
		if fm != nil {
			return fm
		}
	}
	return nil
}
