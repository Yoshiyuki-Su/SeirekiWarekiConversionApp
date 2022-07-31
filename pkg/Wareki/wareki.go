package wareki

import (
	"fmt"
	"strconv"
	"time"
)

type warekiStart struct {
	StartDate time.Time
}

var jst = time.FixedZone("JST", 9*60*60)

var warekiStartList = map[string]warekiStart{
	"昭和": {
		StartDate: time.Date(1926, time.December, 25, 0, 0, 0, 0, jst),
	},
	"平成": {
		StartDate: time.Date(1989, time.January, 8, 0, 0, 0, 0, jst),
	},
	"令和": {
		StartDate: time.Date(2019, time.May, 1, 0, 0, 0, 0, jst),
	},
}

func firstYearCheck(year int) string {
	if year == 1 {
		return "元"
	}
	return strconv.Itoa(year)
}

func ConvertToWareki(t *time.Time) (string, error) {

	if !warekiStartList["令和"].StartDate.After(*t) {

		era_year := warekiStartList["令和"].StartDate.Year()
		set_year := t.Year()
		year := (set_year - era_year) + 1

		return "令和" + firstYearCheck(year) + "年", nil
	} else if !warekiStartList["平成"].StartDate.After(*t) {

		era_year := warekiStartList["令和"].StartDate.Year()
		set_year := t.Year()
		year := (set_year - era_year) + 1

		return "平成" + firstYearCheck(year) + "年", nil
	} else if !warekiStartList["昭和"].StartDate.After(*t) {

		era_year := warekiStartList["令和"].StartDate.Year()
		set_year := t.Year()
		year := (set_year - era_year) + 1

		return "昭和" + firstYearCheck(year) + "年", nil
	} else {
		return "", fmt.Errorf("%s に満たされる和暦は存在しません。", t)
	}
}
