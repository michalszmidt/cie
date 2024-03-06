package processing

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	. "github.com/arran4/golang-ical"
	. "github.com/jucardi/go-streams/v2/streams"
	. "github.com/michalszmidt/cie/objects"
)

func Process(path *string, out *string, cie *CIE) uint64 {
	var counter uint64 = 0

	file, err := os.Open(*path)
	file2, err := os.Create(*out)
	defer file.Close()
	defer file2.Close()

	calendar, _ := ParseCalendar(file)
	newCalendar := NewCalendar()

	events := calendar.Events()

	if err != nil {
		fmt.Println("{}", err)
	}

	// For each event
	FromArray(events, 0).
		Filter(
			func(event *VEvent) bool {
				// For each event
				// Filter one by checking if set of bolleans has true and false or only false\
				// if there is true event has to be omited
				setOfBools := Map(
					FromArray(cie.ToRemove.Rules, 0),
					// for each deletion rule
					// create cheking regex
					// if regex returns true, set will point to filter event
					func(rule RemoveRule) bool {
						rgx := regexp.MustCompile(rule.Regex)
						var value_to_check string

						// get value of key that matches one to check
						var i int
						for _, property := range event.ComponentBase.Properties {
							if strings.Compare(property.BaseProperty.IANAToken, rule.KeyName) == 0 {
								value_to_check = property.Value
								break
							}
							i++
						}
						// if there was no maching key to pattern
						// assume that event should not be filtered
						if i < len(event.ComponentBase.Properties) {
							return rgx.MatchString(value_to_check)
						}

						// if regex sets true, event will be added to set
						return false
					},
				).
					Stream().
					ToDistinct()

				// check if set is empty
				if setOfBools.Contains(true) {
					counter++
					return false
				}
				return true
			}).ForEach(func(event *VEvent) {
		newCalendar.AddVEvent(event)
	})

	newCalendar.SerializeTo(file2)
	return counter
}
