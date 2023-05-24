
package main

import (
    "log"
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) != 2 {
        log.Fatal("Uso: sound <dispositivo>")
    }

    dispositivo := os.Args[1]

    comandos := map[string]string{
        "hdmi":   "alsa_output.pci-0000_01_00.1.hdmi-stereo",
        "hyperx": "alsa_output.usb-Kingston_HyperX_Cloud_Flight_Wireless-00.analog-stereo",
    }

    cmd, ok := comandos[dispositivo]
    if !ok {
        log.Fatalf("Dispositivo %s no encontrado", dispositivo)
    }

    command := exec.Command("pactl", "set-default-sink", cmd)
    err := command.Run()

    if err != nil {
        log.Fatal(err)
    }
}
