# BME280 — Section 4: Data Readout

To read out data after a conversion, it is strongly recommended to use a burst read and not address every register individually. This will prevent a possible mix-up of bytes belonging to different measurements and reduce interface traffic. Note that in I²C mode, even when pressure was not measured, reading the unused registers is faster than reading temperature and humidity data separately.

Data readout is done by starting a burst read from 0xF7 to 0xFC (temperature and pressure) or from 0xF7 to 0xFE (temperature, pressure and humidity). The data are read out in an unsigned 20-bit format both for pressure and for temperature and in an unsigned 16-bit format for humidity. It is strongly recommended to use the BME280 API, available from Bosch Sensortec, for readout and compensation. For details on memory map and interfaces, please consult chapters 5 and 6 respectively.

After the uncompensated values for pressure, temperature and humidity 'ut', 'up' and 'uh' have been read, the actual humidity, pressure and temperature needs to be calculated using the compensation parameters stored in the device. The procedure is elaborated in chapter 4.2.

## 4.1 Data Register Shadowing

In normal mode, the timing of measurements is not necessarily synchronized to the readout by the user. This means that new measurement results may become available while the user is reading the results from the previous measurement. In this case, shadowing is performed in order to guarantee data consistency. Shadowing will only work if all data registers are read in a single burst read. Therefore, the user must use burst reads if he does not synchronize data readout with the measurement cycle. Using several independent read commands may result in inconsistent data.

If a new measurement is finished and the data registers are still being read, the new measurement results are transferred into shadow data registers. The content of shadow registers is transferred into data registers as soon as the user ends the burst read, even if not all data registers were read.

The end of the burst read is marked by the rising edge of CSB pin in SPI case or by the recognition of a stop condition in I2C case. After the end of the burst read, all user data registers are updated at once.

## 4.2 Output Compensation

The BME280 output consists of the ADC output values. However, each sensing element behaves differently. Therefore, the actual pressure and temperature must be calculated using a set of calibration parameters. In this chapter, the method to read out the trimming values will be given. The recommended calculation uses fixed point arithmetic and is given in chapter 4.2.3.

In high-level languages like Matlab™ or LabVIEW™, fixed-point code may not be well supported. In this case the floating-point code in appendix 8.1 can be used as an alternative.

For 8-bit micro controllers, the variable size may be limited. In this case a simplified 32 bit integer code with reduced accuracy is given in appendix 8.2.

### 4.2.1 Computational Requirements

The table below gives an overview of the number of clock cycles needed for compensation on a 32 bit Cortex-M3 micro controller with GCC optimization level -O2. This controller does not feature a floating point unit, thus all floating-point calculations are emulated. Floating point is only recommended for PC application, where an FPU is present and these calculations are performed drastically faster.

**Table 15: Computational requirements for compensation formulas**

| Compensation of | 32 bit integer (ARM Cortex-M3) | 64 bit integer (ARM Cortex-M3) | Double precision (ARM Cortex-M3) |
|----------------|-------------------------------|-------------------------------|----------------------------------|
| Humidity | ~83 | — | ~2400 ¹¹ |
| Temperature | ~46 | — | ~5400 ¹¹ |
| Pressure | ~112 ¹² | ~1400 | ~2900 ¹¹ |

¹¹ Use only recommended for high-level programming languages like Matlab™ or LabVIEW™

¹² Use only recommended for 8-bit micro controllers

### 4.2.2 Trimming Parameter Readout

The trimming parameters are programmed into the devices' non-volatile memory (NVM) during production and cannot be altered by the customer. Each compensation word is a 16-bit signed or unsigned integer value stored in two's complement. As the memory is organized into 8-bit words, two words must always be combined in order to represent the compensation word. The 8-bit registers are named calib00…calib41 and are stored at memory addresses 0x88…0xA1 and 0xE1…0xE7. The corresponding compensation words are named dig_T# for temperature compensation related values, dig_P# for pressure related values and dig_H# for humidity related values. The mapping is seen in Table 16.

**Table 16: Compensation parameter storage, naming and data type**

| Register Address | Register content | Data type |
|-----------------|-----------------|-----------|
| 0x88 / 0x89 | dig_T1 [7:0] / [15:8] | unsigned short |
| 0x8A / 0x8B | dig_T2 [7:0] / [15:8] | signed short |
| 0x8C / 0x8D | dig_T3 [7:0] / [15:8] | signed short |
| 0x8E / 0x8F | dig_P1 [7:0] / [15:8] | unsigned short |
| 0x90 / 0x91 | dig_P2 [7:0] / [15:8] | signed short |
| 0x92 / 0x93 | dig_P3 [7:0] / [15:8] | signed short |
| 0x94 / 0x95 | dig_P4 [7:0] / [15:8] | signed short |
| 0x96 / 0x97 | dig_P5 [7:0] / [15:8] | signed short |
| 0x98 / 0x99 | dig_P6 [7:0] / [15:8] | signed short |
| 0x9A / 0x9B | dig_P7 [7:0] / [15:8] | signed short |
| 0x9C / 0x9D | dig_P8 [7:0] / [15:8] | signed short |
| 0x9E / 0x9F | dig_P9 [7:0] / [15:8] | signed short |
| 0xA1 | dig_H1 [7:0] | unsigned char |
| 0xE1 / 0xE2 | dig_H2 [7:0] / [15:8] | signed short |
| 0xE3 | dig_H3 [7:0] | unsigned char |
| 0xE4 / 0xE5[3:0] | dig_H4 [11:4] / [3:0] | signed short |
| 0xE5[7:4] / 0xE6 | dig_H5 [3:0] / [11:4] | signed short |
| 0xE7 | dig_H6 | signed char |

### 4.2.3 Compensation Formulas

Please note that it is strongly advised to use the API available from Bosch Sensortec to perform readout and compensation. If this is not wanted, the code below can be applied at the user's risk. Both pressure and temperature values are expected to be received in 20 bit format, positive, stored in a 32 bit signed integer. Humidity is expected to be received in 16 bit format, positive, stored in a 32 bit signed integer.

The variable `t_fine` (signed 32 bit) carries a fine resolution temperature value over to the pressure and humidity compensation formula and could be implemented as a global variable.

- `BME280_S32_t` — 32 bit signed integer variable type (typically `long signed int`)
- `BME280_U32_t` — 32 bit unsigned integer variable type (typically `long unsigned int`)
- `BME280_S64_t` — 64 bit signed integer variable type (typically `long long signed int`)

For best possible calculation accuracy in pressure, 64 bit integer support is needed. If this is not possible on your platform, please see appendix 8.2 for a 32 bit alternative.

Code revision: rev. 1.1

#### Temperature Compensation (32-bit integer)

```c
// Returns temperature in DegC, resolution is 0.01 DegC. Output value of "5123" equals 51.23 DegC.
// t_fine carries fine temperature as global value
BME280_S32_t t_fine;
BME280_S32_t BME280_compensate_T_int32(BME280_S32_t adc_T)
{
  BME280_S32_t var1, var2, T;
  var1  = ((((adc_T>>3) – ((BME280_S32_t)dig_T1<<1))) * ((BME280_S32_t)dig_T2)) >> 11;
  var2  = (((((adc_T>>4) – ((BME280_S32_t)dig_T1)) * ((adc_T>>4) – ((BME280_S32_t)dig_T1)))
    >> 12) *
    ((BME280_S32_t)dig_T3)) >> 14;
  t_fine = var1 + var2;
  T  = (t_fine * 5 + 128) >> 8;
  return T;
}
```

#### Pressure Compensation (64-bit integer)

```c
// Returns pressure in Pa as unsigned 32 bit integer in Q24.8 format (24 integer bits and 8 fractional bits).
// Output value of "24674867" represents 24674867/256 = 96386.2 Pa = 963.862 hPa
BME280_U32_t BME280_compensate_P_int64(BME280_S32_t adc_P)
{
  BME280_S64_t var1, var2, p;
  var1 = ((BME280_S64_t)t_fine) – 128000;
  var2 = var1 * var1 * (BME280_S64_t)dig_P6;
  var2 = var2 + ((var1*(BME280_S64_t)dig_P5)<<17);
  var2 = var2 + (((BME280_S64_t)dig_P4)<<35);
  var1 = ((var1 * var1 * (BME280_S64_t)dig_P3)>>8) + ((var1 * (BME280_S64_t)dig_P2)<<12);
  var1 = (((((BME280_S64_t)1)<<47)+var1))*((BME280_S64_t)dig_P1)>>33;
  if (var1 == 0)
  {
    return 0; // avoid exception caused by division by zero
  }
  p = 1048576-adc_P;
  p = (((p<<31)-var2)*3125)/var1;
  var1 = (((BME280_S64_t)dig_P9) * (p>>13) * (p>>13)) >> 25;
  var2 = (((BME280_S64_t)dig_P8) * p) >> 19;
  p = ((p + var1 + var2) >> 8) + (((BME280_S64_t)dig_P7)<<4);
  return (BME280_U32_t)p;
}
```

#### Humidity Compensation (32-bit integer)

```c
// Returns humidity in %RH as unsigned 32 bit integer in Q22.10 format (22 integer and 10 fractional bits).
// Output value of "47445" represents 47445/1024 = 46.333 %RH
BME280_U32_t bme280_compensate_H_int32(BME280_S32_t adc_H)
{
  BME280_S32_t v_x1_u32r;

  v_x1_u32r = (t_fine – ((BME280_S32_t)76800));

  v_x1_u32r = (((((adc_H << 14) – (((BME280_S32_t)dig_H4) << 20) – (((BME280_S32_t)dig_H5) *
    v_x1_u32r)) + ((BME280_S32_t)16384)) >> 15) * (((((((v_x1_u32r *
    ((BME280_S32_t)dig_H6)) >> 10) * (((v_x1_u32r * ((BME280_S32_t)dig_H3)) >> 11) +
    ((BME280_S32_t)32768))) >> 10) + ((BME280_S32_t)2097152)) * ((BME280_S32_t)dig_H2) +
    8192) >> 14));

  v_x1_u32r = (v_x1_u32r – (((((v_x1_u32r >> 15) * (v_x1_u32r >> 15)) >> 7) *
    ((BME280_S32_t)dig_H1)) >> 4));

  v_x1_u32r = (v_x1_u32r < 0 ? 0 : v_x1_u32r);
  v_x1_u32r = (v_x1_u32r > 419430400 ? 419430400 : v_x1_u32r);
  return (BME280_U32_t)(v_x1_u32r>>12);
}
```
