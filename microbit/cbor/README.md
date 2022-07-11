# Tinygo on Microbit CBOR test

(using Arch Linux)

- `cd` to this directory
- `tinygo flash -target microbit`

See https://community.tmpdir.org/t/tinygo-on-mcus/553 for discussion

Currently getting the following with tinygo v0.24.0:

```
[cbrake@ceres tinygo_experiments]$ tg_flash_microbit_cbor
# github.com/fxamacker/cbor/v2
../../../home/cbrake/go/pkg/mod/github.com/fxamacker/cbor/v2@v2.4.0/decode.go:1231:17: MakeMapWithSize not declared by package reflect
```

## Reference

- https://tinygo.org/docs/reference/microcontrollers/microbit/
- https://tinygo.org/docs/reference/microcontrollers/machine/microbit/
