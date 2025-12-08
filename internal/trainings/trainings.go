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
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("data parse error: expected 3 parts (steps,activityType,duration), got %d in '%s'", len(parts), datastring)
	}

	var steps int
	steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("failed to parse steps '%s': %w", parts[0], err)
	}
	if steps <= 0 {
		return errors.New("the number of steps must be greater than 0")
	}

	t.Steps = steps

	trainingType := strings.TrimSpace(parts[1])
	if trainingType == "" {
		return errors.New("training type cannot be empty")
	}

	t.TrainingType = trainingType

	var dur time.Duration
	dur, err = time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("failed to parse duration '%s': %w", parts[2], err)
	}

	if dur <= 0 {
		return errors.New("duration must be greater than 0")
	}

	t.Duration = dur
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	switch t.TrainingType {
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
