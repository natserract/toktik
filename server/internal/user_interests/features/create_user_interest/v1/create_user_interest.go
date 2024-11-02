package v1

type CreateUserMetadata struct {
	VideoID    string
	Title      string
	AuthorName string
	CreateTime int64
}

type CreateUserInterest struct {
	PageContent string
	Metadata    CreateUserMetadata
}
