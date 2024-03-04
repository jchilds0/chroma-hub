package graphics

import "fmt"

type Attr struct {
    Name    string 
    Value   string
    Visible int
}

func NewAttr(name, value string) *Attr {
    return &Attr{Name: name, Value: value}
}

func (attr *Attr) String() string {
    return fmt.Sprintf("{'name': '%s', 'value': '%s', 'visible': %d}", attr.Name, attr.Value, attr.Visible)
}
