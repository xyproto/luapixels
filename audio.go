package luapixels

import (
	"fmt"
	"math"
	"time"
	"unsafe"

	"github.com/gen2brain/malgo"
)

var (
	context *malgo.AllocatedContext
	device  malgo.Device
)

// InitAudio initializes the miniaudio context.
func InitAudio() error {
	var err error
	context, err = malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	return err
}

func PlaySound(frequency float32, duration int) error {
	fmt.Printf("play sound frequency %f duration %d\n", frequency, duration)

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = 44100

	sampleRate := uint32(44100)
	bufferSamples := uint32(duration) * sampleRate / 1000

	samples := make([]int16, bufferSamples)
	for i := uint32(0); i < bufferSamples; i++ {
		samples[i] = int16(math.Sin(float64(i)*2*math.Pi*float64(frequency)/float64(sampleRate)) * 32767)
	}

	sampleIndex := uint32(0)
	onSendSamples := func(outputBuffer, _ []byte, framecount uint32) {
		for i := uint32(0); i < framecount; i++ {
			if sampleIndex < bufferSamples {
				sample := samples[sampleIndex]
				// Copy the sample to the audio buffer
				copy(outputBuffer[i*2:], (*[2]byte)(unsafe.Pointer(&sample))[:])
				sampleIndex++
			} else {
				// Fill the rest of the buffer with silence if we've played all samples
				var x int16 = 0
				copy(outputBuffer[i*2:], (*[2]byte)(unsafe.Pointer(&x))[:])
			}
		}
	}

	deviceCallbacks := malgo.DeviceCallbacks{
		Data: onSendSamples,
	}

	device, err := malgo.InitDevice(context.Context, deviceConfig, deviceCallbacks)
	if err != nil {
		return err
	}
	defer device.Uninit()

	err = device.Start()
	if err != nil {
		return err
	}

	// Wait for the duration of the sound
	time.Sleep(time.Duration(duration) * time.Millisecond)

	return nil
}
