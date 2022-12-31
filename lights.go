package main

import (
	"log"
	"strings"

	"github.com/amimof/huego"
)

func setLightsState(nameOfLightOrGroup string, state huego.State) {
	bridge := getBridge()

	groups, err := bridge.GetGroups()
	if err != nil {
		log.Fatal(err)
	}

	for _, group := range groups {
		if strings.Contains(strings.ToLower(group.Name), strings.ToLower(nameOfLightOrGroup)) {
			go func(group huego.Group) { group.SetState(state) }(group)
		}
	}

	lights, err := bridge.GetLights()
	if err != nil {
		log.Fatal(err)
	}

	for _, light := range lights {
		if strings.Contains(strings.ToLower(light.Name), strings.ToLower(nameOfLightOrGroup)) {
			go func(light huego.Light) { light.SetState(state) }(light)
		}
	}
}

func Toggle(nameOfLightOrGroup string, isOn bool) {
	state := huego.State{
		On: isOn,
	}

	setLightsState(nameOfLightOrGroup, state)
}

func Brighten(nameOfLight string, percent int) {
	if percent > 100 {
		percent = 100
	}

	if percent < 0 {
		percent = 0
	}

	brightnessValue := ((float64(254) * float64(percent)) / float64(100))
	state := huego.State{
		On:  true,
		Bri: uint8(brightnessValue),
	}

	setLightsState(nameOfLight, state)
}
