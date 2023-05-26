package pkg

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(file *os.File) (*CanModel, error)
}

type Writer interface {
	Write(file *os.File, canModel *CanModel) error
}

// CanModel represents the CAN model.
type CanModel struct {
	Version           string                `json:"version"`
	BusSpeed          uint32                `json:"bus_speed"`
	Nodes             map[string]*Node      `json:"nodes"`
	Messages          map[string]*Message   `json:"messages"`
	NodeAttributes    map[string]*Attribute `json:"node_attributes"`
	MessageAttributes map[string]*Attribute `json:"message_attributes"`
	SignalAttributes  map[string]*Attribute `json:"signal_attributes"`
}

// Validate validates the CAN model.
func (c *CanModel) Validate() error {
	for attName, att := range c.NodeAttributes {
		if err := att.validate(attName, attributeKindNode); err != nil {
			return fmt.Errorf("node attribute %s: %w", attName, err)
		}
	}
	for attName, att := range c.MessageAttributes {
		if err := att.validate(attName, attributeKindMessage); err != nil {
			return fmt.Errorf("message attribute %s: %w", attName, err)
		}
	}
	for attName, att := range c.SignalAttributes {
		if err := att.validate(attName, attributeKindSignal); err != nil {
			return fmt.Errorf("signal attribute %s: %w", attName, err)
		}
	}

	for nodeName, node := range c.Nodes {
		if err := node.validate(nodeName, c.NodeAttributes); err != nil {
			return err
		}
	}

	for msgName, msg := range c.Messages {
		if err := msg.validate(msgName, c.MessageAttributes, c.SignalAttributes); err != nil {
			return err
		}
	}

	return nil
}

func (c *CanModel) getAttributes() []*Attribute {
	attributes := []*Attribute{}

	for _, att := range c.NodeAttributes {
		attributes = append(attributes, att)
	}
	for _, att := range c.MessageAttributes {
		attributes = append(attributes, att)
	}
	for _, att := range c.SignalAttributes {
		attributes = append(attributes, att)
	}

	return attributes
}

type nodeAttAssignment struct {
	attAssignmentVal
	nodeName string
}

func (c *CanModel) getNodeAttAssignments() []*nodeAttAssignment {
	assignments := []*nodeAttAssignment{}

	for _, node := range c.Nodes {
		for _, ass := range node.getAttAssignmentValues() {
			assignments = append(assignments, &nodeAttAssignment{
				attAssignmentVal: *ass,
				nodeName:         node.name,
			})
		}
	}

	return assignments
}

type messageAttAssignment struct {
	attAssignmentVal
	messageID uint32
}

func (c *CanModel) getMessageAttAssignments() []*messageAttAssignment {
	assignments := []*messageAttAssignment{}

	for _, msg := range c.Messages {
		for _, ass := range msg.getAttAssignmentValues() {
			assignments = append(assignments, &messageAttAssignment{
				attAssignmentVal: *ass,
				messageID:        msg.ID,
			})
		}
	}

	return assignments
}

type signalAttAssignment struct {
	attAssignmentVal
	messageID  uint32
	signalName string
}

func (c *CanModel) getSignalAttAssignments() []*signalAttAssignment {
	assignments := []*signalAttAssignment{}

	for _, msg := range c.Messages {
		for sigName, sig := range msg.Signals {
			for _, ass := range sig.getAttAssignmentValues() {
				assignments = append(assignments, &signalAttAssignment{
					attAssignmentVal: *ass,
					messageID:        msg.ID,
					signalName:       sigName,
				})
			}
		}
	}

	return assignments
}