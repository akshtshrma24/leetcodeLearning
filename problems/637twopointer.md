# 637

real end game

```py
class Solution:
    """
    @param word: a non-empty string
    @param abbr: an abbreviation
    @return: true if string matches with the given abbr or false
    """
    def valid_word_abbreviation(self, word: str, abbr: str) -> bool:
        # write your code here

        word_pointer = 0
        abbr_pointer = 0

        while word_pointer < len(word) and abbr_pointer < len(abbr):
            if abbr[abbr_pointer].isdigit():
                number = ""

                if abbr[abbr_pointer] == "0":
                    return False
                
                while abbr_pointer < len(abbr) and abbr[abbr_pointer].isdigit():
                    number += abbr[abbr_pointer]
                    abbr_pointer += 1
                
                print(number)
                word_pointer += int(number)
            else:
                if word[word_pointer] != abbr[abbr_pointer]:
                    return False
                
                word_pointer += 1
                abbr_pointer += 1
            
        return word_pointer == len(word) and abbr_pointer == len(abbr)
```