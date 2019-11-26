CURRENT_PROJECT_DIR=${PWD##*/}
echo $CURRENT_PROJECT_DIR
docker image build -t sandwhichmaker$CURRENT_PROJECT_DIR:latest
docker container run --name sandwhichmaker$CURRENT_PROJECT_DIR:latest