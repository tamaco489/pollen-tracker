package domain

import "time"

// PollenType は花粉種別
type PollenType string

const (
	PollenTypeCedar   PollenType = "CEDAR"
	PollenTypeCypress PollenType = "CYPRESS"
	PollenTypeGrass   PollenType = "GRASS"
	PollenTypeRagweed PollenType = "RAGWEED"
	PollenTypeMugwort PollenType = "MUGWORT"
)

// PollenForecast は花粉予報のエンティティ
type PollenForecast struct {
	Date       time.Time
	Level      int
	PollenType PollenType
	SeasonInfo SeasonInfo
}

// SeasonInfo は花粉種別のシーズン情報
type SeasonInfo struct {
	Peak            string
	Characteristics string
	Region          string
}

// SeasonCalendar は花粉種別の静的シーズンカレンダー
var SeasonCalendar = map[PollenType]SeasonInfo{
	PollenTypeCedar:   {Peak: "2月〜4月", Characteristics: "目のかゆみ・鼻水・くしゃみ", Region: "関東・近畿・東北"},
	PollenTypeCypress: {Peak: "3月〜5月", Characteristics: "鼻水・鼻づまり・目のかゆみ", Region: "関東・近畿"},
	PollenTypeGrass:   {Peak: "5月〜10月", Characteristics: "くしゃみ・鼻水・目のかゆみ", Region: "全国"},
	PollenTypeRagweed: {Peak: "8月〜10月", Characteristics: "くしゃみ・目のかゆみ", Region: "関東・中部"},
	PollenTypeMugwort: {Peak: "7月〜10月", Characteristics: "くしゃみ・鼻水", Region: "全国"},
}

// PlantCodeToType は Google Pollen API のプラントコードを PollenType にマッピングする
var PlantCodeToType = map[string]PollenType{
	"CEDAR":          PollenTypeCedar,
	"JAPANESE_CEDAR": PollenTypeCedar,
	"CYPRESS":        PollenTypeCypress,
	"CYPRESS_PINE":   PollenTypeCypress,
	"GRASSES":        PollenTypeGrass,
	"GRASS":          PollenTypeGrass,
	"RAGWEED":        PollenTypeRagweed,
	"MUGWORT":        PollenTypeMugwort,
}
