# BME280 — Section 5: Global Memory Map and Register Description

## 5.1 General Remarks

The entire communication with the device is performed by reading from and writing to registers. Registers have a width of 8 bits. There are several registers which are reserved; they should not be written to and no specific value is guaranteed when they are read. For details on the interface, consult chapter 6.

## 5.2 Register Compatibility to BMP280

The BME280 is downward register compatible to the BMP280, which means that the pressure and temperature control and readout is identical to BMP280. However, the following exceptions have to be considered:

**Table 17: Register incompatibilities between BMP280 and BME280**

| Register | Bits | Content | BMP280 | BME280 |
|----------|------|---------|--------|--------|
| 0xD0 "id" | 7:0 | chip_id | Read value is 0x56 / 0x57 (samples) or 0x58 (mass production) | Read value is 0x60 |
| 0xF5 "config" | 7:5 | t_sb | '110': 2000 ms; '111': 4000 ms | '110': 10 ms; '111': 20 ms |
| 0xF7…0xF9 "press" | 19:0 | press | Resolution (16…20 bit) depends only on osrs_p | Without filter, resolution depends on osrs_p; when using filter, resolution is always 20 bit |
| 0xFA…0xFC "temp" | 19:0 | temp | Resolution (16…20 bit) only depends on osrs_t | Without filter, resolution depends on osrs_t; when using filter, resolution is always 20 bit |

## 5.3 Memory Map

The memory map is given in Table 18 below. Reserved registers are not shown.

**Table 18: Memory map**

| Register Name | Address | bit7 | bit6 | bit5 | bit4 | bit3 | bit2 | bit1 | bit0 | Reset state |
|---------------|---------|------|------|------|------|------|------|------|------|-------------|
| hum_lsb | 0xFE | hum_lsb<7:0> | | | | | | | | 0x00 |
| hum_msb | 0xFD | hum_msb<7:0> | | | | | | | | 0x80 |
| temp_xlsb | 0xFC | temp_xlsb<7:4> | | | | 0 | 0 | 0 | 0 | 0x00 |
| temp_lsb | 0xFB | temp_lsb<7:0> | | | | | | | | 0x00 |
| temp_msb | 0xFA | temp_msb<7:0> | | | | | | | | 0x80 |
| press_xlsb | 0xF9 | press_xlsb<7:4> | | | | 0 | 0 | 0 | 0 | 0x00 |
| press_lsb | 0xF8 | press_lsb<7:0> | | | | | | | | 0x00 |
| press_msb | 0xF7 | press_msb<7:0> | | | | | | | | 0x80 |
| config | 0xF5 | t_sb[2:0] | | | filter[2:0] | | | spi3w_en[0] | 0x00 |
| ctrl_meas | 0xF4 | osrs_t[2:0] | | | osrs_p[2:0] | | | mode[1:0] | 0x00 |
| status | 0xF3 | | | | measuring[0] | | | | im_update[0] | 0x00 |
| ctrl_hum | 0xF2 | | | | | | osrs_h[2:0] | | 0x00 |
| calib26..calib41 | 0xE1…0xF0 | calibration data | | | | | | | | individual |
| reset | 0xE0 | reset[7:0] | | | | | | | | 0x00 |
| id | 0xD0 | chip_id[7:0] | | | | | | | | 0x60 |
| calib00..calib25 | 0x88…0xA1 | calibration data | | | | | | | | individual |

**Register types:**

- **Calibration data** (0x88…0xA1, 0xE1…0xF0): read only
- **Chip ID** (0xD0): read only
- **Reset** (0xE0): write only
- **Status registers** (0xF3): read only
- **Control registers** (0xF2, 0xF4, 0xF5): read / write
- **Data registers** (0xF7…0xFE): read only
- **Reserved registers**: do not change

## 5.4 Register Description

### 5.4.1 Register 0xD0 "id"

The "id" register contains the chip identification number `chip_id[7:0]`, which is **0x60**. This number can be read as soon as the device finished the power-on-reset.

### 5.4.2 Register 0xE0 "reset"

The "reset" register contains the soft reset word `reset[7:0]`. If the value **0xB6** is written to the register, the device is reset using the complete power-on-reset procedure. Writing other values than 0xB6 has no effect. The readout value is always 0x00.

### 5.4.3 Register 0xF2 "ctrl_hum"

The "ctrl_hum" register sets the humidity data acquisition options of the device. Changes to this register only become effective after a write operation to "ctrl_meas".

**Table 19: Register 0xF2 "ctrl_hum"**

| Register 0xF2 "ctrl_hum" | Name | Description |
|--------------------------|------|-------------|
| Bit 2, 1, 0 | osrs_h[2:0] | Controls oversampling of humidity data. See Table 20 for settings and chapter 3.4.1 for details. |

**Table 20: Register settings osrs_h**

| osrs_h[2:0] | Humidity oversampling |
|-------------|----------------------|
| 000 | Skipped (output set to 0x8000) |
| 001 | oversampling ×1 |
| 010 | oversampling ×2 |
| 011 | oversampling ×4 |
| 100 | oversampling ×8 |
| 101, others | oversampling ×16 |

### 5.4.4 Register 0xF3 "status"

The "status" register contains two bits which indicate the status of the device.

**Table 21: Register 0xF3 "status"**

| Register 0xF3 "status" | Name | Description |
|------------------------|------|-------------|
| Bit 3 | measuring[0] | Automatically set to '1' whenever a conversion is running and back to '0' when the results have been transferred to the data registers. |
| Bit 0 | im_update[0] | Automatically set to '1' when the NVM data are being copied to image registers and back to '0' when the copying is done. The data are copied at power-on-reset and before every conversion. |

### 5.4.5 Register 0xF4 "ctrl_meas"

The "ctrl_meas" register sets the pressure and temperature data acquisition options of the device. The register needs to be written after changing "ctrl_hum" for the changes to become effective.

**Table 22: Register 0xF4 "ctrl_meas"**

| Register 0xF4 "ctrl_meas" | Name | Description |
|---------------------------|------|-------------|
| Bit 7, 6, 5 | osrs_t[2:0] | Controls oversampling of temperature data. See Table 24 for settings and chapter 3.4.3 for details. |
| Bit 4, 3, 2 | osrs_p[2:0] | Controls oversampling of pressure data. See Table 23 for settings and chapter 3.4.2 for details. |
| Bit 1, 0 | mode[1:0] | Controls the sensor mode of the device. See Table 25 for settings and chapter 3.3 for details. |

**Table 23: Register settings osrs_p**

| osrs_p[2:0] | Pressure oversampling |
|-------------|----------------------|
| 000 | Skipped (output set to 0x80000) |
| 001 | oversampling ×1 |
| 010 | oversampling ×2 |
| 011 | oversampling ×4 |
| 100 | oversampling ×8 |
| 101, others | oversampling ×16 |

**Table 24: Register settings osrs_t**

| osrs_t[2:0] | Temperature oversampling |
|-------------|--------------------------|
| 000 | Skipped (output set to 0x80000) |
| 001 | oversampling ×1 |
| 010 | oversampling ×2 |
| 011 | oversampling ×4 |
| 100 | oversampling ×8 |
| 101, others | oversampling ×16 |

**Table 25: Register settings mode**

| mode[1:0] | Mode |
|-----------|------|
| 00 | Sleep mode |
| 01 and 10 | Forced mode |
| 11 | Normal mode |

### 5.4.6 Register 0xF5 "config"

The "config" register sets the rate, filter and interface options of the device. Writes to the "config" register in normal mode may be ignored. In sleep mode writes are not ignored.

**Table 26: Register 0xF5 "config"**

| Register 0xF5 "config" | Name | Description |
|------------------------|------|-------------|
| Bit 7, 6, 5 | t_sb[2:0] | Controls inactive duration tstandby in normal mode. See Table 27 for settings and chapter 3.3.4 for details. |
| Bit 4, 3, 2 | filter[2:0] | Controls the time constant of the IIR filter. See Table 28 for settings and chapter 3.4.4 for details. |
| Bit 0 | spi3w_en[0] | Enables 3-wire SPI interface when set to '1'. See chapter 6.3 for details. |

**Table 27: t_sb settings**

| t_sb[2:0] | tstandby [ms] |
|-----------|---------------|
| 000 | 0.5 |
| 001 | 62.5 |
| 010 | 125 |
| 011 | 250 |
| 100 | 500 |
| 101 | 1000 |
| 110 | 10 |
| 111 | 20 |

**Table 28: Filter settings**

| filter[2:0] | Filter coefficient |
|-------------|-------------------|
| 000 | Filter off |
| 001 | 2 |
| 010 | 4 |
| 011 | 8 |
| 100, others | 16 |

### 5.4.7 Register 0xF7…0xF9 "press" (_msb, _lsb, _xlsb)

The "press" register contains the raw pressure measurement output data up[19:0]. For details on how to read out the pressure and temperature information from the device, please consult chapter 4.

**Table 29: Register 0xF7…0xF9 "press"**

| Register 0xF7…0xF9 "press" | Name | Description |
|----------------------------|------|-------------|
| 0xF7 | press_msb[7:0] | Contains the MSB part up[19:12] of the raw pressure measurement output data. |
| 0xF8 | press_lsb[7:0] | Contains the LSB part up[11:4] of the raw pressure measurement output data. |
| 0xF9 (bit 7, 6, 5, 4) | press_xlsb[3:0] | Contains the XLSB part up[3:0] of the raw pressure measurement output data. Contents depend on temperature resolution. |

### 5.4.8 Register 0xFA…0xFC "temp" (_msb, _lsb, _xlsb)

The "temp" register contains the raw temperature measurement output data ut[19:0]. For details on how to read out the pressure and temperature information from the device, please consult chapter 4.

**Table 30: Register 0xFA…0xFC "temp"**

| Register 0xFA…0xFC "temp" | Name | Description |
|---------------------------|------|-------------|
| 0xFA | temp_msb[7:0] | Contains the MSB part ut[19:12] of the raw temperature measurement output data. |
| 0xFB | temp_lsb[7:0] | Contains the LSB part ut[11:4] of the raw temperature measurement output data. |
| 0xFC (bit 7, 6, 5, 4) | temp_xlsb[3:0] | Contains the XLSB part ut[3:0] of the raw temperature measurement output data. Contents depend on pressure resolution. |

### 5.4.9 Register 0xFD…0xFE "hum" (_msb, _lsb)

The "hum" register contains the raw humidity measurement output data uh[15:0]. For details on how to read out the humidity information from the device, please consult chapter 4.

**Table 31: Register 0xFD…0xFE "hum"**

| Register 0xFD…0xFE "hum" | Name | Description |
|--------------------------|------|-------------|
| 0xFD | hum_msb[7:0] | Contains the MSB part uh[15:8] of the raw humidity measurement output data. |
| 0xFE | hum_lsb[7:0] | Contains the LSB part uh[7:0] of the raw humidity measurement output data. |
