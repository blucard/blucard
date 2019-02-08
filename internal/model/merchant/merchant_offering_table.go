package merchant

type OfferingTier struct {
	name string
}

func NewOfferingTier(name string) OfferingTier {
	return OfferingTier{name: name}
}

func (tier OfferingTier) Name() string {
	return tier.name
}

// Retrieves the offerings associated with each merchant for a single tier
type MerchantOfferingTable interface {
	GetOfferings(merchantId MerchantId) []MerchantOffering
}

// A MerchantOfferingTable capable of retrieving offering information at all tiers
type TieredMerchantOfferingTable interface {
	GetTiers() []OfferingTier
	GetOfferings(merchantId MerchantId, tier OfferingTier) []MerchantOffering
}