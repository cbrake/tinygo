# shellcheck disable=SC2046,SC2016
TG_BASE=$(/bin/readlink -f $(dirname '${0}'))

tg_pb_generate() {
  protoc --proto_path=pb --go_out=. --go-vtproto_out=. --go-vtproto_opt=features=marshal+unmarshal+size pb/*.proto
}

tg_flash_microbit() {
  (cd "$TG_BASE/microbit/simple" && tinygo flash -size short -target microbit)
}

tg_flash_microbit_pb() {
  (cd "$TG_BASE/microbit/pb" && tinygo flash -size short -target microbit)
}

tg_flash_microbit_cbor() {
  (cd "$TG_BASE/microbit/cbor" && tinygo flash -size short -target microbit)
}

tg_flash_microbit_karmem() {
  (cd "$TG_BASE/microbit/karmem" &&
    karmem build --golang -o "./" point.km &&
    tinygo flash -size short -target microbit)
}

tg_flash_uno() {
  (cd "$TG_BASE/arduino-uno" && tinygo flash -size short -target arduino)
}
