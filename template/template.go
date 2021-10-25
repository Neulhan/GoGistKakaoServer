package template


type KakaoCard interface {
	CardToJSON() map[string]map[string]string
}

// KakaoCard 카카오 스킬 서버 이미지 카드
type SimpleImageCard struct {
	Text string
}

func (s SimpleImageCard) CardToJSON() map[string]map[string]string {
	return map[string]map[string]string{
		"simpleImage": {
			"imageUrl": s.Text,
			"altText":  "image",
		},
	}
}

// KakaoCard 카카오 스킬 서버 이미지 카드
type SimpleTextCard struct {
	Text string
}

func (s SimpleTextCard) CardToJSON() map[string]map[string]string {
	return map[string]map[string]string {
		"simpleText": {
			"text": s.Text,
		},
	}
}


// KakaoTemplate 카카오 스킬 서버 템플릿
type KakaoTemplate struct {
	Outputs []map[string]map[string]string `json:"outputs"`
}

// KakaoRes 카카오 스킬 서버 리스폰스
type KakaoRes struct {
	Version  string        `json:"version"`
	Template KakaoTemplate `json:"template"`
}

// JSONResMaker json response maker
func JSONResMaker(kakaoCards []KakaoCard) (res *KakaoRes) {
	var jsonCardList []map[string]map[string]string

	for _, card := range kakaoCards {
		jsonCardList = append(jsonCardList, card.CardToJSON())
	}

	template := KakaoTemplate{Outputs: jsonCardList}
	res = &KakaoRes{Version: "2.0", Template: template}
	return
}
