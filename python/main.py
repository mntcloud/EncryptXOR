# EncryptXOR.py
# Getting text input
text = input('-> ').encode('ascii')
# Translating sequence of bytes in a list of intergers
raw = list(text)

# Main process: encrypting
for i in range(len(raw)):
    # 0x2F is our key
    # raw[i] is a character
    raw[i] ^= 0x2F

# Decoding our encrypted text in string
# And print in output result
print(bytes(raw).decode(encoding='utf-8'))