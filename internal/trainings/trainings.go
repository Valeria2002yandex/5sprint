package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	Personal     personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		err := errors.New("data parse error")
		return err
	}

	var steps int
	steps, err = strconv.Atoi(parts[0])
	if err != nil {
		err = fmt.Errorf("failed to parse steps '%s': %w", parts[0], err)
		return
	}
	if steps <= 0 {
		err = errors.New("the number of steps is less than or equal to 0")
		return err
	}

	t.Steps = steps

	var dur time.Duration
	dur, err = time.ParseDuration(parts[2])
	if err != nil {
		return err
	}

	if dur <= 0 {
		err = errors.New("duration must be greater than 0")
		return err
	}

	t.Duration = dur
	return
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(t.Steps, float64(t.Personal.Height))
	speed := spentenergy.MeanSpeed(t.Steps, float64(t.Personal.Height), t.Duration)

	var activity string

	switch activity {
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Height, t.Personal.Weight, t.Duration)

		if err != nil {
			return "", err
		}

		result := fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.Duration.Hours(), dist, speed, calories)

		return result, nil

	case "Ходьба":

		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Height, t.Personal.Weight, t.Duration)
		if err != nil {
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.Duration.Hours(), dist, speed, calories)

		return result, nil

	default:

		err := errors.New("неизвестный тип тренировки")
		return "", err
	}

}
