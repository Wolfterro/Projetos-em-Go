package libs

import (
	"os"

	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
)

// BITDEPTH should be 2 most of the time
const BITDEPTH = 2

// MP3Content struct for the MP3 player instance
type MP3Content struct {
	SampleRate	int
	ChannelNum	int
	BitDepth	int
	BufferSize	int
	Path		string
	Filename	string

	Decoder		*minimp3.Decoder
	Data		[]byte
	Context		*oto.Context
	Player		*oto.Player
}

// SetMP3PathAndFilename sets the mp3 file path and base filename
func (c *MP3Content) SetMP3PathAndFilename(path string) {
	c.Path = path
	c.Filename = GetFilename(path)
}

// SetMP3Data sets all necessary information to build the mp3 player
func (c *MP3Content) SetMP3Data(customSampleRate int) {
	dataBytes, err := GetMP3FileContent(c.Path)
	if err != nil {
		PrintErrorMessage(err.Error())
		os.Exit(1)
	}

	dec, data, err := minimp3.DecodeFull(dataBytes)
	if err != nil {
		PrintErrorMessage(err.Error())
		os.Exit(1)
	}

	if customSampleRate > 0 {
		c.SampleRate = customSampleRate
	} else {
		c.SampleRate = dec.SampleRate
	}

	c.ChannelNum = dec.Channels
	c.BitDepth = BITDEPTH
	c.BufferSize = c.SampleRate
	c.Data = data
}

// SetMP3Context sets the mp3 context for the player
func (c *MP3Content) SetMP3Context() {
	context, err := oto.NewContext(c.SampleRate, c.ChannelNum, c.BitDepth, c.BufferSize)
	if err != nil {
		PrintErrorMessage(err.Error())
		os.Exit(1)
	}

	c.Context = context
}

// SetMP3Player sets the final mp3 player
func (c *MP3Content) SetMP3Player() {
	c.Player = c.Context.NewPlayer()
}

// BuildMP3Player wraps up all the necessary steps to build the mp3 player
func (c *MP3Content) BuildMP3Player(path string, customSampleRate int) {
	c.SetMP3PathAndFilename(path)
	c.SetMP3Data(customSampleRate)
	c.SetMP3Context()
	c.SetMP3Player()
}

// Play plays the mp3 decoded data with player
func (c *MP3Content) Play() {
	c.Player.Write(c.Data)
}
