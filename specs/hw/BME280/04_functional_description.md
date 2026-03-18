# BME280 — Section 3: Functional Description

## 3.1 Block Diagram

The BME280 contains the following functional blocks:

```
VDD        VDDIO
 |              |
 +---[Voltage   |
 |    reference]|
 |              |
 +---[Voltage regulator (analog & digital)]
 |
 +---[Pressure sensing element] → [Pressure front-end] → [ADC] → [Logic] ←→ [Interface] → SDI
 |                                                                                           SDO
 +---[Humidity sensing element] → [Humidity front-end]                                      SCK
 |                                                                                           CSB
 +---[Temperature sensing element] → [Temperature front-end]
 |
 +---[OSC]
 +---[POR]
 +---[NVM]
 |
GND
```

The block diagram (Figure 2 in the original datasheet) shows a simplified view of the BME280 with:

- Separate voltage reference and voltage regulator for analog & digital domains
- Three sensing elements (pressure, humidity, temperature) each with their own front-end
- Shared ADC and Logic block
- Interface block with 4 pins: SDI, SDO, SCK, CSB
- Power-on reset (POR) generator
- Oscillator (OSC)
- Non-volatile memory (NVM) for calibration data

## 3.2 Power Management

The BME280 has two distinct power supply pins:

- **VDD** is the main power supply for all internal analog and digital functional blocks
- **VDDIO** is a separate power supply pin used for the supply of the digital interface

A power-on reset (POR) generator is built in; it resets the logic part and the register values after both VDD and VDDIO reach their minimum levels. There are no limitations on slope and sequence of raising the VDD and VDDIO levels. After powering up, the sensor settles in sleep mode (described in chapter 3.3.2).

**Important:** It is prohibited to keep any interface pin (SDI, SDO, SCK or CSB) at a logical high level when VDDIO is switched off. Such a configuration can permanently damage the device due to an excessive current flow through the ESD protection diodes.

If VDDIO is supplied, but VDD is not, the interface pins are kept at a high-Z level. The bus can therefore already be used freely before the BME280 VDD supply is established.

Resetting the sensor is possible by cycling VDD level or by writing a soft reset command. Cycling the VDDIO level will not cause a reset.

## 3.3 Sensor Modes

The BME280 offers three sensor modes: sleep mode, forced mode and normal mode. These can be selected using the `mode[1:0]` setting (see chapter 5.4.5). The available modes are:

- **Sleep mode:** no operation, all registers accessible, lowest power, selected after startup
- **Forced mode:** perform one measurement, store results and return to sleep mode
- **Normal mode:** perpetual cycling of measurements and inactive periods

### 3.3.1 Sensor Mode Transitions

The supported mode transitions are:

```
Power OFF (VDD or VDDIO = 0)
       |
       v (VDD and VDDIO supplied)
     Sleep  <──────────────── Mode[1:0] = 00
       |                            ^
       | Mode[1:0] = 01/10/11       |
       v                      Mode[1:0] = 01
    Normal ────────────────> Forced
  (cyclic standby           (one measurement
 and measurement             period)
    periods)
```

If the device is currently performing a measurement, execution of mode switching commands is delayed until the end of the currently running measurement period. Further mode change commands or other write commands to the register `ctrl_hum` are ignored until the mode change command has been executed.

### 3.3.2 Sleep Mode

Sleep mode is entered by default after power on reset. In sleep mode, no measurements are performed and power consumption (IDDSM) is at a minimum. All registers are accessible; Chip-ID and compensation coefficients can be read. There are no special restrictions on interface timings.

### 3.3.3 Forced Mode

In forced mode, a single measurement is performed in accordance to the selected measurement and filter options. When the measurement is finished, the sensor returns to sleep mode and the measurement results can be obtained from the data registers. For a next measurement, forced mode needs to be selected again. This is similar to BMP180 operation. Using forced mode is recommended for applications which require low sampling rate or host-based synchronization.

**Forced mode timing:**
```
                    cycle time = rate of force mode
                    |←─────────────────────────────→|
current:
  IDDP ─────────────────────────────────────────────────────
  IDDT ─                                                   ─
  IDDH ─                                                   ─
  IDDSB─             ──────────────────────────────────────
  IDDSL─
       POR  Write     Mode[1:0]=01   Data    Mode[1:0]=01
            settings               readout
```

The measurement sequence within one cycle: Temperature → Pressure → Humidity

### 3.3.4 Normal Mode

Normal mode comprises an automated perpetual cycling between an (active) measurement period and an (inactive) standby period.

The measurements are performed in accordance to the selected measurement and filter options. The standby time is determined by the setting `t_sb[2:0]` and can be set to between 0.5 and 1000 ms according to Table 27.

The total cycle time depends on the sum of the active time (see chapter 9) and standby time tstandby. The current in the standby period (IDDSB) is slightly higher than in sleep mode. After setting the measurement and filter options and enabling normal mode, the last measurement results can always be obtained at the data registers without the need of further write accesses.

Using normal mode is recommended when using the IIR filter. This is useful for applications in which short-term disturbances (e.g. blowing into the sensor) should be filtered.

**Normal mode timing:**
```
                    cycle time = tmeasure + tstandby
                    |←────────────→|←──────────→|
current:
  IDDP ─────────────────────────────────────────────────────
  IDDT ─                          ─
  IDDH ─                          ─
  IDDSB─       ─────────────────────────────────────────────
  IDDSL─
       POR   Write   Mode[1:0]=11   Data readout
             settings               when needed
```

## 3.4 Measurement Flow

The BME280 measurement period consists of a temperature, pressure and humidity measurement with selectable oversampling. After the measurement period, the pressure and temperature data can be passed through an optional IIR filter, which removes short-term fluctuations in pressure (e.g. caused by slamming a door). For humidity, such a filter is not needed and has not been implemented.

**Measurement cycle flow:**

```
Start measurement cycle
        |
        v
Measure temperature
(oversampling set by osrs_t; skip if osrs_t = 0)
        |
        v
Measure pressure
(oversampling set by osrs_p; skip if osrs_p = 0)
        |
        v
IIR filter enabled?
   Yes /        \ No
      v           v
IIR filter        Measure humidity
initialised?      (oversampling set by osrs_h;
   Yes / No       skip if osrs_h = 0)
      |    |           |
      |    v           |
      |  Copy ADC      |
      |  values to     |
      |  filter memory |
      |  (initialises) |
      v                |
Update filter memory    |
using filter memory,   |
ADC value and filter   |
coefficient            |
      |                |
      v                |
Copy filter memory     |
to output registers    |
      |                |
      +────────────────+
      v
End measurement cycle
```

### 3.4.1 Humidity Measurement

The humidity measurement can be enabled or skipped. When enabled, several oversampling options exist. The humidity measurement is controlled by the `osrs_h[2:0]` setting, which is detailed in chapter 5.4.3. For the humidity measurement, oversampling is possible to reduce the noise. The resolution of the humidity measurement is fixed at 16 bit ADC output.

### 3.4.2 Pressure Measurement

Pressure measurement can be enabled or skipped. When enabled, several oversampling options exist. The pressure measurement is controlled by the `osrs_p[2:0]` setting which is detailed in chapter 5.4.5. For the pressure measurement, oversampling is possible to reduce the noise. The resolution of the pressure data depends on the IIR filter (see chapter 3.4.4) and the oversampling setting (see chapter 5.4.5):

- When the IIR filter is enabled, the pressure resolution is 20 bit.
- When the IIR filter is disabled, the pressure resolution is 16 + (osrs_p – 1) bit, e.g. 18 bit when osrs_p is set to '3'.

### 3.4.3 Temperature Measurement

Temperature measurement can be enabled or skipped. Skipping the measurement could be useful to measure pressure extremely rapidly. When enabled, several oversampling options exist. The temperature measurement is controlled by the `osrs_t[2:0]` setting which is detailed in chapter 5.4.5. For the temperature measurement, oversampling is possible to reduce the noise.

The resolution of the temperature data depends on the IIR filter (see chapter 3.4.4) and the oversampling setting (see chapter 5.4.5):

- When the IIR filter is enabled, the temperature resolution is 20 bit.
- When the IIR filter is disabled, the temperature resolution is 16 + (osrs_t – 1) bit, e.g. 18 bit when osrs_t is set to '3'.

### 3.4.4 IIR Filter

The humidity value inside the sensor does not fluctuate rapidly and does not require low pass filtering. However, the environmental pressure is subject to many short-term changes, caused e.g. by slamming of a door or window, or wind blowing into the sensor. To suppress these disturbances in the output data without causing additional interface traffic and processor work load, the BME280 features an internal IIR filter. It effectively reduces the bandwidth of the temperature and pressure output signals and increases the resolution of the pressure and temperature output data to 20 bit.

The output of a next measurement step is filtered using the following formula:

```
data_filtered = (data_filtered_old × (filter_coefficient − 1) + data_ADC) / filter_coefficient
```

Where:
- `data_filtered_old` is the data coming from the current filter memory
- `data_ADC` is the data coming from current ADC acquisition
- `data_filtered` is the new value of filter memory and the value that will be sent to the output registers

The IIR filter can be configured to different filter coefficients, which slows down the response to the sensor inputs. Note that the response time with enabled IIR filter depends on the number of samples generated, which means that the data output rate must be known to calculate the actual response time. For register configuration, please refer to Table 28. A sample response time calculation is shown in chapter 9.4.

**Table 6: Filter settings**

| Filter coefficient | Samples to reach ≥75% of step response |
|-------------------|----------------------------------------|
| Filter off | 1 |
| 2 | 2 |
| 4 | 5 |
| 8 | 11 |
| 16 | 22 |

When writing to the register filter, the filter is reset. The next ADC values will pass through the filter unchanged and become the initial memory values for the filter. If temperature or pressure measurements are skipped, the corresponding filter memory will be kept unchanged even though the output registers are set to 0x80000. When the previously skipped measurement is re-enabled, the output will be filtered using the filter memory from the last time when the measurement was not skipped. If this is not desired, please write to the filter register in order to re-initialize the filter.

> Note: Since the BME280 does not sample continuously, filtering can suffer from signals with a frequency higher than the sampling rate of the sensor. E.g. environmental fluctuations caused by windows being opened and closed might have a frequency <5 Hz. Consequently, a sampling rate of ODR = 10 Hz is sufficient to obey the Nyquist theorem.

## 3.5 Recommended Modes of Operation

The different oversampling options, filter settings and sensor modes result in a large number of possible settings. In this chapter, a number of settings recommended for various scenarios are presented.

### 3.5.1 Weather Monitoring

**Description:** Only a very low data rate is needed. Power consumption is minimal. Noise of pressure values is of no concern. Humidity, pressure and temperature are monitored.

**Table 7: Settings and performance for weather monitoring**

| Parameter | Value |
|-----------|-------|
| Sensor mode | forced mode, 1 sample / minute |
| Oversampling settings | pressure ×1, temperature ×1, humidity ×1 |
| IIR filter settings | filter off |
| **Performance** | |
| Current consumption | 0.16 µA |
| RMS Noise | 3.3 Pa / 30 cm, 0.07 %RH |
| Data output rate | 1/60 Hz |

### 3.5.2 Humidity Sensing

**Description:** A low data rate is needed. Power consumption is minimal. Forced mode is used to minimize power consumption and to synchronize readout, but using normal mode would also be possible.

**Table 8: Settings and performance for humidity sensing**

| Parameter | Value |
|-----------|-------|
| Sensor mode | forced mode, 1 sample / second |
| Oversampling settings | pressure ×0, temperature ×1, humidity ×1 |
| IIR filter settings | filter off |
| **Performance** | |
| Current consumption | 2.9 µA |
| RMS Noise | 0.07 %RH |
| Data output rate | 1 Hz |

### 3.5.3 Indoor Navigation

**Description:** Lowest possible altitude noise is needed. A very low bandwidth is preferred. Increased power consumption is tolerated. Humidity is measured to help detect room changes. This setting is suggested for the Android settings 'SENSOR_DELAY_NORMAL' and 'SENSOR_DELAY_UI'.

**Table 9: Settings and performance for indoor navigation**

| Parameter | Value |
|-----------|-------|
| Sensor mode | normal mode, tstandby = 0.5 ms |
| Oversampling settings | pressure ×16, temperature ×2, humidity ×1 |
| IIR filter settings | filter coefficient 16 |
| **Performance** | |
| Current consumption | 633 µA |
| RMS Noise | 0.2 Pa / 1.7 cm |
| Data output rate | 25 Hz |
| Filter bandwidth | 0.53 Hz |
| Response time (75%) | 0.9 s |

### 3.5.4 Gaming

**Description:** Low altitude noise is needed. The required bandwidth is ~2 Hz in order to respond quickly to altitude changes (e.g. be able to dodge a flying monster in a game). Increased power consumption is tolerated. Humidity sensor is disabled. This setting is suggested for the Android settings 'SENSOR_DELAY_GAMING' and 'SENSOR_DELAY_FASTEST'.

**Table 10: Settings and performance for gaming**

| Parameter | Value |
|-----------|-------|
| Sensor mode | normal mode, tstandby = 0.5 ms |
| Oversampling settings | pressure ×4, temperature ×1, humidity ×0 |
| IIR filter settings | filter coefficient 16 |
| **Performance** | |
| Current consumption | 581 µA |
| RMS Noise | 0.3 Pa / 2.5 cm |
| Data output rate | 83 Hz |
| Filter bandwidth | 1.75 Hz |
| Response time (75%) | 0.3 s |

## 3.6 Noise

The noise depends on the oversampling and, for pressure and temperature, on the filter setting used. The stated values were determined in a controlled environment and are based on the average standard deviation of 32 consecutive measurement points taken at highest sampling speed. This is needed in order to exclude long term drifts from the noise measurement. The noise depends both on humidity/pressure oversampling and temperature oversampling, since the temperature value is used for humidity/pressure temperature compensation. The oversampling combinations used below result in an optimal power to noise ratio.

**Table 11: Noise and current for humidity**

| Humidity / temperature oversampling setting | Typical RMS noise in humidity [%RH] at 25 °C | Typ. current [µA] at 1 Hz forced mode, 25 °C, humidity and temperature measurement, incl. IDDSM |
|---------------------------------------------|----------------------------------------------|--------------------------------------------------------------------------------------------------|
| ×1 / ×1 | 0.07 | 1.8 |
| ×2 / ×1 | 0.05 | 2.5 |
| ×4 / ×1 | 0.04 | 3.8 |
| ×8 / ×1 | 0.03 | 6.5 |
| ×16 / ×1 | 0.02 | 11.7 |

**Table 12: Noise and current for pressure**

Typical RMS noise in pressure [Pa] at 25 °C:

| Pressure / temperature oversampling setting | IIR filter off | IIR filter 2 | IIR filter 4 | IIR filter 8 | IIR filter 16 | Typ. current [µA] at 1 Hz forced mode, 25 °C, pressure and temp, incl. IDDSM |
|---------------------------------------------|---------------|--------------|--------------|--------------|---------------|------------------------------------------------------------------------------|
| ×1 / ×1 | 3.3 | 1.9 | 1.2 | 0.9 | 0.4 | 2.8 |
| ×2 / ×1 | 2.6 | 1.5 | 1.0 | 0.6 | 0.4 | 4.2 |
| ×4 / ×1 | 2.1 | 1.2 | 0.8 | 0.5 | 0.3 | 7.1 |
| ×8 / ×1 | 1.6 | 1.0 | 0.6 | 0.4 | 0.2 | 12.8 |
| ×16 / ×2 | 1.3 | 0.8 | 0.5 | 0.4 | 0.2 | 24.9 |

**Table 13: Temperature dependence of pressure noise**

| Temperature | Typical change in noise compared to 25 °C |
|-------------|-------------------------------------------|
| -10 °C | +25 % |
| 25 °C | ±0 % |
| 75 °C | -5 % |

**Table 14: Noise in temperature**

| Temperature oversampling setting | Typical RMS noise in temperature [°C] at 25 °C |
|---------------------------------|------------------------------------------------|
| ×1 | 0.005 |
| ×2 | 0.004 |
| ×4 | 0.003 |
| ×8 | 0.003 |
| ×16 | 0.002 |
