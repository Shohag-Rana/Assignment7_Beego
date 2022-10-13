package structure

type Cats_Category []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Cats_Category_Details []struct {
	Breeds []struct {
		Weight struct {
			Imperial string `json:"imperial"`
			Metric   string `json:"metric"`
		} `json:"weight"`
		ID           string `json:"id"`
		Name         string `json:"name"`
		Temperament  string `json:"temperament"`
		Origin       string `json:"origin"`
		Description  string `json:"description"`
		LifeSpan     string `json:"life_span"`
		WikipediaURL string `json:"wikipedia_url"`
	} `json:"breeds"`
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
