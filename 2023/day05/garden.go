package day05

import (
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

type almanac struct {
	Seeds                 []int64
	SeedsPart2            []int64
	SeedToSoil            almaMap
	SoilToFertilizer      almaMap
	FertilizerToWater     almaMap
	WaterToLight          almaMap
	LightToTemperature    almaMap
	TemperatureToHumidity almaMap
	HumidityToLocation    almaMap
}

type almaMap struct {
	AlmaValues []almaValues
}

func (am *almaMap) getDestForSrc(src int64) (dest int64) {
	dest = src

	for _, values := range am.AlmaValues {
		if src < values.SrcRangeStart || src > values.SrcRangeStart+values.Range-1 {
			continue
		}

		if values.SrcRangeStart < values.DestRangeStart {
			dest = src + (values.DestRangeStart - values.SrcRangeStart)

			break
		}

		if values.SrcRangeStart > values.DestRangeStart {
			dest = src - (values.SrcRangeStart - values.DestRangeStart)

			break
		}
	}

	return dest
}

type almaValues struct {
	DestRangeStart int64
	SrcRangeStart  int64
	Range          int64
}

func BothParts() (part1 int64, part2 int64) {
	var (
		file []byte
		alma almanac
		err  error
	)

	start := time.Now()

	file, err = os.ReadFile("./2023/day05/input.txt")
	if err != nil {
		log.Fatalf("error while reading file: %s", err)
	}

	lines := strings.FieldsFunc(string(file), func(r rune) bool {
		return r == '\n'
	})

	alma, err = prepareData(lines)
	if err != nil {
		log.Fatalf("error while preparing data: %s", err)
	}

	part1, part2 = findLowestLocationId(alma)

	elapsed := time.Since(start)
	log.Printf("execution took %s", elapsed)

	return part1, part2
}

func prepareData(lines []string) (alma almanac, err error) {
	var (
		seeds        []int64
		seedsPart2   []int64
		section      string
		newAlmaValue almaValues
		seedToSoil,
		soilToFertilizer,
		fertilizerToWater,
		waterToLight,
		lightToTemperature,
		temperatureToHumidity,
		humidityToLocation []almaValues
	)

	for _, line := range lines {
		if unicode.IsLetter(rune(line[0])) {
			switch true {
			case strings.Contains(line, "seeds"):
				splitLine := strings.Split(line, ":")
				seeds, err = util.FindNumbers(splitLine[1])
				if err != nil {
					return alma, err
				}

				seedsPart2 = buildSeedsForPart2(seeds)

			case strings.Contains(line, "seed-to-soil"):
				section = "seed-to-soil"
			case strings.Contains(line, "soil-to-fertilizer"):
				section = "soil-to-fertilizer"
			case strings.Contains(line, "fertilizer-to-water"):
				section = "fertilizer-to-water"
			case strings.Contains(line, "water-to-light"):
				section = "water-to-light"
			case strings.Contains(line, "light-to-temperature"):
				section = "light-to-temperature"
			case strings.Contains(line, "temperature-to-humidity"):
				section = "temperature-to-humidity"
			case strings.Contains(line, "humidity-to-location"):
				section = "humidity-to-location"
			}
		}

		if unicode.IsNumber(rune(line[0])) {
			switch section {
			case "seed-to-soil":
				newAlmaValue, err = buildAlmaMaps(line)
				if err != nil {
					return alma, err
				}

				seedToSoil = append(seedToSoil, newAlmaValue)
			case "soil-to-fertilizer":
				newAlmaValue, err = buildAlmaMaps(line)
				if err != nil {
					return alma, err
				}

				soilToFertilizer = append(soilToFertilizer, newAlmaValue)
			case "fertilizer-to-water":
				newAlmaValue, err = buildAlmaMaps(line)
				if err != nil {
					return alma, err
				}

				fertilizerToWater = append(fertilizerToWater, newAlmaValue)
			case "water-to-light":
				newAlmaValue, err = buildAlmaMaps(line)
				if err != nil {
					return alma, err
				}

				waterToLight = append(waterToLight, newAlmaValue)
			case "light-to-temperature":
				newAlmaValue, err = buildAlmaMaps(line)
				if err != nil {
					return alma, err
				}

				lightToTemperature = append(lightToTemperature, newAlmaValue)
			case "temperature-to-humidity":
				newAlmaValue, err = buildAlmaMaps(line)
				if err != nil {
					return alma, err
				}

				temperatureToHumidity = append(temperatureToHumidity, newAlmaValue)
			case "humidity-to-location":
				newAlmaValue, err = buildAlmaMaps(line)
				if err != nil {
					return alma, err
				}

				humidityToLocation = append(humidityToLocation, newAlmaValue)
			}
		}
	}

	alma.Seeds = seeds
	alma.SeedsPart2 = seedsPart2
	alma.SeedToSoil.AlmaValues = seedToSoil
	alma.SoilToFertilizer.AlmaValues = soilToFertilizer
	alma.FertilizerToWater.AlmaValues = fertilizerToWater
	alma.WaterToLight.AlmaValues = waterToLight
	alma.LightToTemperature.AlmaValues = lightToTemperature
	alma.TemperatureToHumidity.AlmaValues = temperatureToHumidity
	alma.HumidityToLocation.AlmaValues = humidityToLocation

	return alma, err
}

func buildAlmaMaps(line string) (newAlmaValue almaValues, err error) {
	var numbers []int64

	numbers, err = util.FindNumbers(line)
	if err != nil {
		return newAlmaValue, err
	}

	newAlmaValue.DestRangeStart = numbers[0]
	newAlmaValue.SrcRangeStart = numbers[1]
	newAlmaValue.Range = numbers[2]

	return newAlmaValue, err
}

func buildSeedsForPart2(seeds []int64) (seedsPart2 []int64) {
	for i, seed := range seeds {
		if i%2 == 0 {
			start := seed
			seedRange := seeds[i+1]

			var j int64 = 0
			for j < seedRange {
				seedsPart2 = append(seedsPart2, start+j)
				j++
			}

			log.Printf("seed batch %d of 10 for part 2 ready\n", (i/2)+1)
		}
	}

	return seedsPart2
}

func findLowestLocationId(alma almanac) (int64, int64) {
	var (
		lowestLocationPart1, lowestLocationPart2 int64
	)

	for _, seed := range alma.Seeds {
		soil := alma.SeedToSoil.getDestForSrc(seed)
		fertilizer := alma.SoilToFertilizer.getDestForSrc(soil)
		water := alma.FertilizerToWater.getDestForSrc(fertilizer)
		light := alma.WaterToLight.getDestForSrc(water)
		temperature := alma.LightToTemperature.getDestForSrc(light)
		humidity := alma.TemperatureToHumidity.getDestForSrc(temperature)
		location := alma.HumidityToLocation.getDestForSrc(humidity)

		if lowestLocationPart1 == 0 {
			lowestLocationPart1 = location
		}

		if lowestLocationPart1 > location {
			lowestLocationPart1 = location
			log.Printf("new lowest location id: %d\n", lowestLocationPart1)
		}
	}

	log.Printf("lowest location id for part 1: %d\n", lowestLocationPart1)

	log.Println("===============")
	log.Println("finished part 1")
	log.Println("starting part 2")
	log.Println("===============")

	for _, seed := range alma.SeedsPart2 {
		soil := alma.SeedToSoil.getDestForSrc(seed)
		fertilizer := alma.SoilToFertilizer.getDestForSrc(soil)
		water := alma.FertilizerToWater.getDestForSrc(fertilizer)
		light := alma.WaterToLight.getDestForSrc(water)
		temperature := alma.LightToTemperature.getDestForSrc(light)
		humidity := alma.TemperatureToHumidity.getDestForSrc(temperature)
		location := alma.HumidityToLocation.getDestForSrc(humidity)

		if lowestLocationPart2 == 0 {
			lowestLocationPart2 = location
		}

		if lowestLocationPart2 > location {
			lowestLocationPart2 = location
			log.Printf("new lowest location id for part 2: %d\n", lowestLocationPart2)
		}
	}

	return lowestLocationPart1, lowestLocationPart2
}
