package data_struct

import "sort"

type Dob struct {
	Day   int
	Month int
	Year  int
}

type Person struct {
	Name  string
	Email string
	Dob   Dob
}

func SortPersons(list []Person, sortBy string) {
	switch sortBy {

	case "name":
		sort.Slice(list, func(i, j int) bool {
			return list[i].Name < list[j].Name
		})

	case "date":
		sort.Slice(list, func(i, j int) bool {
			if list[i].Dob.Year != list[j].Dob.Year {
				return list[i].Dob.Year < list[j].Dob.Year
			}
			if list[i].Dob.Month != list[j].Dob.Month {
				return list[i].Dob.Month < list[j].Dob.Month
			}
			return list[i].Dob.Day < list[j].Dob.Day
		})
	}
}