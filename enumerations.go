package odexchange

// No-Bid Reason codes Section 5.22
type NoBidReason uint

const (
	NoBidUnknownError             NoBidReason = 0
	NoBidTechnicalError           NoBidReason = 1
	NoBidInvalidRequest           NoBidReason = 2
	NoBidKnownWebSpider           NoBidReason = 3
	NoBidSuspectedNonHumanTraffic NoBidReason = 4
	NoBidCloudDataCenterOrProxyIP NoBidReason = 5
	NoBidUnsupportedDevice        NoBidReason = 6
	NoBidBlockedPublisherOrSite   NoBidReason = 7
	NoBidUnmatchedUser            NoBidReason = 8
)

type Interstital int

const (
	NonInterstital Interstital = 0
	FullScreen     Interstital = 1
)

type ClickBrowser int

const (
	ClickEmbedded ClickBrowser = 0
	ClickNative   ClickBrowser = 1
)

type IPLocationService uint

const (
	IP2Location IPLocationService = 1
	Neustar     IPLocationService = 2
	MaxMind     IPLocationService = 3
)

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

type Relation uint

const (
	RelationIndirect Relation = 0
	RelationDirect   Relation = 1
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
