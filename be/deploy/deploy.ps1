echo "Deploying..."

ssh commonuser@192.168.0.105 "mkdir C:\Users\commonuser\Desktop\TMS"
ssh commonuser@192.168.0.105 "mkdir C:\Users\commonuser\Desktop\TMS\BE"

ssh commonuser@192.168.0.105 "mkdir C:\Users\commonuser\Desktop\be-temp"
ssh commonuser@192.168.0.105 "mkdir C:\Users\commonuser\Desktop\be-temp\backup"

scp 192.168.0.102:/C:\Users\commonuser\Desktop\bin_repo\REL_BE_0_3_1.tar.gz 192.168.0.105:/C:\Users\commonuser\Desktop\be-temp\backup


ssh commonuser@192.168.0.105 "cd C:\Users\commonuser\Desktop\be-temp\backup && tar -xvzf REL_BE_0_3_1.tar.gz"

ssh commonuser@192.168.0.105 "powershell -Command "Copy-Item -Path "C:/Users/commonuser/Desktop/be-temp/backup/be-test-server/*" -Destination "C:/Users/commonuser/Desktop/TMS/BE" -Recurse""

ssh commonuser@192.168.0.105 "powershell -Command "npx kill-port 8080""

ssh commonuser@192.168.0.105 "cd C:\Users\commonuser\Desktop\TMS && echo cd C:\Users\commonuser\Desktop\TMS\BE\bin\src > runbe.ps1 && echo ./backend.exe >> runbe.ps1" 
## copy config folder over to TMS folder
# ssh commonuser@192.168.0.101 "powershell -Command "Copy-Item -Path "C:/Users/commonuser/Desktop/be-temp/backup/be-test-server/config" -Destination "C:/Users/commonuser/Desktop/TMS/BE" -Recurse""
