package template

// KakaoCard 카카오 스킬 서버 이미지 카드
type KakaoCard struct {
	SimpleImage map[string]string `json:"simpleImage"`
}

// KakaoTemplate 카카오 스킬 서버 템플릿
type KakaoTemplate struct {
	Outputs []KakaoCard `json:"outputs"`
}

// KakaoRes 카카오 스킬 서버 리스폰스
type KakaoRes struct {
	Version  string        `json:"version"`
	Template KakaoTemplate `json:"template"`
}

// JSONResMaker json response maker
func JSONResMaker(cards []map[string]string) (res *KakaoRes) {
	var kakaocards []KakaoCard

	for _, card := range cards {
		kakaocard := KakaoCard{SimpleImage: card}
		kakaocards = append(kakaocards, kakaocard)
	}

	template := KakaoTemplate{Outputs: kakaocards}

	res = &KakaoRes{Version: "2.0", Template: template}
	return
}
