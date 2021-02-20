package main

import (
	"machine"
	"time"

	weatherstation "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch7/weather-station"
	"tinygo.org/x/drivers/bme280"
	"tinygo.org/x/drivers/st7735"
)

func printError(message string, err error) {
	println(message, err.Error())
	time.Sleep(time.Second)
}

func main() {
	time.Sleep(5 * time.Second)

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
	})

	resetPin := machine.D6
	dcPin := machine.D5
	csPin := machine.D7
	backLightPin := machine.D2

	display := st7735.New(machine.SPI0, resetPin, dcPin, csPin, backLightPin)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_180,
	})

	machine.I2C0.Configure(machine.I2CConfig{})
	sensor := bme280.New(machine.I2C0)
	sensor.Configure()

	weatherStation := weatherstation.New(sensor, display)

	weatherStation.CheckSensorConnectivity()

	for {

		temperature, pressure, humidity, altitude, err := weatherStation.ReadData()
		if err != nil {
			printError("could not read sensor data:", err)
		}

		weatherStation.DisplayData(temperature, pressure, humidity, altitude)

		time.Sleep(2 * time.Second)

	}

}
