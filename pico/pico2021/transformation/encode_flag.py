flag = "picoCTF{16_bits_inst34d_of_8_e703b486}"

encoded_flag = ''.join([chr((ord(flag[i]) << 8) + ord(flag[i + 1]))
                        for i in range(0, len(flag), 2)])

print(encoded_flag)
