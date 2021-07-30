package charts

import (
	"fmt"
	"testing"

	"github.com/xxjwxc/gocharts/opts"
)

func TestLine(t *testing.T) {
	line := NewLine(WithTitleOpts(opts.Title{Title: "折线图堆叠"}),
		WithTooltipOpts(opts.Tooltip{Show: &True, Trigger: "axis"}),
		WithLegendOpts(opts.Legend{Show: &True, Data: []string{"邮件营销", "联盟广告", "视频广告", "直接访问", "搜索引擎"}}),
		WithGridOpts(opts.Grid{Left: "3%", Right: "4%", Bottom: "3%", ContainLabel: true}),
		WithToolboxOpts(opts.Toolbox{Feature: &opts.ToolBoxFeature{SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{}}}),
		WithXAxisOpts(opts.XAxis{Type: "category", BoundaryGap: &False, Data: []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}}),
		WithYAxisOpts(opts.YAxis{Type: "value"}),
	)

	mp := make(map[string]interface{})
	mp["周一"] = 120
	mp["周二"] = 132
	mp["周三"] = 101
	mp["周四"] = 134
	mp["周五"] = 90
	mp["周六"] = 230
	mp["周日"] = 210
	line.AddSeries("邮件营销", mp, "", WithStackOpts("总量"))
	mp = make(map[string]interface{})
	mp["周一"] = 220
	mp["周二"] = 182
	mp["周三"] = 191
	mp["周四"] = 234
	mp["周五"] = 290
	mp["周六"] = 330
	mp["周日"] = 310
	line.AddSeries("联盟广告", mp, "", WithStackOpts("总量"))

	mp = make(map[string]interface{})
	mp["周一"] = 150
	mp["周二"] = 232
	mp["周三"] = 210
	mp["周四"] = 154
	mp["周五"] = 190
	mp["周六"] = 330
	mp["周日"] = 410
	line.AddSeries("视频广告", mp, "", WithStackOpts("总量"))

	mp = make(map[string]interface{})
	mp["周一"] = 320
	mp["周二"] = 332
	mp["周三"] = 301
	mp["周四"] = 334
	mp["周五"] = 390
	mp["周六"] = 330
	mp["周日"] = 320
	line.AddSeries("直接访问", mp, "", WithStackOpts("总量"))

	mp = make(map[string]interface{})
	mp["周一"] = 820
	mp["周二"] = 932
	mp["周三"] = 901
	mp["周四"] = 934
	mp["周五"] = 1290
	mp["周六"] = 1330
	mp["周日"] = 1320
	line.AddSeries("搜索引擎", mp, "", WithStackOpts("总量"))

	fmt.Println(line.Unmarshal())

}
