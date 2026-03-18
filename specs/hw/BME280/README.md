# BME280 — Combined Humidity, Pressure and Temperature Sensor

**Source:** BST-BME280-DS002, Document Revision 1.24, February 2024
**Document Number:** BST-BME280-DS001-24
**Sales Part Number (SPN):** 0 273 141 185
**Manufacturer:** Bosch Sensortec GmbH, Gerhard-Kindler-Straße 9, 72770 Reutlingen, Germany

## General Description

The BME280 is a combined digital humidity, pressure and temperature sensor based on proven sensing principles. The sensor module is housed in an extremely compact metal-lid LGA package with a footprint of only 2.5 × 2.5 mm² with a height of 0.93 mm. Its small dimensions and its low power consumption allow the implementation in battery driven devices such as handsets, GPS modules or watches.

## Key Features

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

### Key Parameters — Humidity Sensor

- **Response time (τ63%):** 1 s
- **Accuracy tolerance:** ±3 % relative humidity
- **Hysteresis:** ±1 % relative humidity

### Key Parameters — Pressure Sensor

- **RMS Noise:** 0.2 Pa, equiv. to 1.7 cm
- **Offset temperature coefficient:** ±1.5 Pa/K, equiv. to ±12.6 cm at 1 °C temperature change

## Table of Contents

| File | Content |
|------|---------|
| [01_overview.md](01_overview.md) | General description, typical applications, target devices |
| [02_specifications.md](02_specifications.md) | Section 1: Electrical specs, humidity, pressure, temperature parameters |
| [03_absolute_maximum_ratings.md](03_absolute_maximum_ratings.md) | Section 2: Absolute maximum ratings |
| [04_functional_description.md](04_functional_description.md) | Section 3: Block diagram, power management, sensor modes, measurement flow, recommended modes, noise |
| [05_data_readout.md](05_data_readout.md) | Section 4: Data register shadowing, output compensation, trimming parameters, compensation formulas |
| [06_registers.md](06_registers.md) | Section 5: Memory map and register description |
| [07_digital_interfaces.md](07_digital_interfaces.md) | Section 6: I²C and SPI interfaces, timings |
| [08_pinout.md](08_pinout.md) | Section 7: Pin-out, connection diagrams, package info, soldering guidelines |
| [09_appendix_a_compensation.md](09_appendix_a_compensation.md) | Section 8: Alternative compensation formulas (double precision floating point, 32-bit fixed point) |
| [10_appendix_b_timing.md](10_appendix_b_timing.md) | Section 9: Measurement time and current calculation |
| [11_seengreat_module.md](11_seengreat_module.md) | **SeenGreat Breakout Board** — module specs, pin description, BeagleBone Black wiring |
