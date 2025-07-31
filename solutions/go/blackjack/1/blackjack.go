package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card{
        case "ace": return	11	
    	case "eight": return	8
		case "two": return	2	
    	case "nine": return	9
		case "three": return	3	
    	case "ten": return	10
		case "four": return	4	
    	case "jack": return	10
		case "five": return	5	
    	case "queen": return	10
		case "six": return	6	
    	case "king"	: return 10
		case "seven": return	7	
    	case "other": return	0
    }
    return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    val1 := ParseCard(card1)
    val2 := ParseCard(card2)
    dealerVal := ParseCard(dealerCard)
	switch {
        case val1 == 11 && val2 == 11: return "P"
        case val1 + val2 == 21:{
            if dealerVal < 10{
                return "W"
            } else {
                return "S"
            }
        }
        case 17 <= val1 + val2 && val1 + val2 <= 20: return "S"
        case 12 <= val1 + val2 && val1 + val2 <= 16:{
            if dealerVal >= 7{
                return "H"
            }else {
                return "S"
            }
        }
        default: return "H"
    }
    return "H"
}
