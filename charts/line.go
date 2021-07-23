package charts

import (
	"github.com/xxjwxc/gochart/types"
)

// Line represents a line chart.
type Line struct {
	BaseConfiguration
}

// Type returns the chart type.
func (Line) Type() string { return types.ChartLine }

// NewLine creates a new line chart.
func NewLine(options ...GlobalOpts) *Line {
	c := &Line{}

	for _, opt := range options {
		opt(&c.BaseConfiguration)
	}
	return c
}

// AddSeries adds the new series.
func (c *Line) AddSeries(name string, data map[string]interface{}, def interface{}, options ...SeriesOpts) *Line {
	var list []string
	if c.XAxis.Type == "category" {
		list = c.XAxis.Data.([]string)
	} else {
		list = c.YAxis.Data.([]string)
	}
	series := &SingleSeries{Name: name, Type: types.ChartLine}
	series.configureSeriesOpts(options...)

	var tmp []interface{}
	for _, v := range list {
		if _, ok := data[v]; ok {
			tmp = append(tmp, data[v])
		} else { // 设置默认值
			tmp = append(tmp, def)
		}
	}
	series.Data = tmp
	c.MultiSeries = append(c.MultiSeries, *series)
	return c
}
