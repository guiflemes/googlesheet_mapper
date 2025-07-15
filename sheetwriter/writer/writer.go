package writer

import (
	"fmt"
	"time"
)

type Text struct {
	Text string
}

func (t Text) Value() string {
	return t.Text
}

type Number struct {
	V float64
}

func (n Number) Value() string {
	return fmt.Sprintf("%f", n.V)
}

type Date struct {
	D time.Time
}

func (d Date) Value() string {
	return d.D.Format("2006-01-02")
}
