package domain

type Transport interface {
	Inventoriable
	BeginnerSimplicity() int
}

type ElectricTransport interface {
	Transport
	EnergyConsumer
}
