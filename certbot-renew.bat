set isNginxStarting = "false"

@REM go run 결과를 isNginxStarting 변수에 저장
FOR /F "tokens=* USEBACKQ" %%a IN (`go run %~dp0is-nginx-starting\is-nginx-starting.go`) DO (
SET isNginxStarting=%%a
)
@REM echo isNginxStarting
@REM ECHO %isNginxStarting%

@REM nginx 가 실행중이라면 nginx 종료하기
pushd C:\nginx-1.20.1
if "%isNginxStarting%" == "true" nginx.exe -s quit 
popd

@REM certbot renew 진행하기
certbot renew

@REM nginx 가 실행중이었다면 nginx 재실행하기
if "%isNginxStarting%" == "true" %~dp0nginx-start.vbs

pause