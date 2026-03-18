# BME280 — Section 7: Pin-out and Connection Diagram

## 7.1 Pin-out

The BME280 uses an 8-pin LGA package. The pin numbering is performed in the **clockwise direction** when seen in top view and **counter-clockwise** when seen in bottom view (which is untypical).

**Top view (pads not visible):**
```
        8(VDD)    7(GND)
                          Vent hole
 5(SDO)              6(VDDIO)
        Pin 1
        marker
 4(SCK)              3(SDI)
        1(GND)    2(CSB)
```

**Bottom view (pads visible):**
```
        8(VDD)    7(GND)
 5(SDO)              6(VDDIO)
        Pin 1
        marker
 4(SCK)              3(SDI)
        1(GND)    2(CSB)
```

**Table 35: Pin description**

| Pin | Name | I/O Type | Description | SPI 4W | SPI 3W | I²C |
|-----|------|----------|-------------|--------|--------|-----|
| 1 | GND | Supply | Ground | GND | GND | GND |
| 2 | CSB | In | Chip select | CSB | CSB | VDDIO |
| 3 | SDI | In/Out | Serial data input | SDI | SDI/SDO | SDA |
| 4 | SCK | In | Serial clock input | SCK | SCK | SCL |
| 5 | SDO | In/Out | Serial data output | SDO | DNC | GND for default address |
| 6 | VDDIO | Supply | Digital / Interface supply | VDDIO | VDDIO | VDDIO |
| 7 | GND | Supply | Ground | GND | GND | GND |
| 8 | VDD | Supply | Analog supply | VDD | VDD | VDD |

## 7.2 Connection Diagram — I²C

```
VDD         VDDIO
 |              |
 |   ┌──────────┤
 |   |    R1    |   R2
 |   |   ┌──┐  ├──┌──┐── SDA
 |   |   └──┘  |  └──┘── SCL
 |   |          |
 |  [8]VDD [7]GND [6]VDDIO
 |         BME280
 |  [1]GND [2]CSB [3]SDI [4]SCK
 |   |     |      |      |
GND  |   VDDIO   SDA    SCL
     |
     C1     C2
     |       |
    GND     GND
```

**Notes:**
- The recommended value for C1, C2 is **100 nF**
- The value for the pull-up resistors R1, R2 should be based on the interface timing and the bus load; a normal value is **4.7 kΩ**
- A direct connection between CSB and VDDIO is required
- I²C address bit 0: SDO → GND = '0' (address 0x76); SDO → VDDIO = '1' (address 0x77)

## 7.3 Connection Diagram — 4-Wire SPI

```
VDD        VDDIO
 |              |
 |  [8]VDD [7]GND [6]VDDIO
 |         BME280
 |  [1]GND [2]CSB [3]SDI [4]SCK [5]SDO
 |   |     |      |      |      |
GND  |    CSB    SDI    SCK    SDO
     C1     C2
     |       |
    GND     GND
```

**Note:** The recommended value for C1, C2 is **100 nF**

## 7.4 Connection Diagram — 3-Wire SPI

```
VDD        VDDIO
 |              |
 |  [8]VDD [7]GND [6]VDDIO
 |         BME280
 |  [1]GND [2]CSB [3]SDI [4]SCK [5]SDO
 |   |     |      |      |      |
GND  |    CSB  SDI/SDO  SCK    (DNC)
     C1     C2
     |       |
    GND     GND
```

**Note:** The recommended value for C1, C2 is **100 nF**. In 3-wire mode, SDO pin is not connected (DNC).

## 7.5 Package Dimensions

The BME280 is housed in an LGA (Land Grid Array) package. Key dimensions:

- Footprint: **2.5 mm × 2.5 mm**
- Height: **0.93 mm** (metal lid)

(Refer to Figure 20 in the original datasheet for exact dimensional drawings with tolerances for top, bottom, and side views.)

## 7.6 Landing Pattern Recommendation

For the design of the landing pattern, the following is recommended:

- Red areas demarcate exposed PCB metal pads.
- In case of a solder mask defined (SMD) PCB process, the land dimensions should be defined by solder mask openings. The underlying metal pads are larger than these openings.
- In case of a non solder mask defined (NSMD) PCB process, the land dimensions should be defined in the metal layer. The mask openings are larger than these metal pads.

(Refer to Figure 21 in the original datasheet for exact recommended landing pattern dimensions in top view.)

## 7.7 Marking

### 7.7.1 Mass Production Devices

**Table 36: Marking of mass production parts**

| Symbol | Description |
|--------|-------------|
| CCC | Lot counter: 3 alphanumeric digits, variable to generate mass production trace-code |
| T | Product number: 1 alphanumeric digit, fixed to identify product type. T = "U". "U" is associated with the product BME280 (part number 0 273 141 185) |
| L | Sub-contractor ID: 1 alphanumeric digit, variable to identify sub-contractor (L = "P") |

Marking layout (top view):
```
    Vent
    hole
  5 | 8
    | 7
    | 6
CCC T L
  4 | 3
    | 2
    | 1  ← Pin 1 marker
```

### 7.7.2 Engineering Samples

**Table 37: Marking of engineering samples**

| Symbol | Description |
|--------|-------------|
| XX | Sample ID: 2 alphanumeric digits, variable to generate trace-code |
| N | Eng. Sample ID: 1 alphanumeric digit, fixed to identify engineering sample. N = "*" or "e" or "E" |
| CC | Counter ID: 2 alphanumeric digits, variable to generate trace-code |

## 7.8 Soldering Guidelines and Reconditioning Recommendations

The moisture sensitivity level of the BME280 sensors corresponds to **JEDEC Level 1**, see also:
- IPC/JEDEC J-STD-020C "Joint Industry Standard: Moisture/Reflow Sensitivity Classification for non-hermetic Solid State Surface Mount Devices"
- IPC/JEDEC J-STD-033A "Joint Industry Standard: Handling, Packing, Shipping and Use of Moisture/Reflow Sensitive Surface Mount Devices"

The sensor fulfils the lead-free soldering requirements of the above-mentioned IPC/JEDEC standard, i.e. reflow soldering with a **peak temperature up to 260 °C**. The minimum height of the solder after reflow shall be at least **50 µm**. This is required for good mechanical decoupling between the sensor device and the printed circuit board (PCB).

## 7.9 Reconditioning Procedure

After exposing the device to operating conditions which exceed the limits specified in section 1.2, e.g. after reflow, the humidity sensor may possess an additional offset. Therefore the following reconditioning procedure is mandatory to restore the calibration state:

**Option 1:**
1. Dry-Baking: 120 °C at <5% rH for 2 h
2. Re-Hydration: 70 °C at 75% rH for 6 h

**Option 2:**
1. Dry-Baking: 120 °C at <5% rH for 2 h
2. Re-Hydration: 25 °C at 75% rH for 24 h

**Option 3 (after solder reflow only):**
1. Do not perform Dry-Baking
2. Ambient Re-Hydration: ~25 °C at >40% rH for >5 days

## 7.10 Tape and Reel Specification

### 7.10.1 Dimensions

Quantity per reel: **10 kpcs** (10,000 pieces per reel).

(Refer to Figure 23 in the original datasheet for tape and reel dimensions.)

### 7.10.2 Orientation Within the Reel

(Refer to Figure 24 in the original datasheet for the orientation of the BME280 within the tape.)

## 7.11 Mounting and Assembly Recommendations

In order to achieve the specified performance for your design, the following recommendations and the "Handling, soldering & mounting instructions BME280" should be taken into consideration when mounting a pressure sensor on a printed-circuit board (PCB):

- The clearance above the metal lid shall be **0.1 mm at minimum**.
- For the device housing appropriate venting needs to be provided in case the ambient pressure shall be measured.
- Liquids shall not come into direct contact with the device.
- During operation the sensor chip is sensitive to light, which can influence the accuracy of the measurement (photo-current of silicon). The position of the vent hole minimizes the light exposure of the sensor chip. Nevertheless, Bosch Sensortec recommends avoiding the exposure of BME280 to strong light sources.
- Soldering may not be done using vapor phase processes since the sensor will be damaged by the liquids used in these processes.

## 7.12 Environmental Safety

### 7.12.1 RoHS

The BME280 sensor meets the requirements of the EC restriction of hazardous substances (RoHS) directive:
- RoHS–Directive 2011/65/EU and its amendments, including the amendment 2015/863/EU on the restriction of the use of certain hazardous substances in electrical and electronic equipment.

### 7.12.2 Halogen Content

The BME280 is **halogen-free**. For more details on the analysis results please contact your Bosch Sensortec representative.

### 7.12.3 Internal Package Structure

Within the scope of Bosch Sensortec's ambition to improve its products and secure the mass product supply, Bosch Sensortec qualifies additional sources (e.g. 2nd source) for the package of the BME280.

While Bosch Sensortec took care that all of the technical packages parameters described above are 100% identical for all sources, there can be differences in the chemical content and the internal structure between the different package sources.

However, as secured by the extensive product qualification process of Bosch Sensortec, this has no impact to the usage or to the quality of the BME280 product.
