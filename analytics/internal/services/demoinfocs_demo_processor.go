package services

import (
	"context"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
	"io"
	"scopegg2-analytics/internal/collections"
	"scopegg2-analytics/internal/handlers"
	"scopegg2-shared/dto"
)

type DemoReader interface {
	ReadDemo(ctx context.Context, demoPath string) (io.ReadCloser, error)
}

type demoInfoCsDemoProcessor struct {
	demoReader DemoReader
}

func NewDemoInfoCsDemoProcessor(
	demoReader DemoReader,
) handlers.DemoProcessor {
	return &demoInfoCsDemoProcessor{
		demoReader: demoReader,
	}
}

func (d *demoInfoCsDemoProcessor) ProcessDemo(
	ctx context.Context,
	demoPath string,
	steamId uint64,
) (*collections.Highlights, error) {
	reader, err := d.demoReader.ReadDemo(ctx, demoPath)
	if err != nil {
		return &collections.Highlights{}, err
	}

	defer reader.Close()

	highlights, err := d.parseDemo(reader, steamId)
	if err != nil {
		return &collections.Highlights{}, err
	}

	return highlights.FromKills(3), nil
}

func (d *demoInfoCsDemoProcessor) parseDemo(demo io.ReadCloser, steamId uint64) (*collections.Highlights, error) {
	parser := demoinfocs.NewParser(demo)
	defer parser.Close()

	gameState := parser.GameState()
	var highlights collections.Highlights
	highlights.Init()

	currentRound := -1

	parser.RegisterEventHandler(func(e events.RoundStart) {
		currentRound = gameState.TotalRoundsPlayed()
	})

	parser.RegisterEventHandler(func(kill events.Kill) {
		if kill.Killer == nil || kill.Killer.SteamID64 != steamId {
			return
		}
		if currentRound >= 0 {
			highlights.Add(currentRound, dto.Kill{
				IsWallBang:     kill.IsWallBang(),
				IsNoScope:      kill.NoScope,
				IsHeadshot:     kill.IsHeadshot,
				AttackerBlind:  kill.AttackerBlind,
				IsThroughSmoke: kill.ThroughSmoke,
				AssistedFlash:  kill.AssistedFlash,
			})
		}
	})

	err := parser.ParseToEnd()
	if err != nil {
		return &highlights, err
	}

	return &highlights, nil
}
