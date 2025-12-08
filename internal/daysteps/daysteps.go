package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return fmt.Errorf("data parse error: expected 3 parts (steps,name,duration), got %d", len(parts))
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

	ds.Steps = steps

	var dur time.Duration
	dur, err = time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("failed to parse duration '%s': %w", parts[2], err)
	}

	if dur <= 0 {
		return errors.New("duration must be greater than 0")
	}

	ds.Duration = dur

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	dist := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Height, ds.Personal.Weight, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, calories)

	return result, nil
}
