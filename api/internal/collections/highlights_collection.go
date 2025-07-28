package collections

import (
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

type Highlights struct {
	data map[int][]events.Kill
}

func NewHighlights(data map[int][]events.Kill) *Highlights {
	return &Highlights{
		data: data,
	}
}

func (h *Highlights) When(condition bool, fn func(*Highlights)) *Highlights {
	if condition {
		fn(h)
	}
	return h
}

func (h *Highlights) FromKills(minKills int) *Highlights {
	for round, kills := range h.data {
		if len(kills) < minKills {
			delete(h.data, round)
		}
	}

	return h
}

func (h *Highlights) HeadShotsOnly() *Highlights {
	return h.filter(func(k events.Kill) bool { return k.IsHeadshot })
}

func (h *Highlights) WallbangsOnly() *Highlights {
	return h.filter(func(k events.Kill) bool { return k.IsWallBang() })
}

func (h *Highlights) NoScopesOnly() *Highlights {
	return h.filter(func(k events.Kill) bool { return k.NoScope })
}

func (h *Highlights) TroughSmokesOnly() *Highlights {
	return h.filter(func(k events.Kill) bool { return k.ThroughSmoke })
}

func (h *Highlights) filter(predicate func(events.Kill) bool) *Highlights {
	for round, kills := range h.data {
		filtered := kills[:0]
		for _, kill := range kills {
			if predicate(kill) {
				filtered = append(filtered, kill)
			}
		}
		if len(filtered) != len(kills) {
			delete(h.data, round)
		} else {
			h.data[round] = filtered
		}
	}

	return h
}
