package ethermine_test

import (
	"github.com/tier2pool/tier2pool/internal/datasource/ethermine"
	"testing"
)

const GenesisAddress = "0x0000000000000000000000000000000000000000"

func TestGetMinerDashboard(t *testing.T) {
	dashboard, err := ethermine.GetMinerDashboard(GenesisAddress)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", dashboard)
}

func TestGetMinerStats(t *testing.T) {
	dashboard, err := ethermine.GetMinerStats(GenesisAddress)
	if err != nil && err.Error() != "NO DATA" {
		t.Fatal(err)
	}
	t.Logf("%+v", dashboard)
}
