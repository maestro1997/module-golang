package downcase

func Downcase(s string) (string,error){
    out := []rune(s)
    i:=0
    for ; i < len(out); i++ {
        if out[i] <= 90 && out[i] >= 65 {
  	    out[i] = out[i] + 32
	} 
    }
    return string(out),nil
}
