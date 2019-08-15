package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Caesar struct {
    s string
}

func NewCaesar() Cipher {
    var c Cipher
    c = Caesar{""}
    return c
}

func NewShift(int n) {


}

func (c Caesar) Encode_s(s string) string {
    enc := []rune(s)
    j := 0
    i := 0
    flag:=true
    for ; i < len(s); i++ {
	flag = false
        if enc[i] <= 90 && enc[i] >= 65 {
            enc[i] += 32
	    flag = true
	} else {
            if enc[i] >= 97 && enc[i] <= 122 {
                flag = true
	    }
	}
	if flag {
	    enc[i] +=3
	    if enc[i] >122 {
	        enc[i] -= 26
	    }
            enc[j] = enc[i]
	    j++
	}
    }
    return string(enc[0:j])
}

func (c Caesar) Decode(s string) string {
    dec := []rune(s)
    i := 0
    for ; i < len(dec); i++ {
        dec[i] = dec[i] - 3
	if dec[i] < 97 {
            dec[i] += 26
	}
    }
    return string(dec)
}
