# BME280 — Section 1: Specification

If not stated otherwise:

- All values are valid over the full voltage range
- All minimum/maximum values are given for the full accuracy temperature range
- Minimum/maximum values of drifts, offsets and temperature coefficients are ±3σ values over lifetime
- Typical values of drifts, offsets and temperature coefficients are ±1σ values over lifetime
- Typical values of currents and state machine timings are determined at 25 °C
- Minimum/maximum values of currents are determined using corner lots over complete temperature range
- Minimum/maximum values of state machine timings are determined using corner lots over 0…+65 °C temperature range

The specification tables are split into humidity, pressure, and temperature part of BME280.

## 1.1 General Electrical Specification

**Table 1: Electrical parameter specification**

| Parameter | Symbol | Condition | Min | Typ | Max | Unit |
|-----------|--------|-----------|-----|-----|-----|------|
| Supply Voltage Internal Domains | VDD | ripple max. 50 mVpp | 1.71 | 1.8 | 3.6 | V |
| Supply Voltage I/O Domain | VDDIO | | 1.2 | 1.8 | 3.6 | V |
| Sleep current | IDDSL | Max value at 85 °C | | 0.1 | 0.3 | µA |
| Standby current (inactive period of normal mode) | IDDSB | | | 0.2 | 0.5 | µA |
| Current during humidity measurement | IDDH | Max value at -40 °C | | 340 | | µA |
| Current during pressure measurement | IDDP | Max value at 85 °C | | 714 | | µA |
| Current during temperature measurement | IDDT | Max value at 85 °C | | 350 | | µA |
| Start-up time | tstartup | Time to first communication after both VDD > 1.58 V and VDDIO > 0.65 V | | 2 | | ms |
| Power supply rejection ratio (DC) | PSRR | full VDD range | | ±0.01 / ±5 | | %RH/V / Pa/V |
| Standby time accuracy | Δtstandby | | | ±5 | ±25 | % |

## 1.2 Humidity Parameter Specification

**Table 2: Humidity parameter specification**

| Parameter | Symbol | Condition | Min | Typ | Max | Unit |
|-----------|--------|-----------|-----|-----|-----|------|
| Operating range ¹ | RH | | -40 | 25 | 85 | °C |
| | | | 0 | | 100 | %RH |
| Supply current | IDD,H | 1 Hz forced mode, humidity and temperature | | 1.8 | 2.8 | µA |
| Absolute accuracy tolerance | AH | 20…80 %RH, 25 °C, including hysteresis | | | ±3 | %RH |
| Hysteresis ² | HH | 10…90 %RH, 25 °C | | | ±1 | %RH |
| Nonlinearity ³ | NLH | 10…90 %RH, 25 °C | | | 1 | %RH |
| Response time to complete 63% of step ⁴ | τ63% RH | 90…0 or 0…90 %RH, 25°C | | 1 | | s |
| Resolution | | Highest oversampling, see chapter 3.6 | | 0.008 | | %RH |
| Noise in humidity (RMS) | NH | 90…0 or 0…90 %RH, 25°C | | 0.02 | | %RH |
| Long term stability | ΔHstab | 10…90 %RH, 25 °C | | | 0.5 | %RH/year |

**Notes:**

¹ When exceeding the operating range (e.g. for soldering), humidity sensing performance is temporarily degraded and reconditioning is recommended as described in section 7.8. Operating range only for non-condensing environment.

² For hysteresis measurement the sequence 10→30→50→70→90→70→50→30→10 %RH is used. The hysteresis is defined as the difference between measurements of the humidity up / down branch and the averaged curve of both branches.

³ Non-linear contributions to the sensor data are corrected during the calculation of the relative humidity by the compensation formulas described in section 4.2.3.

⁴ The air-flow in direction to the vent-hole of the device has to be dimensioned in a way that a sufficient air exchange inside to outside will be possible. To observe effects on the response time-scale of the device an air-flow velocity of approx. 1 m/s is needed.

**Figure 1: Humidity sensor operating range**

The operating range for the humidity sensor extends from 0 to 100 %RH for temperatures between -40 °C and +85 °C. For temperatures below 0 °C and above 60 °C, the operating range is reduced (see datasheet Figure 1 for the exact boundary curve).

## 1.3 Pressure Sensor Specification

**Table 3: Pressure parameter specification**

| Parameter | Symbol | Condition | Min | Typ | Max | Unit |
|-----------|--------|-----------|-----|-----|-----|------|
| Operating temperature range | TA | operational | -40 | 25 | +85 | °C |
| | | full accuracy | 0 | | +65 | °C |
| Operating pressure range | P | full accuracy | 300 | | 1100 | hPa |
| Supply current | IDD,LP | 1 Hz forced mode, pressure and temperature, lowest power | | 2.8 | 4.2 | µA |
| Temperature coefficient of offset ⁵ | TCOP | 25…65 °C, 900 hPa | | ±1.5 | | Pa/K |
| | | | | ±12.6 | | cm/K |
| Absolute accuracy pressure | AP ext | 300…1100 hPa, -20…0 °C | | | ±1.7 | hPa |
| | AP,full | 300…1100 hPa, 0…65 °C | | | ±1.0 | hPa |
| | AP | 1100…1250 hPa, 25…40 °C | | | ±1.5 | hPa |
| Relative accuracy pressure | Arel | VDD = 3.3 V, 700…900 hPa, 25…40 °C | | ±0.12 | | hPa |
| Resolution of pressure output data | RP | Highest oversampling | | 0.18 | | Pa |
| Noise in pressure | NP,fullBW | Full bandwidth, highest oversampling, see chapter 3.6 | | 1.3 | | Pa |
| | | | | 11 | | cm |
| | NP,filtered | Reduced bandwidth, highest oversampling, see chapter 3.6 | | 0.2 | | Pa |
| | | | | 1.7 | | cm |
| Solder drift | | Minimum solder height 50 µm | -0.5 | | +2.0 | hPa |
| Long term stability ⁶ | ΔPstab | per year | | | ±1.0 | hPa |
| Possible sampling rate | fsample_P | Lowest oversampling, see chapter 9.2 | | 157 | 182 | Hz |

**Notes:**

⁵ When changing temperature by e.g. 10 °C at constant pressure / altitude, the measured pressure / altitude will change by (10 × TCOP).

⁶ Long term stability is specified in the full accuracy operating pressure range 0…65 °C.

## 1.4 Temperature Sensor Specification

**Table 4: Temperature parameter specification**

| Parameter | Symbol | Condition | Min | Typ | Max | Unit |
|-----------|--------|-----------|-----|-----|-----|------|
| Operating range | T | Operational | -40 | 25 | 85 | °C |
| | | Full accuracy | 0 | | 65 | °C |
| Supply current | IDD,T | 1 Hz forced mode, temperature measurement only | | 1.0 | | µA |
| Absolute accuracy temperature ⁷ | AT,25 | 25 °C | | | ±0.5 | °C |
| | AT,full | 0…65 °C | | | ±0.5 | °C |
| | AT,ext ⁸ | -20…0 °C | | | ±1.25 | °C |
| | AT,ext ⁹ | -40…-20 °C | | | ±1.5 | °C |
| Output resolution | RT | API output resolution | | 0.01 | | °C |
| RMS noise | NT | Lowest oversampling | | 0.005 | | °C |

**Notes:**

⁷ Temperature measured by the internal temperature sensor. This temperature value depends on the PCB temperature, sensor element self-heating and ambient temperature and is typically above ambient temperature.

⁸ Target values & not guaranteed.

⁹ Target values & not guaranteed.
