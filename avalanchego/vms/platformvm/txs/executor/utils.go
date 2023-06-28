package executor

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/platformvm/config"
)

// The value of currentTimestamp is used to return new inflation settings over time
func GetCurrentInflationSettings(currentTimestamp time.Time, networkID uint32, config *config.Config) (uint64, uint64, uint64, uint32, time.Duration, time.Duration, time.Duration, time.Time) {
	switch networkID {
	case constants.FlareID:
		return 10 * units.MegaAvax, // minValidatorStake
			50 * units.MegaAvax, // maxValidatorStake
			1 * units.KiloAvax, // minDelegatorStake
			0, // minDelegationFee
			2 * 7 * 24 * time.Hour, // minStakeDuration
			365 * 24 * time.Hour, // maxStakeDuration
			3 * 24 * time.Hour, // minFutureStartTimeOffset
			time.Date(2023, time.July, 5, 15, 0, 0, 0, time.UTC) // minStakeStartTime
	case constants.CostwoID:
		return 100 * units.KiloAvax,
			50 * units.MegaAvax,
			1 * units.KiloAvax,
			0,
			2 * 7 * 24 * time.Hour,
			365 * 24 * time.Hour,
			MaxFutureStartTime,
			time.Date(2023, time.May, 25, 15, 0, 0, 0, time.UTC)
	case constants.StagingID:
		return 100 * units.KiloAvax,
			50 * units.MegaAvax,
			1 * units.KiloAvax,
			0,
			2 * 7 * 24 * time.Hour,
			365 * 24 * time.Hour,
			MaxFutureStartTime,
			time.Date(2023, time.May, 10, 15, 0, 0, 0, time.UTC)
	case constants.LocalFlareID:
		return 10 * units.KiloAvax,
			50 * units.MegaAvax,
			10 * units.KiloAvax,
			0,
			2 * 7 * 24 * time.Hour,
			365 * 24 * time.Hour,
			24 * time.Hour,
			time.Date(2023, time.April, 10, 15, 0, 0, 0, time.UTC)
	default:
		return config.MinValidatorStake,
			config.MaxValidatorStake,
			config.MinDelegatorStake,
			config.MinDelegationFee,
			config.MinStakeDuration,
			config.MaxStakeDuration,
			MaxFutureStartTime,
			time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	}
}
