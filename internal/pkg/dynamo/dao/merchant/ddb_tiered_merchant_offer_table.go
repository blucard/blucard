package merchant

import(
	"github.com/blucard/blucard/internal/model/merchant"
)

// Schema for a table with a single entry containing a map from tier to delegate table name
const tieredMerchantOfferTableSchema = struct {
	tableName string
	hashKey string
	mapName string
	mapKey string
}{
	tableName: "TieredMerchantOfferTable",
	hashKey: "hashKey",
	mapName: "tieredMerchantOfferMap"
}

// DDB-backed implementation of merchant.TieredMerchantOfferTable
type DDBTieredMerchantOfferTable struct {
	tierToDelegateTable map[string]*DDBMerchantOfferTable
	// TODO: DDB client here
	// TODO: treat the backing map as a cached object that periodically refreshes
}

func NewDDBTieredMerchantOfferTable() *DDBTieredMerchantOfferTable {
	//TODO: Pass in DDB client, use to build backing map
	return nil
}

func (table *DDBTieredMerchantOfferTable) GetTiers() []merchant.OfferTier {
	tiers := make([]merchant.OfferTier, len(table.tierToDelegateTable))
	i := 0
	for tier := range table.tierToDelegateTable {
		tiers[i] = merchant.NewOfferTier(tier)
		i++
	}
	return tiers
}

func (table *DDBTieredMerchantOfferTable) GetOffers(merchantID merchant.MerchantID, tier merchant.OfferTier) []merchant.MerchantOffer {
	return table.tierToDelegateTable[tier.Name].GetOffers(merchantID)
}