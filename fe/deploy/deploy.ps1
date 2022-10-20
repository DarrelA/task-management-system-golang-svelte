ssh commonuser@192.168.0.105 "mkdir C:\Users\commonuser\Desktop\TMS"
ssh commonuser@192.168.0.105 "mkdir C:\Users\commonuser\Desktop\fe-temp"
ssh commonuser@192.168.0.105 "mkdir C:\Users\commonuser\Desktop\fe-temp\backup"

scp 192.168.0.102:/C:\Users\commonuser\Desktop\bin_repo\REL_FE_0_4_1.tar.gz 192.168.0.105:/C:\Users\commonuser\Desktop\fe-temp\backup
ssh commonuser@192.168.0.105 "cd C:\Users\commonuser\Desktop\fe-temp\backup && tar -xvzf REL_FE_0_4_1.tar.gz"

ssh commonuser@192.168.0.105 "powershell -Command "Copy-Item -Path "C:/Users/commonuser/Desktop/fe-temp/backup/fe-test-server/bin/FE_src" -Destination "C:/Users/commonuser/Desktop/TMS/FE" -Recurse""

# ssh commonuser@192.168.0.105 "powershell -Command "npx kill-port 3000""
ssh commonuser@192.168.0.105 "$procid=(Get-Process "node").id && kill $procid"

ssh commonuser@192.168.0.105 "cd C:/Users/commonuser/Desktop/TMS/FE && npm install"

ssh commonuser@192.168.0.105 "cd C:\Users\commonuser\Desktop\TMS && echo cd C:\Users\commonuser\Desktop\TMS\FE >> runfe.ps1 && echo npm run dev >> runfe.ps1"
