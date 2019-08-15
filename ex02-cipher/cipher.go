package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

func (c Cipher) Encode(s string) string {
    
}
