package pclient

func (c *Client) Time() (string, error) {
	resp, err := c.GetReq("time")
	if err != nil {
		return "", err
	}
	return string(resp), nil

}
