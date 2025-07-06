package music

import (
	"bytes"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

type MusicStruct struct {
	Context  *oto.Context
	Player   *oto.Player
	Filepath string
	reset    bool

	MPThree *mp3.Decoder
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

func (music *MusicStruct) Reset() {
	music.reset = true
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
			if music.reset {
				music.Player.Close()
				music.reset = false
			}
		}
		music.Player.Close()
	}()

	return music
}
