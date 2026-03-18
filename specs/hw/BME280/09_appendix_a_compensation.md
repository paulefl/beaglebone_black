# BME280 — Section 8: Appendix A — Alternative Compensation Formulas

Please note that it is strongly advised to use the API available from Bosch Sensortec to perform readout and compensation. If this is not wanted, the code below can be applied at the user's risk.

## 8.1 Compensation Formulas in Double Precision Floating Point

Both pressure and temperature values are expected to be received in 20 bit format, positive, stored in a 32 bit signed integer. Humidity is expected to be received in 16 bit format, positive, stored in a 32 bit signed integer.

The variable `t_fine` (signed 32 bit) carries a fine resolution temperature value over to the pressure compensation formula and could be implemented as a global variable.

- `BME280_S32_t` — 32 bit signed integer variable type (typically `long signed int`)

Compensating the measurement value with double precision gives the best possible accuracy but is only recommended for PC applications.

Code revision: rev. 1.1 (pressure and temperature), rev. 1.0 (humidity).

### Temperature Compensation (double precision)

```c
// Returns temperature in DegC, double precision. Output value of "51.23" equals 51.23 DegC.
// t_fine carries fine temperature as global value
BME280_S32_t t_fine;
double BME280_compensate_T_double(BME280_S32_t adc_T)
{
  double var1, var2, T;
  var1  = (((double)adc_T)/16384.0 – ((double)dig_T1)/1024.0) * ((double)dig_T2);
  var2  = ((((double)adc_T)/131072.0 – ((double)dig_T1)/8192.0) *
    (((double)adc_T)/131072.0 – ((double) dig_T1)/8192.0)) * ((double)dig_T3);
  t_fine = (BME280_S32_t)(var1 + var2);
  T  = (var1 + var2) / 5120.0;
  return T;
}
```

### Pressure Compensation (double precision)

```c
// Returns pressure in Pa as double. Output value of "96386.2" equals 96386.2 Pa = 963.862 hPa
double BME280_compensate_P_double(BME280_S32_t adc_P)
{
  double var1, var2, p;
  var1 = ((double)t_fine/2.0) – 64000.0;
  var2 = var1 * var1 * ((double)dig_P6) / 32768.0;
  var2 = var2 + var1 * ((double)dig_P5) * 2.0;
  var2 = (var2/4.0)+(((double)dig_P4) * 65536.0);
  var1 = (((double)dig_P3) * var1 * var1 / 524288.0 + ((double)dig_P2) * var1) / 524288.0;
  var1 = (1.0 + var1 / 32768.0)*((double)dig_P1);
  if (var1 == 0.0)
  {
    return 0; // avoid exception caused by division by zero
  }
  p = 1048576.0 – (double)adc_P;
  p = (p – (var2 / 4096.0)) * 6250.0 / var1;
  var1 = ((double)dig_P9) * p * p / 2147483648.0;
  var2 = p * ((double)dig_P8) / 32768.0;
  p = p + (var1 + var2 + ((double)dig_P7)) / 16.0;
  return p;
}
```

### Humidity Compensation (double precision)

```c
// Returns humidity in %rH as double. Output value of "46.332" represents 46.332 %rH
double bme280_compensate_H_double(BME280_S32_t adc_H)
{
  double var_H;

  var_H = (((double)t_fine) – 76800.0);
  var_H = (adc_H – (((double)dig_H4) * 64.0 + ((double)dig_H5) / 16384.0 *
    var_H)) * (((double)dig_H2) / 65536.0 * (1.0 + ((double)dig_H6) /
    67108864.0 * var_H *
    (1.0 + ((double)dig_H3) / 67108864.0 * var_H)));

  var_H = var_H * (1.0 – ((double)dig_H1) * var_H / 524288.0);

  if (var_H > 100.0)
    var_H = 100.0;
  else if (var_H < 0.0)
    var_H = 0.0;
  return var_H;
}
```

## 8.2 Pressure Compensation in 32-Bit Fixed Point

Please note that it is strongly advised to use the API available from Bosch Sensortec to perform readout and compensation. If this is not wanted, the code below can be applied at the user's risk. Both pressure and temperature values are expected to be received in 20 bit format, positive, stored in a 32 bit signed integer.

The variable `t_fine` (signed 32 bit) carries a fine resolution temperature value over to the pressure compensation formula and could be implemented as a global variable.

- `BME280_S32_t` — 32 bit signed integer variable type (typically `long signed int`)
- `BME280_U32_t` — 32 bit unsigned integer variable type (typically `long unsigned int`)

Compensating the pressure value with 32 bit integer has an accuracy of typically **1 Pa (1-sigma)**. At high filter levels this adds a significant amount of noise to the output values and reduces their resolution.

### Temperature Compensation (32-bit integer)

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

### Pressure Compensation (32-bit integer)

```c
// Returns pressure in Pa as unsigned 32 bit integer. Output value of "96386" equals 96386 Pa = 963.86 hPa
BME280_U32_t BME280_compensate_P_int32(BME280_S32_t adc_P)
{
  BME280_S32_t var1, var2;
  BME280_U32_t p;
  var1 = (((BME280_S32_t)t_fine)>>1) – (BME280_S32_t)64000;
  var2 = (((var1>>2) * (var1>>2)) >> 11 ) * ((BME280_S32_t)dig_P6);
  var2 = var2 + ((var1*((BME280_S32_t)dig_P5))<<1);
  var2 = (var2>>2)+(((BME280_S32_t)dig_P4)<<16);
  var1 = (((dig_P3 * (((var1>>2) * (var1>>2)) >> 13 )) >> 3) + ((((BME280_S32_t)dig_P2) *
    var1)>>1))>>18;
  var1 =((((32768+var1))*((BME280_S32_t)dig_P1))>>15);
  if (var1 == 0)
  {
    return 0; // avoid exception caused by division by zero
  }
  p = (((BME280_U32_t)(((BME280_S32_t)1048576)-adc_P)-(var2>>12)))*3125;
  if (p < 0x80000000)
  {
    p = (p << 1) / ((BME280_U32_t)var1);
  }
  else
  {
    p = (p / (BME280_U32_t)var1) * 2;
  }
  var1 = (((BME280_S32_t)dig_P9) * ((BME280_S32_t)(((p>>3) * (p>>3))>>13)))>>12;
  var2 = (((BME280_S32_t)(p>>2)) * ((BME280_S32_t)dig_P8))>>13;
  p = (BME280_U32_t)((BME280_S32_t)p + ((var1 + var2 + dig_P7) >> 4));
  return p;
}
```
