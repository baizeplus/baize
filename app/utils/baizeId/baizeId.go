package baizeId

import (
	"crypto/rand"
	"errors"
	"io"
	"sync"
	"time"
)

var (
	n        *Node
	r        = rand.Reader
	stepMask = int64(-1 ^ (-1 << 15))
	encode   = []byte("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
)

type Node struct {
	mu   sync.Mutex
	node int32
	time int64
	step int64
}

func NewNode(node int32) error {
	if node < 0 || node > 8191 {
		return errors.New("node number must be between 0 and 8191")
	}
	n = new(Node)
	n.node = node

	return nil
}
func GetId() string {
	buf := make([]byte, 16)
	if n == nil {
		var rByte = [5]byte{}
		_, _ = io.ReadFull(r, rByte[:])
		micro := time.Now().UnixMicro()
		buf[0] = encode[(micro>>47)&0x1F]
		buf[1] = encode[(micro>>42)&0x1F]
		buf[2] = encode[(micro>>37)&0x1F]
		buf[3] = encode[(micro>>32)&0x1F]
		buf[4] = encode[(micro>>27)&0x1F]
		buf[5] = encode[(micro>>22)&0x1F]
		buf[6] = encode[(micro>>17)&0x1F]
		buf[7] = encode[(micro>>12)&0x1F]
		buf[8] = encode[(micro>>7)&0x1F]
		buf[9] = encode[(micro>>2)&0x1F]
		buf[10] = encode[(byte(micro<<3)|rByte[0]>>5)&0x1F]
		buf[11] = encode[rByte[0]&0x1F]
		buf[12] = encode[rByte[1]&0x1F]
		buf[13] = encode[rByte[2]&0x1F]
		buf[14] = encode[rByte[3]&0x1F]
		buf[15] = encode[rByte[4]&0x1F]
		return string(buf)
	}
	n.mu.Lock()
	now := time.Now().UnixMicro()
	if now == n.time {
		n.step = (n.step + 1) & stepMask

		if n.step == 0 {
			for now <= n.time {
				now = time.Now().UnixMicro()
			}
		}
	} else {
		n.step = 0
	}
	n.time = now
	buf[13] = encode[n.step>>10&0x1F]
	buf[14] = encode[n.step>>5&0x1F]
	buf[15] = encode[n.step&0x1F]
	n.mu.Unlock()
	buf[0] = encode[(now>>47)&0x1F]
	buf[1] = encode[(now>>42)&0x1F]
	buf[2] = encode[(now>>37)&0x1F]
	buf[3] = encode[(now>>32)&0x1F]
	buf[4] = encode[(now>>27)&0x1F]
	buf[5] = encode[(now>>22)&0x1F]
	buf[6] = encode[(now>>17)&0x1F]
	buf[7] = encode[(now>>12)&0x1F]
	buf[8] = encode[(now>>7)&0x1F]
	buf[9] = encode[(now>>2)&0x1F]
	buf[10] = encode[((byte(now<<3))|byte(n.node>>10))&0x1F]
	buf[11] = encode[(n.node>>5)&0x1F]
	buf[12] = encode[(n.node)&0x1F]
	return string(buf)
}
