package cipher

type Cipher interface {
    Encode(string) string
    Decode(string) string
}

type Shift struct {
    shift int32
} 

type Vigenere struct {
    key string
}

func NewVigenere (key string) Cipher {
    flag  := true
    for i:=0; i < len(key); i++ {
       if key[i] < 97 || key[i] > 122 {  // if exist non small-letter character
           return nil
       }	    
       if key[i] != 97 {    // it means that key has at leat two different letters
           flag = false
       }
    }
    if flag {      // if no even two letters in key - its not key for cipher
        return nil
    }
    var c Cipher
    c = Vigenere{key}
    return c
}

func NewCaesar() Cipher {
    var c Cipher
    c = Shift{3}
    return c
}

func NewShift(shift int) Cipher {
    var c Cipher
    if (shift > 25) || (shift == 0) || (shift < -25) {
        return nil
    }
    c = Shift{int32(shift)}
    return c
}

func (vig Vigenere) Encode (s string) string {
    enc := []rune(s)
    j := 0
    i := 0
    flag := true
    key := vig.key
    lenn := len(key)
    for ; i < len(s); i++ {
        flag = false
        if (enc[i] <= 90 && enc[i] >= 65) {
            enc[i] += 32
	    flag = true
	} else {
            if (enc[i] >= 97 && enc[i] <= 122) {
                flag = true
	    }  
	}
	if flag {
	    enc[j] = enc[i] + int32(key[j % lenn] - 97)
            if enc[j] > 122 {
                enc[j] -= 26
	    }
	    j++
    }
    }
    return string(enc[0:j]) 
}

func (vig Vigenere) Decode (s string) string {
    dec := []rune(s)
    j := 0
    i := 0
    flag := true
    key := vig.key
    lenn := len(key)
    for ; i < len(s); i++ {
        flag = false
        if (dec[i] <= 90 && dec[i] >= 65) {
            dec[i] += 32
	    flag = true
	} else {
            if (dec[i] >= 97 && dec[i] <= 122) {
                flag = true
	    }  
	}
	if flag {
            dec[j] = dec[i] - int32(key[j % lenn] - 97)
	    if dec[j] < 97 {
                dec[j] += 26
	    }
	    j++
        }
    }
    return string(dec[0:j])      
}

func (sh Shift) Encode(s string) string {
    return Encode_t(s, sh.shift)
}

func (sh Shift) Decode(s string) string {
    return Decode_t(s,sh.shift)
}

func prepare_str(s string) string{  // "downcasing" and removing non-letter characters
    out := []rune(s)
    i := 0
    j := 0
    for ; i < len(out); i++ {
        if s[i] <= 90 && s[i] >= 65 {
            out[j] = out[i] + 32
	    j++
	} else {
            if s[i] >= 97 && s[i] <= 122 {
                out[j] = out[i]
		j++
	    }
	}
    }
    return string(out[0:j])
}

func Encode_t(s string, shift int32) string {
    s = prepare_str(s)
    return Decode_t(s, -shift)
}

func Decode_t (s string, shift int32) string {
    dec := []rune(s)
    i := 0
    for ; i < len(dec); i++ {
        dec[i] -= shift
	if dec[i] > 122 {
            dec[i] -= 26
	} else {
            if dec[i] < 97 {
                dec[i] += 26
	    }
	}
    }
    return string(dec)
}



