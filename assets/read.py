from pylibdmtx.pylibdmtx import decode
from PIL import Image
import sys

res = decode(Image.open(sys.argv[1]))
print(res[0].data)