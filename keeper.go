package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// keeper.go -> interactioneaza cu store
//          -> are referinta la alte fisiere Keeper din alte module


// modulul types din sdk pune la dispozitii tipurile care trebuiesc folosite pentru interactiunea cu SDK
type Keeper struct {
	// referinta la Keeper ul din modulul bank
	coinKeeper bank.Keeper // modulul bank se foloseste pentru tranzactii si pentru datele clientiilor
	// acces la sdk.KVStore care tine starea aplicatiei
	storeKey  sdk.StoreKey
	 // pentru Amino = encodeaza si decodeaza informatia pentru a o trimite mai departe la SDK
	cdc *codec.Codec
}

// Keeper constructor
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}


// Functii pentru a interactiona cu Keeper ul

// Context: contine informatia blockHeight or  chainId
func (k Keeper) SetWhois(ctx sdk.Context, object string, whois Whois) {
	// trebuie sa tinem in store doar daca obiecul are detinator :)
	if whois.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	// daca are o punem in store, dar inainte informatia trebuie encodata cu Amino
	store.Set([]byte(object), k.cdc.MustMarshalBinaryBare(whois))
}


// Gets the entire Whois metadata struct for an object
func (k Keeper) GetWhois(ctx sdk.Context, object string) Whois {
	// il ia din store
	store := ctx.KVStore(k.storeKey)

	// verifica daca  obiectul nu este prezent in store
	if !k.IsNamePresent(ctx, object) {

		return NewWhois() // daca nu se afla in store, atunci ar trebui sa fie donatie
	}
	// daca e obiectul are nume, atunci decodam si returnam
	bz := store.Get([]byte(object))
	var whois Whois
	k.cdc.MustUnmarshalBinaryBare(bz, &whois)
	return whois
}

// Check if the object is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// SetObject - sets the value string that a name resolves to
func (k Keeper) SetObject(ctx sdk.Context, owner sdk.AccAddress, object string) {
	whois := k.GetWhois(ctx, object)
	whois.Owner = owner
	k.SetWhois(ctx, object, whois)
}

func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetWhois(ctx, name).Owner
}

func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	return k.GetWhois(ctx, name).Price
}

func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whois := k.GetWhois(ctx, name)
	whois.Price = price
	k.SetWhois(ctx, name, whois)
}
