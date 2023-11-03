package processing

import (
	"fmt"
	// "log"
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

	// eventstodelete := make([]int, 0)

	From[*VEvent](events, 0).
		Filter(
			func(event *VEvent) bool {

				setOfTrue := Map(From[RemoveRule](cie.ToRemove.Rules, 0).
					Filter(func(rule RemoveRule) bool {
						for _, property := range event.ComponentBase.Properties {
							if strings.Compare(property.BaseProperty.IANAToken, rule.KeyName) == 0 {
								return true
							}
						}
						return false
					}),
					func(rule RemoveRule) bool {
						rgx := regexp.MustCompile(rule.Regex)
						var name string

						for _, property := range event.ComponentBase.Properties {
							if strings.Compare(property.BaseProperty.IANAToken, rule.KeyName) == 0 {
								name = property.Value
								break
							}
						}
						return rgx.MatchString(name)
					},
				).Stream().ToDistinct()

				if setOfTrue.Contains(true) {
					counter++
					return false
				}
				// fmt.Println(setOfTrue.Contains(true))
				return true
			}).ForEach(func(event *VEvent) {
		newCalendar.AddVEvent(event)
	})

	// for i, event := range events {
	// 	for _, cmpnt := range event.ComponentBase.Properties {
	// 		if cmpnt.BaseProperty.IANAToken == "SUMMARY" {
	// 			if cmpnt.BaseProperty.Value == "AM - Wykład (B/1 Aula Główna)" ||
	// 				cmpnt.BaseProperty.Value == "AM - Ćwiczenia (B/208)" {

	// 				eventstodelete = append(eventstodelete, i)
	// 				fmt.Println("Removed: '", "SUMMARY", cmpnt.BaseProperty.Value, "'")
	// 			}
	// 		}
	// 	}
	// }

	// for _, num := range eventstodelete {
	// 	events = RemoveIndexVEventArr(events, num)
	// }

	// for _, vevent := range events {
	// 	newCalendar.AddVEvent(vevent)
	// }

	// fmt.Println("removed", len(counter))
	newCalendar.SerializeTo(file2)
	return counter
}

// func RemoveIndexVEventArr(s []*VEvent, index int) []*VEvent {
// 	ret := make([]*VEvent, 0)
// 	ret = append(ret, s[:index]...)
// 	return append(ret, s[index+1:]...)
// }
