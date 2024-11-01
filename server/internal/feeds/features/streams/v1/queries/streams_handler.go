package queries

import (
	"context"
	"net/http"
)

type StreamsHandler struct{}

func NewStreamsHandler() *StreamsHandler {
	return &StreamsHandler{}
}

func (c *StreamsHandler) Handle(
	ctx context.Context,
	query *Streams,
) (*http.Request, error ){
	// Find url in cache
	// Retrieve, if err ?
	url := "https://v16m.tiktokcdn-us.com/b024652043bd6d795f40cda2082ae723/6724801f/video/tos/useast5/tos-useast5-pve-0068-tx/oURjRkOALeTXZJ4JGHgqKRLSfWACAAebIaI1Qk/?a=1233&bti=NEBzNTY6QGo6OjZALnAjNDQuYCMxNDNg&ch=0&cr=13&dr=0&er=0&lr=all&net=0&cd=0%7C0%7C0%7C&cv=1&br=1664&bt=832&cs=0&ds=6&ft=aEKfuqT0mGjPD12ePvYJ3wUS_-1tjeF~O5&mime_type=video_mp4&qs=0&rc=Z2Y1NjY1OTloOmg1Njs7OUBpMzUzd3A5cjw4djMzZzczNEAuXjRgL2EuXjQxYTA2NDVeYSNraS1uMmQ0aGVgLS1kMS9zcw%3D%3D&vvpl=1&l=20241101011446A3351F145D72E10A510B&btag=e00088000"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
