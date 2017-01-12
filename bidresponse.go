package odexchange

type BidResponse struct {
	Id       string     `json:"id"`
	SeatBids []*SeatBid `json:"seatbid"`
	BidId    string     `json:"bidid"`
	Currency string     `json:"cur"`

	// Optional feature to allow a bidder to set data in
	// the exchange's cookie. The string must be in base85 cookie
	// safe characters and be in any format. Proper JSON encoding
	// must be used to include "escaped" quotation marks.
	CustomData string `json:"customdata"`

	NoBidReason NoBidReason `json:"nbr"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type SeatBid struct {
	// 1+Bid objects each related to an impression.
	// Multiple bids can relate to the same impression.
	Bid []*Bid `json:"bid"`

	// ID of the buyer seat e.g advertiser, agency
	// on whose behalf this bid is made.
	Seat string `json:"seat"`

	// 0 = Impressions can be won individually.
	// 1 = Impressions must be won or lost as a group.
	WinSelection WinSelection `json:"group"`
}

type Bid struct {
	Id string `json:"id"`

	// ID of the Impression object in the related bid request
	ImpressionId string `json:"impid"`

	// Bid price expressed as CPM although the actual transaction
	Price float32 `json:"price"`

	// ID of a preloaded ad to be served if the bid wins
	AdId string `json:"adid"`

	// Win Notice URL called by the exhcange if the bid wins
	// (not necessarily indicative of a delivered, viewed, or billable add).
	// optional means of serving ad markup.
	WinNoticeURL string `json:"nurl"`

	// Optional means of conveying ad markup in case the bid wins;
	// supersedes the win notice if markup is included in both.
	Admarkup string `json:"adm"`

	// Advertiser domain for block list checking e.g "ford.com".
	// This can be an array of for the case of rotating creatives.
	// Exchanges can mandate that only one domain is allowed.
	AdDomains []string `json:"adomain"`

	// Bundle is a platform-specific application identifier
	// intended to unique to the app and independent of the exchange.
	// On Android, this should be a bundle or package name
	// e.g "com.foo.mygame". On iOS, it is a numeric ID.
	Bundle string `json:"bundle"`

	// URL without cache-busting to an image that is representative
	// of the  content of the campaign for ad quality/safety checking.
	ImageURL string `json:"iurl"`

	// Campaign Id to assist with ad quality checking; the
	// collection of creatives for which iurl should be representative.
	CampaignId string `json:"cid"`

	// Creative Id to assist with ad quality checking.
	CreativeId string `json:"crid"`

	// IAB content categories of the creative.
	ContentCategories []ContentCategory `json:"cat"`

	// Set of attributes describing the creative.
	Attributes []CreativeAttribute `json:"attr"`

	API APIFramework `json:"api"`

	VideoResponseProtocol VideoProtocol `json:"protocol"`

	// Creative Media rating per IQG guidelines.
	MediaRating int `json:"qagmediarating"`

	// Reference to the deal.id from the bid request if this
	// bid pertains to a private marketplace direct deal.
	DealId string `json:"dealid"`

	// Width of the creative in device independent pixels (DIPS).
	Width int `json:"w"`

	// Height of the creative in device independent pixels (DIPS).
	Height int `json:"h"`

	// Advisory as to the number of seconds the bidder is
	// willing to wait between the auction and the actual impression.
	ExpirationAuctionToActualImpression int `json:"exp"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}
