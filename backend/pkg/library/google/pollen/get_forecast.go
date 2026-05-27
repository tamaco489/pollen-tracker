package pollen

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const endpoint = "https://pollen.googleapis.com/v1/forecast:lookup"

// DOC: https://developers.google.com/maps/documentation/pollen/reference/rest/v1/forecast/lookup
// Google Pollen API v1 のレスポンスを Go 構造体にマッピングした非公開型
// 公式 Go SDK は存在しないため、REST レスポンスを手動で定義している

type forecastResponse struct {
	RegionCode string      `json:"regionCode"` // ISO 3166-1 alpha-2 の国・地域コード紛争地域では省略される場合がある
	DailyInfo  []dailyInfo `json:"dailyInfo"`  // 要求日数分の日次予報データ
}

type dailyInfo struct {
	Date      date        `json:"date"`      // 予報日付 (UTC)
	PlantInfo []plantInfo `json:"plantInfo"` // 個別花粉種ごとのデータ最大 15 種
}

type date struct {
	Year  int `json:"year"`  // 年 (1-9999)0 は未指定
	Month int `json:"month"` // 月 (1-12)0 は年のみ表記
	Day   int `json:"day"`   // 日 (1-31)0 は部分日付表記
}

type plantInfo struct {
	Code        string     `json:"code"`        // 植物識別子 (例: JAPANESE_CEDAR, CYPRESS)
	DisplayName string     `json:"displayName"` // 人間が読める植物名 (例: Japanese Cedar)
	InSeason    bool       `json:"inSeason"`    // 現在シーズン中かどうか
	IndexInfo   *indexInfo `json:"indexInfo"`   // UPI の詳細inSeason=false の場合 null になることがある
}

type indexInfo struct {
	Code        string `json:"code"`        // インデックスコード現状 "UPI" (Universal Pollen Index) 固定
	DisplayName string `json:"displayName"` // 人間が読める指数名 (例: Universal Pollen Index)
	Value       int    `json:"value"`       // 花粉飛散レベル (0=None / 1=Very low / 2=Low / 3=Moderate / 4=High / 5=Very high)
	Category    string `json:"category"`    // Value に対応する分類ラベル (None / Low / Moderate / High / Very high)
}

type errorResponse struct {
	Error struct {
		Code    int    `json:"code"`    // HTTP ステータスコードに対応する数値コード
		Message string `json:"message"` // エラーの詳細メッセージ
		Status  string `json:"status"`  // gRPC ステータス文字列 (例: INVALID_ARGUMENT)
	} `json:"error"`
}

// GetForecast は指定座標・日数の花粉予報を取得して ForecastResponse に変換して返す
func (c *pollenClient) GetForecast(ctx context.Context, req *ForecastRequest) (*ForecastResponse, error) {
	params := url.Values{}
	params.Set("key", c.apiKey)
	params.Set("location.latitude", strconv.FormatFloat(req.Lat, 'f', -1, 64))
	params.Set("location.longitude", strconv.FormatFloat(req.Lng, 'f', -1, 64))
	params.Set("days", strconv.Itoa(req.Days))

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint+"?"+params.Encode(),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	res, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		var apiErr errorResponse
		if jsonErr := json.NewDecoder(res.Body).Decode(&apiErr); jsonErr == nil && apiErr.Error.Message != "" {
			return nil, fmt.Errorf("google pollen api error %d: %s", apiErr.Error.Code, apiErr.Error.Message)
		}
		return nil, fmt.Errorf("unexpected status: %d", res.StatusCode)
	}

	var apiResp forecastResponse
	if err := json.NewDecoder(res.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return toForecastResponse(&apiResp), nil
}

// toForecastResponse は Google API の内部レスポンスを ForecastResponse に変換する
func toForecastResponse(apiResp *forecastResponse) *ForecastResponse {
	dailyForecasts := make([]DailyForecast, 0, len(apiResp.DailyInfo))
	for _, day := range apiResp.DailyInfo {
		plants := make([]Plant, 0, len(day.PlantInfo))
		for _, p := range day.PlantInfo {
			level := MinPollenLevel.ToInt()
			if p.IndexInfo != nil {
				if v := PollenLevel(p.IndexInfo.Value); v.IsValid() {
					level = v.ToInt()
				}
			}
			plants = append(plants, Plant{
				Code:     p.Code,
				InSeason: p.InSeason,
				Level:    level,
			})
		}
		dailyForecasts = append(dailyForecasts, DailyForecast{Plants: plants})
	}
	return &ForecastResponse{DailyForecasts: dailyForecasts}
}
