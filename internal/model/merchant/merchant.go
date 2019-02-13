package merchant

// Globally unique identifier for merchants
type MerchantID struct {
	ID string
}

// Identifier for deals/benefits offered to customers by a merchant
// Need only be unique among all offerings associated with a given merchant
type OfferID struct {
	ID string
}

// The (MerchantID, OfferID) pair which uniquely identifies any offering (globally unique)
type MerchantOfferID struct {
	merchantID MerchantID
	offeringID OfferID
}

// Information regarding the deal/benefit
// TODO: What else belongs here?
type OfferDetails struct {
	description string
}

func NewOfferDetails(description string) *OfferDetails {
	return &OfferDetails{description: description}
}

func (offeringDetails *OfferDetails) Description() string {
	return offeringDetails.description
}

// Base-level object for representing a deals/benefits offered by merchants
type MerchantOffer struct {
	id MerchantOfferID
	details *OfferDetails
}

func NewMerchantOffer(id MerchantOfferID, details *OfferDetails) MerchantOffer {
	return MerchantOffer{id: id, details: details}
}

func (merchantOffer *MerchantOffer) ID() MerchantOfferID {
	return merchantOffer.id
}

func (merchantOffer *MerchantOffer) Details() *OfferDetails {
	return merchantOffer.details
}