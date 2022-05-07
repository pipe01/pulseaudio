package pulseaudio

import "io"

type OtherClient struct {
	Index           uint32
	ApplicationName string
	ModuleIndex     uint32
	Driver          string
	PropList        map[string]string
}

// ReadFrom deserializes a client packet from pulseaudio
func (c *OtherClient) ReadFrom(r io.Reader) (int64, error) {
	err := bread(r,
		uint32Tag, &c.Index,
		stringTag, &c.ApplicationName,
		uint32Tag, &c.ModuleIndex,
		stringTag, &c.Driver,
		&c.PropList,
	)

	return 0, err
}

func (c *Client) Clients() ([]OtherClient, error) {
	b, err := c.request(commandGetClientInfoList)
	if err != nil {
		return nil, err
	}
	var clients []OtherClient
	for b.Len() > 0 {
		var client OtherClient
		err = bread(b, &client)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}
