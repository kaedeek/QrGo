import qrcode
import sys

def qr_created(url):
    try:
        qr = qrcode.make(url)
        file_path = "qr.png"
        qr.save(file_path)
        return file_path  # ファイルパスを返す
    except Exception as e:
        return str(e)  # エラーメッセージを文字列で返す
    
if __name__ == "__main__":
    if len(sys.argv) > 1:
        result = qr_created(sys.argv[1])
        print(f"successfully: {result}")
    else:
        print("Specify the URL as an argument.")
