package shortener

type Redirect struct {
	Code      string `json:"code" bson:"code" msgpack:"code"`
	URL       string `json:"url" bson:"url" msgpack:"url" validator:"required=true"`
	CreatedAT string `json:"created_at" bson:"created_at" msgpack:"created_at"`
}
