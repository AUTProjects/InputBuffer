/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-01-2018
 * |
 * | File Name:     generator/bursty.go
 * +===============================================
 */

package generator

import (
	"math/rand"
	"time"

	"github.com/1995parham/InputBuffer.go/switches"
)

// Bursty traffic generator
type Bursty struct {
	p      int
	w      int
	status bool
	rand   *rand.Rand
}

// NewBursty creates bursty traffic generator with
// load p and w packet burst size.
func NewBursty(p int, w int) *Bursty {
	return &Bursty{
		p:      p,
		w:      w,
		status: true,
		rand:   rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Generate adds incomming packets into given switch buffers
// with on/off distribution
// and returns number of generated packets
func (b *Bursty) Generate(sw *switches.Switch) int {
	in := 0

	if b.status {
		in += b.w
		sw.ArriveMany(b.w, b.rand.Intn(sw.N), b.rand.Intn(sw.N))
	}

	if b.status {
		if b.rand.Float64() < (1.0 / float64(b.w)) {
			b.status = false
		}
	} else {
		if b.rand.Float64() < (float64(b.p) / (float64(b.w) * float64(100-b.p))) {
			b.status = true
		}
	}

	return in
}
