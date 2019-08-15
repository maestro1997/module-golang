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
///// Vigenere Encode and Decode functions
func (vig Vigenere) Encode (s string) string {
    s = prepare_str(s)
    return vig.Decode_t(s,-1)
}

func (vig Vigenere) Decode (s string) string {
    return vig.Decode_t(s,1)
}

func (vig Vigenere) Decode_t (s string, sign int32) string {
    dec := []rune(s)
    i := 0
    key := vig.key
    lenn := len(key)
    for ; i < len(s); i++ {
        dec[i] = dec[i] - sign*(int32(key[i % lenn] - 97))
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
////////  Shift Encode and Decode functions
func (sh Shift) Encode(s string) string {
    return Decode_t(prepare_str(s), -sh.shift)
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



