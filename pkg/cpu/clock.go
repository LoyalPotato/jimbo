package cpu

type Clock struct {
	m uint64
	t uint64
}

var GlobalClock *Clock

func (c *Clock) Reset() {
	c.t = 0
	c.m = 0
}

func (c *Clock) Inc(incValues ...uint8) {
	if len(incValues) == 1 {
		c.m += uint64(incValues[0])
	} else {
		c.m++
	}
	c.t = c.m * 4
}
