package services

import (
	"testing"

	"github.com/Fizzbuzz/models"
)

func Test_Should_Be_Multiple(t *testing.T) {

	if isMultiple := isMultiple(4, 2); isMultiple != true {
		t.Error()
	}
}
func Test_Should_Not_Be_Multiple(t *testing.T) {

	if isMultiple := isMultiple(4, 3); isMultiple != false {
		t.Error()
	}
}

func Test_Apply_Rules_When_Multiple_Int1(t *testing.T) {
	if wordReturn := applyRules(12, 4, 5, "fizz", "buzz"); wordReturn != "fizz" {
		t.Error()
	}
}

func Test_Apply_Rules_When_Multiple_Int2(t *testing.T) {
	if wordReturn := applyRules(15, 4, 5, "fizz", "buzz"); wordReturn != "buzz" {
		t.Error()
	}
}

func Test_Apply_Rules_When_Multiple_Int1_And_Int2(t *testing.T) {
	if wordReturn := applyRules(15, 3, 5, "fizz", "buzz"); wordReturn != "fizzbuzz" {
		t.Error()
	}
}

func Test_FizzBuzz_Result_Without_Fizz_Buzz(t *testing.T) {

	fizzbuzzParams := models.FizzBuzzParams{
		Int1:  20,
		Int2:  100,
		Str1:  "fizz",
		Str2:  "buzz",
		Limit: 10,
	}

	resultAttended := "1,2,3,4,5,6,7,8,9,10"
	mapNumberRequest := make(map[models.FizzBuzzParams]int)

	if wordReturn := GetFizzbuzzResult(fizzbuzzParams, mapNumberRequest); wordReturn != resultAttended {
		t.Error()
	}
}

func Test_FizzBuzz_Result_With_Fizz(t *testing.T) {
	fizzbuzzParams := models.FizzBuzzParams{
		Int1:  2,
		Int2:  100,
		Str1:  "fizz",
		Str2:  "buzz",
		Limit: 20,
	}

	resultAttended := "1,fizz,3,fizz,5,fizz,7,fizz,9,fizz,11,fizz,13,fizz,15,fizz,17,fizz,19,fizz"
	mapNumberRequest := make(map[models.FizzBuzzParams]int)

	if wordReturn := GetFizzbuzzResult(fizzbuzzParams, mapNumberRequest); wordReturn != resultAttended {
		t.Error()
	}
}

func Test_FizzBuzz_Result_With_Buzz(t *testing.T) {
	fizzbuzzParams := models.FizzBuzzParams{
		Int1:  100,
		Int2:  2,
		Str1:  "fizz",
		Str2:  "buzz",
		Limit: 20,
	}

	resultAttended := "1,buzz,3,buzz,5,buzz,7,buzz,9,buzz,11,buzz,13,buzz,15,buzz,17,buzz,19,buzz"
	mapNumberRequest := make(map[models.FizzBuzzParams]int)

	if wordReturn := GetFizzbuzzResult(fizzbuzzParams, mapNumberRequest); wordReturn != resultAttended {
		t.Error()
	}
}

func Test_FizzBuzz_Result_With_Fizz_Buzz(t *testing.T) {

	fizzbuzzParams := models.FizzBuzzParams{
		Int1:  2,
		Int2:  5,
		Str1:  "fizz",
		Str2:  "buzz",
		Limit: 25,
	}

	resultAttended := "1,fizz,3,fizz,buzz,fizz,7,fizz,9,fizzbuzz,11,fizz,13,fizz,buzz,fizz,17,fizz,19,fizzbuzz,21,fizz,23,fizz,buzz"
	mapNumberRequest := make(map[models.FizzBuzzParams]int)

	if wordReturn := GetFizzbuzzResult(fizzbuzzParams, mapNumberRequest); wordReturn != resultAttended {
		t.Error()
	}
}

func Test_Most_Request_No_Data(t *testing.T) {

	mapNumberRequest := make(map[models.FizzBuzzParams]int)

	params, numberRequest, err := MostRequestedRequest(mapNumberRequest)
	_, _ = params, numberRequest

	if err == nil {
		t.Error()
	}

}

func Test_Most_Request_With_Data(t *testing.T) {

	fizzBuzzParamsAttended := models.FizzBuzzParams{
		Int1:  2,
		Int2:  5,
		Str1:  "fizz",
		Str2:  "buzz",
		Limit: 25,
	}

	mapNumberRequest := map[models.FizzBuzzParams]int{
		fizzBuzzParamsAttended: 5,
		{
			Int1:  20,
			Int2:  50,
			Str1:  "fizza",
			Str2:  "buzza",
			Limit: 100,
		}: 4,
	}

	params, numberRequest, err := MostRequestedRequest(mapNumberRequest)

	if err != nil {
		t.Error()
	}

	if *params != fizzBuzzParamsAttended {
		t.Error()
	}

	if *numberRequest != 5 {
		t.Error()
	}

}
