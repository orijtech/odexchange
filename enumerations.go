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

type WinSelection uint

const (
	WinIndividually  WinSelection = 0
	WinOrLoseAsGroup WinSelection = 1
)

type ContentCategory string

const (
	CArtsAndEntertainment ContentCategory = "IAB1"
	CBooksAndLiterature   ContentCategory = "IAB1-1"
	CCelebrityFanOrGossip ContentCategory = "IAB1-2"
	CFineArt              ContentCategory = "IAB1-3"
	CHumor                ContentCategory = "IAB1-4"
	CMovies               ContentCategory = "IAB1-5"
	CMusic                ContentCategory = "IAB1-6"
	CTelevision           ContentCategory = "IAB1-7"

	CAutomative           ContentCategory = "IAB2"
	CAutoParts            ContentCategory = "IAB2-1"
	CAutoRepair           ContentCategory = "IAB2-2"
	CBuyingOrSellingCars  ContentCategory = "IAB2-3"
	CCarCulture           ContentCategory = "IAB2-4"
	CCertifiedPreOwned    ContentCategory = "IAB2-5"
	CConvertible          ContentCategory = "IAB2-6"
	CCoupe                ContentCategory = "IAB2-7"
	CCrossover            ContentCategory = "IAB2-8"
	CDiesel               ContentCategory = "IAB2-9"
	CElectricVehicle      ContentCategory = "IAB2-10"
	CHatchback            ContentCategory = "IAB2-11"
	CHybrid               ContentCategory = "IAB2-12"
	CLuxury               ContentCategory = "IAB2-13"
	CMinivan              ContentCategory = "IAB2-14"
	CMotorCycles          ContentCategory = "IAB2-15"
	COffRoadVehicles      ContentCategory = "IAB2-16"
	CPerformanceVehicles  ContentCategory = "IAB2-17"
	CPickup               ContentCategory = "IAB2-18"
	CRoadSideAssistance   ContentCategory = "IAB2-19"
	CSedan                ContentCategory = "IAB2-20"
	CTrucksAndAccessories ContentCategory = "IAB2-21"
	CVintageCars          ContentCategory = "IAB2-22"
	CWagon                ContentCategory = "IAB2-23"

	CBusiness            ContentCategory = "IAB3"
	CAdvertising         ContentCategory = "IAB3-1"
	CAgriculture         ContentCategory = "IAB3-2"
	CBiotechOrBiomedical ContentCategory = "IAB3-3"
	CBusinessSoftware    ContentCategory = "IAB3-4"
	CConstruction        ContentCategory = "IAB3-5"
	CForestry            ContentCategory = "IAB3-6"
	CGovernment          ContentCategory = "IAB3-7"
	CGreenSolutions      ContentCategory = "IAB3-8"
	CHumanResources      ContentCategory = "IAB3-9"
	CLogistics           ContentCategory = "IAB3-10"
	CMarketing           ContentCategory = "IAB3-11"
	CMetals              ContentCategory = "IAB3-12"

	CCareers               ContentCategory = "IAB4"
	CCareerPlanning        ContentCategory = "IAB4-1"
	CCollege               ContentCategory = "IAB4-2"
	CFinancialAid          ContentCategory = "IAB4-3"
	CJobFairs              ContentCategory = "IAB4-4"
	CJobSearch             ContentCategory = "IAB4-5"
	CResumeWritingOrAdvice ContentCategory = "IAB4-6"
	CNursing               ContentCategory = "IAB4-7"
	CScholarships          ContentCategory = "IAB4-8"
	CTelecommuting         ContentCategory = "IAB4-9"
	CUSMilitary            ContentCategory = "IAB4-10"
	CCareerAdvice          ContentCategory = "IAB4-11"

	CEducation                ContentCategory = "IAB5"
	C7_12Education            ContentCategory = "IAB5-1"
	CAdultEducation           ContentCategory = "IAB5-2"
	CArtHistory               ContentCategory = "IAB5-3"
	CCollegeAdministration    ContentCategory = "IAB5-4"
	CCollegeLife              ContentCategory = "IAB5-5"
	CDistanceLearning         ContentCategory = "IAB5-6"
	CEnglishAsASecondLanguage ContentCategory = "IAB5-7"
	CLanguageLearning         ContentCategory = "IAB5-8"
	CGraduateSchool           ContentCategory = "IAB5-9"
	CHomeschooling            ContentCategory = "IAB5-10"
	CHomeworkOrStudyTips      ContentCategory = "IAB5-11"
	CK6Educators              ContentCategory = "IAB5-12"
	CPrivateSchool            ContentCategory = "IAB5-13"
	CSpecialEducation         ContentCategory = "IAB5-14"
	CStudyingBusiness         ContentCategory = "IAB5-15"

	CFamilyAndParenting ContentCategory = "IAB6"
	CAdoption           ContentCategory = "IAB6-1"
	CBabiesAndToddlers  ContentCategory = "IAB6-2"
	CDaycareOrPreSchool ContentCategory = "IAB6-3"
	CFamilyInternet     ContentCategory = "IAB6-4"
	CParenting_K6Kids   ContentCategory = "IAB6-5"
	CParentingTeens     ContentCategory = "IAB6-6"
	CPregnancy          ContentCategory = "IAB6-7"
	CSpecialNeedsKids   ContentCategory = "IAB6-8"
	CElderCare          ContentCategory = "IAB6-9"

	CHealthAndFitness        ContentCategory = "IAB7"
	CExercise                ContentCategory = "IAB7-1"
	CADD                     ContentCategory = "IAB7-2"
	CAIDSOrHIV               ContentCategory = "IAB7-3"
	CAllergies               ContentCategory = "IAB7-4"
	CAlternativeMedicine     ContentCategory = "IAB7-5"
	CArthritis               ContentCategory = "IAB7-6"
	CAsthma                  ContentCategory = "IAB7-7"
	CAutismOrPDD             ContentCategory = "IAB7-8"
	CBiPolarDisorder         ContentCategory = "IAB7-9"
	CBrainTumor              ContentCategory = "IAB7-10"
	CCancer                  ContentCategory = "IAB7-11"
	CCholesterol             ContentCategory = "IAB7-12"
	CChronicFatigueSyndrome  ContentCategory = "IAB7-13"
	CChronicPain             ContentCategory = "IAB7-14"
	CColdAndFlu              ContentCategory = "IAB7-15"
	CDeafness                ContentCategory = "IAB7-16"
	CDentalCare              ContentCategory = "IAB7-17"
	CDepression              ContentCategory = "IAB7-18"
	CDermatology             ContentCategory = "IAB7-19"
	CDiabetes                ContentCategory = "IAB7-20"
	CEpilepsy                ContentCategory = "IAB7-21"
	CGERDOrAcidReflux        ContentCategory = "IAB7-22"
	CHeadachesOrMigraines    ContentCategory = "IAB7-23"
	CHeartDisease            ContentCategory = "IAB7-24"
	CHerbsForHealth          ContentCategory = "IAB7-25"
	CHolisticHealing         ContentCategory = "IAB7-26"
	CIBSOrCrohnsDisease      ContentCategory = "IAB7-27"
	CIncestOrAbuseSupport    ContentCategory = "IAB7-28"
	CIncontinence            ContentCategory = "IAB7-29"
	CInfertility             ContentCategory = "IAB7-30"
	CMensHealth              ContentCategory = "IAB7-31"
	CNutrition               ContentCategory = "IAB7-32"
	COrthopedics             ContentCategory = "IAB7-33"
	CPanicOrAnxietyDisorders ContentCategory = "IAB7-34"
	CPediatrics              ContentCategory = "IAB7-35"
	CPhysicalTherapy         ContentCategory = "IAB7-36"
	CPsychologyOrPsychiatry  ContentCategory = "IAB7-37"
	CSeniorHealth            ContentCategory = "IAB7-38"
	CSexuality               ContentCategory = "IAB7-39"
	CSleepDisorders          ContentCategory = "IAB7-40"
	CSmokingCessation        ContentCategory = "IAB7-41"
	CSubstanceAbuse          ContentCategory = "IAB7-42"
	CThyroidDisease          ContentCategory = "IAB7-43"
	CWeightLoss              ContentCategory = "IAB7-44"
	CWomensHealth            ContentCategory = "IAB7-45"

	CFoodAndDrink          ContentCategory = "IAB8"
	CAmericanCuisine       ContentCategory = "IAB8-1"
	CBarbecuesAndGrilling  ContentCategory = "IAB8-2"
	CCajunOrCreole         ContentCategory = "IAB8-3"
	CChineseCuisine        ContentCategory = "IAB8-4"
	CCocktailsOrBeer       ContentCategory = "IAB8-5"
	CCoffeeOrTea           ContentCategory = "IAB8-6"
	CCuisineSpecific       ContentCategory = "IAB8-7"
	CDessertsAndBaking     ContentCategory = "IAB8-8"
	CDiningOut             ContentCategory = "IAB8-9"
	CFoodAllergies         ContentCategory = "IAB8-10"
	CFrenchCuisine         ContentCategory = "IAB8-11"
	CHealthOrLowFatCooking ContentCategory = "IAB8-12"
	CItalianCuisine        ContentCategory = "IAB8-13"
	CJapaneseCuisine       ContentCategory = "IAB8-14"
	CMexicanCuisine        ContentCategory = "IAB8-15"
	CVegan                 ContentCategory = "IAB8-16"
	CVegetarian            ContentCategory = "IAB8-17"
	CWine                  ContentCategory = "IAB8-18"

	CHobbiesAndInterests   ContentCategory = "IAB9"
	CArtOrTechnology       ContentCategory = "IAB9-1"
	CArtsAndCrafts         ContentCategory = "IAB9-2"
	CBeadwork              ContentCategory = "IAB9-3"
	CBirdWatching          ContentCategory = "IAB9-4"
	CBoardGamesOrPuzzles   ContentCategory = "IAB9-5"
	CCandleAndSoapMaking   ContentCategory = "IAB9-6"
	CCardGames             ContentCategory = "IAB9-7"
	CChess                 ContentCategory = "IAB9-8"
	CCigars                ContentCategory = "IAB9-9"
	CCollecting            ContentCategory = "IAB9-10"
	CComicBooks            ContentCategory = "IAB9-11"
	CDrawingOrSketching    ContentCategory = "IAB9-12"
	CFreelanceWriting      ContentCategory = "IAB9-13"
	CGenealogy             ContentCategory = "IAB9-14"
	CGettingPublished      ContentCategory = "IAB9-15"
	CGuitar                ContentCategory = "IAB9-16"
	CHomeRecording         ContentCategory = "IAB9-17"
	CInvestorsAndPatents   ContentCategory = "IAB9-18"
	CJewelryMaking         ContentCategory = "IAB9-19"
	CMagicAndIllusion      ContentCategory = "IAB9-20"
	CNeedlework            ContentCategory = "IAB9-21"
	CPainting              ContentCategory = "IAB9-22"
	CPhotography           ContentCategory = "IAB9-23"
	CRadio                 ContentCategory = "IAB9-24"
	CRolePlayingGames      ContentCategory = "IAB9-25"
	CSciFiAndFantasy       ContentCategory = "IAB9-26"
	CScrapbooking          ContentCategory = "IAB9-27"
	CScreenwriting         ContentCategory = "IAB9-28"
	CStampsAndCoins        ContentCategory = "IAB9-29"
	CVideoAndComputerGames ContentCategory = "IAB9-30"
	CWoodworking           ContentCategory = "IAB9-31"

	CHomeAndGardening          ContentCategory = "IAB10"
	CAppliances                ContentCategory = "IAB10-1"
	CEntertaining              ContentCategory = "IAB10-2"
	CEnvironmentalSafety       ContentCategory = "IAB10-3"
	CGardening                 ContentCategory = "IAB10-4"
	CHomeRepair                ContentCategory = "IAB10-5"
	CHomeTheater               ContentCategory = "IAB10-6"
	CInteriorDecoratin         ContentCategory = "IAB10-7"
	CLandscaping               ContentCategory = "IAB10-8"
	CRemodelingAndConstruction ContentCategory = "IAB10-9"

	CLawGovernmentAndPolitics ContentCategory = "IAB11"
	CImmigration              ContentCategory = "IAB11-1"
	CLegalIssues              ContentCategory = "IAB11-2"
	CUSGovernmentResources    ContentCategory = "IAB11-3"
	CPolitics                 ContentCategory = "IAB11-4"
	CCommentary               ContentCategory = "IAB11-5"

	CNews              ContentCategory = "IAB12"
	CInternationalNews ContentCategory = "IAB12-1"
	CNationalNews      ContentCategory = "IAB12-2"
	CLocalNews         ContentCategory = "IAB12-3"

	CPersonalFinance      ContentCategory = "IAB13"
	CBeginningInvesting   ContentCategory = "IAB13-1"
	CCreditOrDebtAndLoans ContentCategory = "IAB13-2"
	CFinancialNews        ContentCategory = "IAB13-3"
	CFinancialPlanning    ContentCategory = "IAB13-4"
	CHedgeFund            ContentCategory = "IAB13-5"
	CInsurance            ContentCategory = "IAB13-6"
	CInvesting            ContentCategory = "IAB13-7"
	CMutualFunds          ContentCategory = "IAB13-8"
	COptions              ContentCategory = "IAB13-9"
	CRetirementPlanning   ContentCategory = "IAB13-10"
	CStocks               ContentCategory = "IAB13-11"
	CTaxPlanning          ContentCategory = "IAB13-12"

	CSociety        ContentCategory = "IAB14"
	CDating         ContentCategory = "IAB14-1"
	CDivorceSupport ContentCategory = "IAB14-2"
	CGayLife        ContentCategory = "IAB14-3"
	CMarriage       ContentCategory = "IAB14-4"
	CSeniorLiving   ContentCategory = "IAB14-5"
	CTeens          ContentCategory = "IAB14-6"
	CWeddings       ContentCategory = "IAB14-7"
	CEthnicSpecific ContentCategory = "IAB14-8"

	CScience             ContentCategory = "IAB15"
	CAstrology           ContentCategory = "IAB15-1"
	CBiology             ContentCategory = "IAB15-2"
	CChemistry           ContentCategory = "IAB15-3"
	CGeology             ContentCategory = "IAB15-4"
	CParanormalPhenomena ContentCategory = "IAB15-5"
	CPhysics             ContentCategory = "IAB15-6"
	CSpaceOrAstronomy    ContentCategory = "IAB15-7"
	CGeography           ContentCategory = "IAB15-8"
	CBotany              ContentCategory = "IAB15-9"
	CWeather             ContentCategory = "IAB15-10"

	CPets               ContentCategory = "IAB16"
	CAquariums          ContentCategory = "IAB16-1"
	CBirds              ContentCategory = "IAB16-2"
	CCats               ContentCategory = "IAB16-3"
	CDogs               ContentCategory = "IAB16-4"
	CLargeAnimals       ContentCategory = "IAB16-5"
	CReptiles           ContentCategory = "IAB16-6"
	CVeterinaryMedicine ContentCategory = "IAB16-7"

	CSports                ContentCategory = "IAB17"
	CAutoRacing            ContentCategory = "IAB17-1"
	CBaseball              ContentCategory = "IAB17-2"
	CBicyling              ContentCategory = "IAB17-3"
	CBodyBuilding          ContentCategory = "IAB17-4"
	CBoxing                ContentCategory = "IAB17-5"
	CCaneoingOrKayaking    ContentCategory = "IAB17-6"
	CCheerleading          ContentCategory = "IAB17-7"
	CClimbing              ContentCategory = "IAB17-8"
	CCricket               ContentCategory = "IAB17-9"
	CFigureSkating         ContentCategory = "IAB17-10"
	CFlyFishing            ContentCategory = "IAB17-11"
	CFootball              ContentCategory = "IAB17-12"
	CFreshWaterFishing     ContentCategory = "IAB17-13"
	CGameAndFish           ContentCategory = "IAB17-14"
	CGolf                  ContentCategory = "IAB17-15"
	CHorseRacing           ContentCategory = "IAB17-16"
	CHorses                ContentCategory = "IAB17-17"
	CHuntingOrShooting     ContentCategory = "IAB17-18"
	CInlineSkating         ContentCategory = "IAB17-19"
	CMartialArts           ContentCategory = "IAB17-20"
	CMountainBiking        ContentCategory = "IAB17-21"
	CNASCARRacing          ContentCategory = "IAB17-22"
	COlympics              ContentCategory = "IAB17-23"
	CPaintball             ContentCategory = "IAB17-24"
	CPowerAndMotorCycles   ContentCategory = "IAB17-25"
	CProBasketball         ContentCategory = "IAB17-26"
	CProIceHockey          ContentCategory = "IAB17-27"
	CRodeo                 ContentCategory = "IAB17-28"
	CRugby                 ContentCategory = "IAB17-29"
	CRunningOrJogging      ContentCategory = "IAB17-30"
	CSailing               ContentCategory = "IAB17-31"
	CSaltwaterFishing      ContentCategory = "IAB17-32"
	CScubaDiving           ContentCategory = "IAB17-33"
	CSkateboarding         ContentCategory = "IAB17-34"
	CSkiing                ContentCategory = "IAB17-35"
	CSnowboarding          ContentCategory = "IAB17-36"
	CSurfingOrBodyBoarding ContentCategory = "IAB17-37"
	CSwimming              ContentCategory = "IAB17-38"
	CTableTennisOrPingPong ContentCategory = "IAB17-39"
	CTennis                ContentCategory = "IAB17-40"
	CVolleyball            ContentCategory = "IAB17-41"
	CWalking               ContentCategory = "IAB17-42"
	CWaterskiOrWakeboard   ContentCategory = "IAB17-43"
	CWorldSoccer           ContentCategory = "IAB17-44"

	CStyleAndFashion ContentCategory = "IAB18"
	CBeauty          ContentCategory = "IAB18-1"
	CBodyArt         ContentCategory = "IAB18-2"
	CFashion         ContentCategory = "IAB18-3"
	CJewelry         ContentCategory = "IAB18-4"
	CClothing        ContentCategory = "IAB18-5"
	CAccessories     ContentCategory = "IAB18-6"

	CTechnologyAndComputing ContentCategory = "IAB19"
	C3DGraphics             ContentCategory = "IAB19-1"
	CAnimation              ContentCategory = "IAB19-2"
	CAntivirusSoftware      ContentCategory = "IAB19-3"
	CCOrCPlusPlus           ContentCategory = "IAB19-4"
	CCamerasAndCamcorders   ContentCategory = "IAB19-5"
	CCellPhones             ContentCategory = "IAB19-6"
	CComputerCertification  ContentCategory = "IAB19-7"
	CComputerNetworking     ContentCategory = "IAB19-8"
	CComputerPeripherals    ContentCategory = "IAB19-9"
	CComputerReviews        ContentCategory = "IAB19-10"
	CDataCenters            ContentCategory = "IAB19-11"
	CDatabases              ContentCategory = "IAB19-12"
	CDesktopPublishing      ContentCategory = "IAB19-13"
	CDesktopVideo           ContentCategory = "IAB19-14"
	CEmail                  ContentCategory = "IAB19-15"
	CGraphicsSoftware       ContentCategory = "IAB19-16"
	CHomeVideoOrDVD         ContentCategory = "IAB19-17"
	CInternetTechnology     ContentCategory = "IAB19-18"
	CJava                   ContentCategory = "IAB19-19"
	CJavascript             ContentCategory = "IAB19-20"
	CMacSupport             ContentCategory = "IAB19-21"
	CMP3OrMIDI              ContentCategory = "IAB19-22"
	CNetConferencing        ContentCategory = "IAB19-23"
	CNetForBeginners        ContentCategory = "IAB19-24"
	CNetworkSecurity        ContentCategory = "IAB19-25"
	CPalmtopsOrPDAs         ContentCategory = "IAB19-26"
	CPCSupport              ContentCategory = "IAB19-27"
	CPortable               ContentCategory = "IAB19-28"
	CEntertainment          ContentCategory = "IAB19-29"
	CSharewareOrFreeware    ContentCategory = "IAB19-30"
	CUnix                   ContentCategory = "IAB19-31"
	CVisualBasic            ContentCategory = "IAB19-32"
	CWebClipArt             ContentCategory = "IAB19-33"
	CWebDesignOrHTML        ContentCategory = "IAB19-34"
	CWebSearch              ContentCategory = "IAB19-35"
	CWindows                ContentCategory = "IAB19-36"

	CTravel                  ContentCategory = "IAB20"
	CAdventureTravel         ContentCategory = "IAB20-1"
	CAfrica                  ContentCategory = "IAB20-2"
	CAirTravel               ContentCategory = "IAB20-3"
	CAustraliaAndNewZealand  ContentCategory = "IAB20-4"
	CBedAndBreakfast         ContentCategory = "IAB20-5"
	CBudgetTravel            ContentCategory = "IAB20-6"
	CBusinessTravel          ContentCategory = "IAB20-7"
	CByUSLocale              ContentCategory = "IAB20-8"
	CCamping                 ContentCategory = "IAB20-9"
	CCanada                  ContentCategory = "IAB20-10"
	CCaribbean               ContentCategory = "IAB20-11"
	CCruises                 ContentCategory = "IAB20-12"
	CEasternEurope           ContentCategory = "IAB20-13"
	CEurope                  ContentCategory = "IAB20-14"
	CFrance                  ContentCategory = "IAB20-15"
	CGreece                  ContentCategory = "IAB20-16"
	CHoneymoonsOrGetaways    ContentCategory = "IAB20-17"
	CHotels                  ContentCategory = "IAB20-18"
	CItaly                   ContentCategory = "IAB20-19"
	CJapan                   ContentCategory = "IAB20-20"
	CMexicoAndCentralAmerica ContentCategory = "IAB20-21"
	CNationalParks           ContentCategory = "IAB20-22"
	CSouthAmerica            ContentCategory = "IAB20-23"
	CSpas                    ContentCategory = "IAB20-24"
	CThemeParks              ContentCategory = "IAB20-25"
	CTravelingWithKids       ContentCategory = "IAB20-26"
	CUnitedKingdom           ContentCategory = "IAB20-27"

	CRealEstate          ContentCategory = "IAB21"
	CApartments          ContentCategory = "IAB21-1"
	CArchitects          ContentCategory = "IAB21-2"
	CByingOrSellingHomes ContentCategory = "IAB21-3"

	CShopping            ContentCategory = "IAB22"
	CContestsAndFreebies ContentCategory = "IAB22-1"
	CCouponing           ContentCategory = "IAB22-2"
	CComparison          ContentCategory = "IAB22-3"
	CEngines             ContentCategory = "IAB22-4"

	CReligionAndSpirituality ContentCategory = "IAB23"
	CAlternativeReligions    ContentCategory = "IAB23-1"
	CAtheismOrAgnosticism    ContentCategory = "IAB23-2"
	CBuddhism                ContentCategory = "IAB23-3"
	CCatholicism             ContentCategory = "IAB23-4"
	CChristianity            ContentCategory = "IAB23-5"
	CHinduism                ContentCategory = "IAB23-6"
	CIslam                   ContentCategory = "IAB23-7"
	CJudaism                 ContentCategory = "IAB23-8"
	CLatterDaySaints         ContentCategory = "IAB23-9"
	CPaganOrWiccan           ContentCategory = "IAB23-10"

	CUncategorized                    ContentCategory = "IAB24"
	CNonStandardContent               ContentCategory = "IAB25"
	CUnmoderatedUGC                   ContentCategory = "IAB25-1"
	CExtremeGraphicOrExplicitViolence ContentCategory = "IAB25-2"
	CPornography                      ContentCategory = "IAB25-3"
	CProfaneContent                   ContentCategory = "IAB25-4"
	CHateContent                      ContentCategory = "IAB25-5"
	CUnderConstruction                ContentCategory = "IAB25-6"
	CIncentivized                     ContentCategory = "IAB25-7"

	CIllegalContent        ContentCategory = "IAB26"
	CIllegalContent_       ContentCategory = "IAB26-1"
	CWarez                 ContentCategory = "IAB26-2"
	CSpywareOrMalware      ContentCategory = "IAB26-3"
	CCopyrightInfringement ContentCategory = "IAB26-4"
)

type Linearity uint

const (
	LinearityInStream Linearity = 1
	LinearityOverlay  Linearity = 2
)
