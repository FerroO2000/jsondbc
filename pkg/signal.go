package pkg

import "fmt"

// Signal represents a CAN signal in a message.
type Signal struct {
	attributeAssignment
	Description string             `json:"description,omitempty"`
	MuxSwitch   uint32             `json:"mux_switch,omitempty"`
	StartBit    uint32             `json:"start_bit"`
	Size        uint32             `json:"size"`
	BigEndian   bool               `json:"big_endian,omitempty"`
	Signed      bool               `json:"signed,omitempty"`
	Unit        string             `json:"unit,omitempty"`
	Receivers   []string           `json:"receivers,omitempty"`
	Scale       float64            `json:"scale"`
	Offset      float64            `json:"offset"`
	Min         float64            `json:"min"`
	Max         float64            `json:"max"`
	Bitmap      map[string]uint32  `json:"bitmap,omitempty"`
	MuxGroup    map[string]*Signal `json:"mux_group,omitempty"`

	name          string
	isMultiplexor bool
	isMultiplexed bool
}

func (s *Signal) validate(sigAtt map[string]*Attribute) error {
	if s.Scale == 0 {
		s.Scale = 1
	}

	if err := s.attributeAssignment.validate(sigAtt); err != nil {
		return fmt.Errorf("signal %s: %w", s.name, err)
	}

	return nil
}

// IsBitmap returns true if the signal is a bitmap.
func (s *Signal) IsBitmap() bool {
	return len(s.Bitmap) > 0
}

// IsMultiplexor returns true if the signal is a multiplexor.
func (s *Signal) IsMultiplexor() bool {
	return len(s.MuxGroup) > 0
}

// HasDescription returns true if the signal has a description.
func (s *Signal) HasDescription() bool {
	return len(s.Description) > 0
}
