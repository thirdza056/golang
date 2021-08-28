import os, time

botStart = time.time()
try:os.system('screen -S gobot01 -X quit');os.system('rm -rf clone/gobot01')
except:pass
try:os.system('screen -dmS gobot01');os.system('screen -r gobot01 -X stuff "./new a1 \n"')
except:pass

botStart = time.time()
try:os.system('screen -S gobot02 -X quit');os.system('rm -rf clone/gobot02')
except:pass
try:os.system('screen -dmS gobot02');os.system('screen -r gobot02 -X stuff "./new a2 \n"')
except:pass

botStart = time.time()
try:os.system('screen -S gobot03 -X quit');os.system('rm -rf clone/gobot03')
except:pass
try:os.system('screen -dmS gobot03');os.system('screen -r gobot03 -X stuff "./new a3 \n"')
except:pass

botStart = time.time()
try:os.system('screen -S gobot04 -X quit');os.system('rm -rf clone/gobot04')
except:pass
try:os.system('screen -dmS gobot04');os.system('screen -r gobot04 -X stuff "./new a4 \n"')
except:pass
botStart = time.time()
