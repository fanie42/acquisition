package main

import (
    "fmt"

    "github.com/fanie42/acquisition/internal/dvras/http/rest"
    "github.com/fanie42/acquisition/internal/dvras/inmem"
    "github.com/fanie42/acquisition/internal/dvras/portaudio"
    "github.com/fanie42/acquisition/pkg/dvras"
    "github.com/google/uuid"
    pa "github.com/gordonklaus/portaudio"
)

func main() {
    err := pa.Initialize()
    if err != nil {
        fmt.Printf("failed to initialize portaudio: %v", err)
        return
    }
    defer pa.Terminate()

    repo := inmem.New()
    app := portaudio.New(
        &portaudio.Config{
            SampleRate: 44100,
            DeviceID:   dvras.DeviceID(uuid.New()),
        },
        repo,
    )
    controller := rest.New(app)

    controller.Run()
}
