package merchant

import(
	"github.com/blucard/blucard/internal/model/merchant"
)

const merchantOfferTableSchema = struct {
	hashKey string
	rangeKey string
}{
	hashKey: "merchantId",
	rangeKey: "offerId"
}

// DDB-backed implementation of merchant.MerchantOfferTable
type DDBMerchantOfferTable struct {
	tableName string
	// TODO: DDB client here
}

func NewDDBMerchantOfferTable(tableName string) *DDBMerchantOfferTable {
	// TODO: Pass in backing DDB client here
	return &DDBMerchantOfferTable{tableName: tableName}
}

func (table *DDBMerchantOfferTable) GetOffers(merchantID merchant.MerchantID) []merchant.MerchantOffer {
	// TODO: translate to DDB query
	return nil
}
