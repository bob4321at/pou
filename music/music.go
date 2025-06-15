package music

import (
	"bytes"
	"io"
	"main/utils"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

var AtPeak bool

type VolumePoint struct {
	Frame  int
	Volume float64
}

type MusicStruct struct {
	Context  *oto.Context
	Player   *oto.Player
	Filepath string

	MPThree *mp3.Decoder
	Peaks   []VolumePoint
	AtPeak  bool
}

func (music *MusicStruct) PlaySong(path string) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMP, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic(err)
	}

	music.Player = music.Context.NewPlayer(decodedMP)
	music.Player.Play()
}

func NewMusic(path string) (music MusicStruct) {
	music.Filepath = path

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	music.MPThree, err = mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic(err)
	}

	var data []VolumePoint
	frame := 0
	buf := make([]byte, 4096)

	for {
		n, err := music.MPThree.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		volume := utils.CalculateVolume(buf[:n])
		data = append(data, VolumePoint{Frame: frame, Volume: volume})
		frame++
	}

	for _, frame := range data {
		if frame.Volume > 0.2 {
			music.Peaks = append(music.Peaks, frame)
		}
	}

	op := &oto.NewContextOptions{}

	op.SampleRate = 44100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE
	oldOtoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic(err)
	}
	<-readyChan
	music.Context = oldOtoCtx

	go func() {
		music.PlaySong(path)
		for music.Player.IsPlaying() {
			time.Sleep(time.Millisecond)

			for _, frame := range music.Peaks {
				if frame.Frame >= int(utils.GameTime-2) && frame.Frame <= int(utils.GameTime+2) {
					AtPeak = true
					go func() {
						time.Sleep(time.Second / 7)
						AtPeak = false
					}()
				}
			}
		}
		music.Player.Close()
	}()

	return music
}

var Music = NewMusic("./music/song.mp3")
