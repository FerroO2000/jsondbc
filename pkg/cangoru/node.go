package cangoru

type Node struct {
	Description
	AttributeMap

	Name string

	TxMessages []*Message
	RxSignals  []*Signal
}

func NewNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

func (n *Node) AddTxMessage(txMsg *Message) {
	n.TxMessages = append(n.TxMessages, txMsg)
}

func (n *Node) AddRxSignal(rxSig *Signal) {
	n.RxSignals = append(n.RxSignals, rxSig)
}
