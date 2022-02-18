// TCALC -- simple calculator that also understands time stamps
// SvM 13-FEB-2021 - 14-FEB-2021
//
// a. add error checking instead of just crashing everywhere
// c. also read from stdin
// d. add options to specify output format

package main

import (
	"fmt"
	"go/token"
	"go/types"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func hms2sec(x string) string {
	sec := 0

	re := regexp.MustCompile(`[0-9]+[hms]`)
	xx := re.FindAllString(x, -1)

	for _, xxx := range xx {
		if strings.Contains(xxx, "h") {
			h, _ := strconv.Atoi(strings.TrimRight(xxx, "h"))
			sec += h * 60 * 60
		} else if strings.Contains(xxx, "m") {
			m, _ := strconv.Atoi(strings.TrimRight(xxx, "m"))
			sec += m * 60
		} else if strings.Contains(xxx, "s") {
			s, _ := strconv.Atoi(strings.TrimRight(xxx, "s"))
			sec += s
		}
	}
	return strconv.Itoa(sec)
}

func colons2sec(x string) string {
	sec := 0
	mult := 1

	xx := strings.Split(x, ":")
	for i := len(xx); i > 0; i-- {
		xxx, _ := strconv.Atoi(xx[i-1])
		sec += xxx * mult
		mult *= 60
	}
	return strconv.Itoa(sec)
}

func sec2hms(sec int) string {
	var ret string

	h := sec / 3600
	m := (sec % 3600) / 60
	s := sec % 60

	if h > 0 {
		ret += fmt.Sprintf("%02dh", h)
	}
	if m > 0 {
		ret += fmt.Sprintf("%02dm", m)
	}
	ret += fmt.Sprintf("%02ds", s)
	return ret
}

func main() {
	args := os.Args[1:]

	// convert any hms or ::: args to ints, leave the rest alone
	for i, a := range args {
		if matched, _ := regexp.MatchString(`^[0-9:]+$`, a); matched {
			args[i] = colons2sec(a)
		} else if matched, _ := regexp.MatchString(`^[0-9hms]+$`, a); matched {
			args[i] = hms2sec(a)
		}
	}

	// I don't understand any of this, I just copied it from
	// https://stackoverflow.com/questions/23923383/evaluate-formula-in-go/23923467
	fs := token.NewFileSet()
	tv, err := types.Eval(fs, nil, token.NoPos, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}

	result, _ := strconv.Atoi(tv.Value.String())
	fmt.Println(result, "\t", sec2hms(result))
}
