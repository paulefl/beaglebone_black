# BME280 — Section 6: Digital Interfaces

The BME280 supports the I²C and SPI digital interfaces; it acts as a slave for both protocols. The I²C interface supports the Standard, Fast and High Speed modes. The SPI interface supports both SPI mode '00' (CPOL = CPHA = '0') and mode '11' (CPOL = CPHA = '1') in 4-wire and 3-wire configuration.

The following transactions are supported:

- Single byte write
- Multiple byte write (using pairs of register addresses and register data)
- Single byte read
- Multiple byte read (using a single register address which is auto-incremented)

## 6.1 Interface Selection

Interface selection is done automatically based on CSB (chip select) status.

- If CSB is connected to VDDIO, the I²C interface is active.
- If CSB is pulled down, the SPI interface is activated.

After CSB has been pulled down once (regardless of whether any clock cycle occurred), the I²C interface is disabled until the next power-on-reset. This is done in order to avoid inadvertently decoding SPI traffic to another slave as I²C data. Since the device startup is deferred until both VDD and VDDIO are established, there is no risk of incorrect protocol detection because of the power-up sequence used.

However, if I²C is to be used and CSB is not directly connected to VDDIO but is instead connected to a programmable pin, it must be ensured that this pin already outputs the VDDIO level during power-on-reset of the device. If this is not the case, the device will be locked in SPI mode and not respond to I²C commands.

## 6.2 I²C Interface

The I²C slave interface is compatible with Philips I²C Specification version 2.1. For detailed timings, please review Table 33. All modes (standard, fast, high speed) are supported. SDA and SCL are not pure open-drain. Both pads contain ESD protection diodes to VDDIO and GND. As the device does not perform clock stretching, the SCL structure is a high-Z input without drain capability.

The **7-bit device address** is `111011x`. The 6 MSB bits are fixed. The last bit is changeable by SDO value and can be changed during operation:

- Connecting SDO to GND → slave address **1110110 (0x76)**
- Connecting SDO to VDDIO → slave address **1110111 (0x77)** (same as BMP280's I²C address)

The SDO pin cannot be left floating; if left floating, the I²C address will be undefined.

**I²C interface pins:**
- **SCK:** serial clock (SCL)
- **SDI:** data (SDA)
- **SDO:** Slave address LSB (GND = '0', VDDIO = '1')

CSB must be connected to VDDIO to select I²C interface. SDI is bi-directional with open drain to GND: it must be externally connected to VDDIO via a pull up resistor. Refer to chapter 7 for connection instructions.

**Abbreviations used in I²C protocol:**
- **S** — Start
- **P** — Stop
- **ACKS** — Acknowledge by slave
- **ACKM** — Acknowledge by master
- **NACKM** — Not acknowledge by master

### 6.2.1 I²C Write

Writing is done by sending the slave address in write mode (RW = '0'), resulting in slave address `111011X0` ('X' is determined by state of SDO pin). Then the master sends pairs of register addresses and register data. The transaction is ended by a stop condition.

**I²C multiple byte write (not auto-incremented):**
```
S | Slave Address (1110 11X 0) | ACKS | Reg Addr (A0h) | ACKS | Data A0h | ACKS | ... | Reg Addr (A1h) | ACKS | Data A1h | ACKS | P
```

### 6.2.2 I²C Read

To be able to read registers, first the register address must be sent in write mode (slave address `111011X0`). Then either a stop or a repeated start condition must be generated. After this the slave is addressed in read mode (RW = '1') at address `111011X1`, after which the slave sends out data from auto-incremented register addresses until a NOACKM and stop condition occurs.

**I²C multiple byte read (example: registers 0xF6 and 0xF7):**
```
S | Slave Addr (1110 11X 0) | ACKS | Reg Addr (F6h) | ACKS
S | Slave Addr (1110 11X 1) | ACKS | Data F6h | ACKM | Data F7h | NOACKM | P
```

## 6.3 SPI Interface

The SPI interface is compatible with SPI mode '00' (CPOL = CPHA = '0') and mode '11' (CPOL = CPHA = '1'). The automatic selection between mode '00' and '11' is determined by the value of SCK after the CSB falling edge.

The SPI interface has two modes: 4-wire and 3-wire. The protocol is the same for both. The 3-wire mode is selected by setting '1' to the register `spi3w_en`. The pad SDI is used as a data pad in 3-wire mode.

**SPI interface pins:**
- **CSB:** chip select, active low
- **SCK:** serial clock
- **SDI:** serial data input; data input/output in 3-wire mode
- **SDO:** serial data output; hi-Z in 3-wire mode

Refer to chapter 7 for connection instructions.

CSB is active low and has an integrated pull-up resistor. Data on SDI is latched by the device at SCK rising edge and SDO is changed at SCK falling edge. Communication starts when CSB goes to low and stops when CSB goes to high; during these transitions on CSB, SCK must be stable.

**SPI protocol frame (mode '11', 4-wire):**
```
CSB:  ──┐                                          ┌──
        └──────────────────────────────────────────┘
SCK:         _   _   _   _   _   _   _   _
            | | | | | | | | | | | | | | | |
SDI:  RW AD6 AD5 AD4 AD3 AD2 AD1 AD0 | DI7 DI6 DI5 DI4 DI3 DI2 DI1 DI0
SDO:  tri-state                       | DO7 DO6 DO5 DO4 DO3 DO2 DO1 DO0
```

In SPI mode, only 7 bits of the register addresses are used; the MSB of register address is not used and replaced by a read/write bit (RW = '0' for write and RW = '1' for read).

**Example:** address 0xF7 is accessed by using SPI register address 0x77. For write access, the byte 0x77 is transferred; for read access, the byte 0xF7 is transferred.

### 6.3.1 SPI Write

Writing is done by lowering CSB and sending pairs of control bytes and register data. The control bytes consist of the SPI register address (= full register address without bit 7) and the write command (bit7 = RW = '0'). Several pairs can be written without raising CSB. The transaction is ended by raising CSB.

**SPI multiple byte write (not auto-incremented):**
```
CSB=0 | RW=0 Reg Addr (F4h) | Data F4h | RW=0 Reg Addr (F5h) | Data F5h | CSB=1
```

### 6.3.2 SPI Read

Reading is done by lowering CSB and first sending one control byte. The control byte consists of the SPI register address (= full register address without bit 7) and the read command (bit 7 = RW = '1'). After writing the control byte, data is sent out of the SDO pin (SDI in 3-wire mode); the register address is automatically incremented.

**SPI multiple byte read (example: registers 0xF6 and 0xF7):**
```
CSB=0 | RW=1 Reg Addr (F6h) | Data F6h | Data F7h | CSB=1
```

## 6.4 Interface Parameter Specification

### 6.4.1 General Interface Parameters

**Table 32: Interface parameters**

| Parameter | Symbol | Condition | Min | Typ | Max | Unit |
|-----------|--------|-----------|-----|-----|-----|------|
| Input low level | Vil_si | VDDIO=1.2 V to 3.6 V | | | 80 | %VDDIO |
| Input high level | Vih_si | VDDIO=1.2 V to 3.6 V | 80 | | | %VDDIO |
| Output low level I2C | Vol_SDI | VDDIO=1.62 V, Iol=3 mA | | | 20 | %VDDIO |
| Output low level I2C | Vol_SDI_1.2 | VDDIO=1.20 V, Iol=3 mA | | | 20 | %VDDIO |
| Output low level SPI | Vol_SDO | VDDIO=1.62 V, Iol=1 mA | | | 20 | %VDDIO |
| Output low level SPI | Vol_SDO_1.2 | VDDIO=1.20 V, Iol=1 mA | | | 20 | %VDDIO |
| Output high level | Voh | VDDIO=1.62 V, Ioh=1 mA (SDO, SDI) | 80 | | | %VDDIO |
| Output high level | Voh_1.2 | VDDIO=1.20 V, Ioh=1 mA (SDO, SDI) | 60 | | | %VDDIO |
| Pull-up resistor | Rpull | Internal CSB pull-up resistance to VDDIO | | 23 | | kΩ (70…120 typ range) |
| I2C bus load capacitor | Cb | On SDI and SCK | | | 400 | pF |

Note: The pull-up resistor Rpull has a typical range of 70…120 kΩ (190 kΩ max).

### 6.4.2 I²C Timings

For I²C timings, the following abbreviations are used:
- "S&F mode" = standard and fast mode
- "HS mode" = high speed mode
- Cb = bus capacitance on SDA line

All other naming refers to I²C specification 2.1 (January 2000).

**Table 33: I²C timings**

| Parameter | Symbol | Condition | Min | Typ | Max | Unit |
|-----------|--------|-----------|-----|-----|-----|------|
| SDI setup time | tSU;DAT | S&F Mode | 160 | | 115 | ns |
| SDI setup time | tSU;DAT | HS mode | 30 | | 150 | ns |
| SDI hold time | tHD;DAT | S&F Mode, Cb≤100 pF | 80 | | | ns |
| SDI hold time | tHD;DAT | S&F Mode, Cb≤400 pF | 90 | | | ns |
| SDI hold time | tHD;DAT | HS mode, Cb≤100 pF | 18 | | | ns |
| SDI hold time | tHD;DAT | HS mode, Cb≤400 pF | 24 | | | ns |
| SCK low pulse | tLOW | HS mode, Cb≤100 pF, VDDIO = 1.62 V | 160 | | | ns |
| SCK low pulse | tLOW | HS mode, Cb≤100 pF, VDDIO = 1.2 V | 210 | | | ns |

**Internal added delays:**
- Input delay between SDI and SCK inputs: SDI is more delayed than SCK by typically 100 ns in Standard and Fast Modes and by typically 20 ns in High Speed Mode.
- Output delay from SCK falling edge to SDI output propagation is typically 140 ns in Standard and Fast Modes and typically 70 ns in High Speed Mode.

### 6.4.3 SPI Timings

All timings apply both to 4- and 3-wire SPI.

**Table 34: SPI timings**

| Parameter | Symbol | Condition | Min | Typ | Max | Unit |
|-----------|--------|-----------|-----|-----|-----|------|
| SPI clock input frequency | F_spi | | 0 | | 10 | MHz |
| SCK low pulse | T_low_sck | | 20 | | | ns |
| SCK high pulse | T_high_sck | | 20 | | | ns |
| SDI setup time | T_setup_sdi | | 20 | | | ns |
| SDI hold time | T_hold_sdi | | 20 | | | ns |
| SDO output delay | T_delay_sdo | 25 pF load, VDDIO=1.6 V min | | 20 | 30 | ns |
| SDO output delay | T_delay_sdo | 25 pF load, VDDIO=1.2 V min | | 20 | 40 | ns |
| CSB setup time | T_setup_csb | | 20 | | | ns |
| CSB hold time | T_hold_csb | | 10 | | | ns |
