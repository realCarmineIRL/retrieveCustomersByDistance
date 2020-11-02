package main

import (
	"errors"
	"math"
	"strconv"
)

const pi float64 = 3.141592653589793
const earthRadius float64 = 6371

// Intercom latitude and Longitude
const intercomlat string = "53.339428"
const intercomLon string = "-6.257664"

// CalcDistance get distance between two coordinates
func CalcDistance(lat string, lon string) (float64, error) {

	radlat1, errRadlat1 := convertDegToRad(lat)
	radlat2, errRadlat2 := convertDegToRad(intercomlat)
	radtheta, errRadtheta := calcRadTheta(lon, intercomLon)

	if errRadlat1 != nil || errRadlat2 != nil || errRadtheta != nil {
		return 0, errors.New("CalcDistance found an error")
	}

	dist := math.Acos(math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta))
	dist = dist * earthRadius

	return dist, nil
}

func convertDegToRad(degrees string) (float64, error) {
	deg, err := strconv.ParseFloat(degrees, 64)
	rad := pi * deg / 180

	if err != nil {
		return 0, errors.New("convertDegToRad found an error parsing values")
	}

	return rad, nil
}

func calcRadTheta(lon1 string, lon2 string) (float64, error) {
	l1, errLon1 := strconv.ParseFloat(lon1, 64)
	l2, errLon2 := strconv.ParseFloat(lon2, 64)

	if errLon1 != nil || errLon2 != nil {
		return 0, errors.New("calcRadTheta found an error parsing values")
	}

	theta := l1 - l2

	radTheta := pi * theta / 180

	return radTheta, nil
}
