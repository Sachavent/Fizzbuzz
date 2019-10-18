package services

import (
	"strconv"
	"strings"

	"github.com/Fizzbuzz/models"
)

func isMultiple(nbIteration int, multiple int) bool {
	if nbIteration%multiple == 0 {
		return true
	}

	return false
}

func applyRules(nbIteration int, int1 int, int2 int, str1 string, str2 string) string {
	// Check multiple int1 and int2
	if isMultiple(nbIteration, int1) && isMultiple(nbIteration, int2) {
		return str1 + str2
	}

	// Multiple int1
	if isMultiple(nbIteration, int1) {
		return str1
	}

	// Multiple int2
	if isMultiple(nbIteration, int2) {
		return str2
	}

	return ""
}

func GetFizzbuzzResult(fizzBuzzParams models.FizzBuzzParams, mapNumberRequest map[models.FizzBuzzParams]int) string {

	var result []string

	// Iterate till limit has been not reached
	// Careful: Fizzbuzz start at 1 and not 0
	for i := 1; i < fizzBuzzParams.Limit+1; i++ {

		elt := applyRules(i, fizzBuzzParams.Int1, fizzBuzzParams.Int2, fizzBuzzParams.Str1, fizzBuzzParams.Str2)

		if elt != "" {
			result = append(result, elt)
		} else {
			// Default : return iteration
			result = append(result, strconv.Itoa(i))
		}

	}

	// Increment number of request with this params
	mapNumberRequest[fizzBuzzParams] = mapNumberRequest[fizzBuzzParams] + 1

	return strings.Join(result, ",")
}

func MostRequestedRequest(mapNumberRequest map[models.FizzBuzzParams]int) (*models.FizzBuzzParams, *int, *models.Error) {

	// Check map is not empty
	if len(mapNumberRequest) == 0 {
		return nil, nil, &models.Error{StatusCode: 400, Message: "Fizzbuzz service has been never call before, call it before to get some data"}
	}

	var maxNbRequest int
	var associatedKey models.FizzBuzzParams

	// Iterate on map
	for key, nbRequest := range mapNumberRequest {
		if nbRequest > maxNbRequest {
			maxNbRequest = nbRequest
			associatedKey = key
		}
	}

	return &associatedKey, &maxNbRequest, nil
}
