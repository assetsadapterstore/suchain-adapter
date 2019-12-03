// This file is part of SGU Go Client.
//
// (c) SGU Ecosystem <info@SGU.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Pagination struct {
	Page    int    `url:"page"`
	Limit   int    `url:"limit"`
	//Height  int    `url:"height"`
	//BlockId string `url:"blockId"`
}
