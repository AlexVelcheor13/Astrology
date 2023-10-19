package app

type ApodModel struct {
	Copyright      string `json:"copyright" db:"copyright"`
	Date           string `json:"date" db:"date"`
	Explanation    string `json:"explanation" db:"explanation"`
	HDURL          string `json:"hdurl" db:"hd_url"`
	MediaType      string `json:"media_type" db:"media_type"`
	ServiceVersion string `json:"service_version" db:"service_version"`
	Title          string `json:"title" db:"title"`
	URL            string `json:"url" db:"url"`
	FreshRaw       []byte `json:"-" db:"-"`
}
