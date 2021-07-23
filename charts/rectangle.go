package charts

import "github.com/xxjwxc/gochart/opts"

type Overlaper interface {
	overlap() MultiSeries
}

// XYAxis represent the X and Y axis in the rectangular coordinates.
type XYAxis struct {
	XAxisList []opts.XAxis `json:"xaxis"`
	YAxisList []opts.YAxis `json:"yaxis"`
}

func (xy *XYAxis) initXYAxis() {
	xy.XAxisList = append(xy.XAxisList, opts.XAxis{})
	xy.YAxisList = append(xy.YAxisList, opts.YAxis{})
}

// ExtendXAxis adds new X axes.
func (xy *XYAxis) ExtendXAxis(xAxis ...opts.XAxis) {
	xy.XAxisList = append(xy.XAxisList, xAxis...)
}

// ExtendYAxis adds new Y axes.
func (xy *XYAxis) ExtendYAxis(yAxis ...opts.YAxis) {
	xy.YAxisList = append(xy.YAxisList, yAxis...)
}

// WithXAxisOpts sets the X axis.
func WithXAxisOpts(opt opts.XAxis) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.hasXYAxis = true
		bc.XAxis = opt
	}
}

// WithYAxisOpts sets the Y axis.
func WithYAxisOpts(opt opts.YAxis) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.hasXYAxis = true
		bc.YAxis = opt
	}
}

// RectConfiguration contains options for the rectangular coordinates.
type RectConfiguration struct {
	BaseConfiguration
}

func (rect *RectConfiguration) setRectGlobalOptions(options ...GlobalOpts) {
	rect.BaseConfiguration.setBaseGlobalOptions(options...)
}

// RectChart is a chart in RectChart coordinate.
type RectChart struct {
	RectConfiguration

	xAxisData interface{}
}

func (rc *RectChart) overlap() MultiSeries {
	return rc.MultiSeries
}

// SetGlobalOptions sets options for the RectChart instance.
func (rc *RectChart) SetGlobalOptions(options ...GlobalOpts) *RectChart {
	rc.RectConfiguration.setRectGlobalOptions(options...)
	return rc
}

// Overlap composes multiple charts into one single canvas.
// It is only suited for some of the charts which are in rectangular coordinate.
// Supported charts: Bar/BoxPlot/Line/Scatter/EffectScatter/Kline/HeatMap
func (rc *RectChart) Overlap(a ...Overlaper) {
	for i := 0; i < len(a); i++ {
		rc.MultiSeries = append(rc.MultiSeries, a[i].overlap()...)
	}
}

// Validate validates the given configuration.
func (rc *RectChart) Validate() {
	// Make sure that the data of X axis won't be cleaned for XAxisOpts
	rc.XAxis.Data = rc.xAxisData
	// Make sure that the labels of Y axis show correctly
	rc.YAxis.AxisLabel.Show = true

	// rc.Assets.Validate(rc.AssetsHost)
}
