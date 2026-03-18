# BME280 — Section 9: Appendix B — Measurement Time and Current Calculation

In this chapter, formulas are given to calculate measurement rate, filter bandwidth and current consumption in different settings.

## 9.1 Measurement Time

The active measurement time depends on the selected values for humidity, temperature and pressure oversampling and can be calculated in milliseconds using the formulas below.

**Typical measurement time:**

```
t_measure,typ = 1 + [2 · T_oversampling]_{osrs_t≠0}
              + [2 · P_oversampling + 0.5]_{osrs_p≠0}
              + [2 · H_oversampling + 0.5]_{osrs_h≠0}
```

**Maximum measurement time:**

```
t_measure,max = 1.25 + [2.3 · T_oversampling]_{osrs_t≠0}
              + [2.3 · P_oversampling + 0.575]_{osrs_p≠0}
              + [2.3 · H_oversampling + 0.575]_{osrs_h≠0}
```

**Example:** Using temperature oversampling ×1, pressure oversampling ×4 and no humidity measurement:

```
t_measure,typ = 1 + [2 · 1] + [2 · 4 + 0.5] + [0] = 11.5 ms
t_measure,max = 1.25 + [2.3 · 1] + [2.3 · 4 + 0.575] + [0] = 13.325 ms
```

## 9.2 Measurement Rate in Forced Mode

In forced mode, the measurement rate depends on the rate at which it is forced by the master. The highest possible frequency in Hz can be calculated as:

```
ODR_max,forced = 1000 / t_measure
```

If measurements are forced faster than they can be executed, the data rate saturates at the attainable data rate.

**Example:** With 11.5 ms measurement time:

```
ODR_max,forced = 1000 / 11.5 = 87 Hz
```

## 9.3 Measurement Rate in Normal Mode

The measurement rate in normal mode depends on the measurement time and the standby time and can be calculated in Hz using the following formula:

```
ODR_normal_mode = 1000 / (t_measure + t_standby)
```

The accuracy of tstandby is described in the specification parameter Δtstandby.

**Example:** With 11.5 ms measurement time and standby time of 62.5 ms:

```
ODR_normal_mode = 1000 / (11.5 + 62.5) = 13.51 Hz
```

## 9.4 Response Time Using IIR Filter

When using the IIR filter, the response time of the sensor depends on the selected filter coefficient and the data rate used. It can be calculated using the following formula:

```
t_response, 75% = (1000 · samples_75%) / ODR
```

Where `samples_75%` is the number of samples needed to reach 75% of a step response (see Table 6 in chapter 3.4.4).

**Example:** With a data rate of 13.51 Hz and a filter coefficient of 8 (which requires 11 samples to reach 75% of step response):

```
t_response, 75% = (1000 · 11) / 13.51 = 814 ms
```

## 9.5 Current Consumption

The current consumption depends on the selected oversampling settings, the measurement rate and the sensor mode, but **not** on the IIR filter setting.

**Forced mode current:**

```
I_DD,forced = I_DDSL · (1 − t_measure · ODR/1000)
            + (ODR/1000) · (205 + I_DDT · [2 · T_oversampling]_{osrs_t≠0}
                                 + I_DDP · [2 · P_oversampling + 0.5]_{osrs_p≠0}
                                 + I_DDH · [2 · H_oversampling + 0.5]_{osrs_h≠0})
```

**Normal mode current:**

```
I_DD,normal = I_DDSB · (1 − t_measure · ODR/1000)
            + (ODR/1000) · (205 + I_DDT · [2 · T_oversampling]_{osrs_t≠0}
                                 + I_DDP · [2 · P_oversampling + 0.5]_{osrs_p≠0}
                                 + I_DDH · [2 · H_oversampling + 0.5]_{osrs_h≠0})
```

**Note:** The only difference between forced and normal mode current consumption is that the current for the inactive time is either IDDSL (sleep mode current) or IDDSB (standby current).

**Example:** Using temperature oversampling ×1, pressure oversampling ×4, no humidity measurement, in normal mode at 13.51 Hz:

Where:
- I_DDSB = 0.2 µA (standby current)
- I_DDT = 350 µA/unit (temperature measurement current per 2 ms block)
- I_DDP = 714 µA/unit (pressure measurement current per 2 ms block)
- t_measure = 11.5 ms
- ODR = 13.51 Hz

```
I_DD,normal = 0.2 · (1 − 0.0115 · 13.51) + (13.51/1000) · (205 + 350 · [2·1] + 714 · [2·4 + 0.5] + [0])
            = 0.2 · (0.845) + (13.51/1000) · (205 + 700 + 6069 + 0)
            = 0.169 + 0.01351 · 6974
            = 0.2 + 94.2
            = 94.4 µA
```
