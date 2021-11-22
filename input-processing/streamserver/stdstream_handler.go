// File manages address information for standard stream in
package streamserver

type StdInAddress struct {
	Address string
}

func NewStdInAddress(address string) *StdInAddress {
	return &StdInAddress{Address: address}
}

func (address *StdInAddress) String() string {
	return address.Address
}

func (address *StdInAddress) GetAddress() string {
	return address.Address
}
func (a *StdInAddress) Network() string {
	return "standardio"
}


