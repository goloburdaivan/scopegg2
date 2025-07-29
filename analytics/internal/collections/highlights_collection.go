package collections

import (
	"scopegg2-shared/dto"
)

type Highlights struct {
	data map[int][]dto.Kill
}

func (h *Highlights) Init() {
	h.data = make(map[int][]dto.Kill)
}

func (h *Highlights) Add(round int, kill dto.Kill) {
	h.data[round] = append(h.data[round], kill)
}

func (h *Highlights) GetData() map[int][]dto.Kill {
	return h.data
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
	return h.filter(func(k dto.Kill) bool { return k.IsHeadshot })
}

func (h *Highlights) WallbangsOnly() *Highlights {
	return h.filter(func(k dto.Kill) bool { return k.IsWallBang })
}

func (h *Highlights) NoScopesOnly() *Highlights {
	return h.filter(func(k dto.Kill) bool { return k.IsNoScope })
}

func (h *Highlights) TroughSmokesOnly() *Highlights {
	return h.filter(func(k dto.Kill) bool { return k.IsThroughSmoke })
}

func (h *Highlights) filter(predicate func(dto.Kill) bool) *Highlights {
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
