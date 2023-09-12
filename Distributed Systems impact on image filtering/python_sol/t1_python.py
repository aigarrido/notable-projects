import argparse
parser = argparse.ArgumentParser()
from PIL import Image, ImageFilter

# parsing and formatting
parser.add_argument("-k", "--kernel", help="Kernel")
parser.add_argument("-i", "--input", help="Input Path")
parser.add_argument("-o", "--output", help="Output path")
args = parser.parse_args()


# open image

img = Image.open(args.input)


# kernels:
# sharpen (0, -1, 0, -1, 5, -1, 0, -1, 0)
# ridge (-1, -1, -1, -1, 8, -1, -1, -1, -1)
# gaussian blur (0.0625,0.125,0.0625,0.125,0.25,0.125,0.0625,0.125,0.0625)
# box blur (1/9,1/9,1/9,1/9,1/9,1/9,1/9,1/9,1/9)


kernel_list= [(0, -1, 0, -1, 5, -1, 0, -1, 0),(-1, -1, -1, -1, 8, -1, -1, -1, -1),(1/9,1/9,1/9,1/9,1/9,1/9,1/9,1/9,1/9)]

# kernel 0:sharpen 1:ridge 2:box blur
imgo = img.filter(ImageFilter.Kernel((3, 3),kernel_list[int(args.kernel)], 1, 0))

# save the image

imgo.save(args.output)