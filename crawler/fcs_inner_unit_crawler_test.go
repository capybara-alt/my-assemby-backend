package crawler_test

import (
	"testing"

	"github.com/capybara-alt/my-assemble/crawler"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/stretchr/testify/assert"
)

func TestFCSInnerUnitCrawler(t *testing.T) {
	test := struct {
		name string
		want model.Want[model.CrawlResultJSON]
	}{
		name: "[正常系]FCSページのクロールテスト",
		want: model.Want[model.CrawlResultJSON]{
			Value: model.CrawlResultJSON{
				string(model.FCS_INNER_UNIT_TYPE): {
					"FCS": {
						"FCS-G1/P01": {
							"メーカー":      "ファーロン・ダイナミクス",
							"価格":        "-",
							"近距離アシスト適性": "38",
							"中距離アシスト適性": "27",
							"遠距離アシスト適性": "20",
							"ミサイルロック補正": "79",
							"マルチロック補正":  "40",
							"重量":        "80",
							"EN負荷":      "198",
						},
						"FCS-G2/P05": {
							"メーカー":      "ファーロン・ダイナミクス",
							"価格":        "67,000",
							"レギュレーション":  "1.03.1",
							"近距離アシスト適性": "42 (-3)",
							"中距離アシスト適性": "80",
							"遠距離アシスト適性": "26",
							"ミサイルロック補正": "105",
							"マルチロック補正":  "60",
							"重量":        "100",
							"EN負荷":      "232",
						},
						"FCS-G2/P10SLT": {
							"メーカー":      "ファーロン・ダイナミクス",
							"価格":        "96,000",
							"近距離アシスト適性": "40",
							"中距離アシスト適性": "41",
							"遠距離アシスト適性": "29",
							"ミサイルロック補正": "150",
							"マルチロック補正":  "90",
							"重量":        "120",
							"EN負荷":      "209",
						},
						"FCS-G2/P12SML": {
							"メーカー":      "ファーロン・ダイナミクス",
							"価格":        "141,000",
							"近距離アシスト適性": "28",
							"中距離アシスト適性": "52",
							"遠距離アシスト適性": "30",
							"ミサイルロック補正": "132",
							"マルチロック補正":  "120",
							"重量":        "130",
							"EN負荷":      "278",
						},
						"FC-006 ABBOT": {
							"メーカー":      "ベイラム",
							"価格":        "135,000",
							"レギュレーション":  "1.03.1",
							"近距離アシスト適性": "70 (-13)",
							"中距離アシスト適性": "32",
							"遠距離アシスト適性": "5",
							"ミサイルロック補正": "74",
							"マルチロック補正":  "46",
							"重量":        "90",
							"EN負荷":      "266",
						},
						"FC-008 TALBOT": {
							"メーカー":      "ベイラム",
							"価格":        "155,000",
							"レギュレーション":  "1.03.1",
							"近距離アシスト適性": "63 (-4)",
							"中距離アシスト適性": "54",
							"遠距離アシスト適性": "11",
							"ミサイルロック補正": "103",
							"マルチロック補正":  "62",
							"重量":        "140",
							"EN負荷":      "312",
						},
						"VE-21A": {
							"メーカー":      "アーキバス先進開発局",
							"価格":        "228,000",
							"近距離アシスト適性": "10",
							"中距離アシスト適性": "36",
							"遠距離アシスト適性": "92",
							"ミサイルロック補正": "65",
							"マルチロック補正":  "79",
							"重量":        "85",
							"EN負荷":      "364",
						},
						"VE-21B": {
							"メーカー":      "アーキバス先進開発局",
							"価格":        "315,000",
							"近距離アシスト適性": "15",
							"中距離アシスト適性": "50",
							"遠距離アシスト適性": "80",
							"ミサイルロック補正": "97",
							"マルチロック補正":  "70",
							"重量":        "160",
							"EN負荷":      "388",
						},
						"IA-C01F: OCELLUS": {
							"メーカー":      "ルビコン調査技研",
							"価格":        "367,000",
							"近距離アシスト適性": "90",
							"中距離アシスト適性": "12",
							"遠距離アシスト適性": "3",
							"ミサイルロック補正": "85",
							"マルチロック補正":  "50",
							"重量":        "130",
							"EN負荷":      "292",
						},
						"IB-C03F: WLT 001": {
							"メーカー":      "ルビコン調査技研",
							"価格":        "400,000",
							"近距離アシスト適性": "50",
							"中距離アシスト適性": "72",
							"遠距離アシスト適性": "48",
							"ミサイルロック補正": "102",
							"マルチロック補正":  "66",
							"重量":        "150",
							"EN負荷":      "486",
						},
					},
				},
			},
			ErrMsg: "",
		},
	}
	c := crawler.NewFcsInnerUnitCrawler()
	t.Run(test.name, func(t *testing.T) {
		results, err := c.Crawl()
		assert.Equal(t, err, nil)
		assert.Equal(t, test.want.Value, results)
	})
}
