package main

import (
    "log"
    "os"
    "os/exec"
)

func main() {
    // Obtén los argumentos proporcionados al comando
    if len(os.Args) != 2 {
        log.Fatal("Uso: sound <dispositivo>")
    }

    // Obtén el nombre del dispositivo del argumento
    dispositivo := os.Args[1]

    // Mapa para almacenar los comandos correspondientes a cada dispositivo
    comandos := map[string]string{
        "hdmi":   "alsa_output.pci-0000_01_00.1.hdmi-stereo",
        "hyperx": "alsa_output.usb-Kingston_HyperX_Cloud_Flight_Wireless-00.analog-stereo",
    }

    // Verifica si el dispositivo existe en el mapa
    cmd, ok := comandos[dispositivo]
    if !ok {
        log.Fatalf("Dispositivo %s no encontrado", dispositivo)
    }

    // Ejecuta el comando para cambiar la salida de audio
    command := exec.Command("pactl", "set-default-sink", cmd)
    err := command.Run()
    if err != nil {
        log.Fatal(err)
    }
}
