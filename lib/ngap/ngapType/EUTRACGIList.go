//go:binary-only-package

package ngapType

// Need to import "free5gc/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct EUTRA_CGIList */
/* EUTRACGI */
type EUTRACGIList struct {
	List []EUTRACGI `aper:"valueExt,sizeLB:1,sizeUB:256"`
}
