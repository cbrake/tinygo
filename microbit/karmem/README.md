# Tinygo on Microbit Protobuf test

(using Arch Linux)

- `. envsetup.sh`
- `tg_flash_microbit_karmem`

This compiles but the Flash cost is significant:

```
[cbrake@ceres tinygo_experiments]$ tg_flash_microbit
   code    data     bss |   flash     ram
   4224      24    2232 |    4248    2256
[cbrake@ceres tinygo_experiments]$ tg_flash_microbit_karmem
   code    data     bss |   flash     ram
  40796    1400    2232 |   42196    3632
```

Dumping `-size full`, we get:

```
[cbrake@ceres tinygo_experiments]$ tg_flash_microbit_karmem
   code  rodata    data     bss |   flash     ram | package
------------------------------- | --------------- | -------
  12084       0      34      54 |   12118      88 | (unknown)
    206       0       0       0 |     206       0 | /home/runner/work/tinygo/tinygo/home/runner/work/tinygo/tinygo/build/release/tinygo/lib/compiler-rt-builtins/arm
     12       0       0       0 |      12       0 | /scratch/tinygo_experiments/microbit/usr/lib/tinygo/src/device/arm
      4       0       0       0 |       4       0 | /scratch/tinygo_experiments/microbit/usr/lib/tinygo/src/device/nrf
     64       0       0       0 |      64       0 | /scratch/tinygo_experiments/microbit/usr/lib/tinygo/src/internal/task
     28       0       0       0 |      28       0 | /scratch/tinygo_experiments/usr/lib/tinygo/src/runtime
      6       0       0       0 |       6       0 | C nrfx
      0       0       0    2048 |       0    2048 | C stack
    120       0       0       0 |     120       0 | Go interface assert
    280       0       0       0 |     280       0 | Go interface method
      0     499       0       0 |     499       0 | Go reflect data
   1804       0       0       0 |    1804       0 | build/release/tinygo/lib/compiler-rt-builtins
     80       0       0       0 |      80       0 | build/release/tinygo/lib/picolibc/newlib/libc/string
     74       0       0       0 |      74       0 | device/arm
     40       0       0       0 |      40       0 | device/nrf
      6       0       0       0 |       6       0 | errors
   4624     147       0       0 |    4771       0 | fmt
     64       0       0       0 |      64       0 | internal/fmtsort
    106       2       0       0 |     108       0 | internal/itoa
    258      24       0       0 |     282       0 | internal/task
     48       3       0       0 |      51       0 | io/fs
    282      43       8       0 |     333       8 | karmem.org/golang
     66       0       0     130 |      66     130 | machine
    538      23       0       0 |     561       0 | main
    306     256       0       0 |     562       0 | math/bits
    152      39      28       0 |     219      28 | os
   2831     711       0       0 |    3542       0 | reflect
   2944     201       8       0 |    3153       8 | runtime
    106       0       0       0 |     106       0 | runtime/volatile
   6550    4431    1322       0 |   12303    1322 | strconv
     28       0       0       0 |      28       0 | sync
     28       6       0       0 |      34       0 | syscall
    672       0       0       0 |     672       0 | unicode/utf8
------------------------------- | --------------- | -------
  34411    6385    1400    2232 |   42196    3632 | total
```

## Reference

- https://tinygo.org/docs/reference/microcontrollers/microbit/
- https://tinygo.org/docs/reference/microcontrollers/machine/microbit/
