# convert FE_A2_1 file to tar
# cd ./FE_A2_1/FE_src/ && tar -zcvf ../src.tgz . && cd ..

# cp deploy config

mkdir ./fe-test-server

mkdir ./fe-test-server/deploy
mkdir ./fe-test-server/bin

# cp tar file to bin
cp -r ../FE_src ./fe-test-server/bin
cp ../deploy/deploy.ps1 ./fe-test-server/deploy

# tar test-server move to bin_repo
tar -zcvf REL_FE_0_4_1.tar.gz fe-test-server

mv REL_FE_0_4_1.tar.gz C:/Users/commonuser/Desktop/bin_repo
