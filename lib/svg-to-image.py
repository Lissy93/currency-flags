from os import listdir
from os.path import isfile, join
from cairosvg import svg2png

svgSrcPath = '../assets/flags_svg/'
pngDistPath = '../assets/flags_test_png/'

onlyfiles = [f for f in listdir(svgSrcPath) if isfile(join(svgSrcPath, f))]

# If output directory doesn't yet exist, create it
if not os.path.exists(pngDistPath):
  print('Creating directory for results, in '+pngDistPath)
  os.makedirs(pngDistPath)

# For each file name, read file, convert to PNG, and write to output dir
for file in onlyfiles:
  print('Converting '+file+' to PNG...')
  svg2png(url=svgSrcPath+file, write_to=pngDistPath+file+'.png', scale=50)
