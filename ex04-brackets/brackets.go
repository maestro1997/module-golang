package brackets

func Bracket(s string) (bool,error) {
    i:=0
    br1:=0
    br2:=0
    br3:=0
    last:=0
    for ; i < len(s); i++ {
	switch s[i] {
	case '[':
	    br1+=1
	    last = i
        case ']' :
	    br1-=1
	    if (s[last] == '(') || (s[last] == '{') {
	        return false,nil
	    }
	    last = i
	case '(' :
	    br2+=1
	    last = i
        case ')' :
	    br2-=1
            if (s[last] == '{') || (s[last] == '[') {
	        return false,nil
	    }
	    last = i
	case '{':
	    br3+=1
	    last = i
        case '}' :
	    br3-=1
	    if (s[last] == '(') || (s[last] == '[') {
                return false,nil
	    }
	    last = i
        }
	if (br1 < 0) || (br2 < 0) || (br3 < 0) {
            return false,nil
	}
    }
    return  (br1 == 0) && (br2 == 0) && (br3 == 0),nil
}
