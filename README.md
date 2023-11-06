<h1 align="center">lianlinux</h1>

<p align="center">
  <a href="https://github.com/crptmem/lianlinux/stargazers"><img src="https://img.shields.io/github/stars/crptmem/lianlinux?colorA=151515&colorB=B66467&style=for-the-badge&logo=starship"></a>
  <a href="https://github.com/crptmem/lianlinux/issues"><img src="https://img.shields.io/github/issues/crptmem/lianlinux?colorA=151515&colorB=8C977D&style=for-the-badge&logo=bugatti"></a>
  <a href="https://github.com/crptmem/lianlinux/network/members"><img src="https://img.shields.io/github/forks/crptmem/lianlinux?colorA=151515&colorB=D9BC8C&style=for-the-badge&logo=github"></a>
</p>

> 🌈 An app to control Lian Li Hub lights on Linux

# Supported devices
| Product name | Product ID |
|----------|:-------------:|
| LianLi-UNI FAN-SL-v1.8 | 0xa100 |

# Lian Li USB protocol
Color mode switch packet:
| Start (?) | Port ID | First mode byte | Second mode byte |
|----------|:-------------:|-------------:|-------------:|
| E0 | 10-14 | 00-? | 00-? |
