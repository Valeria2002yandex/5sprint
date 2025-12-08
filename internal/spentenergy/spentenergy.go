package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		err := errors.New("the number of steps is negative or equal to zero")
		return 0.0, err
	}

	if weight <= 0 {
		err := errors.New("the user's weight is negative or zero")
		return 0.0, err
	}

	if height <= 0 {
		err := errors.New("the user's height is negative or zero")
		return 0.0, err
	}

	if duration <= time.Duration(0) {
		err := errors.New("the walking time is too short")
		return 0.0, err
	}
	averageSpeed := MeanSpeed(steps, height, duration)

	calories := weight * averageSpeed * duration.Minutes() / minInH

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		err := errors.New("the number of steps is negative or equal to zero")
		return 0.0, err
	}

	if weight <= 0 {
		err := errors.New("the user's weight is negative or zero")
		return 0.0, err
	}

	if height <= 0 {
		err := errors.New("the user's height is negative or zero")
		return 0.0, err
	}

	if duration <= time.Duration(0) {
		err := errors.New("the running time is too short")
		return 0.0, err
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	calories := weight * averageSpeed * duration.Minutes() / minInH
	return calories, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= time.Duration(0) {
		return 0.0
	}
	if steps < 0 {
		return 0.0
	}

	dist := Distance(steps, height)

	return dist / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLenght := height * stepLengthCoefficient

	return float64(steps) * stepLenght / mInKm
}
