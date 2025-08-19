package uploadsv1

type AttachmentType byte

const (
	PhotoAttachment AttachmentType = iota
	VideoAttachment
	AudioAttachment
	FileAttachment

	UnknownAttachment
)

var attachmentTypes = map[AttachmentType]string{
	PhotoAttachment: "photo",
	VideoAttachment: "video",
	AudioAttachment: "audio",
	FileAttachment:  "file",
}

func ParseAttachment(s string) AttachmentType {
	for k, v := range attachmentTypes {
		if v == s {
			return k
		}
	}

	return UnknownAttachment
}

func (t AttachmentType) String() string {
	if name, ok := attachmentTypes[t]; ok {
		return name
	}

	return ""
}

type Endpoint struct {
	Token string `json:"token,omitempty"`
	Url   string `json:"url"`
}
