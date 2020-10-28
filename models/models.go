package models

type Symbol struct {
	OID string `yaml:"OID"`
	// Name yaml.Node `yaml:"name"`
	Name string `yaml:"name"`
}

type MetricTag struct {
	Column Symbol `yaml:"column"`
	Tag    string
}

type BaseMetric struct {
	MIB    string `yaml:"MIB"`
	Symbol `yaml:"symbol"`
}

type TableMetric struct {
	MIB        string      `yaml:"MIB"`
	Table      Symbol      `yaml:"table"`
	Symbols    []Symbol    `yaml:"symbols"`
	MetricTags []MetricTag `yaml:"metric_tags"`
}

type Profile struct {
	Extends     []string      `yaml:"extends"`
	Sysobjectid string        `yaml:"sysobjectid"`
	Metrics     []interface{} `yaml:"metrics"`
}

// NewSymbol creates a new Symbol with a description attached to the name
func NewSymbol(oid, name, desc string) Symbol {
	// var fullName yaml.Node
	// fullName.SetString(name)
	// fullName.LineComment = desc
	// return Symbol{
	// 	OID:  oid,
	// 	Name: fullName,
	// }
	return Symbol{
		OID:  oid,
		Name: name,
	}
}
