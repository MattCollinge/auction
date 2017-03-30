package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AuctionByAuctionIDHandlerFunc turns a function with the right signature into a auction by auction Id handler
type AuctionByAuctionIDHandlerFunc func(AuctionByAuctionIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AuctionByAuctionIDHandlerFunc) Handle(params AuctionByAuctionIDParams) middleware.Responder {
	return fn(params)
}

// AuctionByAuctionIDHandler interface for that can handle valid auction by auction Id params
type AuctionByAuctionIDHandler interface {
	Handle(AuctionByAuctionIDParams) middleware.Responder
}

// NewAuctionByAuctionID creates a new http.Handler for the auction by auction Id operation
func NewAuctionByAuctionID(ctx *middleware.Context, handler AuctionByAuctionIDHandler) *AuctionByAuctionID {
	return &AuctionByAuctionID{Context: ctx, Handler: handler}
}

/*AuctionByAuctionID swagger:route GET /auctions/auctionId/{auctionId} auctionByAuctionId

Returns the auction given by the auctionId

*/
type AuctionByAuctionID struct {
	Context *middleware.Context
	Handler AuctionByAuctionIDHandler
}

func (o *AuctionByAuctionID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewAuctionByAuctionIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}