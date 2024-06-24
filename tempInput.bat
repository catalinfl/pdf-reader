@echo off 
color 07 
title pdfreader

:: START ECHO COMMANDS
:Page1
echo ULTRAS BRAILA  
echo Word document  
echo Stadin alone all alone again  
echo When alone is again and again  
echo Basarabia pamant romanesc  
echo Wismut aue  
echo Wismut aue FCC  
echo Ett  
echo E  
echo Awetawt  
echo Taewwtateaw  
echo teawtaw  
echo etawtwet  
echo awetwett eawt we t 
echo awteat teaw tawe eawt  
echo etawtawt  
echo teawtaw  
echo fdfs  
echo fsf  
echo dsf  
goto :end
:Page2
echo warawrawr  
echo gfdgdfgd  
echo gsafas 
echo pam pam  
echo papmmpmpmpmp 
echo BAMBLONAA4 
:: END ECHO COMMANDS
@REM set /p UserInput=
@REM if "%UserInput%"=="exit" goto :end
:loop
set /p Input=
if "%Input%"=="exit" goto :end
goto :loop

:end
pause
