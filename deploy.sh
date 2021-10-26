git pull
go build
pkill -f GoGistKa
rm /home/gist/GoGistKakaoServer
mv GoGistKakaoServer /home/gist/
nohup /home/gist/GoGistKakaoServer &
