package merchant

type OfferTier struct {
	name string
}

// Retrieves the offers associated with each merchant for a single tier
type MerchantOfferTable interface {
	GetOffers(merchantID MerchantID) []MerchantOffer
}

// A MerchantOfferTable capable of retrieving offer information at all tiers
type TieredMerchantOfferTable interface {
	GetTiers() []OfferTier
	GetOffers(merchantID MerchantID, tier OfferTier) []MerchantOffer
}