package odexchange

type ExchangeExtension interface{}

type BidRequest struct {
	Id string `json:"id"`

	// Atleast one impression is required
	Impressions []*Impression `json:"imp"`

	Site   *Site       `json:"site"`
	App    *App        `json:"app"`
	Device *Device     `json:"device"`
	User   *User       `json:"user"`
	Test   NumericBool `json:"test"`

	AuctionType AuctionType `json:"at"`
	MaxTimeout  int         `json:"tmax"`
	Currencies  []string    `json:"cur"`

	WhitelistedSeats []string    `json:"wseat"`
	AllImpressions   NumericBool `json:"allimps"`

	BlockedAdvertiserCategories []string `json:"bcat"`
	BlockedAdvertiserDomains    []string `json:"badv"`

	// Block list of applications by their platform-specific
	// exchange-independent application identifiers.
	// On Android, these should be bundle or package
	// names (e.g. com.foo.mygmae). On iOS, these are numeric IDs.
	BlockedApps []string    `json:"bapp"`
	Regulations *Regulation `json:"regs"`

	ExchangeExtension *ExchangeExtension `json:"ext"`
}

// NativeImpression ad units are intended to blend seamlessly into
// the surrounding content (e.g a sponsored Twitter or Facebook post).
//
// The Native Subcommittee has developed a companion specification to OpenRTB
// called the Dynamic Native Ads API. It defines the request parameters and
// response markup structure of native ad units.
// It provides the means of transporting request parameters as an opaque string
// so that the ad markup served will be structured according to that specification.
type NativeImpression struct {
	// Request payload complying with the Native Ad Specification.
	Request string `json:"request"`

	// Version of the Dynamic Native Ads API to which request complies, highly
	// recommended for efficient parsing.
	Version string `json:"ver"`

	// List of supported API frameworks for this impression.
	// Refer to List 5.6. If an API is not explicitly listed,
	// it is assumed not to be supported.
	SupportedAPIFrameworks []APIFramework `json:"api"`

	BlockedCreativeAttributes []CreativeAttribute `json:"battr"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

// Site should be included if the ad supported content is a website as
// opposed to a non-browser application. A bid request must not contain both
// a Site and an App object. At a minimum, it is useful to provide a site ID
// or page URL, but this is not strictly required.
type Site struct {
	// Exchange-specific site ID.
	Id string `json:"id"`

	// Site name(may be aliased at the publisher's request).
	Name string `json:"name"`

	// Domain of the site (e.g "mysite.foo.com").
	Domain string `json:"domain"`

	// IAB content categories of the site.
	Categorizes []string `json:"cat"`

	// IAB content categories that describe
	// the current section of the site.
	SectionCategories []string `json:"sectioncat"`

	// IAB content categories that describe
	// the current page or view of the site.
	PageCategories []string `json:"pagecat"`

	// URL of the page where the impression will be shown.
	Page string `json:"page"`

	// Referrer URL that caused navigation to the current page.
	ReferrerURL string `json:"ref"`

	// Search string that caused navigation to the current page.
	SearchString string `json:"search"`

	OptimizedForMobile NumericBool `json:"mobile"`

	HasPrivacyPolicy NumericBool `json:"privacypolicy"`

	Publisher *Publisher `json:"publisher"`

	Content *Content `json:"content"`

	// Comma separated list of keywords about the site.
	Keywords string `json:"keywords"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Format struct {
	// Width in Device Independent Pixels(DIPS).
	WidthDIPS int `json:"w"`

	// Height in Device Independent Pixels(DIPS).
	HeightDIPS int `json:"h"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Audio struct {
	// Content MIME types supported (e.g "audio/mp4")
	SupportedMimeTypes []string `json:"mimes"`

	MinDurationSeconds int             `json:"minduration"`
	MaxDurationSeconds int             `json:"maxduration"`
	SupportedProtocols []VideoProtocol `json:"protocols"`

	Width  int `json:"w"`
	Height int `json:"h"`

	StartDelay int `json:"startdelay"`
	Linearity  int `json:"linearity"`

	Skip NumericBool `json:"skip"`

	MinThresholdForSkipSeconds int `json:"skipmin"`
	SkipAfterSeconds           int `json:"skipafter"`

	Sequence int `json:"sequence"`

	BlockedCreativeAttributes []CreativeAttribute `json:"battr"`

	// The maximum number of ads that can be played in an ad pod.
	MaxSequence int `json:"maxseq"`

	Feed FeedType `json:"feed"`

	Stitched NumericBool `json:"stitched"`

	DownloadableByUser NumericBool `json:"dl"`

	MinBitRateKbps int `json:"minbitrate"`
	MaxBitRateKbps int `json:"maxbitrate"`

	// Indicates if letter-boxing of 4:3
	// content into a 16:9 window is allowed.
	BoxingAllowed NumericBool `json:"boxingallowed"`

	PlaybackMethods []PlaybackMethod `json:"playbackmethod"`

	// Supported delivery methods (e.g streaming, progressive).
	// If none specified, assume all are supported.
	DeliveryMethods []Delivery `json:"delivery"`

	// Ad position on screen.
	Postion AdPosition `json:"pos"`

	// List of supported API frameworks for this impression.
	// Refer to List 5.6. If an API is not explicitly listed,
	// it is assumed not to be supported.
	SupportedAPIFrameworks []APIFramework `json:"api"`

	CompanionAdTypes      []*Banner         `json:"companionad"`
	DAASTCompanionAdTypes []CompanionAdType `json:"companiontype"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Video struct {
	// Content MIME types supported. Popular MIME types may
	// include "video/x-ms-wmv" for Windows Media and
	// "video/x-flv" for Flash Video.
	SupportedMimeTypes []string `json:"mimes"`

	MinDurationSeconds int             `json:"minduration"`
	MaxDurationSeconds int             `json:"maxduration"`
	SupportedProtocols []VideoProtocol `json:"protocols"`

	Width  int `json:"w"`
	Height int `json:"h"`

	StartDelay int `json:"startdelay"`
	Linearity  int `json:"linearity"`

	Skip NumericBool `json:"skip"`

	MinThresholdForSkipSeconds int `json:"skipmin"`
	SkipAfterSeconds           int `json:"skipafter"`

	Sequence int `json:"sequence"`

	BlockedCreativeAttributes []CreativeAttribute `json:"battr"`

	MaxExtendedAdDurationSeconds int `json:"maxextended"`

	MinBitRateKbps int `json:"minbitrate"`
	MaxBitRateKbps int `json:"maxbitrate"`

	// Indicates if letter-boxing of 4:3
	// content into a 16:9 window is allowed.
	BoxingAllowed NumericBool `json:"boxingallowed"`

	PlaybackMethods []PlaybackMethod `json:"playbackmethod"`

	// Supported delivery methods (e.g streaming, progressive).
	// If none specified, assume all are supported.
	DeliveryMethods []Delivery `json:"delivery"`

	// Ad position on screen.
	Postion AdPosition `json:"pos"`

	// List of supported API frameworks for this impression.
	// Refer to List 5.6. If an API is not explicitly listed,
	// it is assumed not to be supported.
	SupportedAPIFrameworks []APIFramework `json:"api"`

	VASTCompanionAdTypes []CompanionAdType `json:"companiontype"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Interstital int

const (
	NonInterstital Interstital = 0
	FullScreen     Interstital = 1
)

type Currency string

type ClickBrowser int

const (
	ClickEmbedded ClickBrowser = 0
	ClickNative   ClickBrowser = 1
)

type Banner struct {
	Id     string `json:"id"`
	Width  int    `json:"w"`
	Height int    `json:"h"`

	Formats []*Format `json:"format"`

	BlockedBannerAdTypes      []BannerAdType      `json:"btype"`
	BlockedCreativeAttributes []CreativeAttribute `json:"battr"`

	Position AdPosition `json:"pos"`

	SupportedMimeTypes []string `json:"mimes"`

	TopFrame         int                   `json:"topframe"`
	ExpandDirections []ExpandableDirection `json:"expdir"`

	// List of supported API frameworks for this impression.
	// Refer to List 5.6. If an API is not explicitly listed,
	// it is assumed not to be supported.
	SupportedAPIFrameworks []APIFramework `json:"api"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Impression struct {
	Id     string  `json:"id"`
	Banner *Banner `json:"banner"`
	Video  *Video  `json:"video"`
	Audio  *Audio  `json:"audio"`

	Native *NativeImpression `json:"native"`

	PrivateMarketPlaceDeals *PrivateMarketPlaceDeal `json:"pmp"`

	DisplayManager        string `json:"displaymanager"`
	DisplayManagerVersion string `json:"displaymanagerver"`

	Interstital Interstital `json:"interstital"`

	TagId string `json:"tagid"`

	// Minimum bid for this impression expressed in CPM.
	BidFloorCPM float64 `json:"bidfloor"`

	// Currency specified using ISO-4217 alpha codes.
	// This may be different from bid currency
	// returned by bidder if this allowed by exchange.
	BidFloorCurrency Currency `json:"bidfloorcur"`

	// Indicates the type of browser opened upon clicking the
	// creative in an app, where:
	//   0 = embedded
	//   1 = native
	// Note that the Safari View Controller in iOS 9.x devices is
	// considered a native browser for purposes of this attribute.
	ClickBrowser ClickBrowser `json:"clickbrowser"`

	Secure NumericBool `json:"secure"`

	// Array of exchange-specific names of supported iframe busters.
	IFrameBuster []string `json:"iframebuster"`

	ExpirationAuctionToActualImpression int `json:"exp"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type App struct {
	// Exchange-specific app ID.
	Id string `json:"id"`

	// App name(may be aliased at the publisher's request).
	Name string `json:"name"`

	// A platform-specific application identifier intended to be
	// unique to the app and independent of the exchange. On
	// Android, this should be a bundle or package name
	// e.g "com.foo.mygame". On iOS, it is a numeric ID.
	Bundle string `json:"bundle"`

	// Domain of the app e.g "mygame.foo.com".
	Domain string `json:"domain"`

	// App store URL for an installed app; for IQG 2.1 compliance.
	StoreURL string `json:"storeurl"`

	// IAB content categories of the site.
	Categorizes []string `json:"cat"`

	// IAB content categories that describe
	// the current section of the site.
	SectionCategories []string `json:"sectioncat"`

	// IAB content categories that describe
	// the current page or view of the site.
	PageCategories []string `json:"pagecat"`

	Version string `json:"ver"`

	HasPrivacyPolicy NumericBool `json:"privacypolicy"`

	// If set signals that the app is a paid version, else is free.
	PaidVersion NumericBool `json:"paid"`

	Publisher *Publisher `json:"publisher"`

	Content *Content `json:"content"`

	Keywords string `json:"keywords"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Publisher struct {
	Id string `json:"id"`

	Name string `json:"name"`

	// Domain of the app e.g "mygame.foo.com".
	Domain string `json:"domain"`

	// IAB content categories of the site.
	Categorizes []string `json:"cat"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Content struct {
	Id            string `json:"id"`
	EpisodeNumber int    `json:"episode"`

	// Content title.
	// Video Examples:
	// + "Seach Committee" (television)
	// + "A New Hope" (movie)
	// + "Endgame" (made for web)
	// Non-Video Example:
	// + "Why an Antarctic Glacier is Melting So Quickly" (Time magazine article)
	Title string `json:"title"`

	Series string `json:"series"`

	// Season example: "Season 3"
	Season string `json:"season"`

	// Artist credited with the content.
	Artist string `json:"artist"`

	// Genre that best describe the content e.g rock, pop etc.
	Genre string `json:"genre"`

	// Album to which the content belongs; typically for audio.
	Album string `json:"album"`

	// ISRC is the International Standard Recording Code
	// conforming to ISO 3901.
	ISRC string `json:"isrc"`

	Producer *Producer `json:"producer"`

	URL string `json:"url"`

	// IAB content categories of the site.
	Categorizes []string `json:"cat"`

	ProductionQuality ProductionQuality `json:"prodq"`

	// Type of content e.g game, video, text etc.
	Context ContentContext `json:"context"`

	// Content rating e.g MPAA.
	ContentRating string `json"contentrating"`

	// User rating of the content e.g number of stars, likes etc.
	UserRating string `json:"userrating"`

	// Media rating per IQG guidelines.
	MediaRating int `json:"qagmediarating"`

	Keywords string `json:"keywords"`

	IsLiveStream NumericBool `json:"livestream"`

	SourceRelationship SourceRelationship `json:"sourcerelationship"`

	DurationSeconds int `json:"len"`

	// Content language using ISO-639-1-alpha-2.
	Language string `json:"language"`

	Embeddable NumericBool `json:"embeddable"`

	// Additional content data. Each Data object
	// represents a different data source.
	Data []*Data `json:"data"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type SourceRelationShip uint

const (
	IndirectSourceRelationShip = 0
	DirectSourceRelationShip   = 1
)

type Producer struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	// Domain of the site (e.g "mysite.foo.com").
	Domain string `json:"domain"`

	// IAB content categories of the site.
	Categorizes []string `json:"cat"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Device struct {
	UserAgent string `json:"ua"`

	CurrentLocation *Geo `json:"geo"`

	LimitAdTracking NumericBool `json:"lmt"`

	// IPv4 address closest to device.
	IP string `json:"ipv4"`

	// IPv6 address closest to device.
	IPV6 string `json:"ipv6"`

	DeviceType DeviceType `json:"devicetype"`

	// Device make e.g "Apple".
	Make string `json:"make"`

	// Device model e.g "iPhone".
	Model string `json:"model"`

	// Device operating system e.g "iOS".
	OS string `json:"os"`

	// Device operating system version e.g "3.1.2"
	OSVersion string `json:"osv"`

	// Hardware version of the device e.g "5S" for iPhone 5S.
	HardwareVersion string `json:"hwv"`

	// Physical height of the screen in pixels.
	Height int `json:"h"`

	// Physical width of the screen in pixels.
	Width int `json:"w"`

	// PPI is the Pixels Per Linear Inch.
	PPI int `json:"ppi"`

	// The ratio of physical pixels to device independent pixels.
	PXRatio float32 `json:"pxratio"`

	SupportsJS NumericBool `json:"js"`

	// Indicates if the geolocation API will be available
	// to Javascript code running in the banner.
	GeoLocationAPIAvailableToJS NumericBool `json:"geofetch"`

	// Version of Flash supported by the browser.
	FlashVersion string `json:"flashver"`

	// Browser language using ISO-639-1-alpha-2.
	Language string `json:"language"`

	// Carrier of ISP(e.g "VERIZON"). "WIFI" is often used in mobile
	// to indicate high bandwidth e.g video friendly vs cellular.
	Carrier string `json:"carrier"`

	ConnectionType ConnectionType `json:"connectiontype"`

	// ID sanctioned for advertiser use in the clear ie not hashed.
	AdvertiserIDUnhashed string `json:"ifa"`

	// Hardware device ID e.g IMEI hashed via SHA1.
	HardwareDeviceIDSHA1 string `json:"didsha1"`

	// Hardware device ID e.g IMEI hashed via MD5.
	HardwareDeviceIDMD5 string `json:"didmd5"`

	// Platform device ID e.g Android ID hashed via SHA1.
	PlatformDeviceIDSHA1 string `json:"dpidsha1"`

	// Platform device ID e.g Android ID hashed via MD5.
	PlatformDeviceIDMD5 string `json:"dpidmd5"`

	// MAC address of the device, hased via SHA1.
	MACSHA1 string `json:"macsha1"`

	// MAC address of the device, hased via MD5.
	MACMD5 string `json:"macmd5"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Geo struct {
	// Latitude from -90.0 to +90.0, where negative is South.
	Lat float32 `json:"lat"`

	// Latitude from -180.0 to +180.0, where negative is West.
	Lon float32 `json:"lon"`

	// Source of location data; recommended when passing lat/lon.
	LocationSource LocationSource `json:"type"`

	// Estimated location accuracy in meters; recommended when
	// lat/lon are specified and derived from a device's location
	// services ie type = 1. Note that this is the accuracy as reported
	// from the device. Consult OS specific documentation for exact
	// interpretation.
	EstimatedAccuracyMeters int `json:"accuracy"`

	// Number of seconds since this geolocation fix was established.
	// Note that devices may cache location data across multiple fetches.
	// Ideally, this value should be from the time the actual fix was taken.
	SecondsSinceLastFix int `json:"lastfix"`

	// Service or provider used to determine
	// geolocation from IP address if applicable.
	IPService int `json:"ipservice"`

	// Country code using ISO-3166-1-alpha-3.
	Country string `json:"country"`

	// Region code using ISO-3166-2: 2-letter state code if USA.
	Region string `json:"region"`

	// Region of a country using FIPS 10-4 notation. While OpenRTB
	// supports this attribute, it was withdrawn by NIST in 2008.
	RegionFIPS104 string `json:"regionfips104"`

	// Google Metro code; similar to but not exactly Nielsen DMAs.
	Metro string `json:"metro"`

	// City using United Nations Code for Trade & Transport locations.
	City string `json:"city"`

	// ZIP or Postal code.
	ZIP string `json:"zip"`

	// Local time as the number +/- of minutes from UTC.
	UTCOffsetMinutes int `json:"utcoffset"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type User struct {
	// Exchange-specific ID for the user.
	// At least one of "id" or "buyeruid" is recommended.
	Id string `json:"id"`

	// Buyer-specific ID for the user as mapped by the exchange for
	// the buyer. At least one of "buyeruid" or "id" is recommended.
	BuyerUID string `json:"buyeruid"`

	YearOfBirth int `json:"yob"`

	Gender string `json:"gender"`

	// Comma separated list of keywords, interest or intent.
	Keywords string `json:"keywords"`

	CustomData string `json:"customdata"`

	// Location of the user's home base defined by a Geo object.
	// This is not necessarily their current location.
	Geo *Geo `json:"geo"`

	// Additional user data. Each Data object
	// represents a different data source.
	Data []*Data `json:"data"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Data struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Segments []*Segment `json:"segment"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Segment struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Value string `json:"value"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Regulation struct {
	COPPA NumericBool `json:"coppa"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type PrivateMarketPlaceDeal struct {
	PrivateAuction NumericBool `json:"private_auction"`

	Deals []*Deal `json:"deals"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type Deal struct {
	Id string `json:"id"`

	// Minimum bid for this impression expressed in CPM.
	BidFloorCPM float64 `json:"bidfloor"`

	// Currency specified using ISO-4217 alpha codes.
	// This may be different from bid currency
	// returned by bidder if this allowed by exchange.
	BidFloorCurrency Currency `json:"bidfloorcur"`

	// Option override of the overall auction type of the bid
	// request, where:
	// + 1 = First Price.
	// + 2 = Second Price Plus.
	// + 3 = the value passed in "bidfloor" is the agreed upon deal
	//       price.
	// Additional auction types can be defined by the exchange.
	AuctionType AuctionType `json:"at"`

	// Whitelist of buyer seat e.g advertisers, agencies
	// allowed to bid on this deal. IDs of seats and knowledge of the
	// buyer's customers to which they refer must be coordinated
	// between bidders and the exchange a priori. Omission implies
	// no seat restrictions.
	WhitelistedSeats []string `json:"wseat"`

	// Advertiser domains allowed to bid on this deal.
	// Omission implies no advertiser restrictions.
	WhitelistedDomains []string `json:"wadomain"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	ExchangeExtension *ExchangeExtension `json:"ext"`
}

type CreativeAttribute uint

const (
	Unknown CreativeAttribute = iota
	AudioAdAutoPlay
	AudioAdUserInitiated
	ExpandableAutomatic
	ExpandableUserInitiatedClick
	ExpandableUserInitiatedRollover
	InBannerVideoAdAutoPlay
	InBannerVideoAdUserInitiated
	Pop
	ProvocativeImagery
	ExtremeAnimation // Shaky, Flashing, Flickering, Extreme Animation, Smileys
	Surveys
	TextOnly
	UserInteractive // e.g Embedded Games.
	AlertStyle
	HasAudioToggle
	ProvidesSkipping
	AdobeFlash
)

type ExpandableDirection uint

const (
	ExpandLeft       ExpandableDirection = 1
	ExpandRight      ExpandableDirection = 2
	ExpandUp         ExpandableDirection = 3
	ExpandDown       ExpandableDirection = 4
	ExpandFullscreen ExpandableDirection = 5
)

type APIFramework uint

const (
	APIVPAID1DOT0 APIFramework = 1
	APIVPAID2DOT0 APIFramework = 2
	APIMRAID1     APIFramework = 3
	APIORMMA      APIFramework = 4
	APIMRAID2     APIFramework = 5
)

type VideoProtocol uint

const (
	ProtoVAST1DOT0        VideoProtocol = 1
	ProtoVAST2DOT0        VideoProtocol = 2
	ProtoVAST3DOT0        VideoProtocol = 3
	ProtoVAST1DOT0Wrapper VideoProtocol = 4
	ProtoVAST2DOT0Wrapper VideoProtocol = 5
	ProtoVAST3DOT0Wrapper VideoProtocol = 6
	ProtoVAST4DOT0        VideoProtocol = 7
	ProtoDAAST1DOT0       VideoProtocol = 8
)

type AdPosition uint

const (
	PosUnknown    AdPosition = 0
	PosAboveFold  AdPosition = 1
	PosDeprecated AdPosition = 2
	PosBelowFold  AdPosition = 3
	PosHeader     AdPosition = 4
	PosFooter     AdPosition = 5
	PosSideBar    AdPosition = 6
	PosFullscreen AdPosition = 7
)

type BannerAdType uint

const (
	AdXHTMLText   BannerAdType = 1
	AdXHTMLBanner BannerAdType = 2
	AdJavascript  BannerAdType = 3
	AdIFrame      BannerAdType = 4
)

type PlaybackMethod uint

const (
	PlaybackAutoPlaySoundOn  PlaybackMethod = 1
	PlaybackAutoPlaySoundOff PlaybackMethod = 2
	PlaybackClickToPlay      PlaybackMethod = 3
	PlaybackMouseOver        PlaybackMethod = 4
)

type CompanionAdType uint

const (
	CompanionStaticResource CompanionAdType = 1
	CompanionHTMLResource   CompanionAdType = 2
	CompanionIFrameResource CompanionAdType = 3
)

type Delivery uint

const (
	DeliveryStreaming   Delivery = 1
	DeliveryProgressive Delivery = 2
)

type FeedType uint

const (
	FeedSingle          FeedType = 1
	FeedMultiple        FeedType = 2
	FeedFMOrAMBroadcast FeedType = 3
	FeedPodcast         FeedType = 4
)

type ContentContext uint

const (
	ContextVideo       ContentContext = 1
	ContextGame        ContentContext = 2
	ContextMusic       ContentContext = 3
	ContextApplication ContentContext = 4
	ContextText        ContentContext = 5
	ContextOther       ContentContext = 6
	ContextUnknown     ContentContext = 7
)

// AuctionType definitions:
// + 1 = First Price.
// + 2 = Second Price Plus.
// + 3 = the value passed in "bidfloor" is the agreed upon deal
//       price.
// Additional auction types can be defined by the exchange.
type AuctionType int

const (
	AuctionFirstPlace  AuctionType = 1
	AuctionSecondPlace AuctionType = 2
	AuctionThirdPlace  AuctionType = 3
)

type ProductionQuality uint

const (
	QualityUnknown                ProductionQuality = 0
	QualityProfessionallyProduced ProductionQuality = 1
	QualityProsumer               ProductionQuality = 2
	QualityUserGeneratedContent   ProductionQuality = 3
)

type SourceRelationship uint

const (
	RelationIndirect SourceRelationship = 0
	RelationDirect   SourceRelationship = 1
)

type LocationSource uint

const (
	SourceGPSLocationServices LocationSource = 1
	SourceIPAddress           LocationSource = 2
	SourceUserProvided        LocationSource = 3
)

type DeviceType uint

const (
	DeviceMobileOrTablet  DeviceType = 1
	DevicePC              DeviceType = 2
	DeviceConnectedTV     DeviceType = 3
	DevicePhone           DeviceType = 4
	DeviceTable           DeviceType = 5
	DeviceConnectedDevice DeviceType = 6
	DeviceSetTopBox       DeviceType = 7
)

type ConnectionType uint

const (
	ConnUnknown                   ConnectionType = 0
	ConnEthernet                  ConnectionType = 1
	ConnWIFI                      ConnectionType = 2
	ConnCellularUnknownGeneration ConnectionType = 3
	ConnCellular2G                ConnectionType = 4
	ConnCellular3G                ConnectionType = 5
	ConnCellular4G                ConnectionType = 6
)
