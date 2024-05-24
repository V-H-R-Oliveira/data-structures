def naive_search(string: str, pattern: str):
    string_length = len(string)
    pattern_length = len(pattern)

    i = 0
    j = 0
    restart_at = 1

    while restart_at < string_length:
        if j  == pattern_length:
            return True

        if i > string_length - 1:
            break

        char_string = string[i]
        char_pattern = pattern[j]

        if char_string == char_pattern:
            i += 1
            j += 1
            continue

        i = restart_at
        j = 0
        restart_at += 1

    return False

# Builds the state machine for efficient pattern matching
def build_pattern_table(pattern: str):
    pattern_length = len(pattern)
    table = [0] * pattern_length

    j = 0
    i = 1

    while True:
        if i == pattern_length - 1:
            char_i = pattern[i]
            char_j = pattern[j]

            if char_i == char_j:
                table[i] = j + 1
                break

            if j == 0:
                break

            prev_table_value = table[j-1]
            j = prev_table_value
            continue

        char_i = pattern[i]
        char_j = pattern[j]

        if char_i != char_j and j > 0:
            j -= 1
            continue

        if char_i != char_j and j == 0:
            table[i] = 0
            i += 1
            continue

        if char_i == char_j:
            table[i] = j + 1
            j += 1
            i += 1

    return table

# Knuth Morris Pratt string matching algorithm
def kmp(s: str, pattern: str):
    string_idx = 0
    pattern_idx = 0

    s_len = len(s)
    pattern_len = len(pattern)
    pattern_table = build_pattern_table(pattern)

    while string_idx < s_len and pattern_idx != pattern_len:
        char_str = s[string_idx]
        char_pattern = pattern[pattern_idx]

        if char_str == char_pattern:
            string_idx += 1
            pattern_idx += 1
            continue

        if char_str != char_pattern and pattern_idx != 0:
            prev_table_value = pattern_table[pattern_idx-1]
            pattern_idx = prev_table_value
            continue

        if char_str != char_pattern and pattern_idx == 0:
            string_idx += 1

    return pattern_idx == pattern_len

if __name__ == '__main__':
    s = "abxabcabcaby"
    pattern = "abcaby"

    print(kmp(s, pattern))
