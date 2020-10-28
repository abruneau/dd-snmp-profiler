package parsers

import (
	"fmt"

	"github.com/abruneau/dd-snmp-profiler/models"
	"github.com/sleepinggenius2/gosmi"
	"gopkg.in/yaml.v3"
)

func parseScalar(moduleName string, node gosmi.SmiNode) models.BaseMetric {
	var metric models.BaseMetric
	metric.MIB = moduleName
	metric.Symbol = models.NewSymbol(node.Oid.String(), node.Name, node.Description)
	return metric
}

func parseTable(moduleName string, table gosmi.Table) models.TableMetric {
	var metric = models.TableMetric{
		MIB:   moduleName,
		Table: models.NewSymbol(table.Oid.String(), table.Name, table.Description),
	}

	for _, column := range table.Columns {
		if column.Oid[len(column.Oid)-1] == 1 {
			metric.MetricTags = append(metric.MetricTags, models.MetricTag{
				Column: models.NewSymbol(column.Oid.String(), column.Name, column.Description),
			})
		} else {
			metric.Symbols = append(metric.Symbols, models.NewSymbol(column.Oid.String(), column.Name, column.Description))
		}
	}

	return metric
}

// ParseMIBModule parses a MIB module to generate a profile
func ParseMIBModule(path, module, output string, debug bool) error {
	gosmi.Init()
	gosmi.AppendPath(path)
	moduleName, err := gosmi.LoadModule(module)
	if err != nil {
		return err
	}
	if debug {
		fmt.Printf("Loaded module %s\n", moduleName)
	}

	m, err := gosmi.GetModule(module)
	if err != nil {
		return err
	}

	var profile models.Profile

	for _, node := range m.GetNodes() {

		if node.Kind.String() == "Scalar" {
			metric := parseScalar(m.Name, node)
			profile.Metrics = append(profile.Metrics, metric)
		}

		if node.Kind.String() == "Table" {
			metric := parseTable(m.Name, node.AsTable())
			profile.Metrics = append(profile.Metrics, metric)
		}
	}

	yProfile, err := yaml.Marshal(&profile)
	if err != nil {
		return err
	}
	err = WriteToFile(output, &yProfile)
	if err != nil {
		return err
	}
	if debug {
		fmt.Printf("Wrote %s profile to %s\n", module, output)
	}

	gosmi.Exit()

	return nil
}
