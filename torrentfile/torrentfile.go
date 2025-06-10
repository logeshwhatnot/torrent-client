package torrentfile

import (
	"github.com/jackpal/bencode-go"
	"io"
)

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

type bencodeInfo struct {
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece_length"`
	Pieces      string `bencode:"pieces"`
	Length      string `bencode:"file"`
}

type bencodeTorrent struct {
	Announce string      `bencode:"announce"`
	info     bencodeInfo `bencode:"info"`
}

// Open parses a torrent file
func Open(r io.Reader) (*bencodeTorrent, error) {
	bt := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bt)
	if err != nil {
		return nil, err
	}
	return &bt, nil
}
