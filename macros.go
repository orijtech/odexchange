package odexchange

// Macros
/*
----------------------------------------------------------------------------------------------------------
 Macro		     | Description
----------------------------------------------------------------------------------------------------------
 ${AUCTION_ID}	     | ID of the bid request; from BidRequest.id attribute
----------------------------------------------------------------------------------------------------------
 ${AUCTION_BID_ID}   | ID of the bid; from BidResponse.bidid attribute
----------------------------------------------------------------------------------------------------------
 ${AUCTION_IMP_ID}   | ID fo the impression just won; from imp.id attribute
----------------------------------------------------------------------------------------------------------
 ${AUCTION_SEAT_ID}  | ID of the bidder seat for whom the bid was made
----------------------------------------------------------------------------------------------------------
 ${AUCTION_AD_ID}    | ID of the ad markup the bidder wishes to serve; from bidid.adid atribute
----------------------------------------------------------------------------------------------------------
 ${AUCTION_PRICE}    | Settlement price using the same currency and units as the bid
----------------------------------------------------------------------------------------------------------
 ${AUCTION_CURRENCY} | The currency used in the bid(explicit or implied); for confirmation purposes only.
----------------------------------------------------------------------------------------------------------
*/
