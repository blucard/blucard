package merchant

import(
	"github.com/blucard/blucard/internal/model/merchant"
)

const (
	// Schema for a table whose sole entry is a map from tier name to delegate table name
	tieredMerchantOfferingTableName = "TieredMerchantOfferingTable"
	tieredMerchantOfferingTableKey = "hashKey"
	tieredMerchantOfferingMapName = "TieredMerchantOfferingMap"
	tieredMerchantOfferingMapKey = "tier"
)

// DDB-backed implementation of merchant.TieredMerchantOfferingTable
type DDBTieredMerchantOfferingTable struct {
	tierToDelegateTable map[string]*DDBMerchantOfferingTable
	// TODO: DDB client here
	// TODO: treat the backing map as a cached object that periodically refreshes
}

func NewDDBTieredMerchantOfferingTable() *DDBTieredMerchantOfferingTable {
	//TODO: Pass in DDB client, use to build backing map
	return nil
}

func (table *DDBTieredMerchantOfferingTable) GetTiers() []merchant.OfferingTier {
	tiers := make([]merchant.OfferingTier, len(table.tierToDelegateTable))
	i := 0
	for tier := range table.tierToDelegateTable {
		tiers[i] = merchant.NewOfferingTier(tier)
		i++
	}
	return tiers
}

func (table *DDBTieredMerchantOfferingTable) GetOfferings(merchantId merchant.MerchantId, tier merchant.OfferingTier) []merchant.MerchantOffering {
	return table.tierToDelegateTable[tier.Name()].GetOfferings(merchantId)
}