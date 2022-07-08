# Tinygo on Arduino UNO

(using Arch Linux)

- `pacman -s avr-gcc avr-libc avrdude tinygo`
- `cd` to this directory
- `tinygo flash -target arduino`

## Reference

- https://tinygo.org/docs/reference/microcontrollers/arduino/
- https://tinygo.org/docs/reference/microcontrollers/machine/arduino/

Note, goroutines are not supported on the AVR backend. See:

https://tinygo.org/docs/reference/usage/important-options/

> - `scheduler`: Use the specified scheduler. The default scheduler varies by
>   platform. For example, AVR currently defaults to none because it has such
>   limited memory while asyncify and tasks are used for other platforms.
>   Normally you do not need to override the default except on AVR where you can
>   optionally select the tasks scheduler if you want concurrency.
>   - `scheduler=tasks` The tasks scheduler is a scheduler much like an RTOS
>     available for non-WASM platforms. This is usually the preferred scheduler.
>   - `scheduler=asyncify` The asyncify scheduler is a scheduler for WASM based
>     off of Binaryen’s Asyncify Pass.
>   - `scheduler=none` The none scheduler disables scheduler support, which
>     means that goroutines and channels are not available. It can be used to
>     reduce firmware size and RAM consumption if goroutines and channels are
>     not needed.
