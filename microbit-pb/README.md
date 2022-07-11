# Tinygo on Microbit Protobuf test

(using Arch Linux)

- `cd` to this directory
- `tinygo flash -target microbit`

See https://community.tmpdir.org/t/tinygo-on-mcus/553 for discussion

Currently getting:

```
[cbrake@ceres tinygo_experiments]$ tg_flash_microbit_pb
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 76 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 104 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 680 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 1364 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 1392 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 1422 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 1452 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 2544 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 3864 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 3892 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 4288 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 5068 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 5096 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 5318 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 5608 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 5636 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 5678 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 5728 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 5756 bytes
tinygo:ld.lld: error: section '.data' will not fit in region 'RAM': overflowed by 5924 bytes
tinygo:ld.lld: error: too many errors emitted, stopping now (use -error-limit=0 to see all errors)
failed to run tool: ld.lld
error: failed to link /tmp/tinygo580269492/main: exit status 1

```

## Reference

- https://tinygo.org/docs/reference/microcontrollers/microbit/
- https://tinygo.org/docs/reference/microcontrollers/machine/microbit/
