package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

func findMin(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}

	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func transformSet(set map[int]int, str string) {
	str = strings.Replace(str, "\r", "", -1)
	var s []string = strings.Split(str, " ")
	key, _ := strconv.Atoi(s[1])
	value, _ := strconv.Atoi(s[0])
	len, _ := strconv.Atoi(s[2])
	var i int = 0
	for i < len {
		set[key+i] = value + i
		i++
	}
}

/*func returnSet(set map[int]int, i int) int {
	if value, ok := set[i]; ok {
		return value
	}
	return i
}*/

func returnCorrespondance(str string, i int) int {
	lignes := strings.Split(str, "\r\n")
	for j, ligne := range lignes {
		if j != 0 {
			var s []string = strings.Split(ligne, " ")
			entry, _ := strconv.Atoi(s[1])
			exit, _ := strconv.Atoi(s[0])
			len, _ := strconv.Atoi(s[2])
			if entry <= i && i < entry+len {
				return exit - entry + i
			}
		}
	}
	return i
}

func part1(str string) int {
	var categories []string = strings.Split(str, "\r\n\r\n")
	var s []string = strings.Split(categories[0], ": ")
	var seeds []string = strings.Split(s[1], " ")
	var locationTab []int
	for _, seed_ := range seeds {
		seed, _ := strconv.Atoi(seed_)
		var soil int = returnCorrespondance(categories[1], seed)
		var fertilizer int = returnCorrespondance(categories[2], soil)
		var water int = returnCorrespondance(categories[3], fertilizer)
		var light int = returnCorrespondance(categories[4], water)
		var temperature int = returnCorrespondance(categories[5], light)
		var humidity int = returnCorrespondance(categories[6], temperature)
		var location int = returnCorrespondance(categories[7], humidity)
		locationTab = append(locationTab, location)
	}
	return findMin(locationTab)
}

func part2(str string) int {
	var categories []string = strings.Split(str, "\r\n\r\n")
	var s []string = strings.Split(categories[0], ": ")
	var seeds []string = strings.Split(s[1], " ")
	var locationTab []int
	for i := 0; i < len(seeds)-1; i += 2 {
		seedi, _ := strconv.Atoi(seeds[i])
		seedi1, _ := strconv.Atoi(seeds[i+1])
		for j := 0; j < seedi1; j++ {
			var soil int = returnCorrespondance(categories[1], seedi+j)
			var fertilizer int = returnCorrespondance(categories[2], soil)
			var water int = returnCorrespondance(categories[3], fertilizer)
			var light int = returnCorrespondance(categories[4], water)
			var temperature int = returnCorrespondance(categories[5], light)
			var humidity int = returnCorrespondance(categories[6], temperature)
			var location int = returnCorrespondance(categories[7], humidity)
			locationTab = append(locationTab, location)
		}
	}
	return findMin(locationTab)
}

/*func part1(str string) int {
	seed_to_soil := make(map[int]int)
	soil_to_fertilizer := make(map[int]int)
	fertilizer_to_water := make(map[int]int)
	water_to_light := make(map[int]int)
	light_to_temperature := make(map[int]int)
	temperature_to_humidity := make(map[int]int)
	humidity_to_location := make(map[int]int)

	var categories []string = strings.Split(str, "\r\n\r\n")
	var s []string = strings.Split(categories[0], ": ")
	var seeds []string = strings.Split(s[1], " ")
	var locationTab []int

	var soils []string = strings.Split(categories[1], "\r\n")
	for i, soil := range soils {
		if i != 0 {
			transformSet(seed_to_soil, soil)
		}
	}

	var fertilizers []string = strings.Split(categories[2], "\r\n")
	for i, fertilizer := range fertilizers {
		if i != 0 {
			transformSet(soil_to_fertilizer, fertilizer)
		}
	}

	var waters []string = strings.Split(categories[3], "\r\n")
	for i, water := range waters {
		if i != 0 {
			transformSet(fertilizer_to_water, water)
		}
	}

	var lights []string = strings.Split(categories[4], "\r\n")
	for i, light := range lights {
		if i != 0 {
			transformSet(water_to_light, light)
		}
	}

	var temperatures []string = strings.Split(categories[5], "\r\n")
	for i, temparature := range temperatures {
		if i != 0 {
			transformSet(light_to_temperature, temparature)
		}
	}

	var humidity []string = strings.Split(categories[6], "\r\n")
	for i, humid := range humidity {
		if i != 0 {
			transformSet(temperature_to_humidity, humid)
		}
	}

	var locations []string = strings.Split(categories[7], "\r\n")
	for i, location := range locations {
		if i != 0 {
			transformSet(humidity_to_location, location)
		}
	}

	for _, seed_ := range seeds {
		seed, _ := strconv.Atoi(seed_)
		var soil int = returnSet(seed_to_soil, seed)
		var fertilizer int = returnSet(soil_to_fertilizer, soil)
		var water int = returnSet(fertilizer_to_water, fertilizer)
		var light int = returnSet(water_to_light, water)
		var temperature int = returnSet(light_to_temperature, light)
		var humidity int = returnSet(temperature_to_humidity, temperature)
		var location int = returnSet(humidity_to_location, humidity)
		locationTab = append(locationTab, location)
	}
	return findMin(locationTab)
}*/

func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
