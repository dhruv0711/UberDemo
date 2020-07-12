package uber

import (
	"math"
	"strconv"
	"strings"
)

func calculateDistance(cab string, customer string) (float64, error) {
	cabArr := strings.Split(cab, ",")
	cablatstring, cablongstring := cabArr[0], cabArr[1]

	cablatitude, err := strconv.ParseFloat(cablatstring, 64)
	if err != nil {
		return math.MaxFloat64, err
	}
	cablong := strings.TrimSpace(cablongstring)
	cablongitude, err := strconv.ParseFloat(cablong, 64)
	if err != nil {
		return 0, err
	}

	customerArr := strings.Split(customer, ",")
	customerlatstring, customerlongstring := customerArr[0], customerArr[1]

	customerlatitude, err := strconv.ParseFloat(customerlatstring, 64)
	if err != nil {
		return 0, err
	}
	customerlong := strings.TrimSpace(customerlongstring)
	customerlongitude, err := strconv.ParseFloat(customerlong, 64)
	if err != nil {
		return 0, err
	}

	LatSquare := (cablatitude - customerlatitude) * (cablatitude - customerlatitude)
	LongSquare := (cablongitude - customerlongitude) * (cablongitude - customerlongitude)

	total := LongSquare + LatSquare
	final := math.Sqrt(total)
	return final, nil

}
