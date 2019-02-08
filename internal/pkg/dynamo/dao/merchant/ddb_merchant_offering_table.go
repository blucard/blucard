package merchant

import(
	"github.com/blucard/blucard/internal/model/merchant"
)

const(
	merchantIdKey = "merchantId" // DDB hash key
	offeringIdKey = "offeringId" // DDB range key
)

// DDB-backed implementation of merchant.MerchantOfferingTable
type DDBMerchantOfferingTable struct {
	tableName string
	// TODO: DDB client here
}

func NewDDBMerchantOfferingTable(tableName string) *DDBMerchantOfferingTable {
	// TODO: Pass in backing DDB client here
	return &DDBMerchantOfferingTable{tableName: tableName}
}

func (table *DDBMerchantOfferingTable) GetOfferings(merchant.MerchantId) []merchant.MerchantOffering {
	// TODO: translate to DDB query
	return nil
}
