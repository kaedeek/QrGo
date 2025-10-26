from VersaLog import *
import qrcode
import sys

logger = VersaLog(enum="detailed", tag="Base", show_tag=True)

def qr_created(url):
    try:
        qr = qrcode.make(url)
        file_path = "qr.png"
        qr.save(file_path)
        return file_path
    except Exception as e:
        return str(e)
    
if __name__ == "__main__":
    if len(sys.argv) > 1:
        result = qr_created(sys.argv[1])
        logger.info(f"successfully: {result}")
    else:
        logger.warning("Specify the URL as an argument.")
