package winact

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"

	"github.com/rs/zerolog/log"
)

// ============================================================================
// Public input plugin interface
// ============================================================================

type WinActInputPlugin struct{}

func init() {
	inputs.Add("winact", func() telegraf.Input {
		return &WinActInputPlugin{}
	})
}

func (input *WinActInputPlugin) Init() error {
	log.Info().Msg("winact plugin started")
	return nil
}

func (input *WinActInputPlugin) Stop() {

}

func (input *WinActInputPlugin) SampleConfig() string {
	return `
[[inputs.winact]]
	# no config
`
}

func (input *WinActInputPlugin) Description() string {
	return "Gather activity information from a Windows machine"
}

func (input *WinActInputPlugin) Gather(a telegraf.Accumulator) error {
	// a.AddFields("winact", ...)
	return nil
}
