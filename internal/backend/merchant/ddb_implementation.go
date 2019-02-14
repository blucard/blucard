package merchant

var merchantOfferTableSchema = struct {
	hashKey  string
	rangeKey string
}{
	hashKey:  "merchantId",
	rangeKey: "offerId",
}

// ddbMerchantOfferTable implements merchantOfferTable using DynamoDB as the backing data store.
type ddbMerchantOfferTable struct {
	tableName string
	// TODO: DDB client here
}

// NewddbMerchantOfferTable creates and initializes a ddbMerchantOfferTable.
func newddbMerchantOfferTable(tableName string) *ddbMerchantOfferTable {
	// TODO: Pass in backing DDB client here
	return &ddbMerchantOfferTable{tableName: tableName}
}

// GetOffers retrieves all the offers associated with the given MerchantID stored in this table.
func (table *ddbMerchantOfferTable) getOffers(merchantID MerchantID) ([]MerchantOffer, error) {
	// TODO: translate to DDB query
	return nil, nil
}

// Schema for a table with a single entry containing a map from tier to delegate table name
var tieredMerchantOfferTableSchema = struct {
	tableName string
	hashKey   string
	mapName   string
	mapKey    string
}{
	tableName: "TieredMerchantOfferTable",
	hashKey:   "hashKey",
	mapName:   "tieredMerchantOfferMap",
}

// DDBTieredMerchantOfferTable implements TieredMerchantOfferTable by using a DDB table which maps
// tier names to the delegate ddbMerchantOfferTables that contain the actual offer information.
type DDBTieredMerchantOfferTable struct {
	tierToDelegateTable map[string]*ddbMerchantOfferTable
	// TODO: DDB client here
	// TODO: treat the backing map as a cached object that periodically refreshes
}

// NewDDBTieredMerchantOfferTable creates and initializes a DDBTieredMerchantOfferTable.
func NewDDBTieredMerchantOfferTable() *DDBTieredMerchantOfferTable {
	//TODO: Pass in DDB client, use to build backing map
	return nil
}

// GetTiers returns all tiers in this table. That is, all keys which map to a delegate ddbMerchantOfferTable.
func (table *DDBTieredMerchantOfferTable) GetTiers() ([]OfferTier, error) {
	tiers := make([]OfferTier, len(table.tierToDelegateTable))
	i := 0
	for tier := range table.tierToDelegateTable {
		tiers[i] = OfferTier(tier)
		i++
	}
	return tiers, nil
}

// GetOffers returns all offers for a given MerchantID at a given OfferTier by first resolving the delegate
// ddbMerchantOfferTable name, then reading from that table.
func (table *DDBTieredMerchantOfferTable) GetOffers(merchantID MerchantID, tier OfferTier) ([]MerchantOffer, error) {
	return table.tierToDelegateTable[string(tier)].getOffers(merchantID)
}
