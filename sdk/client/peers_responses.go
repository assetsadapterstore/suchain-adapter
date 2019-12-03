// This file is part of SGU Go Client.
//
// (c) SGU Ecosystem <info@SGU.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Peer struct {
	Ip      string `json:"ip,omitempty"`
	Port    uint16 `json:"port,omitempty"`
	Version string `json:"version,omitempty"`
	Height  uint32 `json:"height,omitempty"`
	Status  string `json:"status,omitempty"`
	Os      string `json:"os,omitempty"`
	Latency byte   `json:"latency,omitempty"`
}

type Peers struct {
	Meta Meta   `json:"meta,omitempty"`
	Data []Peer `json:"data,omitempty"`
}

type GetPeer struct {
	Meta Meta `json:"meta,omitempty"`
	Data Peer `json:"data,omitempty"`
}
