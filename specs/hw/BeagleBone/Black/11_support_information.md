# Chapter 11: Support Information

> Source: https://docs.beagle.cc/boards/beaglebone/black/ch11.html

## Support Channels

All support for BeagleBone Black design is through the BeagleBoard.org community:
- **Forum:** https://forum.beagleboard.org/
- **Discord:** https://bbb.io/discord

## Hardware Design Files

Available on the eMMC under `documents/hardware/` when connected via USB:

| File | Format |
|------|--------|
| Schematic | PDF |
| Schematic | OrCAD (Cadence Design Entry CIS 16.3) |
| PCB Gerber files | Gerber |
| PCB Layout | Allegro |
| Bill of Material | — |
| System Reference Manual | — |

Latest files: http://www.beagleboard.org/distros

## Software Updates

Always use the latest software. Update instructions: https://elinux.org/BeagleBoneBlack

**Production Board Boot Media:** BeagleBone Black Rev C3a boot media available for download.

## RMA Support

For defective boards: http://beagleboard.org/support/rma

Required information:
- Serial number
- Board revision

Serial number location varies by manufacturing date (three documented locations).

## Troubleshooting HDMI

Reference: http://www.elinux.org/Beagleboard:BeagleBoneBlack_HDMI

### Common Issues

| Issue | Solution |
|-------|----------|
| No display | Check EDID — bad/loose cable, unpowered display |
| Wrong resolution | Display input source not selected; try disconnecting other inputs |
| Image cut off | Overscan enabled on TV — disable or use Normal/Fit mode |
| Display goes blank | Power-save mode — press key/mouse or reset board |
| No audio | Audio only works on CEA/TV resolutions; use HDMI TV with speakers |

### EDID

EDID allows the board to request display capabilities and select the highest compatible resolution. Possible causes of EDID failure:
- Bad or improperly connected cable
- Unpowered display during boot

## Audio

Audio functions only on television (CEA) resolutions. Software defaults to TV resolution providing audio support. Displays require built-in speakers or external speaker connections.
