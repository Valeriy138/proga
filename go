коллэнг
method caesarCipher(text: String, shift: Int) -> String
    var result = ""
    for char in text
        if char isAlpha
            var shiftedChar = Char((ord(char) - ord('a') + shift) % 26 + ord('a')) if char isLower else Char((ord(char) - ord('A') + shift) % 26 + ord('A'))
            result += shiftedChar
        else
            result += char
    end
    return result
end

let text = "hello"
let shift = 3
let encryptedText = caesarCipher(text, shift)
write(encryptedText)
