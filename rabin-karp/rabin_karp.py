PRIME = 3

def rolling_hash(text: str):
    text_hash = 0

    for idx, char in enumerate(text):
        text_hash += ord(char) * (PRIME ** idx)

    return text_hash


def rabin_karp(text: str, pattern: str):
    pattern_hash = rolling_hash(pattern)

    text_length = len(text)
    pattern_length = len(pattern)

    i = 0
    text_window_hash = 0

    while (i + pattern_length - 1) < text_length:
        text_window = text[i:i+pattern_length]

        if i == 0:
            text_window_hash = rolling_hash(text_window)
        else:
            old_char = text[i-1]
            text_window_hash = (text_window_hash - ord(old_char)) // PRIME
            new_char = text[i+pattern_length - 1]
            other_part = PRIME ** (pattern_length - 1)
            text_window_hash += ord(new_char) * other_part

        if text_window_hash == pattern_hash and text_window == pattern:
            return True

        i += 1

    return False

if __name__ == "__main__":
    # pattern = "abc"
    # s = "abedabc"

    # pattern = "aab"
    # s = "aaaaab"

    # pattern = "bce"
    # s = "abcdabce"

    pattern = "dba"
    s = "ccaccaaedba"

    print(f"Is \"{pattern}\" in {s}?:", rabin_karp(s, pattern))