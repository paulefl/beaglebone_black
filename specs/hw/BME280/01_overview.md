# BME280 — Overview

## Digital Humidity, Pressure and Temperature Sensor

### Key Features

- **Package:** 2.5 mm x 2.5 mm x 0.93 mm metal lid LGA
- **Digital interface:** I²C (up to 3.4 MHz) and SPI (3 and 4 wire, up to 10 MHz)
- **Supply voltage:**
  - VDD main supply voltage range: 1.71 V to 3.6 V
  - VDDIO interface voltage range: 1.2 V to 3.6 V
- **Current consumption:**
  - 1.8 µA @ 1 Hz humidity and temperature
  - 2.8 µA @ 1 Hz pressure and temperature
  - 3.6 µA @ 1 Hz humidity, pressure and temperature
  - 0.1 µA in sleep mode
- **Operating range:** -40…+85 °C, 0…100 % rel. humidity, 300…1100 hPa
- Humidity sensor and pressure sensor can be independently enabled / disabled
- Register and performance compatible to Bosch Sensortec BMP280 digital pressure sensor
- RoHS compliant, halogen-free, MSL1

### Key Parameters for Humidity Sensor

- **Response time (τ63%):** 1 s
- **Accuracy tolerance:** ±3 % relative humidity
- **Hysteresis:** ±1 % relative humidity

### Key Parameters for Pressure Sensor

- **RMS Noise:** 0.2 Pa, equiv. to 1.7 cm
- **Offset temperature coefficient:** ±1.5 Pa/K, equiv. to ±12.6 cm at 1 °C temperature change

## Typical Applications

- Context awareness, e.g. skin detection, room change detection
- Fitness monitoring / well-being
- Warning regarding dryness or high temperatures
- Measurement of volume and air flow
- Home automation control
- Control heating, venting, air conditioning (HVAC)
- Internet of things
- Indoor navigation (change of floor detection, elevator detection)
- GPS enhancement (e.g. time-to-first-fix improvement, dead reckoning, slope detection)
- Outdoor navigation, leisure and sports applications
- Weather forecast
- Vertical velocity indication (rise/sink speed)

## Target Devices

- Handsets such as mobile phones, tablet PCs, GPS devices
- Navigation systems
- Gaming, e.g. flying toys
- Camera (DSC, video)
- Home weather stations
- Flying toys
- Watches

## General Description

The BME280 is a combined digital humidity, pressure and temperature sensor based on proven sensing principles. The sensor module is housed in an extremely compact metal-lid LGA package with a footprint of only 2.5 × 2.5 mm² with a height of 0.93 mm. Its small dimensions and its low power consumption allow the implementation in battery driven devices such as handsets, GPS modules or watches. The BME280 is register and performance compatible to the Bosch Sensortec BMP280 digital pressure sensor (see chapter 5.2 for details).

The BME280 achieves high performance in all applications requiring humidity and pressure measurement. These emerging applications of home automation control, in-door navigation, fitness as well as GPS refinement require a high accuracy and a low TCO at the same time.

The humidity sensor provides an extremely fast response time for fast context awareness applications and high overall accuracy over a wide temperature range.

The pressure sensor is an absolute barometric pressure sensor with extremely high accuracy and resolution and drastically lower noise than the Bosch Sensortec BMP180.

The integrated temperature sensor has been optimized for lowest noise and highest resolution. Its output is used for temperature compensation of the pressure and humidity sensors and can also be used for estimation of the ambient temperature.

The sensor provides both SPI and I²C interfaces and can be supplied using 1.71 to 3.6 V for the sensor supply VDD and 1.2 to 3.6 V for the interface supply VDDIO. Measurements can be triggered by the host or performed in regular intervals. When the sensor is disabled, current consumption drops to 0.1 µA.

BME280 can be operated in three power modes:

- Sleep mode
- Normal mode
- Forced mode

In order to tailor data rate, noise, response time and current consumption to the needs of the user, a variety of oversampling modes, filter modes and data rates can be selected.
