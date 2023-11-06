<h1 align="center">lianlinux</h1>

<p align="center">
  <a href="https://github.com/crptmem/lianlinux/stargazers"><img src="https://img.shields.io/github/stars/crptmem/lianlinux?colorA=151515&colorB=B66467&style=for-the-badge&logo=starship"></a>
  <a href="https://github.com/crptmem/lianlinux/issues"><img src="https://img.shields.io/github/issues/crptmem/lianlinux?colorA=151515&colorB=8C977D&style=for-the-badge&logo=bugatti"></a>
  <a href="https://github.com/crptmem/lianlinux/network/members"><img src="https://img.shields.io/github/forks/crptmem/lianlinux?colorA=151515&colorB=D9BC8C&style=for-the-badge&logo=github"></a>
</p>

> 🌈 An app to control Lian Li Hub lights on Linux

> [!IMPORTANT]  
> I only have **LianLi-UNI FAN-SL-v1.8** hub, so I can't support other devices. Contributions are appreciated!

# Install
```sh
$ git clone https://github.com/crptmem/lianlinux.git
$ cd lianlinux
$ go install
```

# Usage
```sh
$ lianlinux
```

# Supported modes
`rainbow`, `morph`, `static`
 
# Supported devices

| Product name | Product ID |
|----------|:-------------:|
| LianLi-UNI FAN-SL-v1.8 | 0xa100 |

# Lian Li USB protocol
## Packets structure
### Color mode switch packet:
| Start (?) | Port ID | First mode byte | Second mode byte |
|----------|:-------------:|:-------------:|:-------------:|
| E0 | 10-14 | 00-? | 00-? |

### Command packet:
| Start (?) | Channel | Command | Argument |
|----------|:-------------:|:-------------:|:-------------:|
| E0 | 10-14 | 10-24 | 00-31 | 00-FF |

## Known bytes
`BaseByte` (start byte) = 0xE0 <br />
`DefaultChannelByte` = 0x10 <br />

## Light bytes
`Rainbow` = 0x05 <br />
`RainbowMorph` = 0x04 <br />
`RainbowSecond` = 0xFF <br />
`StaticFirst` = 0x01 <br />
`StaticSecond` = 0x02 <br />

## Command bytes
`RgbSyncCmd` = 0x30 <br />
└───`Arg` = 0x00 (disable), 0x01 (enable) <br />
`PwmCmd` = 0x31 <br />
`SpeedCmd` = 0x00 <br />
