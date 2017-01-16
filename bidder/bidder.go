package bidder

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"

	"github.com/odeke-em/go-uuid"
	odex "github.com/odeke-em/odexchange"
	"github.com/odeke-em/semalim"
)

type Bidder struct {
	sync.RWMutex
	w  io.Writer
	id string

	curBidRespId uint64
}

func NewBidder() (*Bidder, error) {
	bdr := new(Bidder)
	bdr.id = uuid.NewRandom().String()
	return bdr, nil
}

func (bdr *Bidder) Write(b []byte) (int, error) {
	bdr.Lock()
	w := bdr.w
	bdr.Unlock()

	if w == nil {
		w = os.Stdout
	}

	return w.Write(b)
}

func (b *Bidder) newBidId() string {
	b.Lock()
	b.curBidRespId = atomic.AddUint64(&b.curBidRespId, 1)
	id := b.curBidRespId
	b.Unlock()

	return fmt.Sprintf("%d", id)
}

func (b *Bidder) CurBidId() uint64 {
	b.RLock()
	defer b.RUnlock()

	return b.curBidRespId
}

// Arbitrary timeout
const defaultMaxImpressionProcessingSeconds = 120

type imprJob struct {
	id int

	// We retain a reference to the bid
	// in order to extract extra information that might be useful
	// for us to process the impression.
	bid  *odex.BidRequest
	impr *odex.Impression
}

var errUnimplemented = errors.New("unimplemented")

var _ semalim.Job = (*imprJob)(nil)

func (ij *imprJob) Do() (interface{}, error) {
	// Match the impression with the ad inventory
	return nil, errUnimplemented
}

func matchImpressionToBid(impr *odex.Impression) (*odex.Bid, error) {
	// Fuzzy match the impression and return the matching bid
	bid := &odex.Bid{
		AdId:         "placeholder-ad-id",
		WinNoticeURL: "https://odexchange.com/placeholder-win-notice-url",
		ContentCategories: []odex.ContentCategory{
			odex.CArtsAndEntertainment,
			odex.CAutomative,
			odex.CShopping,
		},
	}

	return bid, nil
}

func (ij *imprJob) Id() interface{} {
	return ij.id
}

func (b *Bidder) checkBidMatch(bid *odex.BidRequest) (*odex.BidResponse, error) {
	// If no user matched
	if matched, err := b.accepted(bid); !matched || err != nil {
		unmatchedResp := &odex.BidResponse{
			Id:    b.newBidId(),
			BidId: bid.Id,

			NoBidReason: odex.NoBidUnmatchedUser,
		}

		return unmatchedResp, nil
	}

	// Otherwise the bid has been matched so now process the bid
	return b.processBid(bid)
}

func (b *Bidder) accepted(bid *odex.BidRequest) (bool, error) {
	if bid.User == nil {
		return false, nil
	}

	// The problem am trying to solve here is to match
	// App, Site, User to content in the inventory database
	// e.g by content topics, locations etc

	return true, nil
}

func (b *Bidder) processBid(bid *odex.BidRequest) (*odex.BidResponse, error) {
	// First step is to see if we can match the user up

	jobsChan := make(chan semalim.Job)

	go func() {
		defer close(jobsChan)

		for i, impr := range bid.Impressions {
			job := &imprJob{id: i, impr: impr, bid: bid}
			jobsChan <- job
		}
	}()

	resChan := semalim.Run(jobsChan, uint64(len(bid.Impressions)))

	var acceptedSeatBids []*odex.SeatBid

	for res := range resChan {
		value, err := res.Value(), res.Err()
		imprIndex := res.Id().(int)
		if err != nil {
			impr := bid.Impressions[imprIndex]
			fmt.Fprintf(b, "[processBid] impressionId: %v bidId: %v err: %v", impr.Id, bid.Id, err)
			continue
		}

		acceptedSeatBids = append(acceptedSeatBids, value.(*odex.SeatBid))
	}

	resp := &odex.BidResponse{
		Id:    b.newBidId(),
		BidId: bid.Id,

		SeatBids: acceptedSeatBids,
	}

	if len(acceptedSeatBids) < 1 {
		resp.NoBidReason = odex.NoBidUnmatchedUser
	}

	return resp, nil
}
