package pulseaudio

import "io"

type SinkInput struct {
	Index            uint32
	Name             string
	ModuleIndex      uint32
	ClientIndex      uint32
	SinkIndex        uint32
	SampleSpec       sampleSpec
	ChannelMap       channelMap
	Cvolume          cvolume
	Latency          uint64
	SinkLatency      uint64
	ResampleMethod   string
	Driver           string
	Muted            bool
	PropList         map[string]string
	Corked           bool
	HasVolume        bool
	IsVolumeWritable bool
	Format           formatInfo
}

// ReadFrom deserializes a sink input packet from pulseaudio
func (s *SinkInput) ReadFrom(r io.Reader) (int64, error) {
	err := bread(r,
		uint32Tag, &s.Index,
		stringTag, &s.Name,
		uint32Tag, &s.ModuleIndex,
		uint32Tag, &s.ClientIndex,
		uint32Tag, &s.SinkIndex,
		&s.SampleSpec,
		&s.ChannelMap,
		&s.Cvolume,
		usecTag, &s.Latency,
		usecTag, &s.SinkLatency,
		stringTag, &s.ResampleMethod,
		stringTag, &s.Driver,
		&s.Muted,
		&s.PropList,
		&s.Corked,
		&s.HasVolume,
		&s.IsVolumeWritable,
		&s.Format,
	)

	return 0, err
}

func (c *Client) SinkInputs() ([]SinkInput, error) {
	b, err := c.request(commandGetSinkInputInfoList)
	if err != nil {
		return nil, err
	}
	var sinkInputs []SinkInput
	for b.Len() > 0 {
		var sinkInput SinkInput
		err = bread(b, &sinkInput)
		if err != nil {
			return nil, err
		}
		sinkInputs = append(sinkInputs, sinkInput)
	}
	return sinkInputs, nil
}

func (c *Client) MoveSinkInput(sinkInputIndex uint32, sinkIndex uint32) error {
	_, err := c.request(commandMoveSinkInput,
		uint32Tag, sinkInputIndex,
		uint32Tag, sinkIndex,
		stringNullTag,
	)
	return err
}
