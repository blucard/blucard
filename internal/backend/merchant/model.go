package merchant

// Merchant is a globally unique identifier for merchants.
type MerchantID string

// OfferID is an identifier for deals/benefits offered to customers by a merchant.
// Need only be unique among all offerings associated with a given merchant.
type OfferID string

// OfferTier is used to determine whether or not a given user can redeem a given offer.
type OfferTier string

// MerchantOfferId is a globally unique (MerchantID, OfferID) pair which identifies any offering.
type MerchantOfferID struct {
	MerchantID MerchantID
	OfferID    OfferID
}

// OfferDetails contains information regarding the deal/benefit, to be presented to customers.
// TODO: What else belongs here?
type OfferDetails struct {
	description string
}

// NewOfferDetails creates and initializes an OfferDetails.
func NewOfferDetails(description string) *OfferDetails {
	return &OfferDetails{description: description}
}

// Description returns the description for the offer described by this OfferDetails.
func (offerDetails *OfferDetails) Description() string {
	return offerDetails.description
}

// MerchantOffer is the base-level object for representing a deals/benefits offered by merchants.
type MerchantOffer struct {
	id      MerchantOfferID
	details *OfferDetails
}

// NewMerchantOffer createds and initializes a MerchantOffer.
func NewMerchantOffer(ID MerchantOfferID, details *OfferDetails) MerchantOffer {
	return MerchantOffer{id: ID, details: details}
}

// ID returns the MerchantOfferID for this MerchantOffer.
func (merchantOffer *MerchantOffer) ID() MerchantOfferID {
	return merchantOffer.id
}

// Details returns the OfferDetails for this MerchantOffer.
func (merchantOffer *MerchantOffer) Details() *OfferDetails {
	return merchantOffer.details
}

// merchantOfferTable retrieves the offers associated with each merchant for a single tier
type merchantOfferTable interface {
	getOffers(merchantID MerchantID) ([]MerchantOffer, error)
}

// TieredMerchantOfferTable is capable of retrieving offer information at all tiers
type TieredMerchantOfferTable interface {
	GetTiers() ([]OfferTier, error)
	GetOffers(merchantID MerchantID, tier OfferTier) ([]MerchantOffer, error)
}
