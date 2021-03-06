package api

import (
	"fmt"
	"github.com/fearful-symmetry/gomsr"
)

func HasSMRR() (bool, error) {
	mtrrcap, err := gomsr.ReadMSR(0, 0xfe)
	if err != nil {
		return false, fmt.Errorf("Cannot access MSR IA32_MTRRCAP: %s", err)
	}

	return (mtrrcap>>11)&1 != 0, nil
}

// MTRR for the SMM code.
type SMRR struct {
	Active   bool
	PhysBase uint64
	PhysMask uint64
}

// Returns SMRR config of the platform
func GetSMRRInfo() (SMRR, error) {
	var ret SMRR

	smrr_physbase, err := gomsr.ReadMSR(0, 0x1f2)
	if err != nil {
		return ret, fmt.Errorf("Cannot access MSR IA32_SMRR_PHYSBASE: %s", err)
	}

	smrr_physmask, err := gomsr.ReadMSR(0, 0x1f3)
	if err != nil {
		return ret, fmt.Errorf("Cannot access MSR IA32_SMRR_PHYSMASK: %s", err)
	}

	ret.Active = (smrr_physmask>>11)&1 != 0
	ret.PhysBase = (smrr_physbase >> 12) & 0xfffff
	ret.PhysMask = (smrr_physmask >> 12) & 0xfffff

	return ret, nil
}

func IA32FeatureControlIsLocked() (bool, error) {
	feat_ctrl, err := gomsr.ReadMSR(0, 0x3a)
	if err != nil {
		return false, fmt.Errorf("Cannot access MSR IA32_FEATURE_CONTROL: %s", err)
	}

	return feat_ctrl&1 != 0, nil
}

func IA32PlatformID() (uint64, error) {
	plt_id, err := gomsr.ReadMSR(0, 0x17)
	if err != nil {
		return 0, fmt.Errorf("Cannot access MSR IA32_PLATFORM_ID: %s", err)
	}

	return plt_id, nil
}

func AllowsVMXInSMX() (bool, error) {
	feat_ctrl, err := gomsr.ReadMSR(0, 0x3a)
	if err != nil {
		return false, fmt.Errorf("Cannot access MSR IA32_FEATURE_CONTROL: %s", err)
	}

	var mask uint64 = (1 << 1) & (1 << 5) & (1 << 6)
	return (mask & feat_ctrl) == mask, nil
}

func TXTLeavesAreEnabled() (bool, error) {
	feat_ctrl, err := gomsr.ReadMSR(0, 0x3a)
	if err != nil {
		return false, fmt.Errorf("Cannot access MSR IA32_FEATURE_CONTROL: %s", err)
	}

	txt_bits := (feat_ctrl >> 8) & 0x1ff
	return (txt_bits&0xff == 0xff) || (txt_bits&0x100 == 0x100), nil
}
