package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var family int

	fmt.Println("Input the number of families : ")
	fmt.Scanln(&family)
	fmt.Println("Input the number of members in the family ( separated by a space) :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	familyMembers := scanner.Text()

	dataBus, err := familyBus(family, familyMembers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Minimum bus required is : ", dataBus)
}

func familyBus(family int, familyMembers string) (int, error) {
	type busData struct {
		Family   int
		Capacity int
	}

	bus := []busData{}

	familyMemberArrString := strings.Fields(familyMembers)
	familyMemberArrInt := []int{}

	for _, val := range familyMemberArrString {
		familyNum, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		familyMemberArrInt = append(familyMemberArrInt, familyNum)
	}

	if family != len(familyMemberArrInt) {
		return 0, fmt.Errorf("\ninput must be equal with count of family")
	}

	sort.Slice(familyMemberArrInt, func(i, j int) bool {
		return familyMemberArrInt[i] > familyMemberArrInt[j]
	})

	for !busIsEmpty(familyMemberArrInt) {
		for key, val := range familyMemberArrInt {
			for val > 4 {
				bus = append(bus, busData{
					Family:   1,
					Capacity: 4,
				})
				val = val - 4
			}

			if val < 4 {
				busKey, isAvailable := func(people int, bus []busData) (int, bool) {
					for key, val := range bus {
						if val.Family < 2 && val.Capacity+people <= 4 {
							return key, true
						}
					}
					return 0, false
				}(val, bus)

				if !isAvailable {
					bus = append(bus, busData{
						Family:   1,
						Capacity: val,
					})

					familyMemberArrInt[key] = 0
				} else {
					bus[busKey] = busData{
						Family:   bus[busKey].Family + 1,
						Capacity: bus[busKey].Capacity + val,
					}

					familyMemberArrInt[key] = 0
				}
			} else if val == 4 {
				bus = append(bus, busData{
					Family:   1,
					Capacity: 4,
				})

				familyMemberArrInt[key] = 0
			}

		}
	}

	return len(bus), nil
}

func busIsEmpty(arr []int) bool {
	for _, val := range arr {
		if val != 0 {
			return false
		}
	}

	return true
}
