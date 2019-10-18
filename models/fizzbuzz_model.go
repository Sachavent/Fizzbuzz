package models

type FizzBuzzParams struct {
	Int1  int    `json:"int1" binding:"required"`
	Int2  int    `json:"int2" binding:"required"`
	Limit int    `json:"limit" binding:"required"`
	Str1  string `json:"str1" binding:"required"`
	Str2  string `json:"str2" binding:"required"`
}
