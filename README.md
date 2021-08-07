# Window 용 Certbot Renew bat 파일
리눅스에서 cron 을 이용해 certbot renew 를 자동으로 동작시킬 수 있는 것처럼, 윈도우에서도 작업스케줄러와 bat 파일을 이용해 certbot renew 를 자동으로 동작시킬 수 있게 작성한 bat 파일 입니다. (웹 서비스로 nginx 를 사용한다고 가정하고 작성되었습니다. apache나 다른 웹서버인 경우 적절히 수정해서 사용하시면 됩니다.)
<br>
<br>

# 사용 방법
1) window 에 golang 설치 (https://golang.org/)
2) 이 프로젝트를 git clone 받는다.
3) certbot-renew.bat, nginx-start.bat 파일에 있는 nginx 설치 경로("C:\nginx-1.20.1")를 본인 PC에 맞게 수정한다.
4) window 작업스케줄러에 certbot-renew.bat 을 등록한다. (최고 관리자 권한으로 실행될 수 있도록 한다.)