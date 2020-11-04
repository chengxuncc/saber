package saber

import (
	"io"
	"net/http"
)

func Curl(url string) *Compound {
	return Do().Curl(url)
}

func (c *Compound) Curl(url string) *Compound {
	return c.Next(func(cmd *Command) error {
		resp, err := c.Script.HttpClient.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		_, err = io.Copy(cmd.GetStdout(), resp.Body)
		if err != nil {
			return err
		}
		return nil
	})
}

func Request(method, url string, hdr http.Header, body io.Reader) *Compound {
	return Do().Request(method, url, hdr, body)
}

func (c *Compound) Request(method, url string, hdr http.Header, body io.Reader) *Compound {
	return c.Next(func(cmd *Command) error {
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			return err
		}
		if len(hdr) > 0 {
			req.Header = hdr
		}
		resp, err := c.Script.HttpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		_, err = io.Copy(cmd.GetStdout(), resp.Body)
		if err != nil {
			return err
		}
		return nil
	})
}
