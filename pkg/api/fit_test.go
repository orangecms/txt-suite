package api

import (
	"golang.org/x/exp/mmap"
	"testing"
)

func TestExtractFit(t *testing.T) {
	file, err := mmap.Open("/home/riot/work_stuff/coreboot_wege100s_systemboot_tboot.rom")
	filesize := file.Len()
	data := make([]byte, filesize)
	file.ReadAt(data, 0)
	fitTable, err := ExtractFit(data)
	if err != nil {
		t.Errorf("ExtractFit() failed: %v", err)
	}

	for _, item := range fitTable {
		item.FancyPrint()
	}

}
