package dto

type Kill struct {
	IsWallBang     bool `json:"is_wall_bang"`
	IsNoScope      bool `json:"is_no_scope"`
	IsHeadshot     bool `json:"is_head_shot"`
	AttackerBlind  bool `json:"attacker_blind"`
	IsThroughSmoke bool `json:"is_through_smoke"`
	AssistedFlash  bool `json:"assisted_flash"`
}
