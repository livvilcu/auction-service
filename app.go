package auctionservice

import (
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/libs/log"
	"os"
	sdk "github.com/cosmos/cosmos-sdk/types"
)
// app.go -> se vor defini modulele
//        -> va decide ce va face aplicatia cand va primi o tranzactie
const appName = "auctionService"

var (
	DefaultCLIHome = os.ExpandEnv("$HOME/.nscli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.nsd")
	ModuleBasics = module.NewBasicManager()
)


type auctionServiceApp struct {
	*bam.BaseApp
	cdc *codec.Codec
}


// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}

// constructor 
func NewAuctionServiceApp(logger log.Logger, db dbm.DB) *auctionServiceApp {

	// va fi distribuit in mai multe module
	cdc := MakeCodec()

	// baseApp  interactioneaza cu Tenderming prin protocolul ABCI
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &auctionServiceApp{
		BaseApp: bApp,
		cdc:     cdc,
	}

	return app
}
