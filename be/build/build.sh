mkdir ../../be-test-server

mkdir ../../be-test-server/deploy
mkdir ../../be-test-server/bin
mkdir ../../be-test-server/config

cd C:/Users/commonuser/Desktop/release/REL_BE_0_3_1
cp ./config/.env ../be-test-server/config
cp ./config/schema.sql ../be-test-server/config

cp -r ./src ../be-test-server/bin
cp ./deploy/deploy.ps1 ../be-test-server/deploy

# tar be-test-server move to bin_repo
tar -zcvf REL_BE_0_3_1.tar.gz ../be-test-server

mv REL_BE_0_3_1.tar.gz C:/Users/commonuser/Desktop/bin_repo
