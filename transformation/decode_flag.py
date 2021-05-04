def read_encoded_string() -> str:
    encoded_string = ""
    with open("enc", "r") as f:
        encoded_string = f.readline()
    return encoded_string


def get_flag(encoded_string: str) -> str:
    flag = ''
    for i in range(len(encoded_string)):
        flag += chr(ord(encoded_string[i]) >> 8)
        flag += chr((ord(encoded_string[i])) -
                    ((ord(encoded_string[i]) >> 8) << 8))

    return flag


def main():
    flag = get_flag(read_encoded_string())
    print("Flag:", flag)


if __name__ == '__main__':
    main()
