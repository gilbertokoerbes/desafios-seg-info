def string_to_hex(input_string):
    # Convert the input string to its hexadecimal representation
    hex_string = input_string.encode('utf-8').hex()
    return hex_string

def hex_xor(hex_string, hex_key):
    # Ensure the key and string are of the same length
    if len(hex_string) != len(hex_key):
        raise ValueError("Key and string must be of the same length for XOR operation")

    # Perform the XOR operation
    result = ''.join(hex(int(a, 16) ^ int(b, 16))[2:] for a, b in zip(hex_string, hex_key))

    return result

# Your plaintext string
plaintext = "deposite 800"

# Your hexadecimal key in plain text
hex_key_text = "7c46a71d60f6445bdfdfb144"

# Convert the key to binary
key_binary = bytes.fromhex(hex_key_text).hex()

# Convert the plaintext to hexadecimal
hex_plaintext = string_to_hex(plaintext)

# Perform XOR operation
result = hex_xor(hex_plaintext, key_binary)

print("Original Text:", plaintext)
print("Hexadecimal Key:", hex_key_text)
print("Hexadecimal Plaintext:", hex_plaintext)
print("XOR Result:", result.upper())
